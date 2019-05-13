package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	cc "github.com/rls/ping-api/conn/cache"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/router"
)

func setup() {
	os.Setenv("PING_API_CONSUL_URL", "127.0.0.1:8500")
	os.Setenv("PING_API_CONSUL_PATH", "ping-api")
	config.Init()
	cc.ConnectRedis()
}
func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func TestLocationsSave(t *testing.T) {
	eps := router.Route()
	srv := httptest.NewServer(eps)
	defer srv.Close()

	payloads := map[string][]byte{
		"invalid_user_id": []byte(`{
			"user_id": "",
			"locations": {
					"lat": 23.831146,
					"lon": 90.425085
				}
		}`),
		"invalid_locations": []byte(`{
			"user_id": "1",
			"locations": {
					"lat": 23.831146
				}
		}`),
		"invalid_referrer": []byte(`{
			"user_id": "1",
			"locations": {
					"lat": 23.831146,
					"lon": 90.425085
				}
		}`),
		"valid": []byte(`{
			"user_id": "1",
			"locations": {
					"lat": 23.831146,
					"lon": 90.425085
				}
		}`),
	}

	url := srv.URL + "/v1/locations"

	for _, testcase := range []struct {
		method, url, name string
		body              []byte
		want              int
	}{
		{"POST", url, "invalid_user_id_test", payloads["invalid_user_id"], http.StatusBadRequest},
		{"POST", url, "invalid_locations_test", payloads["invalid_locations"], http.StatusBadRequest},
		{"POST", url, "invalid_referrer_test", payloads["invalid_referrer"], http.StatusBadRequest},
		{"POST", url, "valid_payload_test", payloads["valid"], http.StatusOK},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, bytes.NewBuffer(testcase.body))

		if testcase.name != "invalid_referrer_test" {
			req.Header.Set("RLS-Referrer", "test.widespace.com")
		}

		resp, _ := http.DefaultClient.Do(req)

		if want, have := testcase.want, resp.StatusCode; want != have {
			t.Errorf("%s %s %s: want %d, have %d", testcase.name, testcase.method, testcase.url, want, have)
		}
	}
}
