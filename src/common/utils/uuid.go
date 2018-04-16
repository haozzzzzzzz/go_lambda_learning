package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

func UUID() string {

	value, _ := uuid.NewV4()

	return strings.ToLower(
		strings.Replace(value.String(), "-", "", -1),
	)
}

func UniqueID() string {
	strUUID := UUID()
	return fmt.Sprintf("%x", md5.Sum([]byte(strUUID)))
}

func VerifyCode6() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}
