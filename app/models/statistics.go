package models

import (
	//"fmt"
//"github.com/revel/revel"
//"regexp"
)

type PageView struct {
	Id 				int
	Hits            int
	Datetime        int64
	Url				string
	HostIp          string
}
