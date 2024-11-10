package messagecode

const (
	// Internal message code
	CODE_SUCCESS                = "success"
	CODE_INTERNAL_ERR           = "internal_server_error"
	CODE_FOREIGN_KEY_CONSTRAINT = "foreign_key_constraint"
	CODE_UNIQUE_CONSTRAINT      = "unique_constraint"
	CODE_CHECK_CONSTRAINT       = "check_constraint"
	// Validate message code
	CODE_PARAM_INVALID = "invalid_parameter"
	// Status message code
	CODE_UPDATE_FAIL    = "update_fail"
	CODE_UPDATE_SUCCESS = "update_success"

	// Counter message code
	CODE_COUNTER_NAME_REQUIRED = "counter_name_required"
	CODE_COUNTER_NOT_FOUND     = "counter_not_found"
)
