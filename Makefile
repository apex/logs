
all: client.go

client.go: $(SCHEMA)
	@echo "==> create $@"
	@rpc-go-client \
		-schema $(SCHEMA) \
		-package logs \
		| gofmt > $@

clean:
	@rm -f client.go
.PHONY: clean