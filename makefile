#!/bin/bash

export db_user='postgres'
export db_name='postgres'
export db_password='password'


# docker build -t golang:v1  --build-arg db_user=${db_user} --build-arg db_name=${db_name} --build-arg db_password=${db_password}  .
function build()(
	docker build -t golang:v1  .
	
)
function run()(
	docker run -d --name go -w /app golang:v1
)

function debug()(
	docker run -d --name go -v $(pwd)/root golang:v1 sleep 1000
)


function exec()(
	docker exec -it go sh
)

function clean()(
	docker rm -f go
)
function cleanI()(
	docker rmi -f baurzhansiit/golang:v1
	docker rmi -f golang:v1
	)

function push(){
	docker tag golang:v1 baurzhansiit/golang:v1
	docker push baurzhansiit/golang:v1
}


"$@"