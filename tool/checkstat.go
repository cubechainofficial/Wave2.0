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
var	addr="CW4dPvQ24RGgMk8NKEj6F1UkYka9sxYTdM"
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
	addr:="CW4dPvQ24RGgMk8NKEj6F1UkYka9sxYTdM"
	aStatistic:=make(map[string]core.StatisticData)
	aStatistic[addr]=core.StatisticData{3851,addr,111355.5024,111356.5025,1.0001,111357.5026,304,1,"T"}
	echo (aStatistic[addr])

	i:=3867
	aStatistic=core.GetStaticData(i,aStatistic)
	echo (aStatistic[addr])
	
}





