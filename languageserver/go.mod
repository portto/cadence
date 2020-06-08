module github.com/portto/cadence/languageserver

go 1.13

require (
	github.com/mattn/go-isatty v0.0.12
	github.com/portto/blocto-flow-go-sdk v0.4.0
	github.com/portto/cadence v0.4.0
	github.com/sourcegraph/jsonrpc2 v0.0.0-20191222043438-96c4efab7ee2
	google.golang.org/grpc v1.29.1
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)

replace github.com/portto/cadence => ../
