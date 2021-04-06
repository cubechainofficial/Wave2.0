package main

import (
	"fmt"
	"encoding/gob"
	 "../config"
	 "../core"
)

var echo=fmt.Println
var Configure config.Configuration

func init() {
	Configure=config.LoadConfiguration("../config/cubechain.conf")
	core.Configure=Configure
	core.CubenoSet()
	echo (core.CubeSetNum)
	
	gob.Register(&core.TxData{})
	gob.Register(&core.TxBST{})
	gob.Register(map[string]string{})
}


func main() {
	//rno:=[]int{41756, 43086}
	rno:=[]int{44163}
	 

	for _,cubeno:=range rno {
		var c core.Cube
		c.Cubeno=cubeno
		echo (c.FilePath())
		c.RepairInfo(cubeno)
	}
}






