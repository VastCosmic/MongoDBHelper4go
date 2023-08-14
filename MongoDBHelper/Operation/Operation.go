package Operation

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var mu sync.Mutex // 互斥锁，保护对 client 的访问

// InsertOne 插入一条数据
func InsertOne(client *mongo.Client, dbName string, collectionName string, data interface{}) (*mongo.InsertOneResult, error) {
	mu.Lock()
	defer mu.Unlock()

	col := client.Database(dbName).Collection(collectionName)

	insertResult, err := col.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("InsertOne failed, err: ", err)
		return nil, err
	}

	fmt.Printf("Inserted ID: %v\n", insertResult.InsertedID)
	return insertResult, nil
}

// InsertMany 批量插入数据
func InsertMany(client *mongo.Client, dbName string, collectionName string, data []interface{}) (*mongo.InsertManyResult, error) {
	mu.Lock()
	defer mu.Unlock()

	col := client.Database(dbName).Collection(collectionName)

	insertResult, err := col.InsertMany(context.Background(), data)
	if err != nil {
		fmt.Println("BatchInsert failed, err: ", err)
		return nil, err
	}

	fmt.Printf("Inserted IDs: %v\n", insertResult.InsertedIDs)
	return insertResult, nil
}
