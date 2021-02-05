package randfloat

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"server/algoBase"
	"server/algoBase/sortFloat"
	"server/utils"
	"time"
)

func CreateRandomInt(c *fiber.Ctx) error {

	ds, err := utils.DirSize("./file/float")
	if err != nil {
		return err
	}
	if ds > 1024000000 {
		fmt.Println(ds)
		return fiber.NewError(fiber.StatusConflict, "dir int is too big, please delete some sample first")
	}
	time, err, path := utils.GenRanFloat()
	if err != nil {
		fmt.Println(err)
	}

	return c.SendString(fmt.Sprintf("time:%v \npath:%s", time, path))
}

func GetFilesFormFolder(c *fiber.Ctx) error {
	files, err := ioutil.ReadDir("./file/float")
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

func SortFloat(c *fiber.Ctx) error {
	algo := c.Params("algo")
	fileName := c.Params("filename")

	if utils.FileExists(fmt.Sprintf("./sortedfile/float/%s", fileName)) {
		return fiber.NewError(fiber.StatusConflict, "file already exist")
	}
	startReadFile := time.Now()
	floatArr := utils.FileToFloatArr(fileName)

	fileReadTime := time.Since(startReadFile)
	startAlgo := time.Now()

	if algo == "default" {
		algoBase.DefaultSortFloat(floatArr)
	} else if algo == "singleMerge" {
		floatArr= sortFloat.SingleMergeSort(floatArr)
	}else if algo == "multiMerge" {
		floatArr= sortFloat.RunMultiMergeSort(floatArr)
	} else {
		return fiber.NewError(fiber.StatusConflict, "alog not found")
	}
	algoRunTime := time.Since(startAlgo)

	startWrite := time.Now()
	sortedFile, err := utils.WriteFloatToFile(fileName, floatArr)
	if err != nil {
		log.Println(err)
	}
	writeRunTime := time.Since(startWrite)

	return c.JSON(fiber.Map{
		"fileName":     sortedFile,
		"fileReadTime": fileReadTime,
		"algoRunTime":  algoRunTime,
		"writeRunTime": writeRunTime,
	})

}

func DeleteIntFile(c *fiber.Ctx) error {
	path := fmt.Sprintf("./file/float/%s", c.Params("filepath"))
	err := utils.DeleteFile(path)
	if err != nil {
		return err
	}
	return nil
}
