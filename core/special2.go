package core

import (
	"strconv"
	"strings"
	"os"
	"math"
)

var AddrStatistics	map[string]StatisticData

type TxStatistics struct {
	Cubeno		int	
	AddrIndex	map[string]StatisticData
}

type StatisticData struct {
	Cubeno		int	
	Addr		string
	Balance		float64
	RBalance	float64
	SBalance	float64
	TBalance	float64
	Txcnt		int
	Tkcnt		int
	Pos			string
}

func MakeStat(filepath string) {
	var bStatistic TxStatistic


	aStatistic:=make(map[string]StatisticData)
	bStatistic.AddrIndex=make(map[string]string)	
	
	ch:=CubeHeight()
	echo(ch)
	path:=filepath+filepathSeparator+"special"

	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	
	for i:=1;i<ch;i++ {
		aStatistic=GetStaticData(i,aStatistic)
	}
	
	for k,v := range aStatistic {
		result:=""
		result=strconv.FormatFloat(v.Balance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.RBalance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.SBalance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.TBalance,'f',-1,64)+","
		result+=strconv.Itoa(v.Txcnt)+","
		result+=strconv.Itoa(v.Tkcnt)+","
		result+=v.Pos
		bStatistic.AddrIndex[k]=result
	}
	
	bStatistic.Cubeno=ch-1
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	err:=FileWrite(pathfile,bStatistic)
	Err(err,0)	
}

func MakeStatRange(filepath string,start,end int) {
	var bStatistic TxStatistic


	aStatistic:=make(map[string]StatisticData)
	bStatistic.AddrIndex=make(map[string]string)	
	
	ch:=CubeHeight()
	echo(strconv.Itoa(start)+"-"+strconv.Itoa(end))
	path:=filepath+filepathSeparator+"special"

	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	if end>ch {
		end=ch
	}

	for i:=start;i<=end;i++ {
		aStatistic=GetStaticData(i,aStatistic)
	}
	
	for k,v := range aStatistic {
		result:=""
		result=strconv.FormatFloat(v.Balance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.RBalance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.SBalance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.TBalance,'f',-1,64)+","
		result+=strconv.Itoa(v.Txcnt)+","
		result+=strconv.Itoa(v.Tkcnt)+","
		result+=v.Pos
		bStatistic.AddrIndex[k]=result
	}
	
	bStatistic.Cubeno=end
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	err:=FileWrite(pathfile,bStatistic)
	Err(err,0)	
}

func MakeStatMerge(filepath,filename string,mergename []string) {
	var bStatistic TxStatistic
	bStatistic.AddrIndex=make(map[string]string)	
	
	path:=filepath+filepathSeparator+"special"

	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}

	rpathfile:=""
	for m,v:=range mergename {
		var aStatistic TxStatistic
		aStatistic.AddrIndex=make(map[string]string)	
		rpathfile=path+filepathSeparator+v
		err:=FileRead(rpathfile,&aStatistic)
		Err(err,0)	
		
		if m==0 {
			bStatistic=aStatistic
			continue
		}
		for k,val:=range aStatistic.AddrIndex {
			if _, ok := bStatistic.AddrIndex[k]; ok==false {
				bStatistic.AddrIndex[k]=val
			} else {					
				bStatistic.AddrIndex[k]=StatisticStrSum(bStatistic.AddrIndex[k],val)
			}
			if aStatistic.Cubeno>bStatistic.Cubeno {
				bStatistic.Cubeno=aStatistic.Cubeno
			}
		}
	}

	pathfile:=path+filepathSeparator+filename
	err:=FileWrite(pathfile,bStatistic)
	Err(err,0)	
}



func MakeStatAddr(filepath,addr string) {
	var bStatistic TxStatistic

	aStatistic:=make(map[string]StatisticData)
	bStatistic.AddrIndex=make(map[string]string)	
	
	ch:=CubeHeight()
	echo(ch)
	echo(addr)
	path:=filepath+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	
	p:=0
	for i:=1;i<ch;i++ {
		aStatistic=GetStaticDataAddr(i,addr,aStatistic)
		p=aStatistic[addr].Cubeno
		if p==i {
			echo(strconv.Itoa(i)+":",aStatistic[addr])
		}
	}
	
	for k,v := range aStatistic {
		result:=""
		result=strconv.FormatFloat(v.Balance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.RBalance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.SBalance,'f',-1,64)+","
		result+=strconv.FormatFloat(v.TBalance,'f',-1,64)+","
		result+=strconv.Itoa(v.Txcnt)+","
		result+=strconv.Itoa(v.Tkcnt)+","
		result+=v.Pos
		bStatistic.AddrIndex[k]=result
	}
	
	bStatistic.Cubeno=ch-1
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	err:=FileWrite(pathfile,bStatistic)
	Err(err,0)	
}

func GetStaticDataAddr(cubeno int,addr string,aStatistic map[string]StatisticData) map[string]StatisticData {
	var iData []TxData
	var sdata1 StatisticData
	var sdata2 StatisticData

	var c Cube 
	c.Cubeno=cubeno
	c.Read()

	for i:=-1;i<27;i++ {
		if i==Configure.Indexing || i==Configure.Statistics || i==Configure.Format || i==Configure.Edit {
		} else {
			if i==-1 {
				tx1,tx2:=GetCubePoh(cubeno)
				iData=append(iData,tx1)
				iData=append(iData,tx2)
			} else {
				iData=BlockTxData(c.Blocks[i].Data)
			}
			for _,v := range iData {
				sdata1=StatisticData{}
				sdata2=StatisticData{}
				if v.Datatype!="NULL" && v.Datatype!="Data" && v.Datatype!="Contract" {
					if v.Datatype=="QUB" {
						if v.From==addr {
							if _, ok := aStatistic[v.From]; ok==false {
								aStatistic[v.From]=StatisticData{}
							}
							sdata1=aStatistic[v.From]
							sdata1.Cubeno=cubeno
							sdata1.SBalance=sdata1.SBalance+v.Amount+v.Fee+v.Tax
							sdata1.SBalance=math.Round(sdata1.SBalance*100000000)/100000000
							sdata1.Txcnt++
						}
						if v.To==addr {
							if _, ok := aStatistic[v.To]; ok==false {
								aStatistic[v.To]=StatisticData{}
							}
							sdata2=aStatistic[v.To]
							sdata2.Cubeno=cubeno
							sdata2.RBalance=sdata2.RBalance+v.Amount+v.Tax
							sdata2.RBalance=math.Round(sdata2.RBalance*100000000)/100000000
							sdata2.Txcnt++
						}
					} else {
						if v.From==addr {
							if _, ok := aStatistic[v.From]; ok==false {
								aStatistic[v.From]=StatisticData{}
							}
							sdata1=aStatistic[v.From]	
							sdata1.Cubeno=cubeno
							sdata1.Tkcnt++
						}
						if v.To==addr {
							if _, ok := aStatistic[v.To]; ok==false {
								aStatistic[v.To]=StatisticData{}
							}
							sdata2=aStatistic[v.To]
							sdata2.Cubeno=cubeno
							sdata2.Tkcnt++
						}
					}
					sdata1.Balance=sdata1.RBalance-sdata1.SBalance
					sdata1.TBalance=sdata1.RBalance+sdata1.SBalance
					sdata1.Balance=math.Round(sdata1.Balance*100000000)/100000000
					sdata1.TBalance=math.Round(sdata1.TBalance*100000000)/100000000
					if sdata1.Balance>=5000.0 {
						sdata1.Pos="T"
					}	
					sdata2.Balance=sdata2.RBalance-sdata2.SBalance
					sdata2.TBalance=sdata2.RBalance+sdata2.SBalance
					sdata2.Balance=math.Round(sdata2.Balance*100000000)/100000000
					sdata2.TBalance=math.Round(sdata2.TBalance*100000000)/100000000

					if sdata2.Balance>=5000.0 {
						sdata2.Pos="T"
					}
					aStatistic[v.From]=sdata1
					aStatistic[v.To]=sdata2
				}
			}
		}
	}
	return aStatistic
}

func GetStaticData(cubeno int,aStatistic map[string]StatisticData) map[string]StatisticData {
	var iData []TxData
	var sdata1 StatisticData
	var sdata2 StatisticData

	var c Cube 
	c.Cubeno=cubeno
	c.Read()

	for i:=-1;i<27;i++ {
		if i==Configure.Indexing || i==Configure.Statistics || i==Configure.Format || i==Configure.Edit {
		} else {
			if i==-1 {
				tx1,tx2:=GetCubePoh(cubeno)
				iData=append(iData,tx1)
				iData=append(iData,tx2)
			} else {
				iData=BlockTxData(c.Blocks[i].Data)
			}
			for _,v := range iData {
				sdata1=StatisticData{}
				sdata2=StatisticData{}
				if v.Datatype!="NULL" && v.Datatype!="Data" && v.Datatype!="Contract" {
					if v.Datatype=="QUB" {
						if len(v.From)==34 && v.From[0:31]!="C"+strings.Repeat("0",30) {
							if _, ok := aStatistic[v.From]; ok==false {
								aStatistic[v.From]=StatisticData{}
							}
							sdata1=aStatistic[v.From]
							sdata1.Cubeno=cubeno
							sdata1.SBalance=sdata1.SBalance+v.Amount+v.Fee+v.Tax
							sdata1.SBalance=math.Round(sdata1.SBalance*100000000)/100000000
							sdata1.Txcnt++
						}
						if len(v.To)==34 && v.To[0:31]!="C"+strings.Repeat("0",30) {
							if _, ok := aStatistic[v.To]; ok==false {
								aStatistic[v.To]=StatisticData{}
							}
							sdata2=aStatistic[v.To]
							sdata2.Cubeno=cubeno
							sdata2.RBalance=sdata2.RBalance+v.Amount+v.Tax
							sdata2.RBalance=math.Round(sdata2.RBalance*100000000)/100000000
							sdata2.Txcnt++
						}
					} else {
						if len(v.From)==34 && v.From[0:31]!="C"+strings.Repeat("0",30) {
							if _, ok := aStatistic[v.From]; ok==false {
								aStatistic[v.From]=StatisticData{}
							}
							sdata1=aStatistic[v.From]	
							sdata1.Cubeno=cubeno
							sdata1.Tkcnt++
						}
						if len(v.To)==34 && v.To[0:31]!="C"+strings.Repeat("0",30) {
							if _, ok := aStatistic[v.To]; ok==false {
								aStatistic[v.To]=StatisticData{}
							}
							sdata2=aStatistic[v.To]
							sdata2.Cubeno=cubeno
							sdata2.Tkcnt++
						}
					}
					sdata1.Balance=sdata1.RBalance-sdata1.SBalance
					sdata1.TBalance=sdata1.RBalance+sdata1.SBalance
					if sdata1.Balance>=5000.0 {
						sdata1.Pos="T"
					}	
					sdata2.Balance=sdata2.RBalance-sdata2.SBalance
					sdata2.TBalance=sdata2.RBalance+sdata2.SBalance
					if sdata2.Balance>=5000.0 {
						sdata2.Pos="T"
					}
					aStatistic[v.From]=sdata1
					aStatistic[v.To]=sdata2
				}
			}
		}
	}
	return aStatistic
}

func TAllIndexing(cubeno int) map[string]string {
	var aIndexing TxIndexing
	aIndexing.AddrIndex=make(map[string]string)	

	if cubeno==0 {
		cubeno=CubeHeight()-1
	}
	if cubeno<=0 {
		cubeno=1
	}

	path:=Configure.Datafolder+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Indexing.cbs"
	rpathfile:=pathfile
	pathfile1:=""
	pathfile2:=""
	if DirExist(pathfile) {
		for i:=SpecialCnt-1;i>0;i-- {
			if i==1 {
				pathfile1=path+filepathSeparator+"Indexing.cbs"
			} else {
				pathfile1=path+filepathSeparator+"Indexing"+strconv.Itoa(i)+".cbs"
			}
			pathfile2=path+filepathSeparator+"Indexing"+strconv.Itoa(i+1)+".cbs"
			if DirExist(pathfile1) {
				if i>3 { rpathfile=pathfile1 }
				FileCopy(pathfile1,pathfile2)
			}
		}
		err:=FileRead(rpathfile,&aIndexing)
		Err(err,0)	
	}
	if aIndexing.Cubeno>=cubeno {
	} else {
		for i:=aIndexing.Cubeno+1;i<cubeno+1;i++ {
			if CheckConfirm(i)==false {
				cubeno=i-1
				break;
			}
			tmpIndexing:=CubeIndexing(i)
			for k,v := range tmpIndexing {
				if aIndexing.AddrIndex[k]=="" {
					aIndexing.AddrIndex[k]=v+","
				} else if v>"1" {
					aIndexing.AddrIndex[k]+=v+","
				}
			}
		}
		if(aIndexing.Cubeno!=cubeno) {
			aIndexing.Cubeno=cubeno
			err:=FileWrite(pathfile,aIndexing)
			if err!=nil && cubeno%10000==0 {
				pathfile2=path+filepathSeparator+"Indexing_"+strconv.Itoa(cubeno)+".cbs"
				FileCopy(pathfile,pathfile2)
			}
		}
	}
	return aIndexing.AddrIndex
}

func TAllStatistic(cubeno int) map[string]string {
	var aStatistic TxStatistic
	aStatistic.AddrIndex=make(map[string]string)	
	vStatistic:=make(map[string]StatisticData)

	if cubeno==0 {
		cubeno=CubeHeight()-1
	}
	if cubeno<=0 {
		cubeno=1
	}

	path:=Configure.Datafolder+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	rpathfile:=pathfile
	pathfile1:=""
	pathfile2:=""
	if DirExist(pathfile) {
		for i:=SpecialCnt-1;i>0;i-- {
			if i==1 {
				pathfile1=path+filepathSeparator+"Statistic.cbs"
			} else {
				pathfile1=path+filepathSeparator+"Statistic"+strconv.Itoa(i)+".cbs"
			}
			pathfile2=path+filepathSeparator+"Statistic"+strconv.Itoa(i+1)+".cbs"
			if DirExist(pathfile1) {
				if i>3 { rpathfile=pathfile1 }
				FileCopy(pathfile1,pathfile2)
			}
		}
		err:=FileRead(rpathfile,&aStatistic)
		Err(err,0)	
	}
	if aStatistic.Cubeno>=cubeno {
	} else {
		for k,v := range aStatistic.AddrIndex {
			if len(k)==34 && k[0:31]!="C"+strings.Repeat("0",30) {
				vStatistic[k]=StatisticStrToData(v)
			}
		}
		for i:=aStatistic.Cubeno+1;i<cubeno+1;i++ {
			if CheckConfirm(i)==false {
				cubeno=i-1
				break;
			}
			vStatistic=GetStaticData(i,vStatistic)
			for k,v := range vStatistic {
				result:=""
				result=strconv.FormatFloat(v.Balance,'f',-1,64)+","
				result+=strconv.FormatFloat(v.RBalance,'f',-1,64)+","
				result+=strconv.FormatFloat(v.SBalance,'f',-1,64)+","
				result+=strconv.FormatFloat(v.TBalance,'f',-1,64)+","
				result+=strconv.Itoa(v.Txcnt)+","
				result+=strconv.Itoa(v.Tkcnt)+","
				result+=v.Pos
				aStatistic.AddrIndex[k]=result
			}			
		}
		if(aStatistic.Cubeno!=cubeno) {
			aStatistic.Cubeno=cubeno
			err:=FileWrite(pathfile,aStatistic)
			if err!=nil {
				pathfile2=path+filepathSeparator+"Statistic_"+strconv.Itoa(cubeno)+".cbs"
				FileCopy(pathfile,pathfile2)
			}
		}
	}
	return aStatistic.AddrIndex
}

func AllIndexing(cubeno int) map[string]string {
	var aIndexing TxIndexing
	aIndexing.AddrIndex=make(map[string]string)	

	if cubeno==0 {
		cubeno=CubeHeight()-1
	}
	if cubeno<=0 {
		cubeno=1
	}

	path:=Configure.Datafolder+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Indexing.cbs"
	rpathfile:=pathfile
	pathfile1:=""
	pathfile2:=""
	if DirExist(pathfile) {
		for i:=SpecialCnt-1;i>0;i-- {
			if i==1 {
				pathfile1=path+filepathSeparator+"Indexing.cbs"
			} else {
				pathfile1=path+filepathSeparator+"Indexing"+strconv.Itoa(i)+".cbs"
			}
			pathfile2=path+filepathSeparator+"Indexing"+strconv.Itoa(i+1)+".cbs"
			if DirExist(pathfile1) {
				if i>3 { rpathfile=pathfile1 }
				FileCopy(pathfile1,pathfile2)
			}
		}
		err:=FileRead(rpathfile,&aIndexing)
		Err(err,0)	
	}
	if aIndexing.Cubeno>=cubeno {
	} else {
		for i:=aIndexing.Cubeno+1;i<cubeno+1;i++ {
			if CheckConfirm(i)==false {
				cubeno=i-1
				break;
			}
			tmpIndexing:=CubeIndexing(i)
			for k,v := range tmpIndexing {
				if aIndexing.AddrIndex[k]=="" {
					aIndexing.AddrIndex[k]=v+","
				} else if v>"1" {
					aIndexing.AddrIndex[k]+=v+","
				}
			}
		}
		if(aIndexing.Cubeno!=cubeno) {
			aIndexing.Cubeno=cubeno
			err:=FileWrite(pathfile,aIndexing)
			if err!=nil && cubeno%10000==0 {
				pathfile2=path+filepathSeparator+"Indexing_"+strconv.Itoa(cubeno)+".cbs"
				FileCopy(pathfile,pathfile2)
			}
		}
	}
	return aIndexing.AddrIndex
}

func AllStatistic(cubeno int) map[string]string {
	var aStatistic TxStatistic
	aStatistic.AddrIndex=make(map[string]string)	
	vStatistic:=make(map[string]StatisticData)

	if cubeno==0 {
		cubeno=CubeHeight()-1
	}
	if cubeno<=0 {
		cubeno=1
	}

	path:=Configure.Datafolder+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	rpathfile:=pathfile
	pathfile1:=""
	pathfile2:=""
	if DirExist(pathfile) {
		for i:=SpecialCnt-1;i>0;i-- {
			if i==1 {
				pathfile1=path+filepathSeparator+"Statistic.cbs"
			} else {
				pathfile1=path+filepathSeparator+"Statistic"+strconv.Itoa(i)+".cbs"
			}
			pathfile2=path+filepathSeparator+"Statistic"+strconv.Itoa(i+1)+".cbs"
			if DirExist(pathfile1) {
				//if i>3 { rpathfile=pathfile1 }
				FileCopy(pathfile1,pathfile2)
			}
		}
		err:=FileRead(rpathfile,&aStatistic)
		Err(err,0)	
	}
	if aStatistic.Cubeno>=cubeno {
	} else {
		for k,v := range aStatistic.AddrIndex {
			if len(k)==34 && k[0:31]!="C"+strings.Repeat("0",30) {
				vStatistic[k]=StatisticStrToData(v)
			}
		}
		for i:=aStatistic.Cubeno+1;i<cubeno+1;i++ {
			if CheckConfirm(i)==false {
				cubeno=i-1
				break;
			}
			vStatistic=GetStaticData(i,vStatistic)
			for k,v := range vStatistic {
				result:=""
				result=strconv.FormatFloat(v.Balance,'f',-1,64)+","
				result+=strconv.FormatFloat(v.RBalance,'f',-1,64)+","
				result+=strconv.FormatFloat(v.SBalance,'f',-1,64)+","
				result+=strconv.FormatFloat(v.TBalance,'f',-1,64)+","
				result+=strconv.Itoa(v.Txcnt)+","
				result+=strconv.Itoa(v.Tkcnt)+","
				result+=v.Pos
				aStatistic.AddrIndex[k]=result
			}			
		}
		if(aStatistic.Cubeno!=cubeno) {
			aStatistic.Cubeno=cubeno
			err:=FileWrite(pathfile,aStatistic)
			if err!=nil {
				pathfile2=path+filepathSeparator+"Statistic_"+strconv.Itoa(cubeno)+".cbs"
				FileCopy(pathfile,pathfile2)
			}
		}
	}
	return aStatistic.AddrIndex
}

func CubeStatisticAdd(cubeno int,cubeStatistic map[string]string) map[string]string {
	var iData []TxData
	var statAddr map[string]bool
	statAddr=make(map[string]bool)

	if cubeno<=0 {
		return cubeStatistic
	}

	var c Cube 
	c.Cubeno=cubeno
	c.Read()
	for i:=-1;i<27;i++ {
		if i==Configure.Indexing || i==Configure.Statistics || i==Configure.Format || i==Configure.Edit {
		} else {
			if i==-1 {
				tx1,tx2:=GetCubePoh(cubeno)
				iData=append(iData,tx1)
				iData=append(iData,tx2)
			} else {
				iData=BlockTxData(c.Blocks[i].Data)
			}
			for _,v := range iData {
				if v.Datatype!="Contract" && v.Datatype!="Data" {
					if(len(v.From)==34) {
						if v.From[0:31]!="C"+strings.Repeat("0",30) {
							statAddr[v.From]=true
						}
					}
					if(len(v.To)==34) {
						if v.To[0:31]!="C"+strings.Repeat("0",30) {
							statAddr[v.To]=true
						}
					}
				}
			}
		}
	}
	for k,_ := range statAddr {
		cStatisticValue:=StatisticAdd(k,cubeno)
		cubeStatistic[k]=cStatisticValue
	}
	return cubeStatistic
}

func StatisticAdd(addr string,cubeno int) string {
	result:=""
	pos:="F"
	balance,rbalance,sbalance,tbalance,txcnt,tkcnt,gcubeno:=GetStaticVar(addr)
	if gcubeno<cubeno {
		for i:=gcubeno+1;i<=cubeno;i++ {
			rbalance1,sbalance1,txcnt1,tkcnt1:=GetStaticVarCube(addr,i)
			sbalance+=sbalance1
			sbalance=math.Round(sbalance*100000000)/100000000
			rbalance+=rbalance1
			rbalance=math.Round(rbalance*100000000)/100000000
			txcnt+=txcnt1
			tkcnt+=tkcnt1
		}
	}
	balance=rbalance-sbalance
	balance=math.Round(balance*100000000)/100000000
	tbalance=rbalance+sbalance
	tbalance=math.Round(tbalance*100000000)/100000000
	if balance>=5000.0 {
		pos="T"
	}
	result=strconv.FormatFloat(balance,'f',-1,64)+","
	result+=strconv.FormatFloat(rbalance,'f',-1,64)+","
	result+=strconv.FormatFloat(sbalance,'f',-1,64)+","
	result+=strconv.FormatFloat(tbalance,'f',-1,64)+","
	result+=strconv.Itoa(txcnt)+","
	result+=strconv.Itoa(tkcnt)+","
	result+=pos
	return result
}

func StatisticStrToVar(item string) (float64,float64,float64,float64,int,int) {
	gsv:=strings.Split(item,",")
	balance,_:=strconv.ParseFloat(gsv[0],64)
	rbalance,_:=strconv.ParseFloat(gsv[1],64)
	sbalance,_:=strconv.ParseFloat(gsv[2],64)
	tbalance,_:=strconv.ParseFloat(gsv[3],64)
	txcnt,_:=strconv.Atoi(gsv[4])
	tkcnt,_:=strconv.Atoi(gsv[5])
	return balance,rbalance,sbalance,tbalance,txcnt,tkcnt
}

func StatisticStrToData(item string) StatisticData {
	var vStatistic StatisticData
	vStatistic.Balance,vStatistic.RBalance,vStatistic.SBalance,vStatistic.TBalance,vStatistic.Txcnt,vStatistic.Tkcnt=StatisticStrToVar(item)
	return vStatistic
}

func StatisticStrSum(item,item2 string) string {
	var vStatistic StatisticData
	var vStatistic2 StatisticData
	var rStatistic StatisticData
	vStatistic.Balance,vStatistic.RBalance,vStatistic.SBalance,vStatistic.TBalance,vStatistic.Txcnt,vStatistic.Tkcnt=StatisticStrToVar(item)
	vStatistic2.Balance,vStatistic2.RBalance,vStatistic2.SBalance,vStatistic2.TBalance,vStatistic2.Txcnt,vStatistic2.Tkcnt=StatisticStrToVar(item2)

	rStatistic.Balance=vStatistic.Balance+vStatistic2.Balance
	rStatistic.RBalance=vStatistic.RBalance+vStatistic2.RBalance
	rStatistic.SBalance=vStatistic.SBalance+vStatistic2.SBalance
	rStatistic.TBalance=vStatistic.TBalance+vStatistic2.TBalance
	rStatistic.Txcnt=vStatistic.Txcnt+vStatistic2.Txcnt
	rStatistic.Tkcnt=vStatistic.Tkcnt+vStatistic2.Tkcnt

	return StatisticString(rStatistic)
}


func StatisticString(rStatistic StatisticData) string{
	rStatistic.Balance=math.Round(rStatistic.Balance*100000000)/100000000
	rStatistic.RBalance=math.Round(rStatistic.RBalance*100000000)/100000000
	rStatistic.SBalance=math.Round(rStatistic.SBalance*100000000)/100000000
	rStatistic.TBalance=math.Round(rStatistic.TBalance*100000000)/100000000
	pos:="F"
	if rStatistic.Balance>=5000.0 {
		pos="T"
	}
	result:=""
	result=strconv.FormatFloat(rStatistic.Balance,'f',-1,64)+","
	result+=strconv.FormatFloat(rStatistic.RBalance,'f',-1,64)+","
	result+=strconv.FormatFloat(rStatistic.SBalance,'f',-1,64)+","
	result+=strconv.FormatFloat(rStatistic.TBalance,'f',-1,64)+","
	result+=strconv.Itoa(rStatistic.Txcnt)+","
	result+=strconv.Itoa(rStatistic.Tkcnt)+","
	result+=pos
	return result
}


