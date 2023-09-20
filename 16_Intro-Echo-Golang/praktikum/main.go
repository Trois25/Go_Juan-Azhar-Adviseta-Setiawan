package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "User Found",
				"data":    user,
			})
		} else {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "User Not Found",
			})
		}
	}
	return nil
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	for i, user := range users {
		if user.Id == id {
			// Hapus pengguna dari slice users menggunakan indeks i
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "User Deleted",
				"data":    user,
			})
		}else{
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "User Not Found",
			})
		}
	}

	return nil
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
    updatedUser := User{} 
    
    for i, user := range users {
        if user.Id == id {
			c.Bind(&updatedUser)

			updatedUser.Id = id

            users[i] = updatedUser

            return c.JSON(http.StatusOK, map[string]interface{}{
                "message": "Success Update User",
                "data":    updatedUser,
            })
        }else{
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "User Not Found",
			})
		}
    }
	return nil    
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
