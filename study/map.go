package study

import (
	"fmt"
	"sort"
)

func TestMap() {
	m1 := map[int]string{1: "周润发", 2: "丁勇岱"}
	m2 := make(map[int]string)
	m3 := make(map[int]string, 10)
	m2[3] = "me"
	delete(m1, 1)
	fmt.Printf("m1:%v len=%v\n", m1, len(m1))
	fmt.Printf("m2:%v len=%v\n", m2, len(m2))
	fmt.Printf("m3:%v len=%v\n", m3, len(m3))

	for key, val := range m1 {
		fmt.Printf("node:{key:%v,val:%v}", key, val)
	}

}

func TestMapSort() {
	m := map[int]string{3: "google", 2: "baidu", 1: "bing", 4: "duckduckgo"}
	for key, val := range m {
		fmt.Printf("{key:%v, val:%v} ", key, val)
	}
	println()
	SortByKey(&m)
	SortByVal(&m)
}

func SortByKey(mp *map[int]string) {
	m := *mp
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("{key:%v, val:%v} ", key, m[key])
	}
	println()
}

type pair struct {
	key int
	val string
}
type pairList []pair

func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].val < p[j].val }

func SortByVal(mp *map[int]string) {
	m := *mp
	pairs := make(pairList, 0, len(m))
	for key, val := range m {
		pairs = append(pairs, pair{key: key, val: val})
	}
	sort.Sort(pairs)
	for _, p := range pairs {
		fmt.Printf("{key: %v, val: %v} ", p.key, p.val)
	}
}

func TestHashMap() {
	hashMap := NewHashMap[int, string]()
	hashMap.Put(1, "一")
	hashMap.Put(2, "二")
	hashMap.Put(3, "三")
	fmt.Println(hashMap)

	m := map[int]string{}
	m[1] = "一"
	m[2] = "二"
	m[3] = "三"
	fmt.Println(m)
}

type HashMap[K int | string, V any] struct {
	data map[K]V
}

func NewHashMap[K int | string, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{data: make(map[K]V, 1)}
}

func (hashMap *HashMap[K, V]) Put(key K, val V) {
	hashMap.data[key] = val
}

func (hashMap *HashMap[K, V]) Remove(key K) {
	delete(hashMap.data, key)
}

func (hashMap *HashMap[K, V]) Get(key K) (V, bool) {
	val, ok := hashMap.data[key]
	return val, ok
}

func (hashMap *HashMap[K, V]) String() string {
	return fmt.Sprint(hashMap.data)
}
