.PHONY: examples clean test

EXDIRS = $(wildcard _examples/*/*)
%.ex_run : %
	@echo Building $<
	@cd $< && go run *.go && touch .ex_run
EXRUN = $(addsuffix /.ex_run, $(EXDIRS))

examples: $(EXRUN)
	
clean: 
	rm -f $(EXRUN)


GOPKGS = $(shell go list ./...)
test:
	@echo "mode: count" > coverage-all.out
	@$(foreach PKG, $(GOPKGS), \
		go test -coverprofile=coverage.out $(PKG) || exit 1; \
		tail -n +2 coverage.out >> coverage-all.out; )
	@go tool cover -func=coverage-all.out
