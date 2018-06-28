package handleShared

import (
	"strings"
	"fmt"
	"os"
	"log"
	"time"
	"reflect"
)

func RecordStrInfos(fFmt, infos string) {
	// 判断文件夹是否存在，不在则创建
	_filedir := "../log"
	exist, err := PathExists(_filedir)
	if err != nil {
		fmt.Println("---> handleShared.RecordStrInfos-file err: ", err)
		return
	}
	if exist {
		fmt.Println("---> handleShared.RecordStrInfos-file has exist: ", _filedir)
	} else {
		fmt.Println("---> handleShared.RecordStrInfos-file not exist!")
		//创建文件夹
		err := os.Mkdir(_filedir, os.ModePerm)
		if err != nil {
		fmt.Println("---> handleShared.RecordStrInfos-mkdir file failed：", err)
		} else {
			fmt.Println("---> handleShared.RecordStrInfos-mkdir file success!")
			
		}
	}
	//获取当前时间年月日(string)
	tNow := time.Now()
	fmt.Println("---> tNow.Date()", tNow.String()[:19])
	timeYMD := tNow.String()[:10]
	timeHMS := tNow.String()[11:19]
	
	fmt.Println("---> tNow.format ",reflect.TypeOf(timeYMD), timeYMD, timeHMS)
	filename := strings.Join([]string{_filedir, "/", timeYMD, fFmt}, "")
	
	successed := FileWriteInfos(filename, infos)
	if successed {
		fmt.Println("---> handleShared.RecordStrInfos-file successed!")	
	}

}

func FileWriteInfos(filename, infos string) bool {
	//可写&追加方式打开文件
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	CheckErr(err)
	defer fd.Close()

	//写字符串到文件(string--->byte)
	rInfos := strings.Join([]string{infos, "\n"}, "")
	strWritten, err := fd.Write([]byte(rInfos))
	successed := CheckErr(err)
	if !successed {
		log.Println("---> not successed: ", successed)
		return false
	}
	log.Println("---> Has Written infos: ", strWritten)
	return true

}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func CheckErr(err error) bool {
	if err != nil {
		log.Println("---> Error: ", err)		
	}
	return true
}