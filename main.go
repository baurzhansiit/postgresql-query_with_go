package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
)

type DataBase struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (db *DataBase) infoDb(host, user, dbname, password string) *DataBase {
	db.host = host
	db.port = 5432
	db.user = user
	db.password = password
	db.dbname = dbname
	return db
}

type Query struct {
	Id   int `json:int`
	Name int `json:"name"`
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) db_query(query string) *sql.Rows {
	user, err := os.ReadFile("./db_creds/username")
	errorHandler(err)
	password, err := os.ReadFile("./db_creds/password")
	errorHandler(err)
	dbname, err := os.ReadFile("./db_creds/database")
	errorHandler(err)

	db.infoDb("postgres-svc", string(user), string(dbname), string(password))
	// db.infoDb("localhost")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.user, db.password, db.dbname)

	db_query, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db_query.Close()
	// Checking connection
	err = db_query.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected/quering  %s database \n", db.dbname)
	sql_query, err := db_query.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return sql_query
}

func (db *DataBase) pending_batch(row *sql.Rows) int {
	var name int
	for row.Next() {
		var q Query
		err := row.Scan(&q.Id)
		if err != nil {
			log.Fatal(err)
		}
		row.Close()
		name = q.Id
	}
	return name
}

func check_results(res int) {
	date := time.Now()
	switch res {
	case 0:
		fmt.Printf("Passed synced: %d\n", res)

	default:
		// send notifications
		//  pass logs
		fmt.Printf("%s WARRNING: 'Not Passed', res: %d\n", date.Format("2006-01-02 15:04:05 Mon"), res)
	}
}

func job() {
	database := new(DataBase)
	sql_query := database.db_query("Select count(*) from clm_resln_email_send_status where email_send_stat = 'PENDING';")
	sql_query2 := database.db_query("Select count(distinct email_send_stat_id) from clm_resln_email_send_status where email_send_stat = 'PENDING';")
	sql_query3 := database.db_query("Select count(*) from onchain_sync where synced = 'N';")

	res := database.pending_batch(sql_query)
	check_results(res)
	res2 := database.pending_batch(sql_query2)
	check_results(res2)
	res3 := database.pending_batch(sql_query3)
	check_results(res3)

}

func runCronJob(t int) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(t).Second().Do(func() {
		job()
	})

	s.StartBlocking()
}

func main() {
	runCronJob(5)
}
