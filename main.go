package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"encoding/gob"
	 "./config"
	 "./core"
)

var echo=fmt.Println
var Configure config.Configuration
var mstr="miningtesting!!..."
var	addr="CLQUKEdCeWmPzAmyJdHzo9cTBrq2JCBbPC"
//var	addr="CQUMoMjngW9tBZf7vkVeuGJ4iyTxcnWXDk"
//var	addr="CQqvmQrQb4wt9CzEspYSWhTETt5nkRCo7Q"


func init() {
	Configure=config.LoadConfiguration("./config/cubechain.conf")
	core.Configure=Configure
	core.CubenoSet()
	echo (core.CubeSetNum)
	
	if core.GenFile=="" {
		path:="./config/genfile"
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
	//cubemining3()
	//quickmining()
	//blockscan()
	//blockscan2()
	//api_CubeBalance()
	//txCheck()
	//Genfile()

	//api_GetIndex()


	//cubemining3()
	
	
	go quickmining2()
	core.ServerRun()


	

	/*
	cube.GetPrevHash()


	hashnip:=core.NodeSend("cubehash","0&cubeno=13200")
	echo ("hash")
	echo (hashnip)

	echo ("")
	api_CubeBalance()
	api_Balance()
	echo ("")
	api_GetIndex()
	echo ("")
	api_GetCount()
	echo ("")
	echo ("")
	//api_GetList()
	//api_GetListDetail()
	//core.ServerRun()
	//api_GetBlock()
	api_GetBlockTx()
	echo ("")
	echo ("")
	blockscan3(1)
	*/
	//cubeIndexing(1)
	//cubeStatistic(2)
	//core.AllIndexing(9)
	//core.CubeStatistic(1)
	//core.AllStatistic(9)
	//core.GetIndexing()
	//core.GetStatistic()

}

func quickmining() {
	tickChan:=time.Tick(time.Duration(Configure.Blocktime)*time.Second)
	echo("Cubechain start!")
	cubemining2()
	for {
		select {
		case <-tickChan:
			cubemining2()
		}
	}
	echo("Cubechain end!")
}

func quickmining2() {
	tickChan:=time.Tick(2*time.Second)
	tickChan2:=time.Tick(time.Duration(Configure.Blocktime+1)*time.Second)
	echo("Cubechain start!")
	
	startTime:=time.Now()
    exeTime:=startTime.Add(-time.Duration(Configure.Blocktime)*time.Second)
    exeTime=exeTime.Add(-5*time.Second)
    cubeDuration:=time.Duration(Configure.Blocktime) * time.Second
	//exeDuration:=time.Duration(1*time.Second)
	
	for {
		select {
		case <-tickChan:
			if  time.Since(exeTime)>=cubeDuration {
				exeTime=time.Now()
				cubemining2()
			}
		case <-tickChan2:
			go core.AllIndexing(0)
			go core.AllStatistic(0)
		}
	}
	echo("Cubechain end!")
}

func Genfile() {
	g:=core.GenesisTx(1)
	echo (g)
}


func api_CubeBalance() float64 {
	result:=core.CubeBalance(addr,2)
	echo (result)
	return result
}

func api_Balance() float64 {
	result:=core.GetBalance(addr)
	echo (result)
	return result
}


func api_GetIndex() {
	result:=core.GetIndexBlock(addr)
	echo (result)
}

func api_GetCount() int {
	result:=core.GetTransactionCount(addr)
	echo (result)
	return result
}

func api_GetList() string {
	result:=core.GetTransactionList(addr)
	echo (result)
	return result
}

func api_GetListDetail() string {
	result:=core.GetTxListDetail(addr)
	echo (result)
	return result
}

func api_GetDetail(hash string) string {
	result,r:=core.GetTransactionDetail(hash)
	echo (result)
	echo (r)
	return result.String()
}

func api_GetBlock() string {
	result:=core.GetBlock(5,1)
	echo (result)
	return result
}

func api_GetBlockTx() string {
	result:=core.GetBlockTx(5,1)
	echo (result)
	return result
}


func cubeIndexing(cubeno int) {
	ci:=core.CubeIndexing(cubeno)
	echo (ci)
}

func cubeStatistic(cubeno int) {
	ci:=core.CubeStatistic(cubeno)
	echo (ci)
}


func working() {
	var block core.Block
	block.Cubeno=651
	block.Blockno=2
	ph:=block.GetPrevHash()
	echo (ph)
}

func cubedown() {
	//var c core.Cube
	core.CubeDownload(600)

}

func cubemining() { 
	var c core.Cube
	ch:=core.CubeHeight()+1
	echo (ch)
	c.Input(ch)
}

func cubemining2() { 
	var c core.Cube
	ch:=core.CubeHeight()
	ch2:=core.GetCubeHeight3()
	ch3,_:=strconv.Atoi(ch2)

	if ch3>ch {
		for i:=ch;i<=ch3;i++ {
			c.Cubeno=i
			c.CHash=""
			c.Download()
		}
	} else {
		if ch>3 {
			c.Cubeno=ch-2
			c.CHash=""
			c.Download()
			
			c.Cubeno=ch-1
			c.CHash=""
			c.Download()
		}

		echo (ch)
		c.InputChanel(ch)
	}
}

func cubemining3() { 
	startTime:= time.Now()
	cubemining2()
	endTime:=time.Now()
	durTime:=endTime.Sub(startTime)	

	echo (durTime)
}

func blockscan() {
	b:=core.BlockScan(11,23)
	b.Print()
	echo(b.Mine)
}

func blockscan2() {
	b:=core.BlockScan(5,6)
	c:=core.BlockTxData(b.Data)
	d:=core.TreeDeserialize(b.Data)

	d.TreePrint2()
	echo(b.Mine)
	echo(c)
}

func blockscan3(cubeno int) {
	for i:=1;i<28;i++ {
		b:=core.BlockScan(cubeno,i)
		c:=core.BlockTxData(b.Data)
		d:=core.TreeDeserialize(b.Data)
		echo(i)
		d.TreePrint2()
		echo(b.Mine)
		echo(c)
	}
}


func txCheck() {
	t,_:=core.TxpoolToTr("",0)
	t.TreePrint2()
	Bd:=core.GetBytes(t)
	echo (Bd)
	tr:=core.TreeDeserialize(Bd)
	tr.TreePrint2()
	c:=core.BlockTxData(Bd)
	echo(c)
	
}


func blockmining1() {
	var b core.Block
	b.Input(1,10)
}
func blockmining2() { 
	var b core.Block
	c:=core.CubeHeight()+1
	echo (c)
	b.Input(c,10)
}





func testcon() {
	r:=core.NodeSend2("pool_result","0&cubeno=1&blockno=3&hashstr=293u89u4832u48eaujfhugjnxcjnujdusifu")
	echo (r)
}

func pohcheck() {
	ph:=core.PohSet(1)
	echo (ph.Cubeno)
}

 
func checking() {
	r:=core.TxPool(7,3)
	bst,_:=core.TxpoolToTr(r,0)
	echo (r)
	bst.TreePrint2()

	b,_:=core.TxBlockData(7,3)
	echo (b)
}


