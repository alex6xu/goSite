package models

import (
"fmt"
//"github.com/revel/revel"
//"regexp"
)

type PageView struct {
	Id 				int
	hits            int
	Datetime        string
	url				string
	ip              string
}
