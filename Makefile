GO ?= CGO_ENABLED=0 go
CPU_NAME := $(shell go run cmd/cpuname/main.go)
BENCH_FILE := benches/$(shell go env GOOS)-$(shell go env GOARCH)-$(CPU_NAME).txt

.PHONY: all
all: tidy test

.PHONY: clean
clean:
	$(GO) clean # remove test results from previous runs so that tests are executed
	-rm \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		govulncheck.sarif \
		junit.xml \
		README.html \
		golangci-lint.json \
		staticcheck.json \
		test.log

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: staticcheck
staticcheck:
	which staticcheck || $(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -version

.PHONY: tidy
tidy: staticcheck
	test -z $(gofmt -l .)
	$(GO) vet
	staticcheck
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest --version
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

cpu.profile:
	$(GO) test -run=^$ -bench=Day10Part1$ -benchmem -memprofile mem.profile -cpuprofile $@

.PHONY: prof
prof: cpu.profile
	$(GO) tool pprof $^

.PHONY: test
test:
	$(GO) test -run=Day -short -vet=all

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml

# write coverage to stdout and to test.log
coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

# Gitlab test report
junit.xml: test.log
	which go-junit-report || $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	go-junit-report -version
	go-junit-report < $< > $@

# Gitlab coverage report
coverage.xml: coverage.txt
	which gocover-cobertura || $(GO) install github.com/boumenot/gocover-cobertura@latest
	gocover-cobertura < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: golangci-lint.json
	which golint-convert || $(GO) install github.com/banyansecurity/golint-convert@latest
	golint-convert < $< > $@

golangci-lint.json:
	-$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --out-format json > $@

staticcheck.json: staticcheck
	-staticcheck -f json > $@

# Gitlab dependency report
govulncheck.sarif:
	which govulncheck || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -version
	govulncheck -format=sarif ./... > $@

$(BENCH_FILE): $(wildcard *.go)
	echo "Running benchmarks and saving to $@..."
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@

README.html: README.adoc
	asciidoc $^

.PHONY: total
total: $(BENCH_FILE)
	awk -f total.awk < $(BENCH_FILE)
