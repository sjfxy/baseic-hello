package main

import (
	"fmt"
	"sin.com/demo/npcs"
)

//这里使用了 go.mod 进行项目的包的管理和依赖关系
// 不再一定按照对应的 目录记录卡

//所以的都是进行按照 go.mod vendor go mod init go mod download go mod why ttib
// 首先在项目目录
func main() {
	mob := npcs.NonPlayerCahracter{Name: "Aftitrf"}
	fmt.Println(mob)

	hortense := npcs.NonPlayerCahracter{Name: "Hortense",
		Loc: npcs.Location{X: 10.0, Y: 10.0, Z: 10.0},
	}
	fmt.Println(hortense)

	fmt.Printf("Aftrrie is %f units from Hortense.\n", mob.DistanceTo(hortense))

}
