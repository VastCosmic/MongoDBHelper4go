package main

import (
	"MongoDB/MongoDBHelper/Connection"
	"MongoDB/MongoDBHelper/Entity"
	"MongoDB/MongoDBHelper/Operation"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	// 初始化连接池
	Connection.ConnectWithPoolSize(1000, 10)
	defer Connection.Disconnect()

	// 创建一个 WaitGroup，用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 创建一个切片，用于存储要插入的数据
	var dataToInsert []interface{}

	// 启动多个 goroutine 并发地生成数据
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 随机生成数据
			var name, gender, creator, updater string
			var age int
			err := faker.FakeData(&name)
			if err != nil {
				fmt.Println("Error generating fake data:", err)
				return
			}
			err = faker.FakeData(&age)
			if err != nil {
				fmt.Println("Error generating fake data:", err)
				return
			}

			// 生成随机数 0 或 1，将其转换为 "man" 或 "female"
			randomNumber := rand.Intn(2) // 生成0或1的随机整数
			if randomNumber == 0 {
				gender = "man"
			} else {
				gender = "female"
			}

			err = faker.FakeData(&creator)
			if err != nil {
				fmt.Println("Error generating fake data:", err)
				return
			}
			err = faker.FakeData(&updater)
			if err != nil {
				fmt.Println("Error generating fake data:", err)
				return
			}

			// 创建一个 Stu 结构体实例
			stu := Entity.Stu{
				BaseObj: Entity.BaseObj{
					Creator:    strings.ToTitle(strings.ToLower(creator[:5])),
					Updater:    strings.ToTitle(strings.ToLower(updater[:5])),
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				},
				Name:   strings.ToTitle(strings.ToLower(name[:5])),
				Age:    age,
				Gender: gender,
			}

			// 将数据添加到批量插入的切片中
			dataToInsert = append(dataToInsert, stu)
		}()
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 批量插入数据
	_, err := Operation.InsertMany(Connection.GetClient(), "batchTest", "data", dataToInsert)
	if err != nil {
		fmt.Println("Error inserting data:", err)
	} else {
		fmt.Println("BatchInsert success.")
	}
}
