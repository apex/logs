
all: docs

docs: $(SCHEMA)
	@echo "==> create $@"
	@rpc-docs

clean:
	@rm -fr docs
.PHONY: clean