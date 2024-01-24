package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func ReadFile(filePath string) ([]Data, error) {
	var (
		data []Data
		err  error
	)

	splitFilePath := strings.SplitN(filePath, ".", -1)
	fileExtension := splitFilePath[len(splitFilePath)-1]

	if fileExtension != "json" {
		return []Data{}, errors.New("Неверный формат файла. Формат должен быть .json")
	}

	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("readFile func os.ReadFile error:", err.Error())
		return []Data{}, errors.New("Файл не найден")
	}

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		log.Println("Неверная структура данных в файле:", err.Error())
		return []Data{}, errors.New("Неверная структура данных в файле")
	}

	if len(data) == 0 {
		return []Data{}, errors.New("Файл пустой")
	}

	return data, nil
}

func Run(numberOfGoroutines int, data []Data) error {

	var (
		wg     sync.WaitGroup
		result int
		mu     sync.Mutex
	)

	wg.Add(numberOfGoroutines)
	processRange := len(data) / numberOfGoroutines

	for i := 0; i < numberOfGoroutines; i++ {
		var lastIndex int
		var firstIndex = i * processRange
		if i == numberOfGoroutines-1 {
			lastIndex = len(data) - 1
		} else {
			lastIndex = (i+1)*processRange - 1
		}
		go func(j int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			result += calculate(data, firstIndex, lastIndex)
		}(i)
	}

	wg.Wait()
	fmt.Println("Result:", result)
	return nil
}

func calculate(data []Data, startingIndex, endingIndex int) int {
	result := 0
	for i := startingIndex; i <= endingIndex; i++ {
		result += data[i].A + data[i].B
	}

	return result
}
