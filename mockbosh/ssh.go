package mockbosh

import (
	"fmt"

	"github.com/pivotal-cf-experimental/cf-webmock/mockhttp"
)

type sshMock struct {
	*mockhttp.MockHttp
}

func StartSSHSession(deploymentName string) *sshMock {
	return &sshMock{
		MockHttp: mockhttp.NewMockedHttpRequest("POST", fmt.Sprintf("/deployments/%s/ssh", deploymentName)),
	}
}

func CleanupSSHSession(deploymentName string) *sshMock {
	return &sshMock{
		MockHttp: mockhttp.NewMockedHttpRequest("POST", fmt.Sprintf("/deployments/%s/ssh", deploymentName)),
	}
}

func (mock *sshMock) RedirectsToTask(taskID int) *mockhttp.MockHttp {
	return mock.RedirectsTo(taskURL(taskID))
}
