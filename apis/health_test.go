package apis_test

import (
	"net/http"
	"testing"

	"github.com/guille1093/G-Base/tests"
)

func TestHealthAPI(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "health status returns 200",
			Method:         http.MethodGet,
			Url:            "/api/health",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"code":200`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
