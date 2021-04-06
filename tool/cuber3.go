package main

import (
	"fmt"
	"os"
	"strconv"
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
	sstart:=os.Args[1]
	send:=os.Args[2]
	start,_:=strconv.Atoi(sstart)
	end,_:=strconv.Atoi(send)

	for cubeno:=start;cubeno<=end;cubeno++ {
		var c core.Cube
		c.Cubeno=cubeno
		c.Read()

		if c.FileSize()<5000.0 {
			echo (cubeno)
			echo (c.FileSize())
			echo (c.FilePath())
			c.DownloadingRpc3()
		} else if c.CHash=="" {
			echo (cubeno)
			echo (c.FilePath())
			c.DownloadingRpc3()
		}
	}

}
