package core

const (

	// Generic Errors
	CodeInternalError = "INTERNAL_ERROR"
	CodeUnknownError  = "UNKNOWN_ERROR"

	// Validation Errors
	CodeValidationError = "VALIDATION_ERROR"
	CodeInvalidInput    = "INVALID_INPUT"

	// Authentication Errors
	CodeUnauthorized = "UNAUTHORIZED"
	CodeForbidden    = "FORBIDDEN"

	// Resource Errors
	CodeNotFound      = "NOT_FOUND"
	CodeAlreadyExists = "ALREADY_EXISTS"

	// Database Errors
	CodeDBError           = "DB_ERROR"
	CodeDBDuplicateKey    = "DB_DUPLICATE_KEY"
	CodeDBForeignKey      = "DB_FOREIGN_KEY"
	CodeDBNoRows          = "DB_NO_ROWS"
	CodeDBConnectionError = "DB_CONNECTION_ERROR"

	// Network Errors
	CodeTimeout = "TIMEOUT"
)
