package koauth

type UserInfo any

type KOAuth interface {
	GetUserInfo(token string) (UserInfo, error)
	GetField(ui UserInfo, field string) (string, error)
}
