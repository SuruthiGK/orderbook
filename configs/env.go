package configs

import (
	"time"
)

//Config The important stuff
type env struct {
	OrderBookRefreshTime          time.Duration
	DSRefreshTime          time.Duration
	OBSRefreshTime          time.Duration
}

//Env an instance to manage environment
var Env *env

func init() {
	Env = &env{
		OrderBookRefreshTime:     5 * time.Second,
		DSRefreshTime:     5 * time.Second,
		OBSRefreshTime:     5 * time.Second,
	}
}
