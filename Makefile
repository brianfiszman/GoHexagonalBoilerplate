.PHONY: run-connector
run-connector: 
	env `cat .env` go run ./pkg/main.go