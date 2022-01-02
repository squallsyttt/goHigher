package lru

import "container/list"

//双向链表节点的数据类型
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// Cache  is an LRU cache. It is not safe for concurrent access
type Cache struct {
	//允许使用的最大内存
	maxBytes int64
	//当前已使用的内存
	nbytes int64
	//这边是go标准库实现的 双向链表 list.List
	ll *list.List
	//键是字符串，值是双向链表中对应节点的指针
	cache map[string]*list.Element
	//optional and executed when an entry is purged
	OnEvicted func(key string, value Value)
}

func New(maxBytes int64,onEvicted func(string,Value)) *Cache{
	return &Cache{
		maxBytes: maxBytes,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value,ok bool){
	if ele,ok := c.cache[key] ;ok{
		c.ll.MoveToFront(ele)
		//这边这个entry没有很看得懂 等会调试的时候打印看下
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) Add(){
// TODO
}

func (c *Cache) Del(){
// TODO
}

func (c *Cache) Len() int {
	return c.ll.Len()
}