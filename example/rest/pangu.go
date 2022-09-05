package rest

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(newServer)
}
