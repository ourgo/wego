// Copyright 2017 The wego Authors. All rights reserved.
//wegoSafeMap is implemented by executing an unexpected method that captures
// a map inside a goroutine .The only way to access the map is via channel,and
// this in itself sufficient to ensure that all map accesses are serialized
//
// 注 : wegoSafeMap : 键是字符串,值是interface{}类型.(by SergeyJay)
//		当存入的值是指针或引用,需保证指向的值是read-only属性,或对它们的访问是串行的.


package utils

type WegoSafeMapper interface {
	Insert(string,interface{})
	Delete(string)
	Find(string) (interface{},bool)
	Len()int
	Update(string,UpdateFunc)
	Close()map[string]interface{}
}

type wegoSafeMap chan commandData

type UpdateFunc func(interface{},bool) interface{}

type commandData struct {
	action 		commandAction
	key 		string
	value 		interface{}
	result 		chan <- interface{}
	data 		chan <- map[string]interface{}
	updater 	UpdateFunc
}

type commandAction int

const (
	remove commandAction = iota
	end
	find
	insert
	length
	update
)

func New()wegoSafeMap  {
	sm := make(wegoSafeMap)
	go sm.run()
	return sm
}

func (sm wegoSafeMap)run()  {
	store := make(map[string]interface{})
	for command := range sm{
		switch command.action {
		case insert:
			store[command.key] = command.value
		case remove:
			delete(store,command.key)
		case find:
			value,found := store[command.key]
			command.result <- findResult{value,found}
		case length:
			command.result <- len(store)
		case update:
			value,found := store[command.key]
			store[command.key] = command.updater(value,found)
		case end:
			close(sm)
			command.data <- store
		}

	}
}

func (sm wegoSafeMap)Insert(key string,value interface{})  {
	sm <- commandData{action:insert,key:key,value:value}
}

func (sm wegoSafeMap)Delete(key string)  {
	sm <- commandData{action:remove,key:key}
}

type findResult struct {
	value 	interface{}
	found 	bool
}

func (sm wegoSafeMap)Find(key string)(value interface{},found bool)  {
	reply := make(chan interface{})
	sm <- commandData{action:find,key:key,result:reply}
	result := (<-reply).(findResult)
	return result.value,result.found
}

func (sm wegoSafeMap)Len()int  {
	reply := make(chan interface{})
	sm <- commandData{action:length,result:reply}
	return (<-reply).(int)
}

func (sm wegoSafeMap)Update(key string,updater UpdateFunc)  {
	sm <- commandData{action:update,key:key,updater:updater}
}

func (sm wegoSafeMap)Close()map[string]interface{}  {
	reply := make(chan map[string]interface{})
	sm <- commandData{action:end,data:reply}
	return <-reply
}


