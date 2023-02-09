# Go utility to run sql query

Prerequizite:
- postgres client
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