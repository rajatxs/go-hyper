package util

import (
	"github.com/rajatxs/go-hyper/log"
)

func Attempt(stmt error) {
	if stmt != nil {
		log.Fatal(stmt)
	}
}
