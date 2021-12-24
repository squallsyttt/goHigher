package main

import "fmt"

//A Getter loads data for key
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc interfaceFunc
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

func GetFromSource(getter Getter, key string) []byte {
	buf, err := getter.Get(key)
	if err == nil {
		return buf
	}
	return nil
}

func test(key string) ([]byte, error) {
	//这边就是一个类型转换
	return []byte(key), nil
}

func test2(key string) ([]byte, error) {
	return []byte(key + "syttt"), nil
}

type DB struct{ url string }

func (db *DB) Query(sql string, args ...string) string {
	return "hello"
}

func (db *DB) Get(key string) ([]byte, error) {
	v := db.Query("SELECT NAME FROM TABLE WHEN NAME= ?", key)
	return []byte(v), nil
}

func main() {
	res1 := GetFromSource(GetterFunc(test), "hello")
	fmt.Println(res1)
	res2 := GetFromSource(GetterFunc(test2), "hello")
	fmt.Println(res2)
	//GetFromSource(new(DB), "hello")
}
