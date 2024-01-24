package main

import (
	"fmt"
	"mechta/service"
)

func main() {
	var (
		filePath           string
		numberOfGoroutines int
		data               []service.Data
		err                error
	)

	for {
		fmt.Println("Введите директорию файла с расширением. Например: D:/files/data.json")
		fmt.Scan(&filePath)

		data, err = service.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
		}
		if len(data) > 0 {
			break
		}
	}

	for {
		fmt.Println("Введите число горутин для запуска")
		_, err = fmt.Scan(&numberOfGoroutines)
		if err != nil {
			fmt.Println("Ожидается число")
		}

		if numberOfGoroutines <= 0 {
			fmt.Println("Введено число меньше или равно 0. Введите число больше 0!")
		}
		if numberOfGoroutines > 0 {
			break
		}

	}

	err = service.Run(numberOfGoroutines, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
