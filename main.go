package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Node struct {
	ID       int    `json:"id"`
	Text     string `json:"Text"`
	Min      string `json:"Min"`
	Value    string `json:"Value"`
	Max      string `json:"Max"`
	ImageURL string `json:"ImageURL"`
	Children []Node `json:"Children"`
}

func main() {
	url := "http://localhost:8085/data.json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	tempCPU1, tempCPU2 := getTemp(body, "cpu")
	tempGPU1, tempGPU2 := getTemp(body, "gpu")
	if tempCPU1 != "" {
		fmt.Println("Температура CPU = ", tempCPU1)
	} else if tempCPU2 != "" {
		fmt.Println("Температура CPU = ", tempCPU2)
	}

	if tempGPU1 != "" {
		fmt.Println("Температура GPU = ", tempGPU1)
	}
	if tempGPU2 != "" {
		fmt.Println("Температура HotSpot = ", tempGPU2)
	}

	var a string
	fmt.Println("ожидание ввода")
	fmt.Scan(&a)
}

func getTemp(body []byte, xpu string) (value1, value2 string) {

	var root Node
	var text1, text2, text3 string
	err := json.Unmarshal(body, &root)
	if err != nil {
		panic(err)
	}

	value1, value2 = "", ""

	if xpu == "cpu" {
		text1 = "images_icon/cpu.png"
		text2 = "Core (Tctl/Tdie)"
		text3 = "CPU Package"
	} else if xpu == "gpu" {
		text1 = "images_icon/nvidia.png"
		text2 = "GPU Core"
		text3 = "GPU Hot Spot"
	}

	// Перебор первого уровня
	for _, child := range root.Children {
		// Перебор второго уровня
		for _, subChild := range child.Children {
			// fmt.Println("child.Children - ", subChild.ImageURL)
			if subChild.ImageURL == text1 {
				// Перебор третьего уровня
				for _, subSubChild := range subChild.Children {
					// fmt.Println("subSubChild.Children - ", subSubChild.ImageURL)

					if subSubChild.Text == "Temperatures" {

						// Перебор четвертого уровня
						for _, subSubSubChild := range subSubChild.Children {
							// fmt.Println("subSubSubChild.Text - ", subSubSubChild.Text)

							if subSubSubChild.Text == text2 {
								// fmt.Println("Значение поля Value:", subSubSubChild.Value)
								value1 = subSubSubChild.Value
								// return value1
							}
							if subSubSubChild.Text == text3 {
								// fmt.Println("Значение поля Value:", subSubSubChild.Value)
								value2 = subSubSubChild.Value
								// return value2
							}
						}
					}
				}
			}
		}
	}
	return value1, value2
}
