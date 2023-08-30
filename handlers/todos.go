package handlers

import (
	"strconv"
	"todoist/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate the Task model (if needed)
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		panic("failed to auto migrate model")
	}
}

func GetAllTodos(c *fiber.Ctx) error {
	var tasks []models.Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving tasks",
			"error":   err.Error(),
		})
	}

	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := db.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating task",
			"error":   err.Error(),
		})
	}

	var tasks []models.Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving tasks",
			"error":   err.Error(),
		})
	}

	return c.JSON(tasks)
}

func CompletedTask(c *fiber.Ctx) error {
	id := c.Params("id")
	taskId, err := strconv.Atoi(id)
	var task models.Task

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task id",
			"error":   err.Error(),
		})
	}
	db.Model(&task).Where("id = ?", taskId).Update("is_completed", true)

	var tasks []models.Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving tasks",
			"error":   err.Error(),
		})
	}

	return c.JSON(tasks)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	taskId, err := strconv.Atoi(id)
	var task models.Task

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task id",
			"error":   err.Error(),
		})
	}
	db.Delete(&task, taskId)

	var tasks []models.Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving tasks",
			"error":   err.Error(),
		})
	}

	return c.JSON(tasks)
}