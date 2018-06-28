package handleShared


/*
#cgo CFLAGS : -I../include
#cgo LDFLAGS : -L../libso -lwiringPi

#include "wiringPi.h"
#include "wiringPiI2C.h"
#include "wiringPiSPI.h"
#include "wiringSerial.h"
#include "wiringShift.h"

*/
import "C"

import (
    "fmt"
)

const (
	ID_PCF8563_DEVICE = 0x51 //PCF8563地址
	ID_PCF8563_SEC=0x02
	ID_PCF8563_MIN=0x03
	ID_PCF8563_H=0x04
	ID_PCF8563_D=0x05
	ID_PCF8563_W=0x06
	ID_PCF8563_M=0x07
	ID_PCF8563_Y=0x08
	ID_PCF8563_MIN_ALARM=0x09

)

//I2C0:EEPROM&PCF8563
//I2C1:RELAY(继电器)
/*
//EEPROM和实时时钟芯片通用底层读写接口
*/
func HandleSharedExecCSoI2C0Write(page, offset, wbyte C.int) {
	fd := C.wiringPiI2CSetupInterface(C.CString("/dev/i2c-0"), page)
	if fd > 0 {
		C.wiringPiI2CWriteReg8(fd, offset, wbyte)
		C.delay(5)
		C.serialClose(fd)
	}
}

func HandleSharedExecCSoI2C0Read(page, offset C.int) int{
        var bytes C.int
	fd := C.wiringPiI2CSetupInterface(C.CString("/dev/i2c-0"), page)
	if fd > 0 {
		bytes = C.wiringPiI2CReadReg8(fd, offset)
	        fmt.Println("---> SoTest ", bytes)
		if bytes == (-1) {
			fmt.Println("---> SoTest Error", bytes)
			bytes = 0
		}
		C.serialClose(fd)
	}
	return int(bytes)
}
/*
//实时时钟校时读写接口
*/
func HandleSharedExecCSoPCFWrite(timeBuff []int) {
	
	for i := 0; i < 7; i++ {
		timeBuff[i] = (((timeBuff[i]/10)<<4)|(timeBuff[i]%10))
		if i == 4 {
			timeBuff[i] = timeBuff[i]
		}
		HandleSharedExecCSoI2C0Write(C.int(ID_PCF8563_DEVICE), C.int(ID_PCF8563_SEC+i), C.int(timeBuff[i]))
	}
	/*
	timeBuff[0] = (((timeBuff[0]/10)<<4)|(timeBuff[0]%10))
	timeBuff[1] = (((timeBuff[1]/10)<<4)|(timeBuff[1]%10))
	timeBuff[2] = (((timeBuff[2]/10)<<4)|(timeBuff[2]%10))
	timeBuff[3] = (((timeBuff[3]/10)<<4)|(timeBuff[3]%10))
	timeBuff[4] = timeBuff[4]
	timeBuff[5] = (((timeBuff[5]/10)<<4)|(timeBuff[5]%10))
	timeBuff[6] = (((timeBuff[6]/10)<<4)|(timeBuff[6]%10))

	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_SEC, timeBuff[0])
	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_MIN, timeBuff[1])
	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_H, timeBuff[2])
	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_D, timeBuff[3])
	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_W, timeBuff[4])
	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_M, timeBuff[5])
	HandleSharedExecCSoI2C0Write(ID_PCF8563_DEVICE, ID_PCF8563_Y, timeBuff[6])
	*/

}

func HandleSharedExecCSoPCFRead() []int{
	timeBuff := make([]int, 7)
	for i := 0; i < 7; i++ {
		timeBuff[i] = HandleSharedExecCSoI2C0Read(C.int(ID_PCF8563_DEVICE), C.int(ID_PCF8563_SEC+i))
	}
	/*
	timeBuff[0] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_SEC)
	timeBuff[1] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_MIN)
	timeBuff[2] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_H)
	timeBuff[3] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_D)
	timeBuff[4] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_W)
	timeBuff[5] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_M)
	timeBuff[6] = HandleSharedExecCSoI2C0Read(ID_PCF8563_DEVICE, ID_PCF8563_Y)
	*/
	timeBuff[0] = ((timeBuff[0]&0x7F)>>4)*10+(timeBuff[0]&0x0F)
	timeBuff[1] = ((timeBuff[1]&0x7F)>>4)*10+(timeBuff[1]&0x0F)
	timeBuff[2] = ((timeBuff[2]&0x3F)>>4)*10+(timeBuff[2]&0x0F)
	timeBuff[3] = ((timeBuff[3]&0x3F)>>4)*10+(timeBuff[3]&0x0F)
	timeBuff[4] = timeBuff[4]&0x07
	timeBuff[5] = ((timeBuff[5]&0x1F)>>4)*10+(timeBuff[5]&0x0F)
	timeBuff[6] = ((timeBuff[6]&0xFF)>>4)*10+(timeBuff[6]&0x0F)

	return timeBuff
}
/*
//校验实时时钟芯片
*/
func HandleSharedExecCSoPCFAlarmWrite(check C.int) {
	HandleSharedExecCSoI2C0Write(C.int(ID_PCF8563_DEVICE), C.int(ID_PCF8563_MIN_ALARM), check)
}

func HandleSharedExecCSoPCFAlarmRead() int {
	tmp := HandleSharedExecCSoI2C0Read(C.int(ID_PCF8563_DEVICE), C.int(ID_PCF8563_MIN_ALARM))
	return tmp
}

//RELAY继电器相关
func HandleSharedExecCSoI2C1Write(page, offset, wbyte C.int) {
	fd := C.wiringPiI2CSetupInterface(C.CString("/dev/i2c-0"), page)
	if fd > 0 {
		C.wiringPiI2CWriteReg8(fd, offset, wbyte)
		C.delay(5)
		C.serialClose(fd)
	}
}


func HandleSharedExecCSoI2C1Read(page, offset C.int) int{
        var bytes C.int
	fd := C.wiringPiI2CSetupInterface(C.CString("/dev/i2c-0"), page)
	if fd > 0 {
		bytes = C.wiringPiI2CReadReg8(fd, offset)
	        fmt.Println("---> SoTest ", bytes)
		if bytes == (-1) {
			fmt.Println("---> SoTest Error", bytes)
			bytes = 0
		}
		C.serialClose(fd)
	}
	return int(bytes)
}

//SPI&GPIO相关接口
func HandleSharedExecCSoSpiInit() {
	C.wiringPiSPISetup(0,10000)
}

func HandleSharedExecCSoGpioInit() {
	C.wiringPiSetupPhys()
}

func HandleSharedExecCSoGpioOutput(pinPhysNum, value C.int) {
	C.pinMode(pinPhysNum, 1)
	C.digitalWrite(pinPhysNum,value)
}

func HandleSharedExecCSoGpioInput(pinPhysNum, upDnMode C.int) int {
	C.pinMode(pinPhysNum, 0)
	C.pullUpDnControl(pinPhysNum,upDnMode)
	return int(C.digitalRead(pinPhysNum))
}

func HandleSharedExecCSoGpio37Blinking() {
	HandleSharedExecCSoGpioInit()
	HandleSharedExecCSoSpiInit()
	for {
		HandleSharedExecCSoGpioOutput(26, 1)
		C.delay(1000)
		HandleSharedExecCSoGpioOutput(26, 0)
		C.delay(1000)
		/*
		HandleSharedExecCSoGpioOutput(12, 0)
		C.delay(1000)
		fmt.Println("---> GpioInput: ", HandleSharedExecCSoGpioInput(36, 1))
		*/
	}
}