package forms_test

import (
	"strings"
	"testing"

	"github.com/guille1093/G-Base/forms"
)

func TestRealtimeSubscribeValidate(t *testing.T) {
	scenarios := []struct {
		clientId    string
		expectError bool
	}{
		{"", true},
		{strings.Repeat("a", 256), true},
		{"test", false},
	}

	for i, s := range scenarios {
		form := forms.NewRealtimeSubscribe()
		form.ClientId = s.clientId

		err := form.Validate()

		hasErr := err != nil
		if hasErr != s.expectError {
			t.Errorf("(%d) Expected hasErr to be %v, got %v (%v)", i, s.expectError, hasErr, err)
		}
	}
}
