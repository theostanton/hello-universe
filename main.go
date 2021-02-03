package main

import (
	"./pagespeed"
	"fmt"
)

func main() {

	var performance string
	var err error

	performance, err = pagespeed.TestSpeedOfAUrl("https://lemonde.fr")

	if err != nil {
		fmt.Printf("Something went wrong err=%v\n", err)
	} else {
		fmt.Printf("Performance=%s\n", performance)
	}

}
