package koauth

type InvalidUserInfoType struct {
	Expected string
	Actual   string
}

func (e InvalidUserInfoType) Error() string {
	return "invalid user info type: expected " + e.Expected + ", actual " + e.Actual
}

type InvalidField struct {
	Field string
}

func (e InvalidField) Error() string {
	return "invalid field: " + e.Field
}
