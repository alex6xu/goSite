package controllers

import (
	"github.com/revel/revel"
	"letsgo/app/models"
	"time"
)


func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptFunc(recordRequest, revel.BEFORE, &App{})
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	// revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	// revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

}

func doNothing(c *revel.Controller) revel.Result {
	return nil
}

func recordRequest(c *GorpController) revel.Result {
	page, err := c.Txn.Select(models.PageView{}, "select * from PageView where path = ?", c.Request.URL.Path)
	if err != nil {
		count, _ := c.Txn.Query("select * from PageView")

		newpage := models.PageView{
			count. + 1,
			1,
			time.Now().String(),
			c.Request.URL.Path,
			c.Request.Host,

		}
		c.Txn.Insert(newpage)
		return nil
	}

	page.hits += 1
	c.Txn.Insert(page)
	return nil
}

func countTimesOfAction(c *revel.Controller)  {

}

func countTimesofCompelete(c *revel.Controller)  {

}