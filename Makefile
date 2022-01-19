test:
	go test -count=2 .

coverage:
	go test -coverprofile=coverage.out .
	go tool cover -html=coverage.out
