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
	txpoolcheck(135195,1)
}



func check() {
	gBlock:=core.CBlockRead(23546,2)
	gBlock.PrintHead()
	core.BlockRead(23546,2,&gBlock)
	gBlock.PrintHead()
}



func getblocktx(cubeno int,blockno int) {
	c:=core.GetBlockTx(cubeno,blockno)
	echo (c)
}


func checkcubesize(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	fs:=c.FileSize()
	if fs>1000.0 {
		echo ("1000")
	}
	echo (c.FileSize())



}

func checkStatAddr() {
	addr:="CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye"
	aStatistic:=make(map[string]core.StatisticData)
	g:=core.GetStaticData(3499,aStatistic)

	echo (g)
	echo (g[addr])
}

func StatisticRead(aStatistic *core.TxStatistic) {
	pathfile:="/home/www/cubepos/pos/tool/special/Statistic.cbs"
	if core.DirExist(pathfile) {
		err:=core.FileRead(pathfile,aStatistic)
		core.Err(err,0)	
	}
}

func GetStatisticAddr(addr string) string {
	var aStatistic core.TxStatistic
	aStatistic.AddrIndex=make(map[string]string)	
	StatisticRead(&aStatistic)
	if aStatistic.AddrIndex[addr]=="" {
		aStatistic.AddrIndex[addr]="0.0,0.0,0.0,0,0,F"
	}
	result:=strconv.Itoa(aStatistic.Cubeno)+"||"+aStatistic.AddrIndex[addr]
	return result
}

func repairchk(cubeno int) { 
	cubeinfo:=core.NodeCube("cuberepair","0&cubeno="+strconv.Itoa(cubeno))
	cubef:=strings.Split(cubeinfo, "|")
	echo(cubeinfo)
	echo(cubef[13])
	echo(len(cubef[13]))
}

func filesizeCheck() {
	var checkCube=[]int{51,57,64,129,143,149,155,160,167,174,180,186,197,203,209,309,314,320,332,338,344,376,382,388,394,406,411,417,640,646,652,658,674,679,691,726,732,814,850,857,867,874,882,894,912,956,984,1010,1020,1027,1050,1056,1062,1068,1074,1080,1085,1098,1104,1110,1116,1122,1128,1133,1140,1147,1152,1157,1163,1169,1175,1182,1188,1207,1214,1229,1244,1251,1257,1280,1286,1380,1402,1409,1416,1423,1434,1468,1501,1514,1752,1992,2004,2010,2016,2063,2068,2074,2080,2085,2092,2098,2103,2115,2177,2183,2378,2419}
	for _,cubeno:=range checkCube {
		echo (filesize(cubeno))
	}
}

func filesize(cubeno int) int64 { 
	path:=core.CubePath(cubeno)
	return core.FileSize(path)
}

func repair(cubeno int) { 
	var c core.Cube
	c.RepairCube(cubeno)
}

func newhash() { 
	echo (core.CallHash(addr,1))
	echo (core.CallHash(addr,5))
	echo (core.PatternHash(addr,3))
	echo (core.PatternHash(addr,4))
	echo (core.PatternHash(addr,5))
	echo (core.PatternHash(addr,6))
	echo (core.PatternHash(addr,7))
	echo (core.PatternHash(addr,8))
}


func cubedownload(cubeno int) { 
	core.CubeDownloadFile(cubeno)
}

func cubeconfirm(cubeno int) bool { 
	result:=core.CheckConfirm(cubeno)
	echo (result)
	return result
}

func checkmining() { 
	var c core.Cube
	ch:=core.CubeHeight()
	ch2:=core.GetCubeHeight3()
	ch3,_:=strconv.Atoi(ch2)

	if ch3>ch {
		for i:=ch;i<=ch3;i++ {
			c.Cubeno=i
			c.CHash=""
			s:=c.Filecast()
			if s=="failure" {
				c.DownloadRpc()
			}
		}
	}
}

func cubeHeader(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	c.Read()
	c.PrintHead()
}

func cubecast(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	c.CHash=""
	echo ("start")
	s:=c.Filecast()
	echo (s)
}

func cubedown(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	c.CHash=""
	c.Download()
}

func makecbs() {
	s1:=core.TAllIndexing(0)
	echo (s1)
	s2:=core.TAllStatistic(0)
	echo (s2)

}

func checkStatic() {
	s:=core.GetStatisticRank("balance","txcnt",100)
	echo (s)
}

func blockread(cubeno int,blockno int) {
	var b core.Block
	core.BlockRead(cubeno,blockno,&b)
	echo (b)
}

func checking() {
	echo(os.Getuid())
	echo(os.Getgid())
	err:= os.Chown("test.txt", 0, 48)
	if err != nil {
		echo(err)
	}
}

func checkingdir() {
	path:="./cc3"
	core.MakePath(path)
}

func cubecheck(cubeno int) {
	var c core.Cube
	c.Cubeno=cubeno
	c.Read()

	c.Print()
	c.FileInfo()
}

func blockcheck(cubeno int,blockno int) {
	var b core.Block
	b.Cubeno=cubeno
	b.Blockno=blockno
	b.Read()

	b.Print()
	b.FileInfo()
}

func blockprvhash(cubeno int,blockno int) {
	var b core.Block
	b.Cubeno=cubeno
	b.Blockno=blockno
	g:=b.GetPrevHash()
	echo(g)
}



func blockhashcheck(cubeno int,blockno int) {
	var b core.Block
	b.Cubeno=cubeno
	b.Blockno=blockno
	b.Read()

	hash:=b.HashString()
	hashs:=core.CallHash(hash,1)
	hashg:=core.CallHash(hashs+strconv.Itoa(b.Nonce),1)
	echo(hash)
	echo(hashs)
	echo(b.Nonce)
	echo(hashg)
}

func txpoolcheck(cubeno int,blockno int) string {
	result:=core.NodeSend("txpool","0&cubeno="+strconv.Itoa(cubeno)+"&blockno="+strconv.Itoa(blockno))
	echo("cubeno="+strconv.Itoa(cubeno)+"&blockno="+strconv.Itoa(blockno)+"\n",result)
	return result
}
