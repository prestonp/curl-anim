VHOST := lol.preston.io
NAME := curl-anim

.PHONY: build
build:
	GOOS=linux go build -o bin/curl-anim .

.PHONY: image
image:
	docker build -t curl-anim .

.PHONY: nginx
nginx:
	docker run --rm -d --name nginx-proxy -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock:ro jwilder/nginx-proxy

.PHONY: run
run:
	docker run --rm -d --name ${NAME} -e VIRTUAL_HOST=${VHOST} -p 8081:8081 ${NAME} 
