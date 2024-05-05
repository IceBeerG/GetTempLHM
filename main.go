package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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

type lhminfo struct {
	level1 string
	level2 string
	level3 string
}

func cpuTemp() (value int) {
	h := lhminfo{"images_icon/cpu.png", "Temperatures", "Core (Tctl/Tdie)"}
	value1 := getValueLHM(h)
	if value1 != "-1" {
		value = takeInt(value1)
	} else {
		h.level3 = "CPU Package"
		value = takeInt(getValueLHM(h))
	}
	return value
}

func gpuCoreTemp() (value int) {
	h := lhminfo{"images_icon/nvidia.png", "Temperatures", "GPU Core"}
	value = takeInt(getValueLHM(h))
	return
}

func gpuHsTemp() (value int) {
	h := lhminfo{"images_icon/nvidia.png", "Temperatures", "GPU Hot Spot"}
	value = takeInt(getValueLHM(h))
	return
}

func gpuFan() (value int) {
	h := lhminfo{"images_icon/nvidia.png", "Fans", "GPU Fan 1"}
	value1 := getValueLHM(h)
	if value1 != "-1" {
		value = takeInt(value1)
	} else {
		h.level3 = "GPU Fan"
		value = takeInt(getValueLHM(h))
	}
	return value
}

func gpuFanPercent() (value int) {
	h := lhminfo{"images_icon/nvidia.png", "Controls", "GPU Fan 1"}
	value1 := getValueLHM(h)
	if value1 != "-1" {
		value = takeInt(value1)
	} else {
		h.level3 = "GPU Fan"
		value = takeInt(getValueLHM(h))
	}
	return value
}

func gpuFan2() (value int) {
	h := lhminfo{"images_icon/nvidia.png", "Fans", "GPU Fan 2"}
	value = takeInt(getValueLHM(h))
	return value
}

func gpuFan2Percent() (value int) {
	h := lhminfo{"images_icon/nvidia.png", "Controls", "GPU Fan 2"}
	value = takeInt(getValueLHM(h))
	return value
}

func printFan(s string, x int) {
	if x != -1 {
		fmt.Println(s, x)
	}
}

func takeInt(valueS string) (valueI int) {
	valueS = strings.Split(valueS, ",")[0]
	valueS = strings.Split(valueS, " ")[0]
	valueI, err := strconv.Atoi(valueS)
	if err != nil {
		fmt.Println("Ошибка конвертирования. ", err)
	}
	return valueI
}

func main() {
	fmt.Println("t CPU =", cpuTemp())
	fmt.Println("t GPU =", gpuCoreTemp())
	fmt.Println("t GPU =", gpuHsTemp())
	printFan("GPU Fan =", gpuFan())
	printFan("GPU Fan =", gpuFanPercent())
	printFan("GPU Fan2 =", gpuFan2())
	printFan("GPU Fan2 =", gpuFan2Percent())
	a := ""
	fmt.Scan(&a)
}

func UnmarshalLHM() (lhmJsonData Node) {
	var body []byte
	urlLHM := "http://localhost:8085/data.json"
	_, err := http.Get(urlLHM)
	if err != nil {
		log.Println("[ERROR] web-сервер LHM недоступен")
		_, err := fmt.Println("web-сервер LHM недоступен")
		if err != nil {
			log.Println("[ERROR] Ошибка отправки сообщения: ", err)
		}
	} else {
		respLHM, err := http.Get(urlLHM)
		if err != nil {
			log.Println(err)
		}
		defer respLHM.Body.Close()
		body, err = io.ReadAll(respLHM.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &lhmJsonData)
		if err != nil {
			fmt.Println(err)
		}
	}
	return lhmJsonData
}

func getValueLHM(h lhminfo) (value string) {
	lhmJsonData := UnmarshalLHM()
	value = "-1"
	// Перебор первого уровня
	for _, child := range lhmJsonData.Children {
		// Перебор второго уровня
		for _, subChild := range child.Children {
			if subChild.ImageURL == h.level1 {
				// Перебор третьего уровня
				for _, subSubChild := range subChild.Children {
					if subSubChild.Text == h.level2 {
						// Перебор четвертого уровня
						for _, subSubSubChild := range subSubChild.Children {
							if subSubSubChild.Text == h.level3 {
								value = subSubSubChild.Value
							}
						}
					}
				}
			}
		}
	}
	return value
}
