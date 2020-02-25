package console

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-app/app/console/command"
	"os"
	"sync"
)

var (
	engine *gin.Engine
	once sync.Once
)

var Commends map[string]map[string]string

func Boot(egn *gin.Engine) {
	once.Do(func() {
		engine = egn
	})
	run()
}

func run() {
	args := os.Args[1:]
	logo()
	commands()
	for c ,i := range command.Commands(){
		Commends[c] = i
	}
	menu(args)
}