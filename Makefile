.PHONY: test test-dfa test-regexp bench clean tools pdf

tools:
	go install golang.org/x/perf/cmd/benchstat@latest

pdf:
	pdflatex readme.tex && pdflatex readme.tex

test-regexp:
	go test -v

test-dfa:
	go test -v -tags=dfa

test: test-regexp test-dfa

bench:
	go test -bench=. -benchmem -count=10 | tee bench-regexp.out
	go test -bench=. -benchmem -count=10 -tags=dfa | tee bench-dfa.out
	benchstat bench-regexp.out bench-dfa.out

clean:
	go clean -testcache
	rm -f readme.aux readme.log readme.out readme.toc readme.pdf

