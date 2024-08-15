package utils

import (
	"github.com/DanPlayer/randomname"
	_ "github.com/DanPlayer/randomname"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}

func GenerateRandomNumberStr() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Intn(900000) + 100000)
}

func RandNickName() string {
	return randomname.GenerateName()
}
