package basic

import (
	"github.com/makepeak/steve-operation/user-service/basic/config"
	"github.com/makepeak/steve-operation/user-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
