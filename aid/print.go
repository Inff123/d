package aid

import (
	"encoding/json"
	"fmt"
	"time"
)

func Print(v ...interface{}) {
	if Config.Output.Level == "prod" {
		return
	}

	fmt.Println(v...)
}

func PrintJSON(v interface{}) {
	if Config.Output.Level == "prod" || Config.Output.Level == "time" {
		return
	}

	json1, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json1))
}

func PrintTime(label string, functions ...func()) {
	current := time.Now()

	for _, f := range functions {
		f()
	}

	if Config.Output.Level == "prod" {
		return
	}

	fmt.Println(label + ":", time.Since(current))
}