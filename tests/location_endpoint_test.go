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
	"github.com/rls/ping-api/conn/queue"
	"github.com/rls/ping-api/pkg/config"
)

var eps http.Handler
var testDomain string

func setup() {
	os.Setenv("PING_API_CONSUL_URL", "127.0.0.1:8500")
	os.Setenv("PING_API_CONSUL_PATH", "ping-api")
	config.Init()
	cc.ConnectRedis()
	queue.ConnectRabbitMQ()
	eps = router.Route()

	testDomain = "test.abcde.com"
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
			"locations":[{
				"user_id": "",
				"lat": 23.831146,
				"lon": 90.425085,
				"client_timestamp_utc": 10990809
			}]
		}`),
		"invalid_locations": []byte(`{
			"locations":[{
				"user_id": "1",
				"lat": 23.831146,
				"client_timestamp_utc": 10990809
			}]
		}`),
		"invalid_referrer": []byte(`{
			"locations": [{
				"user_id": "1",
				"lat": 23.831146,
				"lon": 90.425085,
				"client_timestamp_utc": 10990809
			}]
		}`),
		"missing_client_timestamp_utc": []byte(`{
			"locations": [{
				"user_id": "1",
				"lat": 23.831146,
				"lon": 90.425085,
			}]
		}`),
		"valid": []byte(`{
			"locations":[{
				"user_id": "1",
				"lat": 23.831146,
				"lon": 90.425085,
				"client_timestamp_utc": 10990709
			}, {
				"user_id": "2",
				"lat": 23.834314,
				"lon": 90.422827,
				"client_timestamp_utc": 10990712
			}, {
				"user_id": "3",
				"lat": 23.835899,
				"lon": 90.423106,
				"client_timestamp_utc": 10990715
			}]
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
		{"POST", url, "missing_client_timestamp_utc", payloads["missing_client_timestamp_utc"], http.StatusBadRequest},
		{"POST", url, "valid_payload_test", payloads["valid"], http.StatusOK},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, bytes.NewBuffer(testcase.body))

		if testcase.name != "invalid_referrer_test" {
			req.Header.Set("RLS-Referrer", testDomain)
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
			req.Header.Set("RLS-Referrer", testDomain)
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

func TestLocationsSearch(t *testing.T) {
	srv := httptest.NewServer(eps)
	defer srv.Close()

	qs := map[string]string{
		"missing_lat":           `radius=10&lat=&lon=90.422827&unit=km`,
		"missing_lon":           `radius=10&lat=23.834314&lon=&unit=km`,
		"missing_radius":        `radius=&lat=23.834314&lon=90.422827&unit=km`,
		"invalid_unit":          `radius=10&lat=23.834314&lon=90.422827&unit=`,
		"missing_referrer":      `radius=10&lat=23.834314&lon=90.422827&unit=km`,
		"valid_data":            `radius=1&lat=23.834314&lon=90.422827&unit=km`,
		"valid_data_with_limit": `radius=1&lat=23.834314&lon=90.422827&unit=km&limit=1`,
		"empty_data":            `radius=10&lat=23.934314&lon=90.822827&unit=ft`,
	}

	res := map[string]string{
		"valid_data":            `{"data":[{"user_id":"2","lat":23.83431384237545,"lon":90.42282611131668,"distance":0.0001,"geo_hash":4011392605551972},{"user_id":"3","lat":23.835898043100038,"lon":90.42310506105423,"distance":0.1784,"geo_hash":4011392607786693},{"user_id":"1","lat":23.831145440926278,"lon":90.42508453130722,"distance":0.4207,"geo_hash":4011392603947458}]}`,
		"valid_data_with_limit": `{"data":[{"user_id":"2","lat":23.83431384237545,"lon":90.42282611131668,"distance":0.0001,"geo_hash":4011392605551972}]}`,
		"empty_data":            `{"data":[]}`,
	}

	url := srv.URL + "/v1/locations/users?"

	for _, testcase := range []struct {
		method, url, name, response string
		want                        int
	}{
		{"GET", url + qs["missing_lat"], "missing_lat_test", "", http.StatusBadRequest},
		{"GET", url + qs["missing_lon"], "missing_lon_test", "", http.StatusBadRequest},
		{"GET", url + qs["missing_radius"], "missing_radius_test", "", http.StatusBadRequest},
		{"GET", url + qs["invalid_unit"], "invalid_unit_test", "", http.StatusBadRequest},
		{"GET", url + qs["missing_referrer"], "missing_referrer_test", "", http.StatusBadRequest},
		{"GET", url + qs["valid_data"], "valid_data_test", res["valid_data"], http.StatusOK},
		{"GET", url + qs["valid_data_with_limit"], "valid_data_with_limit_test", res["valid_data_with_limit"], http.StatusOK},
		{"GET", url + qs["empty_data"], "empty_data_test", res["empty_data"], http.StatusNotFound},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, nil)

		if testcase.name != "missing_referrer_test" {
			req.Header.Set("RLS-Referrer", testDomain)
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
			t.Errorf("%s %s: want %q, have %q\n", testcase.method, testcase.url, want, have)
		}
	}
}
