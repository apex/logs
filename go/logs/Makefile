
all: types.go client.go

types.go: $(SCHEMA)
	@echo "==> create $@"
	@apigen \
		--schema $(SCHEMA) \
	  --output go.types > $@

client.go: $(SCHEMA)
	@echo "==> create $@"
	@apigen \
		--schema $(SCHEMA) \
	  --output go.client > $@

clean:
	@rm -f types.go client.go
.PHONY: clean