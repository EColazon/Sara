package handleCmdsManagesSR
/*
流程: 
	1.CmdGeter监听ChCmd
	2.根据不同命令id做不同动作
		2.1DoRedisMessages
		2.2DoStateCallback
		2.3DoSender
	3.DoZigbeeTasks
		3.xDoTasks...

Author:mengfei.wu@foxmail.com
---------start:2018.07.10---------
*/

import (
	"fmt"
	"time"
	"reflect"

	"handleShared"
)

// 实现两个接口s
// HandleCmdGeter:监听channel
// HandleCmdSender:DoZigBeeTasks

type CmdsDistributer interface {
	HandleCmdGeter()
	HandleCmdSender()
}

// channel数据
type CmdChannel struct {
	id		int
	data 	[]int
	snum 	int
}

func (cmd CmdChannel)HandleCmdGeter() {
	fmt.Println("---> CmdGeter Start.")
	go func() {
		for{
			select {

			// fmt.Println("---> buff33: ", len(ChCmd33), cap(ChCmd33), buff33)
			// if buff33["id"] == 73003 {
			// 	fmt.Println("---> id 73003")

			// }
			// case buff2f := <- ChCmd2F:
			// 	fmt.Println("---> buff2f: ", len(ChCmd2F), cap(ChCmd2F), buff2f)
			// 	if buff2f["id"] == 72088 {
			// 		fmt.Println("---> id 72001")
			// 	}
			case buff33 := <- ChCmd33:
				fmt.Println("---> buff33: ", len(ChCmd33), cap(ChCmd33), buff33, time.Now())
				cmd.id, _ = buff33["id"].(int)
				cmd.data, _ = buff33["data"].([]int)
				cmd.snum, _ = buff33["snum"].(int)
				// switch buff33["id"] {
				switch cmd.id {
				case 73000:
					// DoTimeSaveLatLongitude()
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73000")
				
				case 73001:
					// DoTimeSaveHoliday()
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73001")
				
				case 73002:
					// DoTimeSaveSpecial()
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73002")
				
				case 73003:
					// DoTimeSavePwmStage()
					// 将interface{}类型转为[]int类型(通过断言)
					// value, _ := buff33["data"].([]int)
					// snum, _ := buff33["snum"].(int)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73003", reflect.TypeOf(cmd.data), cmd.data[:3], cmd.snum)
				
				case 73004:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73004")
				
				case 73005:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73005")
				
				case 73006:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73006")
				
				case 73007:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73007")
				
				case 73008:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73008")
				
				case 73009:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 73009")
				default:
					fmt.Println("---> id 33 not fetched..")
					return
				}
			case buff2f := <- ChCmd2F:
				fmt.Println("---> buff2f: ", len(ChCmd2F), cap(ChCmd2F), buff2f, time.Now())
				cmd.id, _ = buff2f["id"].(int)
				cmd.data, _ = buff2f["data"].([]int)
				cmd.snum, _ = buff2f["snum"].(int)
				// switch buff2f["id"] {
				switch cmd.id {
				case 72000: // 返回8字节IEEE地址
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72000")
				case 72001:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72001")
				case 72002:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72002")
				case 72003:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72003")
				case 72004:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72004")
				case 72005:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72005")
				case 72006:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72006")
				case 72007:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72007")
				case 72008:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72008")
				case 72009:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72009")
				case 720010:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72010")
				case 72011:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72011")
				case 72012:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72012")
				case 72013:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72013")
				case 72014:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72014")
				case 72015:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72015")
				case 72016:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72016")
				case 72017:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72017")
				case 72018:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72018")
				case 72019:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72019")
				case 72020:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72020")
				case 72021:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72021")
				case 72022:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72022")
				case 72023:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72023")
				case 72024:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72024")
				case 72025:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72025")
				case 72026:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72026")
				case 72027:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72027")
				case 72028:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72028")
				case 72029:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72029")
				case 72030:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72030")
				case 72031:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72031")
				case 72032:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72032")
				case 72033:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72033")
				case 72034:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72034")
				case 72035:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72035")
				case 72036:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72036")
				case 72037:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72037")
				case 72038:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72038")
				case 72039:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72039")
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
				case 72040:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72040")
				case 72041:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72041")
				case 72042:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72042")
				case 72043:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72043")
				case 72044:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72044")
				case 72045:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72045")
				case 72046:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72046")
				case 72047:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72047")
				case 72048:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72048")
				case 72049:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72049")
				case 72050:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72050")
				case 72051:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72051")
				case 72052:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72052")
				case 72053:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72053")
				case 72054:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72054")
				case 72055:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72055")
				case 72056:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72056")
				case 72057:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72057")
				case 72058:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72058")
				case 72059:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72059")
				case 72060:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72060")
				case 72061:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72061")
				case 72062:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72062")
				case 72063:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72063")
				case 72064:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72064")
				case 72065:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72065")
				case 72066:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72066")
				case 72067:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72067")
				case 72068:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72068")
				case 72069:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72069")
				case 72070:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72070")
				case 72071:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72071")
				case 72072:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72072")
				case 72073:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72073")
				case 72074:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72074")
				case 72075:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72075")
				case 72076:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72076")
				case 72077:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72077")
				case 72078:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72078")
				case 72079:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72079")
				case 72080:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72080")
				case 72081:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72081")
				case 72082:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72082")
				case 72083:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72083")
				case 72084:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72084")
				case 72085:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72085")
				case 72086:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72086")
				case 72087:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72087")
				case 72088:
					// value, _ := buff2f["data"].([]int)
					// snum, _ := buff2f["snum"].(int)
					// handleShared.HandleSharedCmdOk(22, value[:8], snum)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72088")
				case 72089:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72089")
				case 72090:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72090")
				case 72091:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72091")
				case 72092:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72092")
				case 72093:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72093")
				case 72094:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72094")
				case 72095:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72095")
				case 72096:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72096")
				case 72097:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72097")
				case 72098:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72098")
				case 72099:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72099")
				case 72100:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72100")
				case 72101:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72101")
				case 72102:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72102")
				case 72103:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72103")
				case 72104:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72104")
				case 72105:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72105")
				case 72106:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72106")
				case 72107:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72107")
				case 72108:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72108")
				case 72109:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72109")
				case 72110:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72110")
				case 72111:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72111")
				case 72112:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72112")
				case 72113:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72113")
				case 72114:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72114")
				case 72115:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72115")
				case 72116:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72116")
				case 72117:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72117")
				case 72118:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72118")
				case 72119:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72119")
				default:
					fmt.Println("---> id 2F not fetched..")
					return
				}
			
			
			}
			// time.Sleep(1*time.Second)
		}
	}()
}

func HandleCmdSender() {
	fmt.Println("---> CmdSender Start.")
}