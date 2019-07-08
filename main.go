package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"flag"
	"goRootRest/goVncRest"
	"goRootRest/helpers"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var redisDB *redis.Client
var mongoDB *mongo.Client

func connectToMongo() (err error) {
	//

	mongoDB, err = mongo.Connect(context.Background(), "mongodb://192.168.192.42:27017", nil)
	if err != nil {
		//

		log.Fatal(err)
	}

	err = mongoDB.Ping(context.Background(), nil)
	if err != nil {
		//

		log.Fatalln("No connection to MONGO. ", err)
	}

	return
}

func connectToRedis() (err error) {
	//

	var redisDBNumInt int = 1
	redisHost, ok := os.LookupEnv("REDIS_PORT_6379_TCP_ADDR")
	if !ok {
		//

		redisHost = "0.0.0.0"
	}

	redisPort, ok := os.LookupEnv("REDIS_PORT_6379_TCP_PORT")
	if !ok {
		//

		redisPort = "6379"
	}

	redisDBNum, ok := os.LookupEnv("REDIS_PORT_6379_DB_NUM")
	if !ok {
		//

		redisDBNumInt = 2

	} else {
		//

		if redisDBNumInt64, err := strconv.ParseInt(redisDBNum, 10, 64); err != nil {
			//

			redisDBNumInt = int(redisDBNumInt64)

		} else {
			//

			redisDBNumInt = 1
		}
	}

	redisDB = redis.NewClient(&redis.Options{
		Addr:         net.JoinHostPort(redisHost, redisPort),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		DB:           redisDBNumInt,
	})

	statusCmd := redisDB.Ping()
	if helpers.IsRedisError(statusCmd) {
		//

		log.Fatalln("No connection to REDIS. ", statusCmd.Err())
		return statusCmd.Err()
	}

	return
}

func main() {
	//

	privateKeyPtr := flag.String("privateKey",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"A private key associated with your node")

	ipPtr := flag.String("ip",
		"0.0.0.0",
		"Your external IP")

	flag.Parse()

	privateKey, err := hex.DecodeString(*privateKeyPtr)
	if err != nil {
		//

		panic("WRONG PRIVATE KEY")
	}

	if err := connectToRedis(); err != nil {
		//

		panic(err.Error())
	}

	defer redisDB.Close()

	if err := connectToMongo(); err != nil {
		//

		panic(err.Error())
	}

	defer mongoDB.Disconnect(context.Background())

	go restServer.Start(redisDB, mongoDB, privateKey, *ipPtr)

	consoleReader := bufio.NewReader(os.Stdin)
	log.Println("Enter text: ") // Just for better testing. Should be refactored for production
	consoleReader.ReadString('\n')
}
