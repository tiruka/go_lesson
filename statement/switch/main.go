package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("Mac")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default!")
	}

	// こちらのパターンでは、変数osはswitch内のみ参照可能
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default!")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
