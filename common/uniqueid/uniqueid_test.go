package uniqueid_test

import (
	"fmt"
	"luck_draw-server/common/uniqueid"
	"testing"
)

func Test_generateUUID(t *testing.T) {
	fmt.Println(uniqueid.GenerateUUID())
	fmt.Println(uniqueid.GenerateUUID())
	fmt.Println(uniqueid.GenerateUUID())
}
