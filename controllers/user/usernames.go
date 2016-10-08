package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	macaron "gopkg.in/macaron.v1"
)

func GetUsernames(ctx *macaron.Context) string {
	// connect to our database server with data source name
	// data source name configuration has the following parameters :
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	// example config :
	// user:password@tcp(127.0.0.1:3306)/database

	conn, err := sql.Open("mysql", "vmp94aeg7w7ey3he:s7wx433pda195t2e@tcp(d5x4ae6ze2og6sjo.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/ezl6t1vp2c6tx218")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// use your own select statement
	// this is just an example statement
	statement, err := conn.Prepare("select name from user limit 10")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query() // execute our select statement

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var names []string

	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}

	conn.Close()

	r := struct {
		Success bool     `json:"success"`
		Names   []string `json:"names"`
	}{
		true,
		names,
	}

	rj, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}

	ctx.Header().Set("Content-Type", "application/json")

	return fmt.Sprintf("%s", rj)
}
