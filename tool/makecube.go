package main

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/gob"
	 "../config"
	 "../core"
)

var echo=fmt.Println
var Configure config.Configuration
var	addr="CLQUKEdCeWmPzAmyJdHzo9cTBrq2JCBbPC"


func init() {
	Configure=config.LoadConfiguration("../config/cubechain.conf")
	core.Configure=Configure
	core.CubenoSet()
	echo (core.CubeSetNum)
	
	if core.GenFile=="" {
		path:="../config/genfile"
		core.GenFile=core.FileReadString(path)
		line:=strings.Split(core.GenFile,"\r\n")
		for _,v:=range line {
			result:=strings.Split(v, "|")
			genno,ok:=strconv.Atoi(result[0])
			if ok==nil {
				core.GenBlock[genno-1]+=v+"\r\n"
			}
		}
	}

	gob.Register(&core.TxData{})
	gob.Register(&core.TxBST{})
	gob.Register(map[string]string{})
	//echo(core.GenBlock)
}


func main() {

	cubemake(24544)

}


func cubemake(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	c.InputChanel(cubeno)
}
