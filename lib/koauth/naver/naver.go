package naver

import (
	"encoding/json"
	"github.com/snowmerak/koauth/lib/koauth"
	"net/http"
	"reflect"
)

const uri = "https://openapi.naver.com/v1/nid/me"

var _ = (koauth.KOAuth)(&Naver{})

type Naver struct {
	client http.Client
}

func New() *Naver {
	return &Naver{
		client: http.Client{},
	}
}

func (n *Naver) GetUserInfo(token string) (ui koauth.UserInfo, err error) {
	nui := UserInfo{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return ui, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return ui, err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&nui)
	if err != nil {
		return ui, err
	}

	return nui, nil
}

func (n *Naver) GetField(ui koauth.UserInfo, field string) (string, error) {
	nui, ok := ui.(UserInfo)
	if !ok {
		return "", koauth.InvalidUserInfoType{
			Expected: "naver.UserInfo",
			Actual:   reflect.TypeOf(ui).String(),
		}
	}

	switch field {
	case FieldId:
		return nui.Response.Id, nil
	case FieldNickname:
		return nui.Response.Nickname, nil
	case FieldName:
		return nui.Response.Name, nil
	case FieldEmail:
		return nui.Response.Email, nil
	case FieldGender:
		return nui.Response.Gender, nil
	case FieldAge:
		return nui.Response.Age, nil
	case FieldBirthday:
		return nui.Response.Birthday, nil
	case FieldProfileImage:
		return nui.Response.ProfileImage, nil
	case FieldBirthYear:
		return nui.Response.BirthYear, nil
	case FieldMobile:
		return nui.Response.Mobile, nil
	}

	return "", koauth.InvalidField{
		Field: field,
	}
}
