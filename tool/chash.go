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
}

func main() {
	cubecheck();
}

func cubecheck() {
	var c core.Cube
	c.Cubeno=6950
	c.Read()

	hstr:=c.HashString()
	result:=core.CallHash(hstr,4)

	chstr:=result+strconv.Itoa(c.Nonce)
	cresult:=core.CallHash(chstr,2)

	echo(hstr)
	echo(result)
	echo(cresult)
	
	echo(c.CHash)
	echo(c.Nonce)
}




