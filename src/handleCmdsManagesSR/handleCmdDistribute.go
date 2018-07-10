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
)

// 实现两个接口
// HandleCmdGeter:监听channel
// HandleCmdSender:DoZigBeeTasks

// type CmdsDistributer interface {
// 	CmdGeter()
// 	CmdSender()
// }

// // channel数据
// type CmdChannel struct {
// 	id	string
// 	data interface{}
// }

func HandleCmdGeter() {
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
				switch buff33["id"] {
				case 73000:
					fmt.Println("---> id 73000")
				
				case 73001:
					fmt.Println("---> id 73001")
				
				case 73002:
					fmt.Println("---> id 73002")
				
				case 73003:
					fmt.Println("---> id 73003")
				
				case 73004:
					fmt.Println("---> id 73004")
				
				case 73005:
					fmt.Println("---> id 73005")
				
				case 73006:
					fmt.Println("---> id 73006")
				
				case 73007:
					fmt.Println("---> id 73007")
				
				case 73008:
					fmt.Println("---> id 73008")
				
				case 73009:
					fmt.Println("---> id 73009")
				default:
					fmt.Println("---> id 33 not fetched..")
					return
				}
			case buff2f := <- ChCmd2F:
				fmt.Println("---> buff2f: ", len(ChCmd2F), cap(ChCmd2F), buff2f, time.Now())
				switch buff2f["id"] {
				case 72000:
					fmt.Println("---> id 72000")
				case 72001:
					fmt.Println("---> id 72001")
				case 72002:
					fmt.Println("---> id 72002")
				case 72003:
					fmt.Println("---> id 72003")
				case 72004:
					fmt.Println("---> id 72004")
				case 72005:
					fmt.Println("---> id 72005")
				case 72006:
					fmt.Println("---> id 72006")
				case 72007:
					fmt.Println("---> id 72007")
				case 72008:
					fmt.Println("---> id 72008")
				case 72009:
					fmt.Println("---> id 72009")
				case 720010:
					fmt.Println("---> id 72010")
				case 72011:
					fmt.Println("---> id 72011")
				case 72012:
					fmt.Println("---> id 72012")
				case 72013:
					fmt.Println("---> id 72013")
				case 72014:
					fmt.Println("---> id 72014")
				case 72015:
					fmt.Println("---> id 72015")
				case 72016:
					fmt.Println("---> id 72016")
				case 72017:
					fmt.Println("---> id 72017")
				case 72018:
					fmt.Println("---> id 72018")
				case 72019:
					fmt.Println("---> id 72019")
				case 72020:
					fmt.Println("---> id 72020")
				case 72021:
					fmt.Println("---> id 72021")
				case 72022:
					fmt.Println("---> id 72022")
				case 72023:
					fmt.Println("---> id 72023")
				case 72024:
					fmt.Println("---> id 72024")
				case 72025:
					fmt.Println("---> id 72025")
				case 72026:
					fmt.Println("---> id 72026")
				case 72027:
					fmt.Println("---> id 72027")
				case 72028:
					fmt.Println("---> id 72028")
				case 72029:
					fmt.Println("---> id 72029")
				case 72030:
					fmt.Println("---> id 72030")
				case 72031:
					fmt.Println("---> id 72031")
				case 72032:
					fmt.Println("---> id 72032")
				case 72033:
					fmt.Println("---> id 72033")
				case 72034:
					fmt.Println("---> id 72034")
				case 72035:
					fmt.Println("---> id 72035")
				case 72036:
					fmt.Println("---> id 72036")
				case 72037:
					fmt.Println("---> id 72037")
				case 72038:
					fmt.Println("---> id 72038")
				case 72039:
					fmt.Println("---> id 72039")
				case 72040:
					fmt.Println("---> id 72040")
				case 72041:
					fmt.Println("---> id 72041")
				case 72042:
					fmt.Println("---> id 72042")
				case 72043:
					fmt.Println("---> id 72043")
				case 72044:
					fmt.Println("---> id 72044")
				case 72045:
					fmt.Println("---> id 72045")
				case 72046:
					fmt.Println("---> id 72046")
				case 72047:
					fmt.Println("---> id 72047")
				case 72048:
					fmt.Println("---> id 72048")
				case 72049:
					fmt.Println("---> id 72049")
				case 72050:
					fmt.Println("---> id 72050")
				case 72051:
					fmt.Println("---> id 72051")
				case 72052:
					fmt.Println("---> id 72052")
				case 72053:
					fmt.Println("---> id 72053")
				case 72054:
					fmt.Println("---> id 72054")
				case 72055:
					fmt.Println("---> id 72055")
				case 72056:
					fmt.Println("---> id 72056")
				case 72057:
					fmt.Println("---> id 72057")
				case 72058:
					fmt.Println("---> id 72058")
				case 72059:
					fmt.Println("---> id 72059")
				case 72060:
					fmt.Println("---> id 72060")
				case 72061:
					fmt.Println("---> id 72061")
				case 72062:
					fmt.Println("---> id 72062")
				case 72063:
					fmt.Println("---> id 72063")
				case 72064:
					fmt.Println("---> id 72064")
				case 72065:
					fmt.Println("---> id 72065")
				case 72066:
					fmt.Println("---> id 72066")
				case 72067:
					fmt.Println("---> id 72067")
				case 72068:
					fmt.Println("---> id 72068")
				case 72069:
					fmt.Println("---> id 72069")
				case 72070:
					fmt.Println("---> id 72070")
				case 72071:
					fmt.Println("---> id 72071")
				case 72072:
					fmt.Println("---> id 72072")
				case 72073:
					fmt.Println("---> id 72073")
				case 72074:
					fmt.Println("---> id 72074")
				case 72075:
					fmt.Println("---> id 72075")
				case 72076:
					fmt.Println("---> id 72076")
				case 72077:
					fmt.Println("---> id 72077")
				case 72078:
					fmt.Println("---> id 72078")
				case 72079:
					fmt.Println("---> id 72079")
				case 72080:
					fmt.Println("---> id 72080")
				case 72081:
					fmt.Println("---> id 72081")
				case 72082:
					fmt.Println("---> id 72082")
				case 72083:
					fmt.Println("---> id 72083")
				case 72084:
					fmt.Println("---> id 72084")
				case 72085:
					fmt.Println("---> id 72085")
				case 72086:
					fmt.Println("---> id 72086")
				case 72087:
					fmt.Println("---> id 72087")
				case 72088:
					fmt.Println("---> id 72088")
				case 72089:
					fmt.Println("---> id 72089")
				case 72090:
					fmt.Println("---> id 72090")
				case 72091:
					fmt.Println("---> id 72091")
				case 72092:
					fmt.Println("---> id 72092")
				case 72093:
					fmt.Println("---> id 72093")
				case 72094:
					fmt.Println("---> id 72094")
				case 72095:
					fmt.Println("---> id 72095")
				case 72096:
					fmt.Println("---> id 72096")
				case 72097:
					fmt.Println("---> id 72097")
				case 72098:
					fmt.Println("---> id 72098")
				case 72099:
					fmt.Println("---> id 72099")
				case 72100:
					fmt.Println("---> id 72100")
				case 72101:
					fmt.Println("---> id 72101")
				case 72102:
					fmt.Println("---> id 72102")
				case 72103:
					fmt.Println("---> id 72103")
				case 72104:
					fmt.Println("---> id 72104")
				case 72105:
					fmt.Println("---> id 72105")
				case 72106:
					fmt.Println("---> id 72106")
				case 72107:
					fmt.Println("---> id 72107")
				case 72108:
					fmt.Println("---> id 72108")
				case 72109:
					fmt.Println("---> id 72109")
				case 72110:
					fmt.Println("---> id 72110")
				case 72111:
					fmt.Println("---> id 72111")
				case 72112:
					fmt.Println("---> id 72112")
				case 72113:
					fmt.Println("---> id 72113")
				case 72114:
					fmt.Println("---> id 72114")
				case 72115:
					fmt.Println("---> id 72115")
				case 72116:
					fmt.Println("---> id 72116")
				case 72117:
					fmt.Println("---> id 72117")
				case 72118:
					fmt.Println("---> id 72118")
				case 72119:
					fmt.Println("---> id 72119")
				}
			
			
			}
			// time.Sleep(1*time.Second)
		}
	}()
}

func HandleCmdSender() {
	fmt.Println("---> CmdSender Start.")
}