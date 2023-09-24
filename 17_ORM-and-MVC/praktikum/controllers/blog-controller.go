package controllers

import (
	"praktikum/config"
	"praktikum/helpers"
	"praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all blogs

func GetBlogsController(c echo.Context) error {

	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error get data"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all blog", blogs))

}

// get blog by id

func GetBlogController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	blog := models.Blog{}
	err := config.DB.First(&blog, id).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Blog not found"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get blog", blog))
}

// create new blog

func CreateBlogController(c echo.Context) error {

	blog := models.Blog{}

	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed creat blog"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success create new blog", blog))

}

// delete blog by id

func DeleteBlogController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	blog := models.Blog{}
	err := config.DB.Delete(&blog, id).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed delete blog"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success delete blog", blog))

}

// update blog by id

func UpdateBlogController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	blog := models.Blog{}
	err := config.DB.First(&blog, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed update blog"))
	}

	update := new(models.Blog)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Blog data not valid"))
	}

	blog.Judul = update.Judul
	blog.Konten = update.Konten

	config.DB.Save(&blog)

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success update blog", blog))

}
