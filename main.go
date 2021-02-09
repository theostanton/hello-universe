package main

import (
	"fmt"
	"log"
	"theo.dev/hello-universe/database"
	"theo.dev/hello-universe/model"
	"theo.dev/hello-universe/model/PerformanceResult"
	"theo.dev/hello-universe/pagespeed"
)

func main() {

	var performance PerformanceResult.PerformanceResult
	var err error

	performance, err = pagespeed.TestSpeedOfAUrl("https://lemonde.fr")

	if err != nil {
		log.Fatal(fmt.Sprintf("Something went wrong err=%v\n", err))
	}

	testResult := model.TestResult{
		Performance: performance,
	}

	id, err := database.Insert(testResult)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to insert err=%v\n", err))
	}

	fmt.Printf("Success inserted id=%d\n", id)

}
