package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"server/services/randfloat"
	"server/services/randint"
	"server/services/randstring"
	"server/utils"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/int/gen", randint.CreateRandomInt)
	app.Get("/int/names", randint.GetFilesFormFolder)
	app.Delete("/int/:filepath", randint.DeleteIntFile)
	app.Get("/int/sort/:filename/:algo", randint.SortInt)
	app.Get("/float/gen", randfloat.CreateRandomInt)
	app.Get("/float/names", randfloat.GetFilesFormFolder)
	app.Delete("/float/:filepath", randfloat.DeleteIntFile)
	app.Get("/float/sort/:filename/:algo", randfloat.SortFloat)
	app.Get("/string/gen", randstring.CreateRandomString)
	app.Get("/string/names", randstring.GetFilesFormFolder)


	app.Get("/gen-string", func(c *fiber.Ctx) error {
		time, err, path := utils.GenRandString()
		if err != nil {
			fmt.Println(err)
		}

		return c.SendString(fmt.Sprintf("time:%v \npath:%s", time, path))
	})

	app.Listen(":5000")
}
