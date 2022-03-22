package main

import (
	"context"
	"fmt"

	"github.com/sojoudian/goRestAPIv2/internal/db"
)

// Run - is goin to be responsible for
// the instantiation and startup of the
// go application
func Run() error {
	fmt.Println("starting up our application")
	fmt.Println("##################################################")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Faild to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged the database!")
	return nil
}

func main() {
	fmt.Println("Go REST API v2")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
