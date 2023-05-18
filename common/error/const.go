package error

var (
	DomainErrorTranslatorDirectories = map[string]*ClientError{
		"CREATE_USER.USER_NOT_AUTHORIZED": NewForbiddenError("you cannot access this resource"),
		"CREATE_USER.USER_EXISTS":         NewInvariantError("user already exists"),
		"CREATE_USER.EMAIL_NOT_VERIFIED":  NewInvariantError("EMAIL_NOT_VERIFIED"),
		"GET_USER.USER_NOT_AUTHORIZED":    NewForbiddenError("you cannot access this resource"),
		"UPDATE_USER.USER_NOT_AUTHORIZED": NewForbiddenError("you cannot access this resource"),

		"CREATE_DRIVER.DRIVER_NOT_AUTHORIZED": NewForbiddenError("you cannot access this resource"),
		"CREATE_DRIVER.DRIVER_EXISTS":         NewInvariantError("user already exists"),
		"CREATE_DRIVER.EMAIL_NOT_VERIFIED":    NewInvariantError("EMAIL_NOT_VERIFIED"),
		"GET_DRIVER.DRIVER_NOT_AUTHORIZED":    NewForbiddenError("you cannot access this resource"),
		"UPDATE_DRIVER.DRIVER_NOT_AUTHORIZED": NewForbiddenError("you cannot access this resource"),
	}
)
