package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CPUData struct {
	ID       int       `json:"id"`
	Text     string    `json:"Text"`
	Min      string    `json:"Min"`
	Value    string    `json:"Value"`
	Max      string    `json:"Max"`
	ImageURL string    `json:"ImageURL"`
	Children []GPUCore `json:"Children"`
}

type GPUCore struct {
	ID       int         `json:"id"`
	Text     string      `json:"Text"`
	Min      string      `json:"Min"`
	Value    string      `json:"Value"`
	Max      string      `json:"Max"`
	ImageURL string      `json:"ImageURL"`
	Children []Processor `json:"Children"`
}

type Processor struct {
	ID       int           `json:"id"`
	Text     string        `json:"Text"`
	Min      string        `json:"Min"`
	Value    string        `json:"Value"`
	Max      string        `json:"Max"`
	ImageURL string        `json:"ImageURL"`
	Children []Temperature `json:"Children"`
}

type Temperature struct {
	ID       int           `json:"id"`
	Text     string        `json:"Text"`
	Min      string        `json:"Min"`
	Value    string        `json:"Value"`
	Max      string        `json:"Max"`
	SensorID string        `json:"SensorId"`
	Type     string        `json:"Type"`
	ImageURL string        `json:"ImageURL"`
	Children []Temperature `json:"Children"`
}

func main() {
	// Создаем HTTP-клиента
	client := http.Client{}

	// Создаем GET-запрос
	req, err := http.NewRequest("GET", "http://localhost:8085/data.json", nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Отправляем запрос и получаем ответ
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка:", resp.Status)
		return
	}

	// Читаем тело ответа в байтовый массив
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return
	}

	// Создаем структуру для разбора JSON
	var cpuData CPUData // Замените YourDataType на свою структуру данных

	// Разбираем JSON
	err = json.Unmarshal(jsonData, &cpuData)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	var cpuPackageValue, gpuCoreValue, gpuCoreHotSpot string
	// Получить значение поля "Value" для CPU package
	// for i; v :=

	cpuPackageValue = cpuData.Children[0].Children[1].Children[3].Children[0].Value
	fmt.Println("Значение поля 'Value' для CPU:", cpuPackageValue)

	// Получить значение поля "Value" для GPU core
	gpuCoreValue = cpuData.Children[0].Children[3].Children[2].Children[0].Value
	fmt.Println("Значение поля 'Value' для GPU core:", gpuCoreValue)

	// Получить значение поля "Value" для GPU Hot Spot
	gpuCoreHotSpot = cpuData.Children[0].Children[3].Children[2].Children[1].Value
	fmt.Println("Значение поля 'Value' для GPU Hot Spot:", gpuCoreHotSpot)
	var a string
	fmt.Scan(&a)
}
