package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"simple-crud-app/api"
	"simple-crud-app/internal/models"
)

func TestFilmUpdate(t *testing.T) {
	dur := "45m3s"
	req := api.ReqFilmUpdate{
		Film: models.Film{
			ID:       2,
			Name:     "UpdatedName",
			Duration: &dur,
		},
	}
	js, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, Gate+"/films/update", bytes.NewBuffer(js))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-AccessKey", "8c0a890940b01cc58fba4dee5c732c11")
	request.Header.Add("X-Token", "a5e5883c-f29b-40f4-939d-ec730461953a")
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
