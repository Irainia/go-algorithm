ensure:
	dep ensure

test: ensure
	go test ./... --cover

test_coverage: ensure
	go test -coverprofile test_coverage.html ./... && go tool cover -html=test_coverage.html

build: ensure
	go build -o ./out/go-algorithm .

run:
	./out/go-algorithm

install: build
	go install
