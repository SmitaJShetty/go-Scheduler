APPNAME= scheduler

.PHONEY: build 
bld: 
		mkdir -p build
		echo $(APPNAME)
		cd cmd/ && go build -o ../build/$(APPNAME)

build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 go build -a -installsuffix cgo -o build/$(APPNAME)
	redis-server &
	./$(APPNAME)


clean: ## Remove all the temporary and build files
	go clean

redis-start:
	docker pull redis
	docker run --name redis-test-instance -p 6379:6379 -d redis

run: bld
	build/$(APPNAME)