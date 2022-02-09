package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type TestCase struct {
	Code     int
	Message  string
	GetQuery func() string
}

func TestAllPosts(t *testing.T) {

	testCases := []TestCase{
		{
			Code:    400,
			Message: "{\"Error\":\"kek\",\"Last\":\"\",\"Last2\":\"\"}",
			GetQuery: func() string {
				q := make(url.Values)
				q.Add("last_id", "142a6bde-70f4-439e-a60f-c0e90d0bc25a")
				q.Add("last_created_at", "2022-01-31T20:55:50MSK")
				return "/?" + q.Encode()
			},
		},
		{
			Code:    400,
			Message: "{\"Error\":\"kek\",\"Last\":\"\",\"Last2\":\"\"}",
			GetQuery: func() string {
				q := make(url.Values)
				q.Add("last_id", "142a6bde-70f4-439e-a60f-c0e90d0bc25a")
				return "/?" + q.Encode()
			},
		},
		{
			Code:    400,
			Message: "{\"Error\":\"kek\",\"Last\":\"\",\"Last2\":\"\"}",
			GetQuery: func() string {
				q := make(url.Values)
				q.Add("last_created_at", "2022-01-31T20:55:50MSK")
				return "/?" + q.Encode()
			},
		},
		{
			Code:    400,
			Message: "{\"Error\":\"cannot be blank\",\"Last\":\"\",\"Last2\":\"\"}",
			GetQuery: func() string {
				return "/"
			},
		},
		{
			Code:    400,
			Message: "{\"Error\":\"kek\",\"Last\":\"\",\"Last2\":\"\"}",
			GetQuery: func() string {
				q := make(url.Values)
				q.Add("last_id", "142a6bde-70f4-439e-a60f-c0e90d0bc25a")
				q.Add("last_created_at", "2022-01-31T20:55:50MosowK")
				return "/?" + q.Encode()
			},
		},
	}

	for i := range testCases {

		e := echo.New()
		query := testCases[i].GetQuery()
		req := httptest.NewRequest(http.MethodPost, query, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if err := AllPosts(c); err != nil {
			t.Error("lel", err)
		}

		res := rec.Result()

		if res.StatusCode != testCases[i].Code {
			t.Errorf("Expected code %d, got code %d", testCases[i].Code, res.StatusCode)
		}
		if rec.Body.String() != testCases[i].Message {
			t.Errorf("Expected body %s, got body %s", testCases[i].Message, rec.Body.String())
		}
	}

}
