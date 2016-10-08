package routes

import (
	"fmt"
	"tryMaria/controllers/user"
)

func init() {
	M.Group("/user", func() {
		M.Get("/usernames", user.GetUsernames)
	}, func() {
		fmt.Println("MIDDLEWARE")
	})
}
