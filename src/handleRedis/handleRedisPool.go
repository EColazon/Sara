package handleRedis
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
//连接池大小
var MaxPoolSize = 20
var redisPool chan redis.Conn

//全局Redis连接游标
var RConn redis.Conn

func init() {
	RConn = InitRedis("tcp", "127.0.0.1:6379")
	fmt.Println("---> RedisPool Inited.")
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