package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	var body []byte
	rezhim := "online"
	if rezhim == "online" {
		url := "http://localhost:8085/data.json"
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		body, err = ioutil.ReadFile("data-4070ti.json")
		if err != nil {
			panic(err)
		}
	}

	tempCPU1, tempCPU2 := getTemp(body, "cpu")

	if tempCPU1 != "" {
		fmt.Println("t CPU = ", tempCPU1)
	} else if tempCPU2 != "" {
		fmt.Println("t CPU = ", tempCPU2)
	}

	// INTEL
	tempGPUi1, tempGPUi2 := getTemp(body, "gpuINTEL")
	if tempGPUi1 != "" {
		fmt.Println("t gpuINTEL = ", tempGPUi1)
	}
	if tempGPUi2 != "" {
		fmt.Println("t gpuINTEL HotSpot = ", tempGPUi2)
	}
	fanI1, fanI2 := getTemp(body, "fanINTEL")
	if fanI1 != "" {
		fmt.Println("gpuINTEL fan1 = ", fanI1)
	}
	if fanI2 != "" {
		fmt.Println("gpuINTEL fan2 = ", fanI2)
	}

	// AMD
	tempGPUa1, tempGPUa2 := getTemp(body, "gpuINTEL")
	if tempGPUa1 != "" {
		fmt.Println("tempGPUa1 = ", tempGPUa1)
	}
	if tempGPUa2 != "" {
		fmt.Println("tempGPUa2 HotSpot = ", tempGPUa2)
	}
	fanA1, fanA2 := getTemp(body, "fanINTEL")
	if fanA1 != "" {
		fmt.Println("gpu fanA1 = ", fanA1)
	}
	if fanA2 != "" {
		fmt.Println("gpu fanA2 = ", fanA2)
	}

	// NVIDIA
	tempGPUnv1, tempGPUnv2 := getTemp(body, "gpuNVidia")
	if tempGPUnv1 != "" {
		fmt.Println("t gpuNVidia = ", tempGPUnv1)
	}
	if tempGPUnv2 != "" {
		fmt.Println("t gpuNVidia HotSpot = ", tempGPUnv2)
	}
	fanNV1, fanNV2 := getTemp(body, "fanNVidia")
	if fanNV1 != "" {
		fmt.Println("gpuNVidia fanNV1 = ", fanNV1)
	}
	if fanNV2 != "" {
		fmt.Println("gpuNVidia fanNV2 = ", fanNV2)
	}

	tempSSD1, tempSSD2 := getTemp(body, "disk")
	if tempSSD1 != "" {
		fmt.Println("t1 disk = ", tempSSD1)
	}
	if tempSSD2 != "" {
		fmt.Println("t2 disk = ", tempSSD2)
	}

	var a string
	fmt.Println("ожидание ввода")
	fmt.Scan(&a)
}

func getTemp(body []byte, xpu string) (value1, value2 string) {

	var root Node
	var text1, text2, text3, text4 string
	err := json.Unmarshal(body, &root)
	if err != nil {
		panic(err)
	}

	value1, value2 = "", ""

	if xpu == "cpu" {
		text1 = "images_icon/cpu.png"
		text4 = "Temperatures"
		text2 = "Core (Tctl/Tdie)"
		text3 = "CPU Package"
	} else if xpu == "gpuNVidia" {
		text1 = "images_icon/nvidia.png"
		text4 = "Temperatures"
		text2 = "GPU Core"
		text3 = "GPU Hot Spot"
	} else if xpu == "fanNVidia" {
		text1 = "images_icon/nvidia.png"
		text4 = "Fans"
		text2 = "GPU Fan 1"
		text3 = "GPU Fan 2"
	} else if xpu == "gpuAMD" {
		text1 = "images_icon/amd.png"
		text4 = "Temperatures"
		text2 = "GPU Core"
		text3 = "GPU Hot Spot"
	} else if xpu == "fanAMD" {
		text1 = "images_icon/amd.png"
		text4 = "Fans"
		text2 = "GPU Fan 1"
		text3 = "GPU Fan 2"
	} else if xpu == "gpuINTEL" {
		text1 = "images_icon/intel.png"
		text4 = "Temperatures"
		text2 = "GPU Core"
		text3 = "GPU Hot Spot"
	} else if xpu == "fanINTEL" {
		text1 = "images_icon/intel.png"
		text4 = "Fans"
		text2 = "GPU Fan 1"
		text3 = "GPU Fan 2"
	} else if xpu == "disk" {
		text1 = "images_icon/hdd.png"
		text4 = "Temperatures"
		text2 = "Temperature"
		text3 = "Temperature 2"
	}

	// Перебор первого уровня
	for _, child := range root.Children {
		// Перебор второго уровня
		for _, subChild := range child.Children {
			// fmt.Println("child.Children - ", subChild.ImageURL)
			if subChild.ImageURL == text1 {
				// Перебор третьего уровня
				for _, subSubChild := range subChild.Children {
					if subSubChild.Text == text4 {
						// Перебор четвертого уровня
						for _, subSubSubChild := range subSubChild.Children {
							if subSubSubChild.Text == text2 {
								if value1 == "" {
									value1 = subSubSubChild.Value
								} else {
									value1 += ", " + subSubSubChild.Value
								}
							}
							if subSubSubChild.Text == text3 {
								if value2 == "" {
									value2 = subSubSubChild.Value
								} else {
									value2 += ", " + subSubSubChild.Value
								}
							}
						}
					}
				}
			}
		}
	}
	return value1, value2
}
