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
	cubeupload(7013);
}


func cubeupload(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	
	c.FileBroadcast();
}




