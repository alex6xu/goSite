package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"letsgo/app/models"
	"letsgo/app/routes"
	//"fmt"
	"log"
	"time"
)

type App struct {
	GorpController
}

func (c App) getUser(username string) *models.User {
	users, err := c.Txn.Select(models.User{}, `select * from User where Username = ?`, username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.User)
}

func (c App) Index() revel.Result {
	//fmt.Print("index start")
	result1, err := Dbm.SelectInt("select sum(hits) from PageView ")
	result2, err := Dbm.SelectInt("select sum(hits) from PageView where Date = ?", time.Now().Format("2006-01-02"))
	result3, err := Dbm.SelectInt("select sum(hits) from PageView where Url = ?", c.Request.URL.Path)
	result4, err := Dbm.SelectInt("select sum(hits) from PageView where Date = ? and Url =?", time.Now().Format("2006-01-02"), c.Request.URL.Path)
	checkErr(err, "database error!")
	//fmt.Print(tdy, err)
	//fmt.Print(tdy)
	all_visits := result1
	today_visits := result2
	this_page_visits := result3
	page_visits_today := result4
	
	return c.Render(all_visits, today_visits, this_page_visits, page_visits_today)
}

func (c App) Login(username, password string, remember bool) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Blog.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.App.Index())
}

func (c App) Register(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
		Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Index())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}

	c.Session["user"] = user.Username
	c.Flash.Success("Welcome, " + user.Name)
	return c.Redirect(routes.App.Profile())
}

func (c App) Profile() revel.Result {
	return c.Render()
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}