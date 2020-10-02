//DataStore local store to save the events fetched from data stream 

package DataStore

import (
	"sync"
)

type DepthEventImage map[int64]WsDepthEvent

// store for DataStore
type Store struct {
	Locker     sync.RWMutex
	DataStoreImages *DepthEventImage
	isReady    bool
}

var St = Store{
	DataStoreImages: &DepthEventImage{},
	isReady:    true,
}

func GetDataStoreImages() *DepthEventImage {
	St.Locker.RLock()
	defer St.Locker.RUnlock()
	return St.DataStoreImages
}

func GetDataStoreImagesCopy() DepthEventImage {
	tempDetails := DepthEventImage{}
	St.Locker.RLock()
	defer St.Locker.RUnlock()
	for key, value := range *St.DataStoreImages {
		tempDetails[key] = value
	}
	return tempDetails
}

func UpdateDataStoreImages(newImages *DepthEventImage) {
	St.Locker.Lock()
	St.DataStoreImages = newImages
	St.Locker.Unlock()
}

func (r DepthEventImage) CreateOrUpdate(update WsDepthEvent) {
	r[update.UpdateID] = update
}
