package test1

import (
	"math/rand"
	"time"
)

func DelSliceARandomElements(s []int) []int {
	l := len(s)
	rand.Seed(time.Now().Unix())
	index := rand.Intn(l)

	res := append(s[:index], s[index+1:]...)

	return res
}
