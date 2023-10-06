package kakao

import (
	"encoding/json"
	"github.com/snowmerak/koauth/lib/koauth"
	"net/http"
	"reflect"
	"strconv"
)

const (
	accessTokenInfoUri = "https://kapi.kakao.com/v1/user/access_token_info"
	userInfoUri        = "https://kapi.kakao.com/v2/user/me"
)

var _ = (koauth.KOAuth)(&Kakao{})

type Kakao struct {
	client http.Client
}

func New() *Kakao {
	return &Kakao{
		client: http.Client{},
	}
}

func (k *Kakao) GetUserInfo(token string) (ui koauth.UserInfo, err error) {
	kui := UserInfo{}

	req, err := http.NewRequest("GET", accessTokenInfoUri, nil)
	if err != nil {
		return ui, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := k.client.Do(req)
	if err != nil {
		return ui, err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&kui)
	if err != nil {
		return ui, err
	}

	req, err = http.NewRequest("GET", userInfoUri, nil)
	if err != nil {
		return ui, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err = k.client.Do(req)
	if err != nil {
		return ui, err
	}

	decoder = json.NewDecoder(resp.Body)
	err = decoder.Decode(&kui)
	if err != nil {
		return ui, err
	}

	return kui, nil
}

func (k *Kakao) GetField(ui koauth.UserInfo, field string) (string, error) {
	kui, ok := ui.(UserInfo)
	if !ok {
		return "", koauth.InvalidUserInfoType{
			Expected: "kakao.UserInfo",
			Actual:   reflect.TypeOf(ui).String(),
		}
	}

	switch field {
	case FieldId:
		return strconv.FormatInt(kui.Id, 10), nil
	case FieldExpiresIn:
		return strconv.Itoa(kui.ExpiresIn), nil
	case FieldAppId:
		return strconv.Itoa(kui.AppId), nil
	case FieldNickname:
		return kui.KakaoAccount.Profile.Nickname, nil
	case FieldThumbnailImage:
		return kui.KakaoAccount.Profile.ThumbnailImage, nil
	case FieldProfileImage:
		return kui.KakaoAccount.Profile.ProfileImage, nil
	case FieldIsDefaultImage:
		return strconv.FormatBool(kui.KakaoAccount.Profile.IsDefaultImage), nil
	case FieldEmail:
		return kui.KakaoAccount.Email, nil
	case FieldName:
		return kui.KakaoAccount.Name, nil
	case FieldAgeRange:
		return kui.KakaoAccount.AgeRange, nil
	case FieldBirthday:
		return kui.KakaoAccount.Birthday, nil
	case FieldGender:
		return kui.KakaoAccount.Gender, nil
	}

	return "", koauth.InvalidField{
		Field: field,
	}
}
