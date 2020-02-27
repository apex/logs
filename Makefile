
all: docs

docs: $(SCHEMA)
	@echo "==> create $@"
	@rpc-md-docs -schema $(SCHEMA)

clean:
	@rm -fr docs
.PHONY: clean