package kakao

type UserInfo struct {
	Id        int64 `json:"id"`
	ExpiresIn int   `json:"expires_in"`
	AppId     int   `json:"app_id"`

	KakaoAccount struct {
		ProfileNeedsAgreement bool `json:"profile_needs_agreement"`
		Profile               struct {
			Nickname       string `json:"nickname"`
			ThumbnailImage string `json:"thumbnail_image_url"`
			ProfileImage   string `json:"profile_image_url"`
			IsDefaultImage bool   `json:"is_default_image"`
		} `json:"profile"`
		EmailNeedsAgreement    bool   `json:"email_needs_agreement"`
		IsEmailValid           bool   `json:"is_email_valid"`
		IsEmailVerified        bool   `json:"is_email_verified"`
		Email                  string `json:"email"`
		NameNeedsAgreement     bool   `json:"name_needs_agreement"`
		Name                   string `json:"name"`
		AgeRangeNeedsAgreement bool   `json:"age_range_needs_agreement"`
		AgeRange               string `json:"age_range"`
		BirthdayNeedsAgreement bool   `json:"birthday_needs_agreement"`
		Birthday               string `json:"birthday"`
		GenderNeedsAgreement   bool   `json:"gender_needs_agreement"`
		Gender                 string `json:"gender"`
	} `json:"kakao_account"`

	Properties struct {
		Nickname       string `json:"nickname"`
		ThumbnailImage string `json:"thumbnail_image"`
		ProfileImage   string `json:"profile_image"`
		CustomField1   string `json:"custom_field1"`
		CustomField2   string `json:"custom_field2"`
	}
}

const (
	FieldId        = "id"
	FieldExpiresIn = "expires_in"
	FieldAppId     = "app_id"

	FieldNickname       = "nickname"
	FieldThumbnailImage = "thumbnail_image"
	FieldProfileImage   = "profile_image"
	FieldIsDefaultImage = "is_default_image"
	FieldEmail          = "email"
	FieldName           = "name"
	FieldAgeRange       = "age_range"
	FieldBirthday       = "birthday"
	FieldGender         = "gender"
)
