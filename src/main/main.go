package main

import (
	"fmt"
	"reflect"
	"handleCmdParsing"
	"timefunc"
	//"time"
	/*
	"handleShared"
	"bytes"
	"strconv"
	*/
	/*
	"chgoroutines"
	"godistrabutor"
	"gyjson"
	"gyrwio"
	"interfaces"
	
	*/
	//"goTcpServer"
	//"goredis"
	//"workerpool"
	//"handleRedis"
	"handleMySql"
)

func main() {
	fmt.Println("Hello Go")
	arraya := []int{1, 2, 3, 4, 56, 7}
	arrayb := [3]int{1, 2, 3}
	slice := make([]int, 3, 10)
	fmt.Println("---> AAA--- > ", reflect.TypeOf(arraya))
	fmt.Println("---> BBB--- > ", reflect.TypeOf(arrayb))
	fmt.Println("---> SSS--- > ", reflect.TypeOf(slice))
	//命令解析模块测试
	timefunc.TimeFunc(handleCmdParsing.Cmd2FParsing)
	timefunc.TimeFunc(handleCmdParsing.Cmd33Parsing)
	// timefunc.TimeFunc(forRange)
	fmt.Println("---Start 33 ---")
	handleCmdParsing.Cmd33Parsing()
	fmt.Println("---Start 2F ---")
	handleCmdParsing.Cmd2FParsing()
	/*

	data := []int{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}

	handleShared.HandleSharedCmdOk(22, data, 0xff)

	//文件读写测试
	tNow := time.Now()
	fmt.Println("---> tNow.Date()", tNow.String()[:19])
	timeYMD := tNow.String()[:10]
	timeHMS := tNow.String()[11:19]
	fmt.Println("---> tNow.format ",reflect.TypeOf(timeYMD), timeYMD, timeHMS)

	handleShared.RecordStrInfos("Rlog", "Hello LaoJiang")
	//调用Linux命令测试
	fmt.Println("---> ", reflect.TypeOf(`dh`))
	handleShared.ExecShell(`df -h`)

	//CRC32测试
	var crc32Buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		crc32Buffer.WriteString(strconv.Itoa(data[i]))
	}
	crc32data := crc32Buffer.String()
	fmt.Println("---> crc32DataJoin: ", crc32data)
	handleShared.ExecCRC32(crc32data)
	
	//MD5测试
	var md5Buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		md5Buffer.WriteString(strconv.Itoa(data[i]))
	}
	md5data := md5Buffer.String()
	fmt.Println("---> md5DataJoin: ", md5data)
	handleShared.ExecMD5(md5data)
	//EEPROM测试
    fmt.Println("---> EEPROM1 : ")
    handleShared.HandleSharedExecCSoI2C0Write(0x54, 0x01, 44)
    bb := handleShared.HandleSharedExecCSoI2C0Read(0x54, 0x01)
	fmt.Println("---> EEPROM2 : ", bb)
	//GoTcpServer测试
	//goTcpServer.GoTcpServer()
	//PCF8563测试
	timeWrite := []int{0,31,11,3,6,12,16}
	handleShared.HandleSharedExecCSoPCFWrite(timeWrite)
	timeBuff := handleShared.HandleSharedExecCSoPCFRead()
	fmt.Println("---> timeBuff: ", timeBuff)

	//GPIO测试
	handleShared.HandleSharedExecCSoGpio37Blinking()

	*/

	//Redis测试
	//goredis.HttpRedisStart()

	/*
	//workerpool测试
	// generate worker to do job
	dispatcher := workerPool.NewDispatcher(3)
	dispatcher.Run()

	// produce job to be done
	producer := workerPool.NewProducer(40)
	producer.Run()

	time.Sleep(time.Second * 2)
	*/
	/*
	//RedisJson测试
	var key01 string
	//var key02 string
	key01 = "Hello02"
	//key02 = "Hello02"
	kvJson := make(map[string]interface{})
	kvJson[key01] = "Redis"
	//kvJson[key02] = "Go"
	handleRedis.HandleRedisJsonInsert(key01, kvJson)
	data := handleRedis.HandleRedisJsonGet(key01)
	fmt.Println("---> Get RedisJson: ", data)

	delv := handleRedis.HandleRedisJsonDel(key01)
	fmt.Println("---> delv: ", delv)
	//handleRedis.HandleRedisJson(key02, kvJson)
	*/
	//MySql测试
	//handleMySql.HandleMySqlPool()
	// handleMySql.HandleDBCreateTable()
	var a = 1.0
	fmt.Println(a + 0.7)
	elecDatas := make([]float64, 5)
	// elecDatas = append(elecDatas, 1.21, 1.31, 1.401, 1.5001, 1.5)
	
	for i := 0; i < 5; i++ {
		elecDatas[i] = float64(i)+100.11
	}
	fmt.Println("---> elecDatas: ", elecDatas)
	
	
	handleMySql.HandleDBElecInsert(2, elecDatas, "dbelec")

	//handleMySql.HandleDBElecDelete(18)
	fmt.Println("---> Select Single.")
	buff := handleMySql.HandleDBElecGetSingle(2)
	fmt.Println("---> buff: ", len(buff), buff)
	fmt.Println("---> Select Manny.")
	handleMySql.HandleDBElecGetManny(1, 10)
	/*
	//chgoroutines.Main()
	//chgoroutines.SelectMain()
	chgoroutines.TikerDemo()
	chgoroutines.ThreadsPools()

	//godistrabutor.Distrabutor01()
	//godistrabutor.Distrabutor02()
	godistrabutor.Distrabutor03()
	gyjson.GYJson01()
	gyjson.GYJson02()
	gyrwio.GYRWIO01()

	interfaces.Interfaces()
	interfaces.InterfacesNew()

	var ary [3]int
	fmt.Println("---> ", ary[0])
	fmt.Println("---> ", ary[len(ary)-1])

	for i, v := range ary {
		//fmt.Println("---> ", i, v)
		fmt.Printf("---> index[%d] value[%d]\n", i, v)
		ary[i] = 1
	}

	fmt.Println("---> ", ary)
	
	fmt.Println("---> ", reflect.TypeOf(ary))

	var ary01 = [...]string{"1", "2"}
	fmt.Println("---> ", reflect.TypeOf(ary01))

	var a *int = test()
	println(a, *a)

	fmt.Println("---> testOpts")
	testOpts("Hello", ary, ary01, a)

	print("---> Go Redis Base Test <---\n")
	goredis.GoRedisBase()
	print("---> Go Redis Json Test <---\n")
	goredis.GoRedisJson()
	*/

}

func forRange() {
	sliceData := []int{1, 2, 3, 4, 5, 6}
	for _, v := range sliceData {
		fmt.Println("Value ---> ", v)
	}
}

func test() *int {
	a := 0x100
	return &a
}

func testOpts(s string, any ...interface{}) {
	fmt.Println("---> ", s, any)
}

