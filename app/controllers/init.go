package controllers

import (
	"github.com/revel/revel"
	"letsgo/app/models"
	"time"
	//"fmt"
)


func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptFunc(recordRequest, revel.BEFORE, &GorpController{})
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
	tday := time.Now().Format("2006-01-02")
	//fmt.Printf(tday)
	page := &models.PageView{}
	err := Dbm.SelectOne(&page, "select * from PageView where Url = ? and Date = ?", c.Request.URL.Path, tday)
	//fmt.Printf("page is %s,  err is %s", page, err)
	if err != nil {
		page = &models.PageView{
			0, 1, tday, c.Request.URL.Path, c.Request.Host,
		}
		Dbm.Insert(page)
		return nil
	}

	page.Hits += 1
	page.Date = tday
	Dbm.Update(page)
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