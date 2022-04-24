package taskUsecase_test

import (
	"encoding/json"
	taskDto "gogolook/dto/task"
	"gogolook/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func marshal(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

var server *gin.Engine

func TestMain(m *testing.M) {
	server = router.SetupRouter()
	m.Run()
}

type testCase struct {
	t            *testing.T
	method       string
	path         string
	input        *strings.Reader
	expectCode   int
	expectResult string
}

func run(tc testCase) {
	req, _ := http.NewRequest(tc.method, tc.path, tc.input)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(tc.t, tc.expectCode, w.Code)
	assert.Contains(tc.t, w.Body.String(), tc.expectResult)
}

func Test_AddNewTaskWithoutParams(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodPost,
		path:         "/task",
		input:        strings.NewReader(""),
		expectCode:   http.StatusUnprocessableEntity,
		expectResult: "",
	}

	run(tc)
}

func Test_AddNewTask(t *testing.T) {

	tc := testCase{
		t:            t,
		method:       http.MethodPost,
		path:         "/task",
		input:        strings.NewReader(marshal(taskDto.AddTaskInput{Name: "buy dinner 1"})),
		expectCode:   http.StatusCreated,
		expectResult: marshal(taskDto.Task{Id: 1, Name: "buy dinner 1", Status: 0}),
	}

	run(tc)
}

func Test_AddSecondTask(t *testing.T) {

	tc := testCase{
		t:            t,
		method:       http.MethodPost,
		path:         "/task",
		input:        strings.NewReader(marshal(taskDto.AddTaskInput{Name: "buy dinner 2"})),
		expectCode:   http.StatusCreated,
		expectResult: marshal(taskDto.Task{Id: 2, Name: "buy dinner 2", Status: 0}),
	}

	run(tc)
}

func Test_GetTasks(t *testing.T) {
	tc := testCase{
		t:          t,
		method:     http.MethodGet,
		path:       "/tasks",
		input:      strings.NewReader(""),
		expectCode: http.StatusOK,
		expectResult: marshal([]taskDto.Task{
			{Id: 2, Name: "buy dinner 2", Status: 0},
			{Id: 1, Name: "buy dinner 1", Status: 0},
		}),
	}

	run(tc)
}

func Test_UpdateTaskWithoutParams(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodPut,
		path:         "/task/0",
		input:        strings.NewReader(marshal(taskDto.EditTaskInput{Status: 1})),
		expectCode:   http.StatusUnprocessableEntity,
		expectResult: "",
	}

	run(tc)
}

func Test_UpdateTaskWithNotExistId(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodPut,
		path:         "/task/-1",
		input:        strings.NewReader(marshal(taskDto.EditTaskInput{Id: -1, Status: 1})),
		expectCode:   http.StatusInternalServerError,
		expectResult: "",
	}

	run(tc)
}

func Test_UpdateTaskStatus(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodPut,
		path:         "/task/1",
		input:        strings.NewReader(marshal(taskDto.EditTaskInput{Id: 1, Status: 1})),
		expectCode:   http.StatusOK,
		expectResult: marshal(taskDto.Task{Id: 1, Name: "buy dinner 1", Status: 1}),
	}

	run(tc)
}

func Test_UpdateTaskName(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodPut,
		path:         "/task/1",
		input:        strings.NewReader(marshal(taskDto.EditTaskInput{Id: 1, Name: "buy dinner updated"})),
		expectCode:   http.StatusOK,
		expectResult: marshal(taskDto.Task{Id: 1, Name: "buy dinner updated", Status: 1}),
	}

	run(tc)
}

func Test_DeleteTaskWithoutParams(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodDelete,
		path:         "/task/xxb",
		input:        strings.NewReader(""),
		expectCode:   http.StatusUnprocessableEntity,
		expectResult: "",
	}

	run(tc)
}

func Test_DeleteTaskWithNotExistId(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodDelete,
		path:         "/task/-1",
		input:        strings.NewReader(""),
		expectCode:   http.StatusInternalServerError,
		expectResult: "",
	}

	run(tc)
}

func Test_DeleteTask(t *testing.T) {
	tc := testCase{
		t:            t,
		method:       http.MethodDelete,
		path:         "/task/2",
		input:        strings.NewReader(""),
		expectCode:   http.StatusOK,
		expectResult: "",
	}

	run(tc)
}
