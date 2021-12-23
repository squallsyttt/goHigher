package main

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

func main() {
	GetFromSource(GetterFunc(test), "hello")
}
