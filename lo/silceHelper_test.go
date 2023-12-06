package lo

import (
	"github.com/samber/lo"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Filter(list, func(x, idx int) bool {
		return x == idx+1
	})
	printResult(list, result)
}

func TestMap(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Map(list, func(x, idx int) int {
		return x * 10
	})
	printResult(list, result)
}

// 做 filter 操作的同时做 map 操作
func TestFilterMap(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.FilterMap(list, func(x, idx int) (int, bool) {
		if x == idx+1 {
			return x * 10, true
		}
		return 0, false
	})
	printResult(list, result)
}

// 平铺扩展
func TestFlatMap(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.FlatMap(list, func(x, idx int) []int {
		return []int{x, x * 10, x * 100}
	})
	printResult(list, result)
}

// 汇总
func TestReduce(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Reduce(list, func(agg, item, idx int) int {
		return agg + item
	}, 0)
	printResult(list, result)
}

func TestReduceRight(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.ReduceRight(list, func(agg, item, idx int) int {
		return agg + item
	}, 0)
	printResult(list, result)
}

func TestForEach(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := 0
	lo.ForEach(list, func(x, idx int) {
		result += x
	})
	printResult(list, result)
}

// 感觉不太常用, 传入的函数必须得要个返回值
// 要是能传入无返回值的函数感觉可能会好一点
func TestTimes(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Times(4, func(idx int) int {
		return list[idx] * 10
	})
	printResult(list, result)
}

func TestUniq(t *testing.T) {
	list := []int{1, 3, 2, 4, 1, 3, 2, 4}
	result := lo.Uniq(list)
	printResult(list, result)
}

// 自定义去重的方法
func TestUniqBy(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.UniqBy(list, func(x int) int {
		return x % 2
	})
	printResult(list, result)
}

// 按逻辑分组, 结果是 map
func TestGroupBy(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.GroupBy(list, func(x int) string {
		return strconv.Itoa(x % 2)
	})
	printResult(list, result)
}

// 按位置分块
func TestChunk(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Chunk(list, 2)
	printResult(list, result)
}

// 按逻辑分块
func TestPartitionBy(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.PartitionBy(list, func(x int) bool {
		return x%2 == 0
	})
	printResult(list, result)
}

// 平铺多级数组, 可能也不太实用
func TestFlatten(t *testing.T) {
	list := [][]int{{1, 2}, {3, 4}}
	result := lo.Flatten(list)
	printResult(list, result)
}

// 交错合并
func TestInterleave(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{4, 5, 6}
	result := lo.Interleave(list1, list2)
	printResult(list1, result)
}

// 随机打乱
// tips: 输入的数组本身会被打乱
func TestShuffle(t *testing.T) {
	list := []int{1, 3, 2, 4}
	listCopy := make([]int, len(list))
	copy(listCopy, list)
	result := lo.Shuffle(listCopy)
	printResult(list, result)
}

// tips: 输入的数组本身也会反转, 只有 Shuffle 和 Reverse 会这样, 需要注意
func TestReverse(t *testing.T) {
	list := []int{1, 3, 2, 4}
	listCopy := make([]int, len(list))
	copy(listCopy, list)
	result := lo.Reverse(listCopy)
	printResult(list, result)
}

// 只能填充实现了 Clone 方法的结构体切片
// 实现逻辑就是创建了一个和带初始化切片一样长的 slice
// 然后每个 slice 元素值都通过 initial.Clone() 方法获取
// 缺点就是, 很多结构都未必会实现 Clonable 接口, 使用上还是有一定限制的
type Ele struct {
	Data string
}

func (e Ele) Clone() Ele {
	return Ele{Data: e.Data}
}
func TestFill(t *testing.T) {
	list := make([]Ele, 3)
	result := lo.Fill(list, Ele{Data: "c"})
	printResult(list, result)
}

// 限制也是要实现 Clonable 接口
func TestRepeat(t *testing.T) {
	result := lo.Repeat(3, Ele{Data: "c"})
	printResult("no input", result)
}

// 这个好用, 因为没有接口限制, 基础类型也能用
func TestRepeatBy(t *testing.T) {
	result := lo.RepeatBy(3, func(idx int) int {
		return 10
	})
	printResult("no input", result)
}

// 将 slice 按元素内容转换为 map
// 依据源码实现, 如果形成的 key 重复, 后面的会覆盖前面的
func TestKeyBy(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.KeyBy(list, func(x int) string {
		return strconv.Itoa(x)
	})
	printResult(list, result)
}

// 同时定义 kv 的生成规则
func TestAssociate(t *testing.T) {
	list1 := []int{1, 3, 2, 4}
	result := lo.Associate(list1, func(x int) (string, int) {
		return strconv.Itoa(x), x * 10
	})
	printResult(list1, result)
}

func TestDrop(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Drop(list, 2)
	printResult(list, result)
}

func TestDropRight(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.DropRight(list, 2)
	printResult(list, result)
}

// 从 slice 头部开始扫描
// 只要满足条件, 就丢弃, 直到条件第一次不被满足
func TestDropWhile(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.DropWhile(list, func(x int) bool {
		return x < 3
	})
	printResult(list, result)
}

func TestDropRightWhile(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.DropRightWhile(list, func(x int) bool {
		return x > 2
	})
	printResult(list, result)
}

// 和 filter 正好相反
func TestReject(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Reject(list, func(x int, idx int) bool {
		return x < 3
	})
	printResult(list, result)
}

func TestCount(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Count(list, 1)
	printResult(list, result)
}

// 自定义计数规则
func TestCountBy(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.CountBy(list, func(x int) bool {
		return x < 3
	})
	printResult(list, result)
}

// 每个元素都统计数目
func TestCountValues(t *testing.T) {
	list := []int{1, 3, 2, 4, 1, 1}
	result := lo.CountValues(list)
	printResult(list, result)
}

// 每一类都做数目统计, 需要定义分类的规则
// 其实相当于 lo.Map + lo.CountValues
func TestCountValuesBy(t *testing.T) {
	list := []int{1, 3, 2, 4, 1, 1}
	result := lo.CountValuesBy(list, func(x int) bool {
		return x < 3
	})
	printResult(list, result)
}

// 获取子区间(start + len), 越界不会 panic, 自动做裁剪
func TestSubset(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Subset(list, 2, 3)
	printResult(list, result)
}

// 获取子区间(start + end), 越界不会 panic, 自动做裁剪
func TestSlice(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Slice(list, 2, 3)
	printResult(list, result)
}

func TestReplace(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.Replace(list, 2, 10, 1)
	printResult(list, result)
}

func TestReplaceAll(t *testing.T) {
	list := []int{1, 3, 2, 4, 2}
	result := lo.ReplaceAll(list, 2, 10)
	printResult(list, result)
}

// 去除切片中的零值
func TestCompact(t *testing.T) {
	list := []int{1, 3, 2, 4, 0, 0, 0}
	result := lo.Compact(list)
	printResult(list, result)
}

func TestIsSorted(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.IsSorted(list)
	printResult(list, result)
}

func TestIsSortedByKey(t *testing.T) {
	list := []int{1, 3, 2, 4}
	result := lo.IsSortedByKey(list, func(x int) int {
		return x * 0
	})
	printResult(list, result)
}
