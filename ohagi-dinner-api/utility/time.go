package utility

import (
	"log"
	"time"
)

func NowInJST() time.Time {
	return time.Now().In(LocationInJST())
}

func LocationInJST() *time.Location {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	return jst
}
