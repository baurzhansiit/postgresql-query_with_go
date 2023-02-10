# Go utility to run sql query
# Script run sql query via postgres database and retrieve int results to trigger warrings

Prerequizite:
- postgres database, client
- golang 

## Step by step
1. Create postgresql database engine by running docker container, official site: www.dockerhub.com 

2. Connect to postgresql by runnig command: 
``` sql postgresql://postgres:password@localhost:5432/postgres ```
3. Run utility: 
``` go run . ```


run with docker 
<!-- build image -->
./makefile build
<!-- create container -->
./makefile run
<!-- for debug purpose  -->
./makefile debug
<!-- when done, clean by running -->
./makefile clean



<!-- for test purpose building table  -->

1. create table clm_resln_email_send_status (id int primary key not null,               email_send_stat varchar(40) not null ); 
1.1 alter table clm_resln_email_send_status add column email_send_stat_id varchar(50) ;
1.2 insert into  clm_resln_email_send_status values(1, 'PENDING');

2. create table onchain_sync (id int primary key not null,synced varchar(40) not null ); 
2.1 alter table onchain_sync add column synced varchar(50) ;
2.2 insert into  onchain_sync values(1, 'N');
