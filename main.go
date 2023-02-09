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

func test() {
	os.Setenv("db_user", "postgres")
	os.Setenv("db_name", "postgres")
	os.Setenv("db_password", "password")
}

type DataBase struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (db *DataBase) infoDb(host string) *DataBase {
	db.host = host
	db.port = 5432
	db.user = os.Getenv("db_user")
	db.password = os.Getenv("db_password")
	db.dbname = os.Getenv("db_name")
	return db
}

type Query struct {
	Id   int `json:int`
	Name int `json:"name"`
}

func (db *DataBase) db_query(query string) *sql.Rows {
	// db.infoDb("172.17.0.2")
	db.infoDb("localhost")

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
	test()
	database := new(DataBase)
	sql_query := database.db_query("Select count(*) from clm_resln_email_send_status where email_send_stat = 'PENDING';")
	sql_query2 := database.db_query("Select count(distinct email_send_stat_id) from clm_resln_email_send_status where email_send_stat = 'PENDING';")
	sql_query3 := database.db_query("Select count(*) from onchain_sync where synced = 'N';")
	// sql_query4 := database.db_query("Select  date_crtd from work_center_email_asgnmt order by date_crtd desc limit 1;")
	// sql_query4 := database.db_query("Select extract(minute from date_crtd - date_rcvd ) from work_center_email_asgnmt order by date_rcvd desc limit 1;")

	res := database.pending_batch(sql_query)
	check_results(res)
	res2 := database.pending_batch(sql_query2)
	check_results(res2)
	res3 := database.pending_batch(sql_query3)
	check_results(res3)
	// res4 := database.pending_batch(sql_query4)
	// fmt.Printf("%d", res4)
	// check_results(res4)
}

func runCronJob(t int) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(t).Minute().Do(func() {
		job()
	})
	s.StartBlocking()
}

func main() {
	runCronJob(5)
}
