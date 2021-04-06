package main

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/gob"
	"os"
	 "../config"
	 "../core"
)

var echo=fmt.Println
var Configure config.Configuration
var	addr="CLQUKEdCeWmPzAmyJdHzo9cTBrq2JCBbPC"
var	filepath1="."
var	filepath2="/data/bdata"
var filepathSeparator="/"

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
	addr:="CUqZZeBnjyKuAyWmkjooRT13zFpDL1fiR7"
	p:=GetStatisticAddrHere(addr)
	echo(p)
}

func GetStatisticAddrHere(addr string) string {
	var aStatistic core.TxStatistic
	aStatistic.AddrIndex=make(map[string]string)	
	StatisticReadHere(&aStatistic)
	if aStatistic.AddrIndex[addr]=="" {
		aStatistic.AddrIndex[addr]="0.0,0.0,0.0,0,0,F"
	}
	result:=strconv.Itoa(aStatistic.Cubeno)+"||"+aStatistic.AddrIndex[addr]
	return result
}

func StatisticReadHere(aStatistic *core.TxStatistic) {
	path:=filepath1+filepathSeparator+"special"
	if core.DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	if core.DirExist(pathfile) {
		err:=core.FileRead(pathfile,aStatistic)
		core.Err(err,0)	
	}
}