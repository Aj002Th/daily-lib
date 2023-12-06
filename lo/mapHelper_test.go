package lo

import (
	"fmt"
	"github.com/samber/lo"
	"testing"
)

var mp = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}

func TestKeys(t *testing.T) {
	result := lo.Keys(mp)
	printResult(mp, result)
}

func TestValues(t *testing.T) {
	result := lo.Values(mp)
	printResult(mp, result)
}

// 默认值
func TestValuesOr(t *testing.T) {
	result := lo.ValueOr(mp, "d", 10)
	printResult(mp, result)
}

// map helper 版本的 filter
func TestPickBy(t *testing.T) {
	result := lo.PickBy(mp, func(key string, value int) bool {
		return value > 1
	})
	printResult(mp, result)
}

func TestPickByKeys(t *testing.T) {
	result := lo.PickByKeys(mp, []string{"a", "b"})
	printResult(mp, result)
}

func TestPickByValues(t *testing.T) {
	result := lo.PickByValues(mp, []int{1, 2})
	printResult(mp, result)
}

// map helper 版本的 reject
func TestOmitBy(t *testing.T) {
	result := lo.OmitBy(mp, func(key string, value int) bool {
		return value > 1
	})
	printResult(mp, result)
}

func TestOmitByKeys(t *testing.T) {
	result := lo.OmitByKeys(mp, []string{"a", "b"})
	printResult(mp, result)
}

func TestOmitByValues(t *testing.T) {
	result := lo.OmitByValues(mp, []int{1, 2})
	printResult(mp, result)
}

// 从 map 变成 pair 切片
func TestEntries(t *testing.T) {
	result := lo.Entries(mp)
	printResult(mp, result)
}

// 从 pair 切片变成 map
func TestFromEntries(t *testing.T) {
	input := []lo.Entry[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	result := lo.FromEntries(input)
	printResult(input, result)
}

// kv 反转, 重复则覆盖
func TestInvert(t *testing.T) {
	result := lo.Invert(mp)
	printResult(mp, result)
}

// 合并多个 map
func TestAssign(t *testing.T) {
	result := lo.Assign(mp, map[string]int{"d": 4})
	printResult(mp, result)
}

// 变换 key, 不仅仅是数值, 类型也能变
// 比较奇怪的是这组 api 注入的函数 value, key 参数顺序反过来
func TestMapKeys(t *testing.T) {
	result := lo.MapKeys(mp, func(value int, key string) string {
		return key + "1"
	})
	printResult(mp, result)
}

// 变换 value
func TestMapValues(t *testing.T) {
	result := lo.MapValues(mp, func(value int, key string) int {
		return value + 1
	})
	printResult(mp, result)
}

// key, value 都变换
func TestMapEntries(t *testing.T) {
	result := lo.MapEntries(mp, func(key string, value int) (string, int) {
		return key + "1", value + 1
	})
	printResult(mp, result)
}

// 自定义 map entry 转换成 slice element 的方法
func TestMapToSlice(t *testing.T) {
	result := lo.MapToSlice(mp, func(key string, value int) string {
		return key + ":" + fmt.Sprint(value) + "!"
	})
	printResult(mp, result)
}
