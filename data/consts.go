package data

import (
	"fmt"

	"github.com/rrsoft/guestbook/utils"
)

var AppStting *utils.Setting

type CommandType int

const (
	Text            CommandType = 1
	StoredProcedure             = 4
	TableDirect                 = 512
)

func init() {
	AppStting = utils.GetSetting()
	fmt.Println(AppStting)
}
