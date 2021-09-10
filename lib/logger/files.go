package logger

import (
	"fmt"
	"os"
)
//定义一些常用的文件操作


/*
golang判断文件或文件夹是否存在的方法为使用os.Stat ()函数返回的错误值进行判断:
如果返回的错误为nil,说明文件或文件夹存在,如果返回的错误类型使用os.IsNotExist 
()判断为true,说明文件或文件夹不存在,如果返回的错误为其它类型,则不确定是否在存在
*/

//用于检查文件是否存在
func checkNotExist(src string) bool{
	//Stat返回一个描述name指定的文件对象的FileInfo,区别于Lstat
	_,err := os.Stat(src)
	return os.IsNotExist(err)
}

//检查权限
func checkPermission(src string) bool{
	//err的数据类型是err error
	_,err :=os.Stat(src)
	//返回一个布尔值说明该错误是否表示因权限不足要求被拒绝
	return os.IsPermission(err)
}

//如果不存在就创建
func isNotExistMkdir(src string) error(){
	if notExist := checkNotExist(src); notExist ==true{
		if err :=mkdir(src);err !=nil{
			return error
		}
	}
	return  nil
}

//MkdirAll使用指定的权限和名称创建一个目录，正常的话返回nil,否则返回error
func mkDir(src  string) error{
	//其中，Perm是权限位
	err := os.MkdirAll(src,os.ModePerm)
	if err != nil{
		return err
	}
	return nil
}

func mustOpen(fileName, dir string) (*os.File,error){
	perm := checkPermission(dir)

	//表明权限不足
	if perm == true{
		return nil,fmt.Errorf("permisson denied dir:%s",dir)
	}

	err := isNotExistMkdir(dir)
	//未创建成功
	if err!=nil{
		return nil,fmt.Errorf("error during make dir %S,err:%s",dir,err)
	}
	f,err := os.OpenFile(dir+string(os.PathSeparator)+fileName,os.O_APPEND|os.O_CREATE|os.O_rdwr,0644)
	if err!=nil{
		return nil,fmt.Errorf('fail to open file,err:%s',err)
	}
	return f, nil
}