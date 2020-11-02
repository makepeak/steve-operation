package basic

import (
	"github.com/makepeak/steve-operation//basic/config"
	"github.com/makepeak/steve-operation/basic/db"
	"github.com/makepeak/steve-operation/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
