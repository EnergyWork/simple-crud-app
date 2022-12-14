package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"simple-crud-app/api"
)

func TestAuthRegister(t *testing.T) {
	reqRegister := api.ReqAuthRegister{
		Login:    "User002",
		Password: "password",
	}
	js, err := json.Marshal(reqRegister)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest(http.MethodPost, Gate+"/auth/registration", bytes.NewBuffer(js))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/json")

	t.Log("request URL:", request.URL)
	t.Log("request Headers:", request.Header)
	t.Log("request Body:", string(js))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Log("response Status:", resp.Status)
	t.Log("response Headers:", resp.Header)
	t.Log("response Body:", string(body))
}
