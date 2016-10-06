package mockbosh

import (
	"fmt"

	"github.com/pivotal-cf-experimental/cf-webmock/http"
)

type taskMock struct {
	*mockhttp.MockHttp
	taskID int
}

func Task(taskID int) *taskMock {
	return &taskMock{
		MockHttp: mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/tasks/%d", taskID)),
		taskID:   taskID,
	}
}

func (t *taskMock) RespondsWithTaskContainingState(provisioningTaskState string) *mockhttp.MockHttp {
	return t.RespondsWithJson(map[string]interface{}{
		"id":    t.taskID,
		"state": provisioningTaskState,
	})
}
