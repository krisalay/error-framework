package pgxadapter

import (
	"testing"

	"github.com/jackc/pgconn"
)

func TestDuplicateKey(t *testing.T) {

	pgErr := &pgconn.PgError{
		Code: "23505",
	}

	adapter := New()

	appErr := adapter.FromError(pgErr)

	if appErr.Code != "DB_DUPLICATE_KEY" {
		t.Fatal("Duplicate key not handled")
	}
}
