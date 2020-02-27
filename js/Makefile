
all: dist dist.browser

dist: src/index.ts
	@tsc -p tsconfig.json

dist.browser: src/index.ts
	@tsc -p tsconfig.browser.json

src/index.ts: $(SCHEMA)
	@echo "==> create $@"
	@rpc-ts-client \
		-schema $(SCHEMA) \
	  > $@

test:
	@ts-node -O '{ "target": "es6" }' test.ts
.PHONY: test

clean:
	@rm -f src/index.ts
	@rm -fr dist.browser
	@rm -fr dist
.PHONY: clean