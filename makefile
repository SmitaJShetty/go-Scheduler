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

db-migrate-up:
	docker run -v postgres-data:/var/lib/postgresql/dat/lo0a migrate/migrate -source file://migrations/ -database postgres://0.0.0.0:5432/scheduler up

db-migrate-down:
	migrate -path ./internal/db/migrations -database postgres://smita:pwd123@0.0.0.0:5432/scheduler?sslmode=disable down

run: bld
	build/$(APPNAME)

compose-upd:
	cd ./deployments && docker-compose up -d 
	db-migrate


compose-down:
	cd ./deployments && docker-compose down