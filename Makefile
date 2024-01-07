GOCMD=go

run-no-lock-no-wg:
	$(GOCMD) run 01-no-locks/no-wg/main.go

run-no-lock-wg:
	$(GOCMD) run 01-no-locks/wg/main.go

run-mutex:
	$(GOCMD) run 02-mutex/main.go

run-rwmutex:
	$(GOCMD) run 03-rwmutex/main.go