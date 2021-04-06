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
	ch:=core.CubeHeight()
	ch2:=core.GetCubeHeight3()

	echo(ch)
	echo(ch2)
	for cubeno:=124867;cubeno<=ch-1;cubeno++ {
		var c core.Cube
		c.Cubeno=cubeno
		c.Read()

		if c.CHash=="" {
			echo (cubeno)
			echo (c.FilePath())
		}
	}

}
