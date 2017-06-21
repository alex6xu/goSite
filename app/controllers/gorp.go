package controllers

import (
	
	"database/sql"
    
    "golang.org/x/crypto/bcrypt"
    
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
	"github.com/revel/modules/db/app"

    "letsgo/app/models"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Username": 20,
		"Name":     100,
	})

	 t = Dbm.AddTable(models.PageView{}).SetKeys(true, "Id")
	 setColumnSizes(t, map[string]int{
		 "Hits":		32,
		 "Datetime":    64,
		 "Url":			200,
		 "HostIp":      32,
	 })

	// t = Dbm.AddTable(models.Booking{}).SetKeys(true, "BookingId")
	// t.ColMap("User").Transient = true
	// t.ColMap("Hotel").Transient = true
	// t.ColMap("CheckInDate").Transient = true
	// t.ColMap("CheckOutDate").Transient = true
	// setColumnSizes(t, map[string]int{
	// 	"CardNumber": 16,
	// 	"NameOnCard": 50,
	// })

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("demo"), bcrypt.DefaultCost)
	demoUser := &models.User{0, "Demo User", "demo", "demo", bcryptPassword}
	if err := Dbm.Insert(demoUser); err != nil {
		panic(err)
	}

	// hotels := []*models.Hotel{
	// 	&models.Hotel{0, "Marriott Courtyard", "Tower Pl, Buckhead", "Atlanta", "GA", "30305", "USA", 120},
	// 	&models.Hotel{0, "W Hotel", "Union Square, Manhattan", "New York", "NY", "10011", "USA", 450},
	// 	&models.Hotel{0, "Hotel Rouge", "1315 16th St NW", "Washington", "DC", "20036", "USA", 250},
	// }
	// for _, hotel := range hotels {
	// 	if err := Dbm.Insert(hotel); err != nil {
	// 		panic(err)
	// 	}
	// }
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
