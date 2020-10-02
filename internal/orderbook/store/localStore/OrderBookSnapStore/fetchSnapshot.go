package OrderBookSnapStore

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/url"
	"net/http"
	"time"
	"fmt"
)

// Do method fetches the depth snapshot
func Do(ctx context.Context) (res *DepthResponse, err error) {

	var payload io.Reader

	values := url.Values{}
	values.Add("symbol", "BNBBTC")
	values.Add("limit", "1000")

	header := MakeHeaders(ctx)
	finalUrl := "https://www.binance.com/api/v1/depth?" + values.Encode()

	var response http.Response

	response, err = Request("GET", finalUrl, header,
		5*time.Second, payload)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		err = errors.New(string(body))
		fmt.Println(err)
		return res, err
	}


	j, err := newJSON(body)
	if err != nil {
		return nil, err
	}
	res = new(DepthResponse)

	
	res.LastUpdateID = j.Get("lastUpdateId").MustInt64()
	bidsLen := len(j.Get("bids").MustArray())
	res.Bids = make([]Bid, bidsLen)
	for i := 0; i < bidsLen; i++ {
		item := j.Get("bids").GetIndex(i)
		res.Bids[i] = Bid{
			Price:    item.GetIndex(0).MustString(),
			Quantity: item.GetIndex(1).MustString(),
		}
	}
	asksLen := len(j.Get("asks").MustArray())
	res.Asks = make([]Ask, asksLen)
	for i := 0; i < asksLen; i++ {
		item := j.Get("asks").GetIndex(i)
		res.Asks[i] = Ask{
			Price:    item.GetIndex(0).MustString(),
			Quantity: item.GetIndex(1).MustString(),
		}
	}
	
	return res, nil
}
