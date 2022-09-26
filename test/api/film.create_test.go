package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"simple-crud-app/api"
)

func TestFilmCreate(t *testing.T) {
	var score uint64 = 80
	var comment string = "OK"
	reqRegister := api.ReqCreateFilm{
		Name:    "FILM NAME",
		Score:   &score,
		Comment: &comment,
	}
	js, err := json.Marshal(reqRegister)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, Gate+"/films/create", bytes.NewBuffer(js))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-AccessKey", "8c0a890940b01cc58fba4dee5c732c11")
	request.Header.Add("X-Token", "c628811a-3206-4664-b9b3-5ef9c46e1b20")
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
