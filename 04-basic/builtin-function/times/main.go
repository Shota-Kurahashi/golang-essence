package times

import (
	"fmt"
	"log"
	"time"
)

func Main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	d, err := time.ParseDuration("3s")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

	d, err = time.ParseDuration("4m")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

	d, err = time.ParseDuration("5h")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

}
