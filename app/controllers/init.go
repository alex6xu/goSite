package controllers

import (
	"github.com/revel/revel"
	"letsgo/app/models"
	"time"
)


func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptFunc(recordRequest, revel.BEFORE, &App{})
	revel.InterceptFunc(doNothing, revel.BEFORE, &App{})
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	// revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	// revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

}

func doNothing(c *revel.Controller) revel.Result {
	return nil
}

func recordRequest(c *revel.Controller) revel.Result {
	page := &models.PageView{}
	err := Dbm.SelectOne(&page, "select * from PageView where path = ?", &page.Id)
	if err != nil {
		//count, _ := Dbm.Query("select * from PageView")

		page = &models.PageView{
			1, 1, time.Now().UnixNano(), c.Request.URL.Path, c.Request.Host,
		}
	}

	Dbm.Insert(page)
	return nil
}

//
//func countTimesOfAction(c *revel.Controller)  {
//
//}
//
//func countTimesofCompelete(c *revel.Controller)  {
//
//}