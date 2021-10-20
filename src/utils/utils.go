package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Join() string {
	return "join func"
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Rand() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}