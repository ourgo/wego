// Copyright 2017 The wego Authors. All rights reserved.

package utils

import (
	"sync"
)

// A Mapper implements the essential behaviors a custom map type.
type Mapper interface {
	Set(interface{}, interface{})
	Get(interface{}) (interface{}, bool)
	Delete(interface{})
	Len() int
	Keys() []interface{}
	Values() []interface{}
}

// SafeMap is a read-write lock synchronized map[interface{}]interface{}.
type SafeMap struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
}

// NewSafeMap returns a pointer to a new SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[interface{}]interface{}),
	}
}

// Set maps the key to the value
func (c *SafeMap) Set(key interface{}, value interface{}) {
	c.lock.Lock()
	c.data[key] = value
	c.lock.Unlock()
}

// Get returns the value associated with the key and whether the key was
// present.
func (c *SafeMap) Get(key interface{}) (interface{}, bool) {
	c.lock.RLock()
	value, present := c.data[key]
	c.lock.RUnlock()
	return value, present
}

// Delete removes the key and associated value.
func (c *SafeMap) Delete(key interface{}) {
	c.lock.Lock()
	delete(c.data, key)
	c.lock.Unlock()
}

// Len reads the length of the map.
func (c *SafeMap) Len() int {
	c.lock.RLock()
	length := len(c.data)
	c.lock.RUnlock()
	return length
}

// Keys returns the slice of keys. Note that the order is not guaranteed.
func (c *SafeMap) Keys() []interface{} {
	c.lock.RLock()
	keys := make([]interface{}, 0, len(c.data))
	for key := range c.data {
		keys = append(keys, key)
	}
	c.lock.RUnlock()
	return keys
}

// Values returns the slice of values. Note that the order is not guaranteed.
func (c *SafeMap) Values() []interface{} {
	c.lock.RLock()
	values := make([]interface{}, 0, len(c.data))
	for _, value := range c.data {
		values = append(values, value)
	}
	c.lock.RUnlock()
	return values
}
