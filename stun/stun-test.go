package stun

import (
    "github.com/ccding/go-stun/stun"
    "log"
)

func TestStun() {
    nat, host, err := stun.NewClient().Discover()
    log.Println(nat, host, err)
}
