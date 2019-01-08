bin-prep:
	GO111MODULE=off go get -u github.com/kevinburke/go-bindata/...

bin-migrations:
	go-bindata -o data/migrate.go -prefix data/migrations/ data/migrations/
	sed -i "s/package main/package data/" data/migrate.go

bin-dist:
	go-bindata -o audiofile/assets.go -prefix dist/ dist/...

# at this time, there's no watch enabled for the go binary
dev-be: bin-prep bin-migrations bin-dist install
	audiofile

# serve is a watch task with built-in node server
dev-fe:
	yarn serve

install:
	yarn build
	go install ./audiofile/

test: bin-migrations bin-dist install
	go test -v -race ./...

