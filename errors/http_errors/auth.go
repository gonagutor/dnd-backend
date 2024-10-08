package http_errors

const (
	BAD_EMAIL             = "BAD_EMAIL"
	BAD_PASSWORD          = "BAD_PASSWORD"
	COULD_NOT_CREATE_USER = "COULD_NOT_CREATE_USER"

	EMAIL_NOT_VERIFIED     = "EMAIL_NOT_VERIFIED"
	EMAIL_ALREADY_VERIFIED = "EMAIL_ALREADY_VERIFIED"
	COULD_NOT_VERIFY_EMAIL = "COULD_NOT_VERIFY_EMAIL"
	BAD_EMAIL_TOKEN        = "BAD_EMAIL_TOKEN"

	COULD_NOT_REVOKE_REFRESH_TOKEN = "COULD_NOT_REVOKE_REFRESH_TOKEN"
	COULD_NOT_GENERATE_TOKEN       = "COULD_NOT_GENERATE_TOKEN"
	BAD_REFRESH_TOKEN              = "BAD_REFRESH_TOKEN"
	BAD_ACCESS_TOKEN               = "BAD_ACCESS_TOKEN"
	REVOKED_REFRESH_TOKEN          = "REVOKED_REFRESH_TOKEN"

	BAD_RECOVER_TOKEN         = "BAD_RECOVER_TOKEN"
	COULD_NOT_UPDATE_PASSWORD = "COULD_NOT_UPDATE_PASSWORD"

	NOT_AN_ADMIN = "NOT_AN_ADMIN"
)
