package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"io/ioutil"
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

	ssnum:=os.Args[1]
	snum,_:=strconv.Atoi(ssnum)
	srange:=10000
	start:=snum*srange+1
	end:=(snum+1)*srange
	
    file, err := os.Open("miner"+strconv.Itoa(start)+".scn")
    if err != nil {
        log.Fatal(err)
    }
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }
	strdata:=core.ByteToStr(data)

	if strings.Index(strdata,"\n")<0 {
		echo("line invalid.")
		return
	}
	cubeminer:=strings.Split(strdata, "\n")

	for cubeno:=start;cubeno<=end;cubeno++ {
		if cubeno>127746 {
			return
		}
		
		if cubeno%100==0 {
			echo("****"+strconv.Itoa(cubeno))
		}

		var c core.Cube
		c.Cubeno=cubeno
		c.Read()
		cMiner:=cubeminer[cubeno-start]
		miner:=cMiner[:34]
		if c.Mine.MineAddr!=miner {
			echo (strconv.Itoa(cubeno)+ " : "+c.Mine.MineAddr+" ===> "+miner)
			echo (c.FilePath())
			c.RepairScan(cubeno)
		}
	}

}
