package mockbosh

import (
	"net/http/httptest"

	"github.com/pivotal-cf-experimental/cf-webmock/mockhttp"
)

func New() *mockhttp.Server {
	return mockhttp.StartServer("mock-bosh", httptest.NewServer)
}

func NewTLS() *mockhttp.Server {
	return mockhttp.StartServer("mock-bosh", httptest.NewTLSServer)
}
