package utils

import (
	"github.com/DanPlayer/randomname"
	_ "github.com/DanPlayer/randomname"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}

func GenerateRandomNumberStr() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(900000) + 100000)
}

func RandNickName() string {
	return randomname.GenerateName()
}
