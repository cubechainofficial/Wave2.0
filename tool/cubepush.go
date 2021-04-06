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

	for i:=50318;i<=ch-1;i++ {
	//i:=49900
		var c core.Cube
		c.Cubeno=i
		c.Read()
		t:=c.BroadcastPush()
		echo (i)
		echo (t)
	}
}