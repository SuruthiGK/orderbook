package OrderBookSnapStore

import (
	"context"
	"fmt"
	"time"
	"orderbook/configs"
)

//Initializing OrderBookSnapStore
func init() {
	fmt.Println("Initializing OrderBookSnap Store ...")
	ctx := context.Background()

	go func() {
		for {
			fmt.Println("OrderBookSnapStore Hydrate Started!")
			InitializeStoreHydrate(ctx)
			fmt.Println("OrderBookSnapStore Hydrated!")
			time.Sleep(configs.Env.OBSRefreshTime)
		}
	}()
}

//InitializeStoreHydrate method hydrates the OrderBookSnapStore
func InitializeStoreHydrate(ctx context.Context) {
	event, err := Do(ctx)	
	if err != nil {
	    fmt.Println(err)
	    return
	}
	event.CreateImage()
}

func (event *DepthResponse) CreateImage() {
	snapStoreImagesCopy := GetSnapStoreImagesCopy()
	snapStoreImagesCopy.CreateOrUpdate(*event)
	UpdateSnapStoreImages(&snapStoreImagesCopy)
}
