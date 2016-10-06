package mockbosh

import "github.com/pivotal-cf-experimental/cf-webmock/mockhttp"

type infoMock struct {
	*mockhttp.MockHttp
}

func Info() *infoMock {
	return &infoMock{
		MockHttp: mockhttp.NewMockedHttpRequest("GET", "/info"),
	}
}

func (m *infoMock) RespondsWithSufficientAPIVersion() *mockhttp.MockHttp {
	return m.RespondsWith(`{"version":"1.3262.0.0 (00000000)"}`)
}
