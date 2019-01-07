bin-prep:
	GO111MODULE=off go get -u github.com/kevinburke/go-bindata/...

bin-migrations:
	go-bindata -o data/migrate.go -prefix data/migrations/ data/migrations/
	sed -i "s/package main/package data/" data/migrate.go

# at this time, there's no watch enabled for the go binary
dev-be:
	go install ./audiofile/ && audiofile

dev-fe:
	yarn serve

install:
	yarn build
	go install ./audiofile/

test:
	go test -v -race ./...

