package constants

const (
	INIT_FAILURE                 string = "Fail to pass config file"
	UNAUTHORIZED_ACCESS          string = "Unauthorized Access"
	EXISTENT_ACCOUNT             string = "User account already exists"
	NONEXISTENT_ACCOUNT          string = "User account does not exist"
	EMAIL_NOT_VALID              string = "Email is not valid"
	WRONG_VALIDATION             string = "Wrong Validation Code"
	WRONG_LOGIN_CREDENTIALS      string = "Email or Password is incorrect"
	FAILURE_EMAILSEND            string = "Fail to send email"
	INVALID_TOKEN_FORMAT         string = "Invalid Token Format"
	No_AUTHORIZATION_HEADER      string = "No Authorization Header provided"
	INVALID_AUTHORIZATION_METHOD string = "Invalid Authorization Method provided"
	INVALID_TOKEN                string = "Invalid Token"
	EXPIRED_TOKEN                string = "Expired Token"
	CLAIMS_GENERATE_FAILURE      string = "Failed to generate claims"
	CLAIMS_PARSE_FAILURE         string = "Failed to parse claims"
	FAILURE_PARSE_JSON           string = "Failed to parse request body"
	VALIDATION_JSON_ERROR        string = "Json validation failed"
	NOT_VERIFIED                 string = "Email is not verified"
	USER_NOT_IN_DATABASE         string = "User is not in database"
	CANNOT_SAVE_USER_IN_DATABASE string = "Cannot save user in database"
	DATABASE_ERROR               string = "Database error"
	LOGIN_EXPIRED                string = "Please Login again"
	CANNOT_PARSE_REQUEST         string = "Request is in wrong format"
)
