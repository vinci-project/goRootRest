package restServer

import (
	"context"
	"fmt"
	"goRootRest/goVncRest/tools"
	"goRootRest/helpers"
	"net"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/valyala/fasthttp"
)

var redisDB *redis.Client
var mongoDB *mongo.Client

func fastHTTPRawHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == "GET" {
		//

		switch string(ctx.Path()) {

		case "/oracle/exchangeRate":
			//

			args := ctx.QueryArgs()
			for errNum, v := range helpers.RequestExchangeRateFields {
				//

				if !args.Has(v) {
					//

					tools.MakeResponse(errNum, ctx)
					return
				}
			}

			redisCmd := redisDB.Get(fmt.Sprintf("%s:%s:%s", "RATE", string(args.Peek("SOURCE")), string(args.Peek("PAIR"))))
			if helpers.IsRedisError(redisCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeExchangeRateResponse(redisCmd.Val(), helpers.StatusOk, ctx)

		case "/oracle/exchangePairs":
			//

			redisCmd := redisDB.ZRange("RATE PAIRS", 0, -1)
			if helpers.IsRedisError(redisCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeExchangeTypesResponse(redisCmd.Val(), helpers.StatusOk, ctx)

		case "/oracle/exchangeRateSourceList":
			//

			redisCmd := redisDB.ZRange("RATE SOURCES", 0, -1)
			if helpers.IsRedisError(redisCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeRateSourcesResponse(redisCmd.Val(), helpers.StatusOk, ctx)

		case "/oracle/stockPrice":
			//

			args := ctx.QueryArgs()
			for errNum, v := range helpers.RequestStockPriceFields {
				//

				if !args.Has(v) {
					//

					tools.MakeResponse(errNum, ctx)
					return
				}
			}

			redisCmd := redisDB.Get(fmt.Sprintf("%s:%s:%s", "STOCK", string(args.Peek("SOURCE")), string(args.Peek("TICKER"))))
			if helpers.IsRedisError(redisCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeStockPriceResponse(redisCmd.Val(), helpers.StatusOk, ctx)

		case "/oracle/stockTickers":
			//

			redisCmd := redisDB.ZRange("STOCK TICKERS", 0, -1)
			if helpers.IsRedisError(redisCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeStockTickersResponse(redisCmd.Val(), helpers.StatusOk, ctx)

		case "/oracle/stockPriceSourceList":
			//

			redisCmd := redisDB.ZRange("STOCK SOURCES", 0, -1)
			if helpers.IsRedisError(redisCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeStockSourcesResponse(redisCmd.Val(), helpers.StatusOk, ctx)

		case "/shards/shardsList":
			db := mongoDB.Database("rootShard")
			collection := db.Collection("SHARD LIST")
			cur, err := collection.Find(context.Background(), nil)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			defer cur.Close(context.Background())
			filter := bson.NewDocument()
			cursor, err := collection.Find(context.Background(), filter)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			var response []string

			for cursor.Next(context.Background()) {
				//

				result := bson.NewDocument()
				err := cursor.Decode(result)
				if err != nil {
					//

					tools.MakeResponse(helpers.StatusDataNotFound, ctx)
					return
				}

				result.Delete("_id")
				response = append(response, result.ToExtJSON(false))
			}

			tools.MakeShardsListResponse(response, helpers.StatusOk, ctx)

		case "/shards/shardNodes":
			args := ctx.QueryArgs()
			for errNum, v := range helpers.RequestShardNodesFields {
				//

				if !args.Has(v) {
					//

					tools.MakeResponse(errNum, ctx)
					return
				}
			}

			db := mongoDB.Database("rootShard")
			collection := db.Collection("SHARD NODES LIST")
			cur, err := collection.Find(context.Background(), nil)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			defer cur.Close(context.Background())
			result := bson.NewDocument()

			sid, err := strconv.ParseInt(string(args.Peek("SID")), 10, 64)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusWrongAttr_SID, ctx)
				return
			}

			filter := bson.NewDocument(bson.EC.Int64("SHARD ID", sid))
			err = collection.FindOne(context.Background(), filter).Decode(result)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			result.Delete("_id")
			tools.MakeDataResponse(result.ToExtJSON(false), helpers.StatusOk, ctx)

		case "/shards/blockHash":
			args := ctx.QueryArgs()
			for errNum, v := range helpers.RequestBlockHashFields {
				//

				if !args.Has(v) {
					//

					tools.MakeResponse(errNum, ctx)
					return
				}
			}

			db := mongoDB.Database("rootShard")
			collection := db.Collection("SHARD BLOCKCHAIN HASH LIST")
			cur, err := collection.Find(context.Background(), nil)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			defer cur.Close(context.Background())
			result := bson.NewDocument()
			bh, err := strconv.ParseInt(string(args.Peek("BHEIGHT")), 10, 64)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusWrongAttr_BHEIGHT, ctx)
				return
			}

			sid, err := strconv.ParseInt(string(args.Peek("SID")), 10, 64)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusWrongAttr_SID, ctx)
				return
			}

			filter := bson.NewDocument(bson.EC.Int64("BHEIGHT", bh), bson.EC.Int64("SHARD ID", sid))
			err = collection.FindOne(context.Background(), filter).Decode(result)
			if err != nil {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			result.Delete("_id")
			tools.MakeDataResponse(result.ToExtJSON(false), helpers.StatusOk, ctx)

		default:
			//

			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

		return
	}

	ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
}

func Start(r *redis.Client, m *mongo.Client, privateKey []byte, ip string) {
	//

	redisDB = r
	mongoDB = m

	server := &fasthttp.Server{
		Handler:          fastHTTPRawHandler,
		DisableKeepalive: true,
	}

	panic(server.ListenAndServe(net.JoinHostPort(ip, "5000")))
}
