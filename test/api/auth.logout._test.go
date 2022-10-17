package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simple-crud-app/api/handlers"
	"testing"
	"time"
)

func TestAuthLogOut(t *testing.T) {
	reqRegister := handlers.ReqAuthLogout{}
	js, err := json.Marshal(reqRegister)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, Gate+"/auth/logout", bytes.NewBuffer(js))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-AccessKey", "8c0a890940b01cc58fba4dee5c732c11")
	request.Header.Add("X-Token", "9f755a15-4cda-476e-8f08-f8b01412e9f2")
	request.Header.Add("X-User", "User002")

	t.Log("request URL:", request.URL)
	t.Log("request Headers:", request.Header)
	t.Log("request Body:", string(js))

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
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
