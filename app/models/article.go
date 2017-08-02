package models

import (
	"fmt"
	//"time"
)


type Article struct {
	Id             			int
	DateCreated             string
	DateUpdated             string
	Title					string
	Content     			string

}

func (a *Article) String() string {
	return fmt.Sprintf("Title(%s)", a.Title)
}

func (a *Article) Preview() string {
	return fmt.Sprint("%s", a.Content[100])
}

//var userRegex = regexp.MustCompile("^\\w*$")
//
//func (user *User) Validate(v *revel.Validation) {
//	v.Check(user.Username,
//		revel.Required{},
//		revel.MaxSize{15},
//		revel.MinSize{4},
//		revel.Match{userRegex},
//	)
//
//	ValidatePassword(v, user.Password).
//		Key("user.Password")
//
//	v.Check(user.Name,
//		revel.Required{},
//		revel.MaxSize{100},
//	)
//}
//
//func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
//	return v.Check(password,
//		revel.Required{},
//		revel.MaxSize{15},
//		revel.MinSize{5},
//	)
//}
