package tokens_test

import (
	"testing"

	"github.com/guille1093/G-Base/tests"
	"github.com/guille1093/G-Base/tokens"
)

func TestNewAdminAuthToken(t *testing.T) {
	app, _ := tests.NewTestApp()
	defer app.Cleanup()

	admin, err := app.Dao().FindAdminByEmail("test@example.com")
	if err != nil {
		t.Fatal(err)
	}

	token, err := tokens.NewAdminAuthToken(app, admin)
	if err != nil {
		t.Fatal(err)
	}

	tokenAdmin, _ := app.Dao().FindAdminByToken(
		token,
		app.Settings().AdminAuthToken.Secret,
	)
	if tokenAdmin == nil || tokenAdmin.Id != admin.Id {
		t.Fatalf("Expected admin %v, got %v", admin, tokenAdmin)
	}
}

func TestNewAdminResetPasswordToken(t *testing.T) {
	app, _ := tests.NewTestApp()
	defer app.Cleanup()

	admin, err := app.Dao().FindAdminByEmail("test@example.com")
	if err != nil {
		t.Fatal(err)
	}

	token, err := tokens.NewAdminResetPasswordToken(app, admin)
	if err != nil {
		t.Fatal(err)
	}

	tokenAdmin, _ := app.Dao().FindAdminByToken(
		token,
		app.Settings().AdminPasswordResetToken.Secret,
	)
	if tokenAdmin == nil || tokenAdmin.Id != admin.Id {
		t.Fatalf("Expected admin %v, got %v", admin, tokenAdmin)
	}
}
