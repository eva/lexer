.PHONY: default
default: \
	build

.PHONY: build
build: \
	test \
	bench

.PHONY: test
test:
	go test -cover ./src/...

.PHONY: bench
bench:
	go test -bench="." -benchmem ./src/...
