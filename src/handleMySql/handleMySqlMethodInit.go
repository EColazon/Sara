package handleMySql

//数据库连接池测试
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "strings"
	
)

//数据库配置
const (
    userName = "root"
    password = "hello520Sara"
    ip = "127.0.0.1"
    port = "3306"
    dbName = "dbtest"
)
// 数据库读写异常报警
var (
	AlarmBuffDB = []int{0x33, 0x01, 0x10, 0x02, 0x00, 0x06, 0x00, 0x01, 0xD0, 0x00, 0x00, 0x00, 0x32, 0x99}
)
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func init()  {
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

    //打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
    DB, _ = sql.Open("mysql", path)
    //设置数据库最大连接数
    DB.SetConnMaxLifetime(100)
    //设置上数据库最大闲置连接数
    DB.SetMaxIdleConns(10)
    //验证连接
    if err := DB.Ping(); err != nil{
        fmt.Println("opon database fail")
        return
    }
    fmt.Println("connnect success")
}

//创建表
func HandleDBCreateTable() {
	
	//创建电参量表
	sqlDBElec := "CREATE TABLE IF NOT EXISTS dbelec(" +
		"uid int AUTO_INCREMENT NOT NULL ," +
		"num int NOT NULL ," +
		"current float NOT NULL ," +
		"volt float NOT NULL ," +
		"pf float NOT NULL ," +
		"power float NOT NULL ," +
		"energy float NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
		") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBElecdrop := "DROP TABLE dbelec"
	DB.Exec(sqlDBElecdrop)	
	fmt.Println("---> sql: 01 ")
	DB.Exec(sqlDBElec)

	//创建节点表
	sqlDBNode := "CREATE TABLE IF NOT EXISTS dbnode(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"stateNode INT NOT NULL ," +
		"timeTrigeH INT NOT NULL ," +
		"timeTrighM INT NOT NULL ," +
		"stateLight INT NOT NULL ," +
		"timeRecoverH INT NOT NULL ," +
		"timeRecoverM INT NOT NULL ," +
		"timeOpenH INT NOT NULL ," +
		"timeOpenM INT NOT NULL ," +
		"timeCloseH INT NOT NULL ," +
		"timeCloseM INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBNodedrop := "DROP TABLE dbinode"
	DB.Exec(sqlDBNodedrop)	
	DB.Exec(sqlDBNode)
	fmt.Println("---> sql: 02 ")
	
	//连续三个相同
	//创建单灯阶段调光表
	sqlDBStagePwm := "CREATE TABLE IF NOT EXISTS dbstagepwm(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"t01 INT NOT NULL ," +
		"t02 INT NOT NULL ," +
		"t03 INT NOT NULL ," +
		"t04 INT NOT NULL ," +
		"t05 INT NOT NULL ," +
		"t06 INT NOT NULL ," +
		"t07 INT NOT NULL ," +
		"t08 INT NOT NULL ," +
		"t09 INT NOT NULL ," +
		"t10 INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBStagePwmdrop := "DROP TABLE dbstagepwm"
	DB.Exec(sqlDBStagePwmdrop)	
	fmt.Println("---> sql: 03 ")
	DB.Exec(sqlDBStagePwm)
	//创建特殊时间策略表
	sqlDBSpecial := "CREATE TABLE IF NOT EXISTS dbspecial(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"t01 INT NOT NULL ," +
		"t02 INT NOT NULL ," +
		"t03 INT NOT NULL ," +
		"t04 INT NOT NULL ," +
		"t05 INT NOT NULL ," +
		"t06 INT NOT NULL ," +
		"t07 INT NOT NULL ," +
		"t08 INT NOT NULL ," +
		"t09 INT NOT NULL ," +
		"t10 INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSpecialdrop := "DROP TABLE dbspecial"
	DB.Exec(sqlDBSpecialdrop)	
	fmt.Println("---> sql: 04 ")
	DB.Exec(sqlDBSpecial)
	//创建节假日时间策略表
	sqlDBHoliday := "CREATE TABLE IF NOT EXISTS dbholiday(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"t01 INT NOT NULL ," +
		"t02 INT NOT NULL ," +
		"t03 INT NOT NULL ," +
		"t04 INT NOT NULL ," +
		"t05 INT NOT NULL ," +
		"t06 INT NOT NULL ," +
		"t07 INT NOT NULL ," +
		"t08 INT NOT NULL ," +
		"t09 INT NOT NULL ," +
		"t10 INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBHolidaydrop := "DROP TABLE dbholiday"
	DB.Exec(sqlDBHolidaydrop)	
	fmt.Println("---> sql: 05 ")
	DB.Exec(sqlDBHoliday)
	
	//创建单灯pwm表
	sqlDBSlPwm := "CREATE TABLE IF NOT EXISTS dbslpwm(" +
		"uid int AUTO_INCREMENT NOT NULL ," +
		"num int NOT NULL ," +
		"t01 float NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
		") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlPwmdrop := "DROP TABLE dbslpwm"
	DB.Exec(sqlDBSlPwmdrop)	
	fmt.Println("---> sql: 06 ")
	DB.Exec(sqlDBSlPwm)

	//创建单灯定时开关时间表
	sqlDBSlTimer := "CREATE TABLE IF NOT EXISTS dbsltimer(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampMOpen INT NOT NULL ," +
		"lampMClose INT NOT NULL ," +
		"lampFOpen INT NOT NULL ," +
		"lampFClose INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlTimerdrop := "DROP TABLE dbsltimer"
	DB.Exec(sqlDBSlTimerdrop)	
	fmt.Println("---> sql: 07 ")
	DB.Exec(sqlDBSlTimer)
	
	// 连续四个相同结构表
	//创建单灯组类型表
	sqlDBSlGType := "CREATE TABLE IF NOT EXISTS dbslgtype(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampGType INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlGTypeDrop := "DROP TABLE dbslgtype"
	DB.Exec(sqlDBSlGTypeDrop)	
	fmt.Println("---> sql: 08 ")
	DB.Exec(sqlDBSlGType)
	//创建单灯组号表
	sqlDBSlGNum := "CREATE TABLE IF NOT EXISTS dbslgnum(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampGNum INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlGNumDrop := "DROP TABLE dbslgnum"
	DB.Exec(sqlDBSlGNumDrop)	
	fmt.Println("---> sql: 09 ")
	DB.Exec(sqlDBSlGNum)
	//创建单灯继电器状态表
	sqlDBSlRelayST := "CREATE TABLE IF NOT EXISTS dbslrelayst(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampRelayST INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlRelaySTDrop := "DROP TABLE dbslrelayst"
	DB.Exec(sqlDBSlRelaySTDrop)	
	fmt.Println("---> sql: 10 ")
	DB.Exec(sqlDBSlRelayST)
	//创建单灯手动开关状态表
	sqlDBSlHandST := "CREATE TABLE IF NOT EXISTS dbslhandst(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"num INT NOT NULL ," +
		"lampHandST INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlHandSTDrop := "DROP TABLE dbslhandst"
	DB.Exec(sqlDBSlHandSTDrop)	
	fmt.Println("---> sql: 11 ")
	DB.Exec(sqlDBSlHandST)

	//创建单灯数据表
	sqlDBSlLamp := "CREATE TABLE IF NOT EXISTS dblamp(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"lampNum INT NOT NULL ," +
		"lampNumGroup INT NOT NULL ," +
		"lAdvV INT NOT NULL ," +
		"lAdvI INT NOT NULL ," +
		"lAdvP INT NOT NULL ," +
		"lAdvPF INT NOT NULL ," +
		"lAuxV INT NOT NULL ," +
		"lAuxI INT NOT NULL ," +
		"lAuxP INT NOT NULL ," +
		"lAuxPF INT NOT NULL ," +
		"lAdvPwm INT NOT NULL ," +
		"lAuxPwm INT NOT NULL ," +
		"lStateRelayBT INT NOT NULL ," +
		"lModeTX INT NOT NULL ," +
		"lModeRX INT NOT NULL ," +
		"lTimeAlarm INT NOT NULL ," +
		"lRelayChange INT NOT NULL ," +
		"lFlagEX INT NOT NULL ," +
		"lStateAlarm INT NOT NULL ," +
		"lBNetAddr INT NOT NULL ," +
		"lAdvPower INT NOT NULL ," +
		"lAuxPower INT NOT NULL ," +
		"lampHigherV INT NOT NULL ," +
		"lampLowerV INT NOT NULL ," +
		"lampHigherI INT NOT NULL ," +
		"lampLowerI INT NOT NULL ," +
		"lampHigherP INT NOT NULL ," +
		"lampLowerP INT NOT NULL ," +
		"lampHigherPF INT NOT NULL ," +
		"lampLowerPF INT NOT NULL ," +
		"lFlagSetNum INT NOT NULL ," +
		"lFlagSetAdu INT NOT NULL ," +
		"lChecksum INT NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBSlLampDrop := "DROP TABLE dblamp"
	DB.Exec(sqlDBSlLampDrop)	
	fmt.Println("---> sql: 12 ")
	DB.Exec(sqlDBSlLamp)


	//创建命令解析异常记录表
	sqlDBLogCmdErr := "CREATE TABLE IF NOT EXISTS dblogcmderr(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogCmdErrDrop := "DROP TABLE dblogcmderr"
	DB.Exec(sqlDBLogCmdErrDrop)	
	fmt.Println("---> sql: 08 ")
	DB.Exec(sqlDBLogCmdErr)

	// 创建命令解析正常记录表 server->RTU
	sqlDBLogCmd00 := "CREATE TABLE IF NOT EXISTS dblogcmdokd0(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogCmd00Drop := "DROP TABLE dblogcmdokd0"
	DB.Exec(sqlDBLogCmd00Drop)	
	fmt.Println("---> sql: 09 ")
	DB.Exec(sqlDBLogCmd00)

	// 创建命令解析正常记录表 RTU->router
	sqlDBLogCmd01 := "CREATE TABLE IF NOT EXISTS dblogcmdokd1(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogCmd01Drop := "DROP TABLE dblogcmdokd1"
	DB.Exec(sqlDBLogCmd01Drop)	
	fmt.Println("---> sql: 10 ")
	DB.Exec(sqlDBLogCmd01)

	// 创建命令解析正常记录表 router->RTU
	sqlDBLogCmd02 := "CREATE TABLE IF NOT EXISTS dblogcmdokd2(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogCmd02Drop := "DROP TABLE dblogcmdokd2"
	DB.Exec(sqlDBLogCmd02Drop)	
	fmt.Println("---> sql: 11 ")
	DB.Exec(sqlDBLogCmd02)

	// 创建命令解析正常记录表 RTU->server
	sqlDBLogCmd03 := "CREATE TABLE IF NOT EXISTS dblogcmdokd3(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogCmd03Drop := "DROP TABLE dblogcmdokd3"
	DB.Exec(sqlDBLogCmd03Drop)	
	fmt.Println("---> sql: 12 ")
	DB.Exec(sqlDBLogCmd03)

	// 创建命令解析正常记录表 RTUDect
	sqlDBLogCmd04 := "CREATE TABLE IF NOT EXISTS dblogcmdokd4(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogCmd04Drop := "DROP TABLE dblogcmdokd4"
	DB.Exec(sqlDBLogCmd04Drop)	
	fmt.Println("---> sql: 13 ")
	DB.Exec(sqlDBLogCmd04)

	// 新建RTU&单灯异常报警表
	sqlDBLogAlarmE := "CREATE TABLE IF NOT EXISTS dblogalarme(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogAlarmEDrop := "DROP TABLE dblogalarme"
	DB.Exec(sqlDBLogAlarmEDrop)	
	fmt.Println("---> sql: 14 ")
	DB.Exec(sqlDBLogAlarmE)

	// 新建心跳包记录表
	sqlDBLogHeartA0 := "CREATE TABLE IF NOT EXISTS dbloghearta0(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogHeartA0Drop := "DROP TABLE dbloghearta0"
	DB.Exec(sqlDBLogHeartA0Drop)	
	fmt.Println("---> sql: 15 ")
	DB.Exec(sqlDBLogHeartA0)

	sqlDBLogHeartA1 := "CREATE TABLE IF NOT EXISTS dbloghearta1(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogHeartA1Drop := "DROP TABLE dbloghearta1"
	DB.Exec(sqlDBLogHeartA1Drop)	
	fmt.Println("---> sql: 16 ")
	DB.Exec(sqlDBLogHeartA1)

	// 新建系统参数记录表
	sqlDBLogSysB := "CREATE TABLE IF NOT EXISTS dblogsysb(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogSysBDrop := "DROP TABLE dblogsysb"
	DB.Exec(sqlDBLogSysBDrop)	
	fmt.Println("---> sql: 17 ")
	DB.Exec(sqlDBLogSysB)

	// 新建程序模块异常记录表
	sqlDBLogModuleC := "CREATE TABLE IF NOT EXISTS dblogmodulec(" +
		"uid INT AUTO_INCREMENT NOT NULL ," +
		"flag INT NOT NULL ," +
		"content CHAR(73) NOT NULL ," +
		// "update_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		// "create_time datetime NOT NULL  DEFAULT CURRENT_TIMESTAMP," +
		"update_time TIMESTAMP," +
		"PRIMARY KEY (uid)" +
	  ") engine=innodb DEFAULT charset=utf8mb4;"
	sqlDBLogModuleCDrop := "DROP TABLE dblogmodulec"
	DB.Exec(sqlDBLogModuleCDrop)	
	fmt.Println("---> sql: 18 ")
	DB.Exec(sqlDBLogModuleC)
}