package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/rls/ping-api/router"

	cc "github.com/rls/ping-api/conn/cache"
	"github.com/rls/ping-api/pkg/config"
)

var eps http.Handler

func setup() {
	os.Setenv("PING_API_CONSUL_URL", "127.0.0.1:8500")
	os.Setenv("PING_API_CONSUL_PATH", "ping-api")
	config.Init()
	cc.ConnectRedis()
	eps = router.Route()
}
func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func TestLocationsSave(t *testing.T) {
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
			req.Header.Set("RLS-Referrer", "test.abcd.com")
		}

		resp, _ := http.DefaultClient.Do(req)

		if want, have := testcase.want, resp.StatusCode; want != have {
			t.Errorf("%s %s %s: want %d, have %d", testcase.name, testcase.method, testcase.url, want, have)
		}
	}
}

func TestLocationsGet(t *testing.T) {
	srv := httptest.NewServer(eps)
	defer srv.Close()

	qs := map[string]string{
		"missing_user_ids":       `user_ids=`,
		"invalid_referrer":       `user_ids=1,2,3`,
		"valid_ids":              `user_ids=1,`,
		"valid_with_invalid_ids": `user_ids=1,99`,
		"all_invalid_ids":        `user_ids=111,99`,
	}

	res := map[string]string{
		"valid_ids":              `{"data":{"1":{"lat":23.831145440926278,"lon":90.42508453130722}}}`,
		"valid_with_invalid_ids": `{"data":{"1":{"lat":23.831145440926278,"lon":90.42508453130722},"99":null}}`,
		"all_invalid_ids":        `{"data":{"111":null,"99":null}}`,
	}

	url := srv.URL + "/v1/locations?"

	for _, testcase := range []struct {
		method, url, name, response string
		want                        int
	}{
		{"GET", url + qs["missing_user_ids"], "missing_user_ids_test", "", http.StatusBadRequest},
		{"GET", url + qs["invalid_referrer"], "invalid_referrer_test", "", http.StatusBadRequest},
		{"GET", url + qs["valid_with_invalid_ids"], "valid_with_invalid_ids_test", res["valid_with_invalid_ids"], http.StatusOK},
		{"GET", url + qs["all_invalid_ids"], "all_invalid_ids_test", res["all_invalid_ids"], http.StatusOK},
		{"GET", url + qs["valid_ids"], "valid_ids_test", res["valid_ids"], http.StatusOK},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, nil)

		if testcase.name != "invalid_referrer_test" {
			req.Header.Set("RLS-Referrer", "test.abcd.com")
		}

		resp, _ := http.DefaultClient.Do(req)

		if want, have := testcase.want, resp.StatusCode; want != have {
			t.Errorf("%s %s %s: want %d, have %d", testcase.name, testcase.method, testcase.url, want, have)
		}

		// skipping response structure test for all except 200
		if testcase.want != resp.StatusCode || testcase.want != http.StatusOK {
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)

		if want, have := strings.TrimSpace(testcase.response), strings.TrimSpace(string(body)); want != have {
			t.Errorf("%s %s: want %q, have %q", testcase.method, testcase.url, want, have)
		}
	}
}
