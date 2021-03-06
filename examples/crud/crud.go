package main

import (
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vingarcia/kissorm"
	"github.com/vingarcia/kissorm/nullable"
)

// User ...
type User struct {
	ID   int    `kissorm:"id"`
	Name string `kissorm:"name"`
	Age  int    `kissorm:"age"`

	// This field will be saved as JSON in the database
	Address Address `kissorm:"address,json"`
}

// PartialUpdateUser ...
type PartialUpdateUser struct {
	ID      int      `kissorm:"id"`
	Name    *string  `kissorm:"name"`
	Age     *int     `kissorm:"age"`
	Address *Address `kissorm:"address,json"`
}

// Address ...
type Address struct {
	State string `json:"state"`
	City  string `json:"city"`
}

func main() {
	ctx := context.Background()
	db, err := kissorm.New("sqlite3", "/tmp/hello.sqlite", kissorm.Config{
		MaxOpenConns: 1,
		TableName:    "users",
	})
	if err != nil {
		panic(err.Error())
	}

	// In the definition below, please note that BLOB is
	// the only type we can use in sqlite for storing JSON.
	err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
	  id INTEGER PRIMARY KEY,
		age INTEGER,
		name TEXT,
		address BLOB
	)`)
	if err != nil {
		panic(err.Error())
	}

	var alison = User{
		Name: "Alison",
		Age:  22,
		Address: Address{
			State: "MG",
		},
	}
	err = db.Insert(ctx, &alison)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Alison ID:", alison.ID)

	// Inserting inline:
	err = db.Insert(ctx, &User{
		Name: "Cristina",
		Age:  27,
		Address: Address{
			State: "SP",
		},
	})
	if err != nil {
		panic(err.Error())
	}

	// Deleting Alison:
	err = db.Delete(ctx, alison.ID)
	if err != nil {
		panic(err.Error())
	}

	// Retrieving Cristina:
	var cris User
	err = db.QueryOne(ctx, &cris, "SELECT * FROM users WHERE name = ? ORDER BY id", "Cristina")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Cristina: %#v\n", cris)

	// Updating all fields from Cristina:
	cris.Name = "Cris"
	err = db.Update(ctx, cris)

	// Changing the age of Cristina but not touching any other fields:

	// Partial update technique 1:
	err = db.Update(ctx, struct {
		ID  int `kissorm:"id"`
		Age int `kissorm:"age"`
	}{ID: cris.ID, Age: 28})
	if err != nil {
		panic(err.Error())
	}

	// Partial update technique 2:
	err = db.Update(ctx, PartialUpdateUser{
		ID:  cris.ID,
		Age: nullable.Int(28),
	})
	if err != nil {
		panic(err.Error())
	}

	// Listing first 10 users from the database
	// (each time you run this example a new Cristina is created)
	//
	// Note: Using this function it is recommended to set a LIMIT, since
	// not doing so can load too many users on your computer's memory or
	// cause an Out Of Memory Kill.
	//
	// If you need to query very big numbers of users we recommend using
	// the `QueryChunks` function.
	var users []User
	err = db.Query(ctx, &users, "SELECT * FROM users LIMIT 10")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Users: %#v\n", users)
}
