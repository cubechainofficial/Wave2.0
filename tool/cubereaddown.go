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
	echo(ch)
	for cubeno:=124000;cubeno<=ch-1;cubeno++ {
		var c core.Cube
		c.Cubeno=cubeno
		c.Read()

		if c.CHash=="" {
			echo (c.FilePath())
			c.DownloadingRpc2()
		} else if c.FileSize()<5000.0 {
			echo (cubeno)
			echo (c.FileSize())
			echo (c.FilePath())
		}
	}
}










