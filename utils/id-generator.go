package utils

import (
	"strings"
	"math/rand"
	"log"
	"strconv"
)

func GenerateId(prefix string) string {
	i := rand.Intn(9)

	log.Println(i)
	s := prefix + "-" + strconv.Itoa(i)

	return strings.ToUpper(s)
}


