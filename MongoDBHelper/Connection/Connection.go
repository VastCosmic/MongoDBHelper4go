package Connection

import (
	"MongoDB/MongoDBHelper/Config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var (
	client *mongo.Client
	mu     sync.Mutex // 互斥锁，保护对 client 的访问
)

func Connect() {
	configInstance := Config.GetInstance()
	host := configInstance.GetHost()
	port := configInstance.GetPort()

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(100) // 设置连接池最大连接数
	clientOptions.SetMinPoolSize(10)  // 设置连接池最小闲置连接数

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Connect failed, err: ", err)
	}
}

// ConnectWithPoolSize 初始化数据库连接，设置连接池大小
func ConnectWithPoolSize(maxPoolSize uint64, minPoolSize uint64) {
	configInstance := Config.GetInstance()
	host := configInstance.GetHost()
	port := configInstance.GetPort()

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(maxPoolSize) // 设置连接池最大连接数
	clientOptions.SetMinPoolSize(minPoolSize) // 设置连接池最小闲置连接数

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Connect failed, err: ", err)
	}
}

func Disconnect() {
	mu.Lock()
	defer mu.Unlock()

	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Println("Disconnect failed, err: ", err)
		}
	}
}

func GetClient() *mongo.Client {
	mu.Lock()
	defer mu.Unlock()

	if client == nil {
		Connect()
	}
	return client
}
