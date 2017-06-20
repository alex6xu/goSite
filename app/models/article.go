package models

import (
	"fmt"
	//"github.com/revel/revel"
	//"regexp"
)

type Review_Article struct {
	Article			int
	Review			int

}

type Review struct {
	Id				int
	User			int
	Content			string

}

type Article struct {
	Id             			int
	DateCreated             string
	DateUpdated             string
	title					string
	content     			string

}

func (a *Article) String() string {
	return fmt.Sprintf("Title(%s)", a.title)
}

func (a *Article) Preview() string {
	return fmt.Sprint("%s", a.content[100])
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
