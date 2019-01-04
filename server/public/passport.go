package main

import (
	"github.com/KylinHe/aliensboot-core"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/passport"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		passport.Module,
	)
}
