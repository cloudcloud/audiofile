
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

