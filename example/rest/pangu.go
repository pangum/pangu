package rest_test

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Musts(newServer)
}
