package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Salary   int    `json:"salary"`
}

var obj []Employee //obj is slice of Employee struct

func getEmployee(c *fiber.Ctx) error {
	obj = append(obj, Employee{ID: 1, Name: "Muhammad", Division: "CAKE Software", Salary: 1234})
	obj = append(obj, Employee{ID: 2, Name: "John", Division: "CAKE Software", Salary: 1264})
	obj = append(obj, Employee{ID: 3, Name: "Alex", Division: "CAKE Software", Salary: 1520})
	obj = append(obj, Employee{ID: 4, Name: "Doe", Division: "CAKE Software", Salary: 1502})

	return c.Status(fiber.StatusOK).JSON(obj)
}

func createEmployee(c *fiber.Ctx) error {
	body := new(Employee)     //getting data from the body
	err := c.BodyParser(body) //parse the body
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	obj1 := Employee{
		ID:       len(obj) + 1,
		Name:     body.Name,
		Division: body.Division,
		Salary:   body.Salary,
	}
	obj = append(obj, obj1)
	return c.Status(fiber.StatusOK).JSON(obj)
}

func getEmployeeById(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id") //return id in string
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	var res Employee
	for _, x := range obj {
		if x.ID == id {
			res = x
			break
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func deleteEmployee(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	var res []Employee
	for _, x := range obj {
		if x.ID != id {
			res = append(res, x)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func updateEmployee(ctx *fiber.Ctx) error {
	body := new(Employee)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	var res Employee
	for _, x := range obj {
		if x.ID == id {
			x.Name = body.Name
			x.Division = body.Division
			x.Salary = body.Salary
			res = x
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Fiber")
	})

	app.Use(logger.New()) //maintain the log
	app.Get("/emp", getEmployee)
	app.Post("/emp", createEmployee)
	app.Get("/emp/:id", getEmployeeById)
	app.Delete("/emp/:id", deleteEmployee)
	app.Patch("/emp/:id", updateEmployee)

	log.Fatal(app.Listen(":1000"))
}
