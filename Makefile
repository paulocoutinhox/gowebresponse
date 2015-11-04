GOFMT=gofmt

GOFILES=\
	gowebresponse.go\

format:
	${GOFMT} -w gowebresponse.go
	${GOFMT} -w gowebresponse_test.go

test:
	go test -v