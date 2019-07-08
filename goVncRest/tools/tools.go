package tools

import (
	"encoding/json"
	"goRootRest/helpers"

	"github.com/valyala/fasthttp"
)

func MakeResponse(statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
}

func MakeDataResponse(data string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody([]byte(data))
}

func MakeRateSourcesResponse(sources []string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	var sourcesMap []map[string]string

	for _, source := range sources {
		//

		var raw map[string]string
		json.Unmarshal([]byte(source), &raw)
		sourcesMap = append(sourcesMap, raw)
	}

	response := helpers.RateSourcesResponse{sourcesMap}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeExchangeTypesResponse(types []string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.ExchangeTypesResponse{types}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeShardsListResponse(shards []string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	var shardsMap []map[string]string

	for _, shard := range shards {
		//

		var raw map[string]string
		json.Unmarshal([]byte(shard), &raw)
		shardsMap = append(shardsMap, raw)
	}

	response := helpers.ShardsListResponse{shardsMap}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeExchangeRateResponse(rate string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	var raw map[string]string
	json.Unmarshal([]byte(rate), &raw)

	response := helpers.ExchangeRateResponse{raw}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeStockTickersResponse(tickers []string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.StockTickersResponse{tickers}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeStockSourcesResponse(sources []string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	var sourcesMap []map[string]string

	for _, source := range sources {
		//

		var raw map[string]string
		json.Unmarshal([]byte(source), &raw)
		sourcesMap = append(sourcesMap, raw)
	}

	response := helpers.StockSourcesResponse{sourcesMap}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeStockPriceResponse(price string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	var raw map[string]string
	json.Unmarshal([]byte(price), &raw)

	response := helpers.StockPriceResponse{raw}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}
