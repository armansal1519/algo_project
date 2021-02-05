package utils

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kshedden/gonpy"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GenRandInt() (time.Duration, error, string) {
	start := time.Now()
	const MaxUint = ^uint(0)
	const MaxInt = int((MaxUint>>1)/2) - 1
	const MinInt = -MaxInt - 2
	rand.Seed(time.Now().UTC().UnixNano())
	//fmt.Println(MaxInt)
	//fmt.Println(MinInt)

	sampleNumber := 1
	fileName := fmt.Sprintf("./file/int/randomIntSample%d.npy", sampleNumber)

	for FileExists(fileName) {
		sampleNumber++
		fileName = fmt.Sprintf("./file/int/randomIntSample%d.npy", sampleNumber)

	}
	w, err := gonpy.NewFileWriter(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		elapsed := time.Since(start)
		return elapsed, err, ""
	}
	//defer f.Close()
	intArr := make([]int32, 10000000)
	for i := 0; i < 10000000; i++ { // Generating...
		temp := rand.Intn(MaxInt-MinInt) + MinInt
		co := rand.Float32() * 2
		temp = int(float32(temp) * co)
		intArr[i] = int32(temp)

	}
	err = w.WriteInt32(intArr)
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}
	elapsed := time.Since(start)
	return elapsed, err, fileName
}

func GenRanFloat() (time.Duration, error, string) {
	start := time.Now()
	max := math.MaxFloat64
	rand.Seed(time.Now().UTC().UnixNano())
	sampleNumber := 1
	fileName := fmt.Sprintf("./file/float/randomFloatSample%d.npy", sampleNumber)
	for FileExists(fileName) {
		sampleNumber++
		fileName = fmt.Sprintf("./file/float/randomFloatSample%d.npy", sampleNumber)
	}
	w, err := gonpy.NewFileWriter(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		elapsed := time.Since(start)
		return elapsed, err, ""
	}
	floatArr := make([]float64, 10000000)
	for i := 0; i < 10000000; i++ { // Generating...
		co := rand.Intn(2) + 1
		var temp float64
		if co == 1 {
			temp = rand.Float64() * max
		} else {
			temp = rand.Float64() * max * (-1)
		}
		floatArr[i] = temp

	}
	err = w.WriteFloat64(floatArr)
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}
	elapsed := time.Since(start)
	return elapsed, err, fileName

}

func WriteFloatToFile(fn string, arr []float64) (string, error) {
	fileName := fmt.Sprintf("./sortedfile/float/%s", fn)
	if FileExists(fileName) {
		return "", errors.New("file already exist")
	}
	w, err := gonpy.NewFileWriter(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	fmt.Println(len(arr))

	err = w.WriteFloat64(arr)
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

	return fileName, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GenRandString() (time.Duration, error, string) {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	sampleNumber := 1
	fileName := fmt.Sprintf("./file/string/randomStringSample%d.txt", sampleNumber)
	for FileExists(fileName) {
		sampleNumber++
		fileName = fmt.Sprintf("./file/string/randomStringSample%d.txt", sampleNumber)
	}
	f, err := os.Create(fileName) // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		elapsed := time.Since(start)
		return elapsed, err, ""
	}
	defer f.Close()
	ii := 0

	for i := 0; i < 10000000; i++ { // Generating...
		l := rand.Intn(79) + 2
		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, l)
		for i := range b {
			b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
		}
		temp := string(b)
		if temp != "" {
			ii++
		}

		if temp == "" {
			fmt.Println(1)
		}

		_, err = f.WriteString(temp + "\n") // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
	fmt.Println(ii)
	elapsed := time.Since(start)
	return elapsed, err, fileName
}

func WriteStringToFile(arr []string) string {
	sampleNumber := 1
	fileName := fmt.Sprintf("./sortedfile/string/sortedStringFile%d.txt", sampleNumber)
	for FileExists(fileName) {
		sampleNumber++
		fileName = fmt.Sprintf("./sortedfile/string/sortedStringFile%d.txt", sampleNumber)
	}
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)

	}
	defer f.Close()

	for _, item := range arr {

		_, err = f.WriteString(item + "\n") // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
	return fileName
}

func WriteIntToFile(fn string, arr []int32) (string, error) {

	fileName := fmt.Sprintf("./sortedfile/int/%s", fn)
	if FileExists(fileName) {
		return "", errors.New("file already exist")
	}

	w, err := gonpy.NewFileWriter(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}

	err = w.WriteInt32(arr)
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

	return fileName, nil

}

func FileToIntArr(fileName string) []int32 {
	r, err := gonpy.NewFileReader(fmt.Sprintf("./file/int/%s", fileName))
	if err != nil {
		log.Println("error while reading int npy")
	}
	result, _ := r.GetInt32()
	return result

}
func FileToFloatArr(fileName string) []float64 {

	r, err := gonpy.NewFileReader(fmt.Sprintf("./file/float/%s", fileName))
	if err != nil {
		log.Println("error while reading float npy")
	}

	floatArr, _ := r.GetFloat64()

	return floatArr

}

func FileToStringArr(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var str []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		str = append(str, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return str

}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func DeleteFile(path string) error {
	e := os.Remove(path)
	if e != nil {
		return e
	}
	return nil
}
