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

1. create table nameOftable (id int primary key not null, nameOfcomun varchar(40) not null ); 
1.1 alter table nameOftable add column nameOfcomun varchar(50) ;
1.2 insert into  nameOftable values(1, 'PENDING');

2. create table nameOftable (id int primary key not null, nameOFcolumn varchar(40) not null ); 
2.1 alter table nameOftable add column nameOFcolumn varchar(50) ;
2.2 insert into  nameOftable values(1, 'N');
