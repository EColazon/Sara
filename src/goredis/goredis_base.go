package goredis

import (
	"fmt"
	"time"
	"net/http"
	"runtime"
	"io"
	"log"
	"github.com/garyburd/redigo/redis"
)

//连接池大小
var MaxPoolSize = 20
var redisPool chan redis.Conn

func GoRedisBase() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("---> Connect to redis error", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", "keyName", "Sara")
	if err != nil {
		fmt.Println("---> redis set failed: ", err)
	}
	
	name, err := redis.String(conn.Do("GET", "keyName"))
	if err != nil {
		fmt.Println("---> redis get failed: ", err)
	} else {
		fmt.Printf("---> Get keyName: %v\n", name)
	}

	fmt.Println("---> Go Redis Timeout Test <---")
	_, err = conn.Do("SET", "keyNameEx", "SaraEx", "EX", 5)
	if err != nil {
		fmt.Println("---> redis set failed: ", err)
	}
	time.Sleep(8 * time.Second)	
	name, err = redis.String(conn.Do("GET", "keyNameEx"))
	if err != nil {
		fmt.Println("---> redis get failed: ", err)
	} else {
		fmt.Printf("---> Get keyNameEx: %v\n", name)
	}
}



func PutRedis(conn redis.Conn) {
	//基于函数和接口间互不信任原则
	if redisPool == nil {
		redisPool = make(chan redis.Conn, MaxPoolSize)
	}
	if len(redisPool) >= MaxPoolSize {
		conn.Close()
		return
	}

	redisPool <- conn
}

func InitRedis(network, address string) redis.Conn {
	//缓冲机制,相当于消息队列
	if len(redisPool) == 0 {
		//长度为0,则定义一个redis.Conn类型长度为MaxPoolSize的channel
		redisPool = make(chan redis.Conn, MaxPoolSize)
		
		go func() {
			for i := 0; i < MaxPoolSize/2; i++ {
				c, err := redis.Dial(network, address)
				if err != nil {
					fmt.Println("---> Redis.Dial err: ", err)
					panic(err)
				}
				fmt.Println("---> Init Redis Successed!", i)
				PutRedis(c)
			}
		}()
	}
	return <- redisPool
}

func RedisServer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	c := InitRedis("tcp", "127.0.0.1:6379")
	dbkey := "info"
	if ok, err := redis.Bool(c.Do("LPUSH", dbkey, "Sara")); ok {
		fmt.Println("---> LPUSH OK")
	} else {
		log.Print("---> LPUSH err: ", err)
	}

	msg := fmt.Sprintf("TimeTotal: %s", time.Now().Sub(startTime))
	io.WriteString(w, msg+"\n\n")
}

func HttpRedisStart() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", RedisServer)
	http.ListenAndServe(":32730", nil)
}
