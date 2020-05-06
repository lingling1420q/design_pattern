package main

import "fmt"

func main() {
	dataMap := make(map[string]string, 10)

	dataMap["hunan"] = "changsha"
	dataMap["hubei"] = "wuhan"
	dataMap["guangdong"] = "guangzhou"
	dataMap["guangxi"] = "guilin"
	dataMap["anhui"] = "hefei"

	// 遍历输出省和省会城市
	for province, city := range dataMap {
		fmt.Printf("%s 的省会是 %s \n", province, city)
	}

	// 获取湖南省的省会
	hunanCity := dataMap["hunan"]
	beijingCity, ok := dataMap["heibei"]

	fmt.Printf("hunan 省的省会城市是 : %s \n",hunanCity)

	if ok {
		fmt.Printf("hebei 省的省会城市是 : %s \n",beijingCity)
	} else {
		fmt.Printf("hebei 省的省会城市不存在 \n")
	}

	delete(dataMap, "hunan")

	hunanCity = dataMap["hunan"]
	fmt.Printf("hunan 省的省会城市是 : %s \n",hunanCity)

}
