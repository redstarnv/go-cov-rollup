# go-cov-rollup

This tool can be used to roll up and summarise coverage reports produced by `go test` runs executed with Go modules enabled.

The problem that occurs in this mode (e.g. when using [go-acc](get github.com/ory/go-acc) tool) is that the same code block may be
declared in the report multiple times with different hit metrics (see `examples` directory for some samples).
When this happens, some code coverage reporting services (e.g. Code Climate) calculate total coverage incorrectly.

This tool processes coverage report and combines multiple definitions of the same code block into a single line. It also handles boolean
coverage results of "set" covermode – when "set" mode is declared in the report files, summarised report will also contain boolean hit flags
only.

## Usage

Installation:

```
go get github.com/redstarnv/go-cov-rollup
```

To run it, pipe the report into stdin - rolled up report goes to stdout:

```
cat coverage.txt | ./go-cov-rollup > coverage-rolledup.txt
```
