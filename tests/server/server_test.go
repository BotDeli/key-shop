package server

import (
	"bytes"
	"encoding/json"
	"io"
	"key-shop/internal/database/sql/postgres"
	"net/http"
	"testing"
)

const address = "http://localhost:8080"

var (
	testedUser = postgres.User{
		Login:    "testLogin",
		Password: "testPassword",
	}
	testedItem = postgres.Item{
		Name:        "testedNameItem",
		Description: "testedDescriptionItem",
		Count:       "100",
		Cost:        "1000",
	}
)

// start main.go

var client = &http.Client{}

func TestServer(t *testing.T) {
	authorizationTests(t, testRegistration)
	sessionKey := authorizationTests(t, testLogin)

	testMyItemIsEmpty(t, sessionKey)

	testAddItem(t, sessionKey)

	testMyItemsIsNotShort(t, sessionKey)

	testCountPages(t)

	testAllItemsIsNotShort(t)

	testDeleteItem(t, sessionKey)

	testExitUser(t, sessionKey)
}

func authorizationTests(t *testing.T, functionTest func(t *testing.T) string) string {
	sessionKey := functionTest(t)
	checkLenSessionKey(t, sessionKey)
	testGetLogin(t, sessionKey)
	return sessionKey
}

func testRegistration(t *testing.T) string {
	return testAuthorization(t, address+"/registration")
}

func testAuthorization(t *testing.T, url string) string {
	testedJson := getTestedJson(t, testedUser)

	request := getNewRequest(t, http.MethodPost, url, bytes.NewBuffer(testedJson))

	setContentTypeJSON(request)

	response := getResponse(t, request)

	bodyBytes := getBodyBytes(t, response)

	defer closeBody(t, response)

	var errResponse struct {
		Err string `json:"ERR"`
	}

	decodeJSON(t, bodyBytes, &errResponse)

	if errResponse.Err != "<nil>" {
		t.Error(errResponse.Err)
	}

	cookies := response.Cookies()
	return cookies[0].Value
}

func getTestedJson(t *testing.T, testStruct interface{}) []byte {
	testedJson, err := json.Marshal(testStruct)
	if err != nil {
		t.Fatal("error marshalling: ", err)
	}
	return testedJson
}

func getNewRequest(t *testing.T, method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return request
}

func setContentTypeJSON(request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
}

func getResponse(t *testing.T, request *http.Request) *http.Response {
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	return response
}

func getBodyBytes(t *testing.T, response *http.Response) []byte {
	body := response.Body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}
	return bodyBytes
}

func closeBody(t *testing.T, response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		t.Error("error close body: ", err)
	}
}

func decodeJSON(t *testing.T, bodyBytes []byte, responseJSON interface{}) {
	err := json.Unmarshal(bodyBytes, &responseJSON)
	if err != nil {
		t.Fatal(err)
	}
}

func checkLenSessionKey(t *testing.T, sessionKey string) {
	if len(sessionKey) != 16 {
		t.Fatal("dont correct session key: ", sessionKey)
	}
}

func testGetLogin(t *testing.T, sessionKey string) {
	request := getNewRequest(t, "GET", address+"/get_login", nil)
	addCookieSessia(request, sessionKey)

	response := getResponse(t, request)
	checkStatusCode(t, response, http.StatusOK)
	defer closeBody(t, response)

	bodyBytes := getBodyBytes(t, response)
	bodyStr := string(bodyBytes)
	loginHandled := bodyStr[10 : len(bodyStr)-2]

	if loginHandled != testedUser.Login {
		t.Errorf("expected login: %s, got %s", testedUser.Login, loginHandled)
	}
}

func checkStatusCode(t *testing.T, response *http.Response, expectedCode int) {
	if response.StatusCode != expectedCode {
		t.Errorf("expected code %d, got %d", expectedCode, response.StatusCode)
	}
}

func addCookieSessia(request *http.Request, sessionKey string) {
	cookie := &http.Cookie{
		Name:     "sessia",
		Value:    sessionKey,
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
	}
	request.AddCookie(cookie)
}

func testLogin(t *testing.T) string {
	return testAuthorization(t, address+"/login")
}

func testMyItemIsEmpty(t *testing.T, sessionKey string) {
	testMyItems(t, sessionKey, checkMyItemsIsEmpty)
}

func testMyItems(t *testing.T, sessionKey string, checkFunction func(*testing.T, string)) {
	request := getNewRequest(t, "GET", address+"/my_items", nil)
	addCookieSessia(request, sessionKey)

	response := getResponse(t, request)
	defer closeBody(t, response)

	bodyBytes := getBodyBytes(t, response)
	bodyStr := string(bodyBytes)
	checkFunction(t, bodyStr)
}

func checkMyItemsIsEmpty(t *testing.T, bodyStr string) {
	if bodyStr != "{\"items\":null}" {
		t.Error("expected items is null, got: ", bodyStr)
	}
}

func testAddItem(t *testing.T, sessionKey string) {
	testedJSON := getTestedJson(t, testedItem)

	request := getNewRequest(t, "POST", address+"/add_item", bytes.NewBuffer(testedJSON))
	addCookieSessia(request, sessionKey)
	setContentTypeJSON(request)

	response := getResponse(t, request)
	defer closeBody(t, response)
	checkStatusCode(t, response, http.StatusCreated)
}

func testMyItemsIsNotShort(t *testing.T, sessionKey string) {
	testMyItems(t, sessionKey, checkItemsIsNotShort)
}

func checkItemsIsNotShort(t *testing.T, bodyStr string) {
	if len(bodyStr) < 20 || bodyStr == "{\"items\":null}" {
		t.Error("expected my items, got is short: ", bodyStr)
	}
}

func testCountPages(t *testing.T) {
	request := getNewRequest(t, "GET", address+"/count_pages", nil)

	response := getResponse(t, request)
	defer closeBody(t, response)
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected: %d, got: %d", http.StatusOK, response.StatusCode)
	}

	bodyBytes := getBodyBytes(t, response)
	bodyStr := string(bodyBytes)

	numberPage := bodyStr[9 : len(bodyStr)-1]
	if numberPage == "" || numberPage == "0" || numberPage == "null" {
		t.Error("got number page: ", numberPage)
	}
}

func testAllItemsIsNotShort(t *testing.T) {
	pageNumber := struct {
		Page int `json:"page"`
	}{1}

	testedJSON := getTestedJson(t, pageNumber)

	request := getNewRequest(t, "POST", address+"/items", bytes.NewBuffer(testedJSON))

	response := getResponse(t, request)
	checkStatusCode(t, response, http.StatusOK)

	bodyBytes := getBodyBytes(t, response)
	bodyStr := string(bodyBytes)

	checkItemsIsNotShort(t, bodyStr)
}

func testDeleteItem(t *testing.T, sessionKey string) {
	testedJSON := getTestedJson(t, testedItem)

	request := getNewRequest(t, "DELETE", address+"/delete_item", bytes.NewBuffer(testedJSON))
	addCookieSessia(request, sessionKey)

	response := getResponse(t, request)
	checkStatusCode(t, response, http.StatusAccepted)
}

func testExitUser(t *testing.T, sessionKey string) {
	request := getNewRequest(t, "POST", address+"/exit", nil)
	addCookieSessia(request, sessionKey)

	response := getResponse(t, request)
	checkStatusCode(t, response, http.StatusAccepted)
}
