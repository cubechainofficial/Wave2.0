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
	//cubeno:=22193
	rno:=[]int{23316,23318,23320,24042,24044,24354}

	for _,cubeno:=range rno {
		var c core.Cube
		c.Cubeno=cubeno
		echo (c.FilePath())
		c.DownloadingRpc2()
	}
}