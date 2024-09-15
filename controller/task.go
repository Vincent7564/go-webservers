package controller

import (
	"fmt"
	"go-webserver/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Is_active   bool   `json:"is_active"`
}

type Response struct {
	ResponseCode    string
	ResponseMessage string
	Result          interface{}
}

func (cx *Controller) InsertData(c *fiber.Ctx) error {
	var task Task

	err := c.BodyParser(&task)

	if err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "400",
				ResponseMessage: "Data invalid, please re-check the request",
				Result:          err.Error(),
			},
		})
	}

	if err := cx.DB.Create(&model.Task{Title: task.Title, Description: task.Description, Is_Active: task.Is_active}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "400",
				ResponseMessage: "Failed to insert data",
				Result:          err.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": Response{
			ResponseCode:    "200",
			ResponseMessage: "Data inserted succesfully",
			Result:          task,
		},
	})
}

func (cx *Controller) DeleteData(c *fiber.Ctx) error {
	var task Task
	err := c.BodyParser(&task)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "400",
				ResponseMessage: "Data invalid, please re-check the request",
				Result:          err.Error(),
			},
		})
	}

	if err := cx.DB.Where("id = ?", task.Id).Delete(&Task{}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "400",
				ResponseMessage: "Data failed to deleted",
				Result:          err.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": Response{
			ResponseCode:    "200",
			ResponseMessage: "Data deleted succesfully",
			Result:          task,
		},
	})
}

func (cx *Controller) UpdateTask(c *fiber.Ctx) error {
	var task Task
	err := c.BodyParser(&task)

	if err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "400",
				ResponseMessage: "Data invalid, please re-check the request",
				Result:          err.Error(),
			},
		})
	}

	if err := cx.DB.Model(&model.Task{}).Where("id = ?", task.Id).Updates(map[string]interface{}{"title": task.Title, "description": task.Description}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "500",
				ResponseMessage: "Update data failed",
				Result:          err.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": Response{
			ResponseCode:    "200",
			ResponseMessage: "Data updated succesfully",
			Result:          task,
		},
	})
}

func (cx *Controller) UpdateIsActive(c *fiber.Ctx) error {
	var task Task
	err := c.BodyParser(&task)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "400",
				ResponseMessage: "Data invalid, please re-check the request",
				Result:          err.Error(),
			},
		})
	}

	if err := cx.DB.Model(&model.Task{}).Where("id = ?", task.Id).Updates(map[string]interface{}{"is_active": task.Is_active}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": Response{
				ResponseCode:    "500",
				ResponseMessage: "Update data failed",
				Result:          err.Error(),
			},
		})
	}
	return c.JSON(fiber.Map{
		"data": Response{
			ResponseCode:    "200",
			ResponseMessage: "Data updated succesfully",
			Result:          task,
		},
	})
}

func (cx *Controller) GetAllTask(c *fiber.Ctx) error {
	var task []Task
	response := Response{}
	response.ResponseCode = "200"
	response.ResponseMessage = "Data retrieved succesfully!"
	if err := cx.DB.Find(&task).Error; err != nil {
		response.ResponseCode = "500"
		response.ResponseMessage = "Internal Server Error"
		response.Result = err
		return c.JSON(fiber.Map{
			"data": response,
		})
	} else {
		response.Result = task
	}
	return c.JSON(fiber.Map{
		"data": response,
	})
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func TestingHandler(c *fiber.Ctx) error {
	printHelloWorld()
	return nil
}
