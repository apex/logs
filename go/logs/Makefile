
all: client.go

client.go: $(SCHEMA)
	@echo "==> create $@"
	@rpc-go-client \
		-schema $(SCHEMA) \
		-package logs \
	  > $@

clean:
	@rm -f client.go
.PHONY: clean