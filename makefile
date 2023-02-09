#!/bin/bash


function build()(
	docker build -t goland:v1 .
)
function run()(
	docker run -d --name go -w /app goland:v1
)

function debug()(
	docker run -d --name go -v $(pwd)/root goland:v1 sleep 1000
)


function exec()(
	docker exec -it go sh
)

function clean()(
	docker rm -f go
)



"$@"