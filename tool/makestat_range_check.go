package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"encoding/gob"
	 "../config"
	 "../core"
)


var echo=fmt.Println
var Configure config.Configuration
var	addr="CLQUKEdCeWmPzAmyJdHzo9cTBrq2JCBbPC"
var	filepath1="."
var	filepath2="/data/bdata"

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
	sstart:=os.Args[1]
	send:=os.Args[2]
	start,_:=strconv.Atoi(sstart)
	end,_:=strconv.Atoi(send)

	core.MakeStatRange(filepath1,start,end)

	file0:=filepath1+"/special/"+"Statistic.cbs"
	file2:=filepath1+"/special/"+"Statistic_"+strconv.Itoa(start)+"-"+strconv.Itoa(end)+".cbs"

	core.FileCopy(file0,file2)
}


