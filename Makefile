GO ?= CGO_ENABLED=0 go
CPU_NAME := $(shell $(GO) run ./cmd/cpuname)
BENCH_FILE := benches/$(shell $(GO) env GOOS)-$(shell $(GO) env GOARCH)-$(CPU_NAME).txt

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: all
all: tidy test ## Run tidy and test

.PHONY: clean
clean: ## Remove generated files
	$(GO) clean
	-rm -f \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		golangci-lint.json \
		govulncheck.sarif \
		junit.xml \
		test.log

.PHONY: bench
bench: ## Run benchmarks
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem

.PHONY: tidy
tidy: ## Format check and lint
	test -z "$$(gofmt -l .)"
	$(GO) vet
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: test
test: ## Run all tests
	$(GO) test -short

$(BENCH_FILE): $(wildcard *.go)
	@mkdir -p benches
	@echo "Running benchmarks and saving to $@..."
ifeq ($(shell $(GO) env GOOS),linux)
	@if [ -d /sys/devices/system/cpu/cpu0/cpufreq ]; then \
		SAVED_GOV=$$(cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor); \
		echo "Setting CPU governor to performance mode..."; \
		for cpu in /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor; do \
			echo performance | sudo tee $$cpu > /dev/null; \
		done; \
		$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@; \
		echo "Restoring CPU governor to $$SAVED_GOV..."; \
		for cpu in /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor; do \
			echo $$SAVED_GOV | sudo tee $$cpu > /dev/null; \
		done; \
	else \
		$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@; \
	fi
else
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@
endif

.PHONY: total
total: $(BENCH_FILE) ## Run benchmarks and show total runtime
	@awk -f total.awk < $(BENCH_FILE)

.PHONY: total-nogc
total-nogc: ## Run benchmarks with GOGC=off and show total runtime
	@mkdir -p benches
	GOGC=off $(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $(BENCH_FILE)
	@awk -f total.awk < $(BENCH_FILE)

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml ## Generate GitLab CI reports

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

junit.xml: test.log
	$(GO) run github.com/jstemmer/go-junit-report/v2@latest < $< > $@

coverage.xml: coverage.txt
	$(GO) run github.com/boumenot/gocover-cobertura@latest < $< > $@

gl-code-quality-report.json: golangci-lint.json
	$(GO) run github.com/banyansecurity/golint-convert@latest < $< > $@

golangci-lint.json:
	-$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --out-format json > $@

govulncheck.sarif:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest -format=sarif ./... > $@
