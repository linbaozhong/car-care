package utils

import (
	"sync"
	"time"
)

type Cache struct {
	name    string
	table   sync.Map
	expired time.Duration
}

// 实例化cache，生命期缺省60秒
func NewCache(name string, expired ...time.Duration) *Cache {
	exp := 60 * time.Second
	
	if len(expired) > 0 {
		exp = expired[0]
	}
	cache := &Cache{
		name:    name,
		expired: exp,
	}
	
	return cache
}

func (c *Cache) Set(key interface{}, value interface{}, expired ...time.Duration) {
	exp := c.expired
	if len(expired) > 0 {
		exp = expired[0]
	}
	c.table.Store(key, value)
	
	time.AfterFunc(exp, func() {
		c.Delete(key)
	})
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	if val, ok := c.table.Load(key); ok {
		return val, ok
	}
	return nil, false
}

func (c *Cache) Delete(key interface{}) {
	c.table.Delete(key)
}

//清空cache
func (c *Cache) Empty() {
	c.table.Range(func(key, value interface{}) bool {
		c.table.Delete(key)
		return true
	})
}
