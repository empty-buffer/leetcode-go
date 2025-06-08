.PHONY: run_test
run_test:
	go test -v $(src) -run $(test)
# go test -v ./easy/minimum_waiting_time -run Test_MinimumWaitingTime
