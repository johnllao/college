build-webserver:
	go build -o bin/webserver github.com/johnllao/college/webserver

run-webserver:
	bin/webserver -root /Users/johnllao/src/college/apps/hello/

build-helloweb:
	go build -o bin/helloweb github.com/johnllao/college/cmd/helloweb

run-helloweb:
	bin/helloweb -root /Users/johnllao/src/college/apps/hello/
