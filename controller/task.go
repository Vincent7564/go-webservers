package controller

import (
	"fmt"
	"go-webserver/model"
	util "go-webserver/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

type Response struct {
	ResponseCode    string
	ResponseMessage string
	Result          interface{}
}

func (cx *Controller) InsertData(c *fiber.Ctx) error {
	var task model.Task

	err := c.BodyParser(&task)

	if err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Invalid Request", err.Error()),
		})
	}

	if err := cx.DB.Create(&model.Task{Title: task.Title, Description: task.Description, IsActive: task.IsActive, CreatedAt: time.Now()}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Failed to Insert Data", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"data": util.GenerateResponse("200", "Data inserted succesfully", task),
	})
}

func (cx *Controller) DeleteData(c *fiber.Ctx) error {
	var task model.Task
	err := c.BodyParser(&task)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Invalid Request", err.Error()),
		})
	}

	if err := cx.DB.Where("id = ?", task.ID).Delete(&model.Task{}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Data failed to delete!", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"data": util.GenerateResponse("200", "Data deleted succesfully", task),
	})
}

func (cx *Controller) UpdateTask(c *fiber.Ctx) error {
	var task model.Task
	err := c.BodyParser(&task)

	if err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Invalid Request", err.Error()),
		})
	}

	if err := cx.DB.Model(&model.Task{}).Where("id = ?", task.ID).Updates(map[string]interface{}{"title": task.Title, "description": task.Description}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Failed to update data", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"data": util.GenerateResponse("200", "Data updated succesfully", task),
	})
}

func (cx *Controller) UpdateIsActive(c *fiber.Ctx) error {
	var task model.Task
	err := c.BodyParser(&task)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Invalid Request", err.Error()),
		})
	}

	if err := cx.DB.Model(&model.Task{}).Where("id = ?", task.ID).Updates(map[string]interface{}{"is_active": task.IsActive}).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Failed to update data", err.Error()),
		})
	}
	return c.JSON(fiber.Map{
		"data": util.GenerateResponse("200", "Data updated succesfully", task),
	})
}

func (cx *Controller) GetAllTask(c *fiber.Ctx) error {
	var task []model.Task
	if err := cx.DB.Find(&task).Error; err != nil {
		return c.JSON(fiber.Map{
			"data": util.GenerateResponse("400", "Invalid Request", err.Error()),
		})
	}

	result := fiber.Map{
		"tasks": task,
		"total": len(task),
	}

	return c.JSON(fiber.Map{
		"data": util.GenerateResponse("200", "Success", result),
	})
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func TestingHandler(c *fiber.Ctx) error {
	printHelloWorld()
	return nil
}
