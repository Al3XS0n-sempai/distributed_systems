package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Al3XS0n-sempai/distributed_systems/internal/repository"
)

// testEnv contains variable required for request and expected http status code
// as key and value expected simple types, such as int or string
// in 'get' tests, value is expected reponse value, if http.StatusOk expected
type testEnv struct {
	key            interface{}
	value          interface{}
	expectedStatus int
}

var (
	testPutOKEnv = testEnv{
		key:            123,
		value:          3000,
		expectedStatus: http.StatusOK,
	}
	testPutNoKeyEnv = testEnv{
		key:            "lol",
		value:          123,
		expectedStatus: http.StatusBadRequest,
	}
	testPutNoValueEnv = testEnv{
		key:            321,
		value:          "kek",
		expectedStatus: http.StatusBadRequest,
	}
)

var (
	testGetOkEnv = testEnv{
		key:            3333,
		value:          3000,
		expectedStatus: http.StatusOK,
	}
	testGetNoKeyEnv = testEnv{
		key:            nil,
		value:          3000,
		expectedStatus: http.StatusBadRequest,
	}
	testGetBadKeyEnv = testEnv{
		key:            1,
		value:          3000,
		expectedStatus: http.StatusNotFound,
	}
)

func TestPut(t *testing.T) {
	var repo = repository.NewInMemoryCache()
	var service = NewSimpleService(repo)

	t.Run("Test /put (OK)", func(t *testing.T) {
		_env := testPutOKEnv
		path := fmt.Sprintf("/put?key=%v&value=%v", _env.key, _env.value)
		request, _ := http.NewRequest(http.MethodPost, path, nil)
		response := httptest.NewRecorder()

		service.setValueHandler(response, request)

		if response.Result().StatusCode != _env.expectedStatus {
			t.Errorf(
				"Expected HTTP status code %d, recieved %d",
				_env.expectedStatus,
				response.Result().StatusCode,
			)
		}
	})

	t.Run("Test /put (No key)", func(t *testing.T) {
		_env := testPutNoKeyEnv
		path := fmt.Sprintf("/put?value=%v", _env.value)
		request, _ := http.NewRequest(http.MethodPost, path, nil)
		response := httptest.NewRecorder()

		service.setValueHandler(response, request)

		if response.Result().StatusCode != _env.expectedStatus {
			t.Errorf(
				"Expected HTTP status code %d, recieved %d",
				_env.expectedStatus,
				response.Result().StatusCode,
			)
		}
	})

	t.Run("Test /put (No value)", func(t *testing.T) {
		_env := testPutNoValueEnv
		path := fmt.Sprintf("/put?key=%v", _env.key)
		request, _ := http.NewRequest(http.MethodPost, path, nil)
		response := httptest.NewRecorder()

		service.setValueHandler(response, request)

		if response.Result().StatusCode != _env.expectedStatus {
			t.Errorf(
				"Expected HTTP status code %d, recieved %d",
				_env.expectedStatus,
				response.Result().StatusCode,
			)
		}
	})
}

func TestGet(t *testing.T) {
	var repo = repository.NewInMemoryCache()
	var service = NewSimpleService(repo)

	t.Run("Test /get (OK)", func(t *testing.T) {
		_env := testGetOkEnv

		setPath := fmt.Sprintf("/put?key=%v&value=%v", _env.key, _env.value)
		getPath := fmt.Sprintf("/get?key=%v", _env.key)
		setRequest, _ := http.NewRequest(http.MethodGet, setPath, nil)
		getRequest, _ := http.NewRequest(http.MethodGet, getPath, nil)
		setResponse := httptest.NewRecorder()
		getResponse := httptest.NewRecorder()

		service.setValueHandler(setResponse, setRequest)
		service.getValueHandler(getResponse, getRequest)

		if getResponse.Result().StatusCode != _env.expectedStatus {
			t.Errorf(
				"Expected HTTP status code %d, recieved %d",
				_env.expectedStatus,
				getResponse.Result().StatusCode,
			)
		}
	})

	t.Run("Test /get (No key)", func(t *testing.T) {
		_env := testGetNoKeyEnv

		path := "/get"
		request, _ := http.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()

		service.getValueHandler(response, request)

		if response.Result().StatusCode != _env.expectedStatus {
			t.Errorf(
				"Expected HTTP status code %d, recieved %d",
				_env.expectedStatus,
				response.Result().StatusCode,
			)
		}
	})

	t.Run("Test /get (Bad key)", func(t *testing.T) {
		_env := testGetBadKeyEnv

		path := fmt.Sprintf("/get?key=%v", _env.key)
		request, _ := http.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()

		service.getValueHandler(response, request)

		if response.Result().StatusCode != _env.expectedStatus {
			t.Errorf(
				"Expected HTTP status code %d, recieved %d",
				_env.expectedStatus,
				response.Result().StatusCode,
			)
		}
	})
}
