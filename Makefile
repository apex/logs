
all: Readme.md

Readme.md: $(SCHEMA)
	@echo "==> create $@"
	@apigen \
		--schema $(SCHEMA) \
	  --output md.docs > $@

clean:
	@rm -f Readme.md
.PHONY: clean