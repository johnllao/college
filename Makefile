build-webserver:
	go build -o bin/webserver github.com/johnllao/college/webserver

run-webserver:
	bin/webserver -root /Users/johnllao/src/college/apps/hello/
