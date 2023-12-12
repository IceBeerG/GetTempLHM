package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Item struct {
	ID       int    `json:"id"`
	Text     string `json:"Text"`
	Min      string `json:"Min"`
	Value    string `json:"Value"`
	Max      string `json:"Max"`
	ImageURL string `json:"ImageURL"`
	Children []Item `json:"Children"`
}

func findValue(items []Item) string {
	for _, item := range items {
		if item.ImageURL == "images_icon/cpu.png" {
			for _, child := range item.Children {
				if child.Text == "Temperatures" {
					for _, grandchild := range child.Children {
						if grandchild.Text == "Core (Tctl/Tdie)" {
							return grandchild.Value
						}
					}
				}
			}
		}
	}

	return ""
}

func main() {
	// Чтение файла
	fileContent, err := ioutil.ReadFile("data-4070ti.json")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Декодирование JSON
	var items []Item
	err = json.Unmarshal(fileContent, &items)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	// Поиск значения
	value := findValue(items)
	fmt.Println("Значение поля Value:", value)
}
