package main

import (
	"encoding/json"
	"github.com/durolavoro/snippet-app/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	Home(rr, r, nil)
	res := rr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var hInfo model.HealthInfo
	err = json.Unmarshal(body, &hInfo)
	if err != nil {
		t.Fatal(err)
	}

	if hInfo.Status != "OK" {
		t.Fatalf("healthInfo.Status want OK, got %s", hInfo.Status)
	}
	if hInfo.Version != "v1" {
		t.Fatalf("healthInfo.Version want v1, got %s", hInfo.Version)
	}
}

func TestShowSnippet(t *testing.T) {
	tt := []struct {
		name string
		path string
		want int
	}{
		{
			name: "without id",
			path: "/snippet",
			want: http.StatusOK,
		},
		{
			name: "with correct id",
			path: "/snippet?id=1",
			want: http.StatusOK,
		},
		{
			name: "with incorrect id",
			path: "/snippet?id=test",
			want: http.StatusBadRequest,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, test.path, nil)

			ShowSnippet(rr, r, nil)
			resp := rr.Result()

			if resp.StatusCode != test.want {
				t.Fatalf("want %d, got %d", test.want, resp.StatusCode)
			}
		})
	}
}
