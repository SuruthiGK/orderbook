package DataStore

import (
	"context"
	"fmt"
	"time"
	"orderbook/configs"
)

//Initializing DataStore
func init() {
	fmt.Println("Initializing Data Store ...")
	ctx := context.Background()

	go func() {
		for {
			fmt.Println("DataStore Hydrate Started!")
			InitializeStoreHydrate(ctx)
			fmt.Println("DataStore Hydrated!")
			time.Sleep(configs.Env.DSRefreshTime)
		}
	}()
}

//InitializeStoreHydrate method to hydrate DataStore
func InitializeStoreHydrate(ctx context.Context) {

	wsDepthHandler := func(event *WsDepthEvent) {
	    event.CreateImage()
	}

	errHandler := func(err error) {
	    fmt.Println(err)
	}
	doneC, stopC, err := WsDepthServe("BNBBTC", wsDepthHandler, errHandler)
	if err != nil {
	    fmt.Println(err)
	    return
	}

	// use stopC to exit
	go func() {
	    time.Sleep(5 * time.Second)
	    stopC <- struct{}{}
	}()
	
	_ = <-doneC
	
}


func (event *WsDepthEvent) CreateImage() {
	dStoreImagesCopy := GetDataStoreImagesCopy()
	dStoreImagesCopy.CreateOrUpdate(*event)
	UpdateDataStoreImages(&dStoreImagesCopy)
}
