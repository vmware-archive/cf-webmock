package mockbosh

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/gomega"
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

func (mock *sshMock) SetSSHResponseCallback(callback func(string, string)) *sshMock {
	mock.SetResponseCallback(func(body []byte) {
		response := map[string]interface{}{}
		Expect(json.Unmarshal(body, &response)).To(Succeed())
		params := response["params"].(map[string]interface{})
		callback(params["user"].(string), params["public_key"].(string))
	})
	return mock
}
