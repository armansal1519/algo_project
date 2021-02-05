package randstring

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"server/utils"
)

func CreateRandomString(c *fiber.Ctx) error {

	ds, err := utils.DirSize("./file/string")
	if err != nil {
		return err
	}
	if ds > 1024000000 {
		fmt.Println(ds)
		return fiber.NewError(fiber.StatusConflict, "dir int is too big, please delete some sample first")
	}
	time, err, path := utils.GenRandString()
	if err != nil {
		fmt.Println(err)
	}

	return c.SendString(fmt.Sprintf("time:%v \npath:%s", time, path))
}

func GetFilesFormFolder(c *fiber.Ctx) error {
	files, err := ioutil.ReadDir("./file/string")
	if err != nil {
		log.Fatal(err)
	}
	type filesStruct struct {
		filesName []string
	}
	var temp []string
	for _, f := range files {
		fmt.Println(f.Name())
		temp = append(temp, f.Name())

	}

	data := filesStruct{filesName: temp}
	fmt.Println(data)
	return c.JSON(fiber.Map{
		"filesName": temp,
	})
}

func DeleteIntFile(c *fiber.Ctx) error {
	path := fmt.Sprintf("./file/string/%s", c.Params("filepath"))
	err := utils.DeleteFile(path)
	if err != nil {
		return err
	}
	return nil
}

//func SortString(c *fiber.Ctx) error {
//	algo := c.Params("algo")
//	fileName := c.Params("filename")
//	if utils.FileExists(fmt.Sprintf("./sortedfile/string/%s", fileName)) {
//		return fiber.NewError(fiber.StatusConflict, "file already exist")
//	}
//	startReadFile := time.Now()
//	intArr := utils.FileToIntArr(fileName)
//	//if err != nil {
//	//	return err
//	//}
//
//	fileReadTime := time.Since(startReadFile)
//	startAlgo := time.Now()
//
//	if algo == "default" {
//		algoBase.SortInt(intArr)
//	} else {
//		return fiber.ErrConflict
//	}
//	algoRunTime := time.Since(startAlgo)
//
//	startWrite := time.Now()
//	sortedFile, _ := utils.WriteIntToFile(fileName, intArr)
//	writeRunTime := time.Since(startWrite)
//
//	return c.JSON(fiber.Map{
//		"fileName":     sortedFile,
//		"fileReadTime": fileReadTime,
//		"algoRunTime":  algoRunTime,
//		"writeRunTime": writeRunTime,
//	})
//
//}