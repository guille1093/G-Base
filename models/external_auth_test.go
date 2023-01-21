package models_test

import (
	"testing"

	"github.com/guille1093/G-Base/models"
)

func TestExternalAuthTableName(t *testing.T) {
	m := models.ExternalAuth{}
	if m.TableName() != "_externalAuths" {
		t.Fatalf("Unexpected table name, got %q", m.TableName())
	}
}
