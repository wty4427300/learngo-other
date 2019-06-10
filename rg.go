package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var redisClient *redis.Pool
func init() {
	maxIdle := 10
	maxActive := 100

	// 建立连接池
	redisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 2 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", "127.0.0.1:6379",
				redis.DialConnectTimeout(2*time.Second),
				redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(2*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}
func main() {
	rc := redisClient.Get()
	// 用完后将连接放回连接池
	_, err := rc.Do("SET", "go_key", "hanhan")
	if err != nil {
		fmt.Println("err while setting:", err)
	}
	a, err1 := redis.String(rc.Do("get", "go_key"))
	if err1 != nil {
		fmt.Println("err while setting:", err)
	}else {
		fmt.Println(a)
	}
	//判断key是否存在
	is_key_exit, err := redis.Bool(rc.Do("EXISTS", "go_key"))
	if err != nil {
		fmt.Println("err while checking keys:", err)
	} else {
		fmt.Println(is_key_exit)
	}
	defer rc.Close()
	//删除key
	_, err = rc.Do("del", "go_key")
	if err != nil {
		fmt.Println("err while deleting:", err)
	}
	//设置key失效时间
	_,err =rc.Do("set","mykey","shabi","ex","5")
	if err != nil {
		fmt.Println("err while deleting:", err)
	}
	a1,err :=redis.String(rc.Do("get","mykey"))
	if err != nil {
		fmt.Println("err while deleting:", err)
	}else {
		fmt.Println(a1)
	}
	time.Sleep(10*time.Second)
	a2,err := redis.String(rc.Do("get","mykey"))
	if err != nil {
		fmt.Println("err while deleting:", err)
	}else {
		fmt.Println(a2)
	}
    //对已有key设置5秒过期时间
    n,err:=rc.Do("expire","go_key","5")
    if err!=nil{
    	fmt.Println("err",err)
	}else if n!=int64(1){
		fmt.Println("failed")
	}
	err=rc.Send("subscribe","__keyevent@0__:expired")

}