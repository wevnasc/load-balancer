# load-balancer
This is a load balancer example, for this example, I am using Nginx as the load balancer and a web server wroten in go 

three replicated machines were created to handle the load and the nginx uses the round robin strategy to determine which machine to hit for each request.

## Stack
- go
- nginx
- docker

## how to run
```sh 
$ docker compose build
$ docker compose up -d
``` 

go to [http://localhost:8080](http://localhost:8080g)