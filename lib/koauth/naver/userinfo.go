package naver

type UserInfo struct {
	ResultCode string `json:"resultcode"`
	Message    string `json:"message"`
	Response   struct {
		Id           string `json:"id"`
		Nickname     string `json:"nickname"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Gender       string `json:"gender"`
		Age          string `json:"age"`
		Birthday     string `json:"birthday"`
		ProfileImage string `json:"profile_image"`
		BirthYear    string `json:"birthyear"`
		Mobile       string `json:"mobile"`
	} `json:"response"`
}

const (
	FieldId           = "id"
	FieldNickname     = "nickname"
	FieldName         = "name"
	FieldEmail        = "email"
	FieldGender       = "gender"
	FieldAge          = "age"
	FieldBirthday     = "birthday"
	FieldProfileImage = "profile_image"
	FieldBirthYear    = "birthyear"
	FieldMobile       = "mobile"
)
