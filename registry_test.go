package dr

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegistry(t *testing.T) {
	r, err := New()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		url     string
		method  string
		code    int
		body    string
		headers map[string][]string
	}{
		{"http://registry.loc/v1/_ping",
			"GET",
			200,
			"",
			map[string][]string{
				"Vary":                      []string{"Accept"},
				"Content-Type":              []string{"application/json"},
				"X-Docker-Registry-Version": []string{"0.6.0"},
			},
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(test.method, test.url, nil)

		rw := httptest.NewRecorder()
		rw.Body = &bytes.Buffer{}

		r.ServeHTTP(rw, req)

		if g, w := rw.Code, test.code; g != w {
			t.Errorf("%s: code = %d, want %d", test.url, g, w)
		}

		if len(test.headers) != len(rw.Header()) {
			t.Errorf("%s: headers = %d, want %d", test.url, len(rw.Header()), len(test.headers))
		}

		if test.body != rw.Body.String() {
			t.Errorf("%s: body = %s, want %s", test.url, rw.Body, test.body)
		}

		for k, vs := range rw.Header() {
			switch tvs, ok := test.headers[k]; {
			case !ok:
				t.Errorf("%s: missing header: %s", k)
			case len(tvs) != len(vs):
				t.Errorf("%s: values = %d, want %d", test.url, len(vs), len(tvs))
			default:
			NextValue:
				for _, tv := range tvs {
					for _, v := range vs {
						if tv == v {
							continue NextValue
						}
					}
					t.Errorf("%s: missing value: %s", test.url, tv)
				}
			}
		}
	}
}
