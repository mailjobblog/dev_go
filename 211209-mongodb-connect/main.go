package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	// 连接 mongodb
	// 注意：这里的账号和密码要改成你自己的
	clientOptions := options.Client().ApplyURI("mongodb+srv://test_user:test_pwd@cluster0.sy0un.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("mongodb 连接失败，error：", err)
	}

	// 定义连接库和连接集合
	collection := client.Database("t_db").Collection("t_coll")

	// 定义要插入的数据
	user := UserData{
		Name:       "张三",
		Age:        20,
	}

	// 测试插入数据
	insert,err := collection.InsertOne(ctx, &user)
	if err != nil {
		log.Fatal("mongodb 数据插入失败，error：", err)
	}

	// 打印执行结果的id
	fmt.Println(insert)
}

// UserData 定义插入数据的结构体
type UserData struct {
	// Id         string `bson:"_id,omitempty" json:"id"` // 这里不设置id，让数据库自动生成
	Name       string `bson:"name" json:"name"`
	Age        int    `bson:"age" json:"age"`
}
