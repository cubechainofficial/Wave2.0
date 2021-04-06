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
	core.MakeStat(filepath1)

	file0:=filepath1+"/special/"+"Statistic.cbs"
	//file1:=filepath1+"/special/"+"Statistic.cbs"
	file2:=filepath1+"/special/"+"Statistic2.cbs"
	file3:=filepath1+"/special/"+"Statistic3.cbs"
	file4:=filepath1+"/special/"+"Statistic4.cbs"
	file5:=filepath1+"/special/"+"Statistic5.cbs"


	core.FileCopy(file0,file2)
	core.FileCopy(file0,file3)
	core.FileCopy(file0,file4)
	core.FileCopy(file0,file5)


	filepath1="/data/bdata"
	file11:=filepath1+"/special/"+"Statistic.cbs"
	file12:=filepath1+"/special/"+"Statistic2.cbs"
	file13:=filepath1+"/special/"+"Statistic3.cbs"
	file14:=filepath1+"/special/"+"Statistic4.cbs"
	file15:=filepath1+"/special/"+"Statistic5.cbs"
	
	core.FileCopy(file0,file11)
	core.FileCopy(file0,file12)
	core.FileCopy(file0,file13)
	core.FileCopy(file0,file14)
	core.FileCopy(file0,file15)


	//addr:="CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye"
	//core.MakeStatAddr(filepath1,addr) 
}


