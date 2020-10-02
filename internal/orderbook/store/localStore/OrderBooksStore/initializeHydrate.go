package OrderBooksStore

import (
	"context"
	"fmt"
	"time"
	"orderbook/configs"
	obs "orderbook/internal/orderbook/store/localStore/OrderBookSnapStore"
	ds "orderbook/internal/orderbook/store/localStore/DataStore"
)

//Initializing OrderBooksStore
func init() {
	fmt.Println("Initializing OrderBooks Store ...")
	ctx := context.Background()

	go func() {
		for {
			fmt.Println("Order Books Store Hydrate Started!")
			InitializeStoreHydrate(ctx)
			fmt.Println("Order Books Store Hydrated!")
			time.Sleep(configs.Env.OrderBookRefreshTime)
		}
	}()
}


//InitializeStoreHydrate method hydrates the OrderBooksStore
func InitializeStoreHydrate(ctx context.Context)  {

	dStoreImagesCopy := ds.GetDataStoreImagesCopy() //map[int64]WsDepthEvent
	snapStoreImagesCopy := obs.GetSnapStoreImagesCopy() // map[int64]DepthResponse
	lastUpdateId := snapStoreImagesCopy.GetLastUpdatedID()

	for UpdateID ,data := range dStoreImagesCopy {
		var bidList []Bid
		var askList []Ask
		if UpdateID <= lastUpdateId {
			continue
		}else {
			if (data.FirstUpdateID <= lastUpdateId+1) && (data.UpdateID >= lastUpdateId+1) {
				for _, bid := range data.Bids{
					if bid.Quantity == "0.00000000" {
						bid.Price = "0.00000000"
					}
					bidList = append(bidList,Bid{Price: bid.Price,Quantity: bid.Quantity}) 
				}
				for _, ask := range data.Asks{
					if ask.Quantity == "0.00000000" {
						ask.Price = "0.00000000"
					}
					askList = append(askList,Ask{Price: ask.Price,Quantity: ask.Quantity}) 
				}	
			
			event := &WsDepthEvent{
				Event: data.Event,
				Time: data.Time,
				Symbol: data.Symbol,
				UpdateID: data.UpdateID,
				FirstUpdateID: data.FirstUpdateID,
				Bids: bidList,
				Asks: askList,
			}
			event.CreateImage()
			}
		}
	}
}

func (event *WsDepthEvent) CreateImage() {

	orderBooksStoreCopy := GetOrderBooksStoreCopy()
	orderBooksStoreCopy.CreateOrUpdate(*event)
	UpdateOrderBooksStore(&orderBooksStoreCopy)
	fmt.Println("#################PRINTING ORDERBOOKS FROM IN-MEMORY STORE##################")
	//Test if the data written to local store
	fmt.Println(GetOrderBooksStoreCopy())
	fmt.Println("###########################################################################")
}
