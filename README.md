# go-playground

Sample code for learning Go. Requires Go 1.27 or later.

| Directory | Topic | Article (Japanese) |
|---|---|---|
| [testing](./testing) | Go testing: table-driven tests, subtests, helpers, httptest, Example tests, coverage, benchmarks, fuzzing, testing/synctest | [Part 1](https://jedipunkz.rocks/post/go-testing-1/) / [Part 2](https://jedipunkz.rocks/post/go-testing-2/) |

## Usage

```sh
cd testing
go test ./...
go test -bench=. -benchmem -run='^$' ./concat/
go test -fuzz=FuzzReverse -fuzztime=30s ./strutil/
```
