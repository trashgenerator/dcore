package server

import (
	"fmt"
	"strconv"
)

const (
	RegPacket1013ID   = 1013
	DeathPacket777ID  = 777
	ConfirmPacket88ID = 88
)

// TODO : расширить адресом поинта. Актуально когда будет более одного Поинта в системе
func BuildRequest1013(key string) []byte {
	return []byte(fmt.Sprintf("1013\t%s\n", key))
}

// TODO : расширить адресом поинта. Актуально когда будет более одного Поинта в системе
func BuildGoodResponse1013(status bool, key string, address string) []byte {
	return []byte(fmt.Sprintf("%d\t%s\t%s\t%s\n", RegPacket1013ID, strconv.FormatBool(status), key, address))
}

func BuildBadResponse1013(status bool) []byte {
	return []byte(fmt.Sprintf("%d\t%s\n", RegPacket1013ID, strconv.FormatBool(status)))
}

func BuildCommand777(status bool) []byte {
	return []byte(fmt.Sprintf("%d\t%s\n", DeathPacket777ID, strconv.FormatBool(status)))
}

func BuildRequest88(key string, addr string) []byte {
	return []byte(fmt.Sprintf("%d\t%s\t%s\n", ConfirmPacket88ID, key, addr))
}
