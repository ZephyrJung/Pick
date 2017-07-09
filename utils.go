package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func listLabels(){
	rs,_ := redis.Dial("tcp",redis_host)
	defer rs.Close()
	rs.Do("SET","name","red")
	g,_ := redis.String(rs.Do("GET","name"))
	fmt.Println(g)
}