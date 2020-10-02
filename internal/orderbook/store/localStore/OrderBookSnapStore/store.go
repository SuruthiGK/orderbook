//OrderBookSnapStore local store to save the events fetched by hitting depth-snapshot API 

package OrderBookSnapStore

import (
	"sync"
)

type DepthSnapImage map[int64]DepthResponse

// store for OrderBookSnapStore
type Store struct {
	Locker     				 sync.RWMutex
	OrderBookSnapStoreImages *DepthSnapImage
	isReady    				 bool
}

var St = Store{
	OrderBookSnapStoreImages: &DepthSnapImage{},
	isReady:    true,
}

func GetSnapStoreImages() *DepthSnapImage {
	St.Locker.RLock()
	defer St.Locker.RUnlock()
	return St.OrderBookSnapStoreImages
}

func GetSnapStoreImagesCopy() DepthSnapImage {
	tempDetails := DepthSnapImage{}
	St.Locker.RLock()
	defer St.Locker.RUnlock()
	for key, value := range *St.OrderBookSnapStoreImages {
		tempDetails[key] = value
	}
	return tempDetails
}

func UpdateSnapStoreImages(newImages *DepthSnapImage) {
	St.Locker.Lock()
	St.OrderBookSnapStoreImages = newImages
	St.Locker.Unlock()
}

func (r DepthSnapImage) CreateOrUpdate(update DepthResponse) {
	r[update.LastUpdateID] = update
}

func (r DepthSnapImage) GetLastUpdatedID() int64{
	var LastUpdateID int64
	for _,v:= range r {
		LastUpdateID =  v.LastUpdateID
	}
	return LastUpdateID
	
}
