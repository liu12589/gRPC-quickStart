package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"sync"
	"useCpu/pb"
)

var ErrAlreadyExists = errors.New("record already exists")
//LaptopStore 是一个用于存储电脑信息的接口
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

//InMemoryLaptopStore 将电脑信息保存到内存中
type InMemoryLaptopStore struct {
	//使用读写互斥锁来处理并发问题
	mutex sync.RWMutex
	data map[string]*pb.Laptop
}

//NewInMemoryLaptopStore 返回一个新创建的 InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return  nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other,nil
}

func (store *InMemoryLaptopStore) Save (laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// 这里存在一个深拷贝问题，将laptop的值复制给other，然后使用store进行存储，store实际上是浅拷贝问题
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil{
		return fmt.Errorf("cannot copy laptop data: %w",err)
	}

	store.data[other.Id] = other
	return nil
}