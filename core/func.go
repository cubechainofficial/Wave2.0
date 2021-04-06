package core

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
    "bytes"
	"net"
	"net/http"
	"io"
	"io/ioutil"
    "path/filepath"
	"mime/multipart"

	"../config"
)


var Configure config.Configuration
var filepathSeparator=string(filepath.Separator)
var echo=fmt.Println
var MineCubeno=0
var TxMine string
var TxMerkleHash string
var Sumfee=0.0
var Txamount=0.0
var Txcnt=0
var Tkcnt=0
var CubeSetNum [27]string
var DebugMode=false
var SpecialCnt=5

var Version="2.01"
var Pratio=Pohr{4.566,4.566,4.566,127.855,54.795}
var Pratio1=Pohr{4.566,4.566,4.566,127.855,54.795}
var Pratio2=Pohr{3.913,3.913,3.913,109.590,73.060}
var Pratio3=Pohr{3.261,3.261,3.261,91.325,91.325}
var Pratio4=Pohr{2.609,2.609,2.609,73.060,109.590}
var Pratio5=Pohr{1.956,1.956,1.956,54.795,127.855}
var Pratio6=Pohr{1.304,1.304,1.304,36.530,146.120}
var Pratio7=Pohr{0.652,0.652,0.652,18.265,164.385}
var Pratio8=Pohr{0,0,0,0,182.650}
var Pratio9=Pohr{0,0,0,0,182.650}
var Pratio10=Pohr{0,0,0,0,182.650}
var exAddr=[]string{"CLwaJn7JWtHPC98mUzybn54guuFiKXtqKu","CbFYBk5Sdor5AKonQbpN3fo7J7KXJYSVFS","CUQX229R16k5SmBduqZU8MwUb667xX6i8o","CXBindaHGHpRdyeY7ERc8LRi7e2uvGYTGq","CML2JEtqArxRoGSAb9CvGaadLFwvhFjh3B","Ce7jPeg69UQhNLYcNaJFVxATGm1xJSyoPG","CaYTrJaEUes2MezYJ4kt7BAoKLCVa5jm8q","CN1qmgcmufAwsgDk1JDoF1B2A9Q5HJCPRR","CYvX6pvgcJpoNWDxdhRHrnMkAhSs9hVyvi","CTFqiRtjJErihp61T7oxhuauL8vRD8sVmy"}
var RepairAddress="CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye"
var mineExAddr=[]string{"CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CMEYDLY8eg8XQ8cGZd7axPCrMV7Z8j49af","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye","CcXFizxhCykQ5M2vuCmTnoK86rJERfyGye"}
var mineExCnt=0



var TxDelim="|"
var BlockDelim="||"
var CubeDelim="|||"
var CNO=1
var PrvCubing Cubing
var CurrCube Cube
var GenFile string
var GenBlock [27]string
var MineDifficulty="Ffffffff"
var MineDifficultyBase="Ffffffff"
var DiffTime=0


type Pair struct {
  Key string
  Value int
}
type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

type PairFloat struct {
  Key string
  Value float64
}
type PairFloatList []PairFloat

func (p PairFloatList) Len() int { return len(p) }
func (p PairFloatList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairFloatList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func PairSortInt(sortmap map[string]int) PairList{
	plist:=make(PairList,len(sortmap))
	p:=0
	filter:=false
	for k,v := range sortmap {
		filter=false
		for _,addr:= range exAddr {
			if addr==k || len(k)!=34 || k[0:31]=="C"+strings.Repeat("0",30) {
				filter=true
				break;
			}
		}
		if filter==false {
			plist[p]=Pair{k,v}
			p++
		}
	}
	sort.Sort(sort.Reverse(plist))
	return plist
}

func PairSortFloat(sortmap map[string]float64) PairFloatList{
	plist:=make(PairFloatList,len(sortmap))
	p:=0
	filter:=false
	for k,v := range sortmap {
		filter=false
		for _,addr:= range exAddr {
			if addr==k || len(k)!=34 || k[0:31]=="C"+strings.Repeat("0",30) {
				filter=true
				break;
			}
		}
		if filter==false {
			plist[p]=PairFloat{k,v}
			p++
		}
	}
	sort.Sort(sort.Reverse(plist))
	return plist
}


func Err(err error, exit int) int {
	if err != nil {
		fmt.Println(err)	
	}
	if exit>=1 {
		os.Exit(exit)
		return 1
	}
	return 0
}

func netError(err error) {
	if err!=nil && err!=io.EOF {
		fmt.Println("Network Error : ", err)
	}
}

func IpCheck() []string {
	host, err := os.Hostname()
	if err != nil {
		return nil
	}
	addrs, err := net.LookupHost(host)
	if err != nil {
		return nil
	}
	addrs=append(addrs,host)
	if len(addrs)==2 {
		addrs2:=make([]string,3)
		addrs2[0]="mac_linux"
		addrs2[1]=addrs[0]
		addrs2[2]=addrs[1]
		addrs=addrs2
	}
	return addrs
}

func GetBytes(key interface{}) []byte {
    var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
		decho(err)
        return nil
    }
    return buf.Bytes()
}

func StrToByte(str string) []byte {
	sb := make([]byte, len(str))
	for k, v := range str {
		sb[k] = byte(v)
	}
	return sb[:]
}

func ByteToStr(bytes []byte) string {
	var str []byte
	for _, v := range bytes {
		if v != 0x0 {
			str = append(str, v)
		}
	}
	return fmt.Sprintf("%s", str)
}

func GetCubeHeight() string {
	result:=NodeSend("cubeheight","0")
	return result
}

func GetCubeHeight2() string {
	result:=NodeSend2("cubeheight","0")
	return result
}

func GetCubeHeight3() string {
	result:=NodeCube("cubeheight","0")
	return result
}

func CubeHeight() int {
	result:=CurrentHeight()
	return result
}

func CurrentHeight() int {
	result:=0
	f:=MaxFind(Configure.Datafolder+filepathSeparator)
	if f=="0" {
		return 1
	}
	f2:=MaxFind(Configure.Datafolder+filepathSeparator+f)
	if f2=="0" {
		return 1
	}
	nint,_:=strconv.ParseUint(f,16,32)
	mint,_:=strconv.ParseUint(f2,16,32)
	result=(int(nint)-1)*Configure.Datanumber+int(mint)
	if FileSearch(FilePath(result),".cub")>"" {
		result++
	}
	return result	
}

/*
func GetTxCount(addr string) int {
	return 1
}
*/


func BlockName(idx int,cno int) string {
	find:=strconv.Itoa(idx)+"_"+strconv.Itoa(cno)+"_"	
    dirname:=FilePath(idx)
	if DirExist(dirname)==false {
		return ""
	}	
	result:=FileBlockSearch(dirname,find)
	return result
}

func BlockRead(index int,cubeno int,object *Block) error {
	var gBlock Block
	gob.Register(gBlock)

	gBlock=CBlockRead(index,cubeno)
	*object=gBlock
	return nil	

	/*
	filename:=BlockName(index,cubeno)
	if filename=="" {
		gBlock=CBlockRead(index,cubeno)
		*object=gBlock
		return nil
	}
	
	datapath:=FilePath(index)+filepathSeparator 

	file,err:=os.Open(datapath+filename)
	if err==nil {
		decoder:=gob.NewDecoder(file)
		err=decoder.Decode(object)
	}
	file.Close()
	
	return err
	*/
}


func CBlockRead(cubeno int,blockno int) Block {
	var c Cube
	CubeRead(cubeno,&c)
	return c.Blocks[blockno-1]
}


func IndexingRead(aIndexing *TxIndexing) {
	path:=Configure.Datafolder+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Indexing.cbs"
	if DirExist(pathfile) {
		err:=FileRead(pathfile,aIndexing)
		Err(err,0)	
	}
}

func StatisticRead(aStatistic *TxStatistic) {
	path:=Configure.Datafolder+filepathSeparator+"special"
	if DirExist(path)==false {
		if err:=os.MkdirAll(path, os.FileMode(0755)); err!=nil {
			echo ("Special block directory not found")
		}	
	}
	pathfile:=path+filepathSeparator+"Statistic.cbs"
	if DirExist(pathfile) {
		err:=FileRead(pathfile,aStatistic)
		Err(err,0)	
	}
}

func BlockScan(cubeno int,blockno int) Block {
	var block Block
	BlockRead(cubeno,blockno,&block)
	return block
}



func ReadBlockHash(index int,cubeno int) string {
	var iBlock Block
	var hash string
	err:=BlockRead(index,cubeno,&iBlock)
	Err(err,0)
	hash=iBlock.Hash
	return hash
}

func ReadBlockPHash(index int,cubeno int) string {
	var iBlock Block
	var hash string
	err:=BlockRead(index,cubeno,&iBlock)
	Err(err,0)
	hash=iBlock.PatternHash
	return hash
}


func CubeRead(index int,object *Cube) error {
	datapath:=FilePath(index)+filepathSeparator 
	filename:=FileSearch(datapath,".cub")
	if filename=="" {
		return nil
	}	
	file,err:=os.Open(datapath+filename)
	if err==nil {
		decoder:=gob.NewDecoder(file)
		err=decoder.Decode(object)
	}
	file.Close()
	return err
}

func CubeFileName(idx int) string {
	find:=".cub"	
    dirname:=FilePath(idx)
	result:=FileSearch(dirname,find)
	return result
}


func CubePath(idx int) string {
	find:=".cub"	
    dirname:=FilePath(idx)
	result:=FileSearch(dirname,find)
	return dirname+filepathSeparator+result
}


func CubeSync() bool {
	c,_:=strconv.Atoi(GetCubeHeight())
	cc:=CubeHeight()-1
	ccube:=cc
	echo(c)
	echo(cc)
	for c>cc {
		ccube=cc
		echo (ccube)
		CubeDownloadFile(ccube)
		cc=CubeHeight()
		if(ccube==cc) {
			echo ("Download failure : "+strconv.Itoa(ccube))
			cc++
		}
	}
	c,_=strconv.Atoi(GetCubeHeight())
	cc=CubeHeight()
	if(cc>c) {
		return true
	}
	return false
}

func SpecialSync() {
	url1:="http://"+Configure.PosServer+":7080/files/special/Indexing.cbs"
	url2:="http://"+Configure.PosServer+":7080/files/special/Statistic.cbs"
	filepath:=Configure.Datafolder+filepathSeparator+"special"+filepathSeparator
	s1,e1:=DownloadFileWithName(filepath,url1,"Indexing.cbs")	
	s2,e2:=DownloadFileWithName(filepath,url2,"Statistic.cbs")	

	echo(s1)
	echo(s2)
	echo(e1)
	echo(e2)
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}


func NodeSend(cmode string,data string) string {
	arr:=IpCheck()
	reader:=strings.NewReader("cmode="+cmode+"&_token=9X1rK2Z2sofIeFpqg6VBXI5aUWsPOfGPGyzzztgu&data="+data+"&mac="+arr[0]+"&ip="+arr[1]+"&hostname="+arr[2]+"&netname="+Configure.Network+"&netset="+Configure.Nettype+"&chaintype="+Configure.Chaintype+"&netport="+strconv.Itoa(Configure.Port)+"&ver="+Version)
	request,_:=http.NewRequest("POST","http://"+Configure.MainServer+"/"+cmode, reader)
	request.Header.Add("content-type","application/x-www-form-urlencoded")
	request.Header.Add("cache-control","no-cache")
	client:=&http.Client{}
	res, err := client.Do(request)
	Err(err,0)
	body, err := ioutil.ReadAll(res.Body)
	Err(err,0)
	defer res.Body.Close()
	s:=string(body)
	return s
}

func NodeSend2(cmode string,data string) string {
	arr:=IpCheck()
	reader:=strings.NewReader("cmode="+cmode+"&_token=9X1rK2Z2sofIeFpqg6VBXI5aUWsPOfGPGyzzztgu&data="+data+"&mac="+arr[0]+"&ip="+arr[1]+"&hostname="+arr[2]+"&netname="+Configure.Network+"&netset="+Configure.Nettype+"&chaintype="+Configure.Chaintype+"&netport="+strconv.Itoa(Configure.Port)+"&ver="+Version)
	request,_:=http.NewRequest("POST","http://"+Configure.PoolServer+"/"+cmode, reader)
	request.Header.Add("content-type","application/x-www-form-urlencoded")
	request.Header.Add("cache-control","no-cache")
	client:=&http.Client{}
	res, err := client.Do(request)
	Err(err,0)
	body, err := ioutil.ReadAll(res.Body)
	Err(err,0)
	defer res.Body.Close()
	s:=string(body)
	return s
}

func NodeCube(cmode string,data string) string {
	arr:=IpCheck()
	reader:=strings.NewReader("cmode="+cmode+"&_token=9X1rK2Z2sofIeFpqg6VBXI5aUWsPOfGPGyzzztgu&data="+data+"&mac="+arr[0]+"&ip="+arr[1]+"&hostname="+arr[2]+"&netname="+Configure.Network+"&netset="+Configure.Nettype+"&chaintype="+Configure.Chaintype+"&netport="+strconv.Itoa(Configure.Port)+"&ver="+Version)
	request,_:=http.NewRequest("POST","http://"+Configure.PosServer+"/"+cmode, reader)
	request.Header.Add("content-type","application/x-www-form-urlencoded")
	request.Header.Add("cache-control","no-cache")
	client:=&http.Client{}
	res, err := client.Do(request)
	Err(err,0)
	body, err := ioutil.ReadAll(res.Body)
	Err(err,0)
	defer res.Body.Close()
	s:=string(body)
	return s
}

func CubeDownload(cubeno int) string {
	decho("Download file")
	cubenum:=strconv.Itoa(cubeno)
    fileUrl:="http://"+Configure.PosServer+"/download/"+cubenum
    filepath,err := DownloadFile(FilePath(cubeno)+filepathSeparator, fileUrl)
    if err != nil {
		decho(err)
    }
	return filepath
}

func CubeDownloadFile(cubeno int) string {
	hashnip:=NodeCube("cubehash","0&cubeno="+strconv.Itoa(cubeno))
	path:=""
	if hashnip=="0,0" || hashnip=="" {
	} else {
		haship:=strings.Split(hashnip, ",")
		if haship[1]>"0" {
			path=CubeDownloadFrom(cubeno,haship[1],haship[0]+".cub")
		}
	}
	return path
}	


func CubeDownloadFrom(cubeno int,ip string,hash string) string {
	decho("Download file")
    fileUrl:="http://"+ip+":"+strconv.Itoa(Configure.Httpport)+"/download?cubeno="+strconv.Itoa(cubeno)
	echo (fileUrl)
    downpath:=FilePath(cubeno)+filepathSeparator
	echo (downpath)
	filepath,err := DownloadFileWithName(downpath,fileUrl,hash)
    if err != nil {
		decho(err)
    }
	return filepath
}

func CubeDownloadRpcFrom(cubeno int,ip string) string {
	decho("Download file")
	//cubeno:=strconv.Itoa(cubeno)
	//filename:="test"+cubenum
    fileUrl:="http://"+ip+":"+strconv.Itoa(Configure.Httpport)
    filepath,err := DownloadRpc(cubeno,fileUrl)
    if err != nil {
		decho(err)
    }
	return filepath
}


func DownloadRpc(cubeno int, url string) (string,error) {
	drpc := Request{Callno:1,Com:"download_cube",Rmsg:"downcube"}
	drpc.Vars["cubeno"]=strconv.Itoa(cubeno)
    dbytes, _ := json.Marshal(drpc)
    buff := bytes.NewBuffer(dbytes) 
	resp,err:=http.Post(url,"application/json",buff)	
    if err != nil {
        return "",err
    }
    defer resp.Body.Close()
	filename:=headerFilename(resp)
	if filename=="untitle.file" {
		return "",nil	
	}
    filepath:=FilePath(cubeno)+filepathSeparator
	MakePath(filepath)
	filepath+=headerFilename(resp)
	out,err:=os.Create(filepath)
    if err!=nil {
        return "",err
    }
    defer out.Close()
    _, err=io.Copy(out,resp.Body)
    return filepath,err
}

func DownloadFileWithName(filepath string, url string,filename string) (string,error) {
    resp,err:=http.Get(url)
    if err != nil {
        return "",err
    }
	filepath+=filename
    defer resp.Body.Close()
	out,err:=os.Create(filepath)
    if err!=nil {
        return "",err
    }
    defer out.Close()
    _, err=io.Copy(out, resp.Body)
    return filepath,err
}

func DownloadFile(filepath string, url string) (string,error) {
    resp,err:=http.Get(url)
    if err != nil {
        return "",err
    }
    defer resp.Body.Close()
	filename:=headerFilename(resp)
	if filename=="untitle.file" {
		return "",nil	
	}

	MakePath(filepath)
	filepath+=headerFilename(resp)
	out,err:=os.Create(filepath)
    if err!=nil {
        return "",err
    }
    defer out.Close()
    _, err=io.Copy(out, resp.Body)
    return filepath,err
}

func headerFilename(resp *http.Response) string {
	filename:="untitle.file"
	decho(resp)

	if resp.Header["Content-Length"][0]=="0" {
	} else if resp.Header["Content-Disposition"][0]>"" {
		filename=resp.Header["Content-Disposition"][0]
		filename=strings.Replace(filename,"attachment;", "",-1)
		filename=strings.Replace(filename," ", "",-1)
		filename=strings.Replace(filename,"filename=", "",-1)
		filename=strings.Replace(filename,"'", "",-1)
		//decho(filename)
	} else {
	}
	return filename
}

func timecheck() int {
	timestamp:=NodeCube("timecheck","0")
	result,_:=strconv.Atoi(timestamp)
	return result
}

func timecube(cubeno int) (int,int) {
	timestamp:=NodeCube("timecube","0&cubeno="+strconv.Itoa(cubeno))
	result:=strings.Split(timestamp,"|")
	timescheck,_:=strconv.Atoi(result[0])
	cubetime,_:=strconv.Atoi(result[1])
	
	return timescheck,cubetime
}

func Timestampcube(cubeno int) (int,int) {
	timestamp:=NodeCube("timestamp","0&cubeno="+strconv.Itoa(cubeno))
	result:=strings.Split(timestamp,"|")
	timescheck,_:=strconv.Atoi(result[0])
	cubetime,_:=strconv.Atoi(result[1])
	return timescheck,cubetime
}

func difficulty() string {
	MineDifficulty:=NodeCube("difficulty","0")
	if len(MineDifficulty)<8 {
		MineDifficulty=MineDifficultyBase
	}
	return MineDifficulty
}	



func CheckConfirm(cubeno int) bool {
	if(CubeHeight()-cubeno>5) {
		return true
	}
	cpath:=FilePath(cubeno)
	path:=cpath+filepathSeparator+"confirm_"+strconv.Itoa(cubeno)
	if DirExist(path)==false {
		return false
	} else {
		return true
	}
}

func RepairFile(cubeno int,str string) {
	cpath:=FilePath(cubeno)
	path:=cpath+filepathSeparator
	filename:="repair_"+strconv.Itoa(cubeno)
	FileLog(path,filename,str)
}

func CompleteFile(cubeno int,str string) {
	cpath:=FilePath(cubeno)
	path:=cpath+filepathSeparator
	filename:="complete_"+strconv.Itoa(cubeno)
	FileLog(path,filename,str)
}

func decho(v interface{}) {
	if DebugMode {
		echo(v)
	}
}