package core

import (
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"strings"
    "io"
    "io/ioutil"
)


func FilePath(idx int) string {
	divn:=idx/Configure.Datanumber
	divm:=idx%Configure.Datanumber
	if divm>0 {
		divn++
	} else if divm==0 {
		divm=Configure.Datanumber
	}
	if divn==0 {
		divn++
		divm=1
	}
	nhex:=fmt.Sprintf("%x",Configure.Datanumber)
	mcnt:=len(nhex)
	nstr:=fmt.Sprintf("%0.5x",divn)
	mstr:=fmt.Sprintf("%0."+strconv.Itoa(mcnt)+"x",divm)
	dirname:=Configure.Datafolder+filepathSeparator+nstr+filepathSeparator+mstr
	
	return dirname
}

func MakeDir(idx int) string {
	dirname:=FilePath(idx)
	result:=MakePath(dirname)
	return result
}

func MakePath(dirname string) string {
	if DirExist(dirname)==false {
		if err:=os.MkdirAll(dirname, 0775); err!=nil {
			return "Directory cannot create."
		} else {
			os.Chown(dirname, 0, 48)
			os.Chmod(dirname, 0775)
		}
	}
	return dirname
}

func FileWrite(path string, object interface{}) error {
	file,err:=os.Create(path)
	if err==nil {
		encoder:=gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func FileRead(path string, object interface{}) error {
	/*
	var block Block
	var cube Cube
	var cubing Cubing
	
	gob.Register(block)  
	gob.Register(cube)
	gob.Register(cubing)  
	*/
	if DirExist(path)==false {
		return nil
	}

	file,err:=os.Open(path)
	if err==nil {
		decoder:=gob.NewDecoder(file)
		err=decoder.Decode(object)
	}
	file.Close()
	return err
}

func FileReadString(path string) string {
    data, err := ioutil.ReadFile(path)
    if err!=nil {
		echo (err)
    }
	return ByteToStr(data)
}


func FileSize(dirpath string) int64 {
	if DirExist(dirpath)==false {
		return 0.0
	}
	file, err := os.Open(dirpath) 
	if err != nil {
		echo (err)
	}
	fi, err := file.Stat()
	if err != nil {
		echo (err)
	}
	file.Close()
	return fi.Size()
}

func FileSearch(dirname string,find string) string{
    result:=""
	if DirExist(dirname)==false {
		return ""
	}

	d,err:=os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
    file, err:=d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi:=range file {
        if fi.Mode().IsRegular() {
			fstr:=fi.Name()
			if strings.Index(fstr,find)>=0 {
				result=fi.Name()
				return result
			}
        }
    }
	return result
}

func FileBlockSearch(dirname string,find string) string{
    result:=""
	if DirExist(dirname)==false {
		return ""
	}

	d,err:=os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
    file, err:=d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi:=range file {
        if fi.Mode().IsRegular() {
			fstr:=fi.Name()
			if strings.Index(fstr,find)>=0 {
				result=fi.Name()
				if len(result)>30 && result[len(result)-4:]==".blk" {
					return result
				}
			}
        }
    }
	return result
}

func DirExist(dirName string) bool{
	result:=true
	_,err:=os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			result=false
		}
	}
	return result
}

func MaxFind(dirpath string) string {
	find:="0"
    d, err:=os.Open(dirpath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
	fi, err:=d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi:=range fi {
        if fi.Mode().IsRegular() {
        } else {
  			if fi.Name()!="special" && fi.Name()>find {
				find=fi.Name()
			}
		}
   }
   return find
}

func CubePathNum(path string) int {
	result:=0
	split:=strings.Split(path, filepathSeparator)
	slen:=len(split)
	nint,_:=strconv.ParseUint(split[slen-2],16,32)
	mint,_:=strconv.ParseUint(split[slen-1],16,32)
	result=(int(nint)-1)*Configure.Datanumber+int(mint)
	return result
}

func PathDelete(path string) error {
	err:=os.RemoveAll(path)
	os.MkdirAll(path,0775)
	return err
}

func FileCopy(path1 string,path2 string) {
    originalFile, err := os.Open(path1)
	Err(err,0)
    defer originalFile.Close()

    newFile, err := os.Create(path2)
	Err(err,0)
    defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	Err(err,0)
    decho(bytesWritten)
    
    err = newFile.Sync()
	Err(err,0)
}


func FileLog(pathdir,pathfile,str string) {
	sf:=""
	if FileSearch(pathdir,pathfile)>"" {
		sf=FileReadString(pathdir+pathfile)
	}
	sf=str+"\n"+sf
	FileWrite(pathdir+pathfile,sf)
}

