# koauth

koauth is a simple OAuth library for naver and kakao.

## Installation

```bash
go get github.com/snowmerak/koauth
```

## Usage

```go
type KOAuth interface {
	GetUserInfo(token string) (UserInfo, error)
	GetField(ui UserInfo, field string) (string, error)
}
```

KOauth is an interface that provides a common interface for these.

### naver

```go
package main

import (
	"fmt"
	"github.com/snowmerak/koauth/lib/koauth/naver"
)

func main() {
	nc := naver.New()

	const token = "<naver-oauth-access-token>"
	ui, err := nc.GetUserInfo(token)
	if err != nil {
		panic(err)
	}

	id, err := nc.GetField(ui, naver.FieldId)
	if err != nil {
		panic(err)
	}

	nickname, err := nc.GetField(ui, naver.FieldNickname)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s(%s)\n", nickname, id)
}
```

#### naver user info field

```go
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
```

### kakao

```go
package main

import (
	"fmt"
	"github.com/snowmerak/koauth/lib/koauth/kakao"
)

func main() {
	kc := kakao.New()

	const token = "<kakao-oauth-access-token>"
	ui, err := kc.GetUserInfo(token)
	if err != nil {
		panic(err)
	}

	id, err := kc.GetField(ui, kakao.FieldId)
	if err != nil {
		panic(err)
	}

	nickname, err := kc.GetField(ui, kakao.FieldNickname)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s(%s)\n", nickname, id)
}
```

#### kakao user info field

```go
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
```
