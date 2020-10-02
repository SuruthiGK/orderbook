package OrderBooksStore

import (
	"sync"
)


type OrderBooks map[int64]WsDepthEvent

// store for OrderBooks
type Store struct {
	Locker     sync.RWMutex
	OrderBooksStore *OrderBooks
	isReady    bool
}

var St = Store{
	OrderBooksStore: &OrderBooks{},
	isReady:    true,
}

func GetOrderBooksStore() *OrderBooks {
	St.Locker.RLock()
	defer St.Locker.RUnlock()
	return St.OrderBooksStore
}

func GetOrderBooksStoreCopy() OrderBooks {
	tempDetails := OrderBooks{}
	St.Locker.RLock()
	defer St.Locker.RUnlock()
	for key, value := range *St.OrderBooksStore {
		tempDetails[key] = value
	}
	return tempDetails
}

func UpdateOrderBooksStore(newImages *OrderBooks) {
	St.Locker.Lock()
	St.OrderBooksStore = newImages
	St.Locker.Unlock()
}

func (r OrderBooks) CreateOrUpdate(update WsDepthEvent) {
	r[update.UpdateID] = update
}
