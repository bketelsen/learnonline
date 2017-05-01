package models

import (
	"log"
	"net/url"
	"strconv"


	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}
func getID(s string) (int, error) {
	//?type=Module&id=4
	u, err := url.Parse(s)
	if err != nil {
		return 0, err
	}
	vals := u.Query()
	ii, ok := vals["id"]
	if !ok {
		return 0, err
	}
	return strconv.Atoi(ii[0])
}
