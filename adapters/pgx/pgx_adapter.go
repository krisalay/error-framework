package pgxadapter

import (
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/krisalay/error-framework/core"
)

type Adapter struct {
	includeConstraint bool
	includeTable      bool
}

func New() *Adapter {
	return &Adapter{
		includeConstraint: true,
		includeTable:      true,
	}
}

func (a *Adapter) WithConstraintDetails(enabled bool) *Adapter {
	a.includeConstraint = enabled
	return a
}

func (a *Adapter) WithTableDetails(enabled bool) *Adapter {
	a.includeTable = enabled
	return a
}

func (a *Adapter) FromError(err error) *core.AppError {

	if err == nil {
		return nil
	}

	// Handle pgx no rows
	if errors.Is(err, pgx.ErrNoRows) {

		return core.New().
			WithMessage("Resource not found").
			WithCode(core.CodeNotFound).
			WithStatus(http.StatusNotFound).
			WithLevel(core.LevelInfo).
			WithSensitive(false).
			Build()
	}

	// Handle PostgreSQL errors
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {

		return a.handlePgError(pgErr, err)
	}

	// Connection or unknown DB error
	return core.New().
		WithMessage("Database error").
		WithCode(core.CodeDBError).
		WithStatus(http.StatusInternalServerError).
		WithLevel(core.LevelError).
		WithInternal(err).
		WithSensitive(true).
		Build()
}

func (a *Adapter) handlePgError(pgErr *pgconn.PgError, original error) *core.AppError {

	details := make(map[string]any)

	if a.includeConstraint && pgErr.ConstraintName != "" {
		details["constraint"] = pgErr.ConstraintName
	}

	if a.includeTable && pgErr.TableName != "" {
		details["table"] = pgErr.TableName
	}

	switch pgErr.Code {

	// unique_violation
	case "23505":

		return core.New().
			WithMessage("Resource already exists").
			WithCode(core.CodeDBDuplicateKey).
			WithStatus(http.StatusConflict).
			WithDetails(details).
			WithLevel(core.LevelWarn).
			WithInternal(original).
			WithSensitive(false).
			Build()

	// foreign_key_violation
	case "23503":

		return core.New().
			WithMessage("Invalid reference").
			WithCode(core.CodeDBForeignKey).
			WithStatus(http.StatusBadRequest).
			WithDetails(details).
			WithLevel(core.LevelWarn).
			WithInternal(original).
			WithSensitive(false).
			Build()

	// connection_exception
	case "08000", "08003", "08006":

		return core.New().
			WithMessage("Database connection error").
			WithCode(core.CodeDBConnectionError).
			WithStatus(http.StatusInternalServerError).
			WithLevel(core.LevelError).
			WithInternal(original).
			WithSensitive(true).
			Build()

	default:

		return core.New().
			WithMessage("Database error").
			WithCode(core.CodeDBError).
			WithStatus(http.StatusInternalServerError).
			WithDetails(details).
			WithLevel(core.LevelError).
			WithInternal(original).
			WithSensitive(true).
			Build()
	}
}
