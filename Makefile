.PHONY: example
example:
	go run ./mockgen/ --source=sample/mockio/interface.go -destination=sample/mockio/mock/mock.go