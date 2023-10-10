package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"project/config"
	"project/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// get all users
func GetUsersController(c echo.Context) error {
  var users []models.User


  if err := config.DB.Find(&users).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success get all users",
    "users":   users,
  })
}


// get user by id
func GetUserController(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, "Invalid user id!")
  }

  
	var users models.User

  if err := config.DB.Where("id = ?", id).Find(&users).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
	return c.JSON(http.StatusOK, map[string]interface{}{
  		"messages": "Success get user by id",
		  "user" : users,
  })	
}


// create new user
func CreateUserController(c echo.Context) error {
  user := models.User{}
  c.Bind(&user)


  if err := config.DB.Save(&user).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success create new user",
    "user":    user,
  })
}


// delete user by id
func DeleteUserController(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, "Invalid user id!")
  }
  

  if err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
	return c.JSON(http.StatusOK, map[string]interface{}{
  		"messages": "Success delete user by id",
  })  
}


// update user by id
func UpdateUserController(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, "Invalid user id!")
  }
  
  // binding data user 
  updatedUser := models.User{}
  if err := c.Bind(&updatedUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

  // mencari apakah terdapat data user dengan id yg diberikan
  var existingUser models.User 
  if err := config.DB.First(&existingUser, id).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
      return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }

  // update data user
  existingUser.Name = updatedUser.Name
  existingUser.Email = updatedUser.Email
  existingUser.Password = updatedUser.Password
  

  if err := config.DB.Save(&existingUser).Where("id = ?", id).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success create new user",
    "user": existingUser,
  })
}
