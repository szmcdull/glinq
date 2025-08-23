//go:build go1.23
// +build go1.23

package giter

import (
	"slices"
	"testing"
)

func TestBasicLinqOperations(t *testing.T) {
	// Test data
	numbers := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Test Where
	evens := Where(numbers, func(x int) bool { return x%2 == 0 })
	evenSlice := ToSlice(evens)
	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(evenSlice, expected) {
		t.Errorf("Where failed: got %v, want %v", evenSlice, expected)
	}

	// Test Select
	doubled := Select(slices.Values([]int{1, 2, 3}), func(x int) int { return x * 2 })
	doubledSlice := ToSlice(doubled)
	expectedDoubled := []int{2, 4, 6}
	if !slices.Equal(doubledSlice, expectedDoubled) {
		t.Errorf("Select failed: got %v, want %v", doubledSlice, expectedDoubled)
	}

	// Test Take
	firstThree := Take(numbers, 3)
	firstThreeSlice := ToSlice(firstThree)
	expectedFirstThree := []int{1, 2, 3}
	if !slices.Equal(firstThreeSlice, expectedFirstThree) {
		t.Errorf("Take failed: got %v, want %v", firstThreeSlice, expectedFirstThree)
	}

	// Test Skip
	skipTwo := Skip(slices.Values([]int{1, 2, 3, 4, 5}), 2)
	skipTwoSlice := ToSlice(skipTwo)
	expectedSkipTwo := []int{3, 4, 5}
	if !slices.Equal(skipTwoSlice, expectedSkipTwo) {
		t.Errorf("Skip failed: got %v, want %v", skipTwoSlice, expectedSkipTwo)
	}
}

func TestAggregateOperations(t *testing.T) {
	numbers := slices.Values([]int{1, 2, 3, 4, 5})

	// Test Count
	count := Count(numbers)
	if count != 5 {
		t.Errorf("Count failed: got %d, want 5", count)
	}

	// Test Any
	hasEven := Any(numbers, func(x int) bool { return x%2 == 0 })
	if !hasEven {
		t.Error("Any failed: should have found even number")
	}

	// Test All
	allPositive := All(numbers, func(x int) bool { return x > 0 })
	if !allPositive {
		t.Error("All failed: all numbers should be positive")
	}

	// Test First
	first := First(numbers)
	if first != 1 {
		t.Errorf("First failed: got %d, want 1", first)
	}

	// Test Last
	last := Last(numbers)
	if last != 5 {
		t.Errorf("Last failed: got %d, want 5", last)
	}
}

func TestSetOperations(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3, 4})
	seq2 := slices.Values([]int{3, 4, 5, 6})

	// Test Union
	union := Union(seq1, seq2)
	unionSlice := ToSlice(union)
	// Union should contain all unique elements
	if len(unionSlice) != 6 {
		t.Errorf("Union failed: got %d elements, want 6", len(unionSlice))
	}

	// Test Intersect
	intersect := Intersect(slices.Values([]int{1, 2, 3, 4}), slices.Values([]int{3, 4, 5, 6}))
	intersectSlice := ToSlice(intersect)
	if len(intersectSlice) != 2 {
		t.Errorf("Intersect failed: got %v, want 2 elements", intersectSlice)
	}
	// Verify the intersection contains the expected elements
	for _, v := range []int{3, 4} {
		if !slices.Contains(intersectSlice, v) {
			t.Errorf("Intersect failed: missing element %d", v)
		}
	}

	// Test Except
	except := Except(slices.Values([]int{1, 2, 3, 4}), slices.Values([]int{3, 4, 5, 6}))
	exceptSlice := ToSlice(except)
	if len(exceptSlice) != 2 {
		t.Errorf("Except failed: got %v, want 2 elements", exceptSlice)
	}
	// Verify the except contains the expected elements
	for _, v := range []int{1, 2} {
		if !slices.Contains(exceptSlice, v) {
			t.Errorf("Except failed: missing element %d", v)
		}
	}
}

func TestDistinct(t *testing.T) {
	numbers := slices.Values([]int{1, 2, 2, 3, 3, 3, 4})
	distinct := Distinct(numbers)
	distinctSlice := ToSlice(distinct)
	expected := []int{1, 2, 3, 4}
	if !slices.Equal(distinctSlice, expected) {
		t.Errorf("Distinct failed: got %v, want %v", distinctSlice, expected)
	}
}

func TestChunk(t *testing.T) {
	numbers := slices.Values([]int{1, 2, 3, 4, 5, 6, 7})
	chunks := Chunk(numbers, 3)

	var allChunks [][]int
	for chunk := range chunks {
		chunkSlice := ToSlice(chunk)
		allChunks = append(allChunks, chunkSlice)
	}

	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
	if len(allChunks) != 3 {
		t.Errorf("Chunk failed: got %d chunks, want 3", len(allChunks))
	}

	for i, chunk := range allChunks {
		if !slices.Equal(chunk, expected[i]) {
			t.Errorf("Chunk %d failed: got %v, want %v", i, chunk, expected[i])
		}
	}
}

func TestGroupBy(t *testing.T) {
	words := slices.Values([]string{"apple", "apricot", "banana", "cherry", "avocado"})

	// Group by first letter
	groups := GroupBy(words, func(s string) byte { return s[0] })

	groupMap := make(map[byte][]string)
	for group := range groups {
		groupMap[group.Key] = ToSlice(group.Values)
	}

	// Check that we have groups for 'a', 'b', 'c'
	if len(groupMap) != 3 {
		t.Errorf("GroupBy failed: got %d groups, want 3", len(groupMap))
	}

	// Check 'a' group
	aGroup := groupMap['a']
	if len(aGroup) != 3 {
		t.Errorf("GroupBy 'a' group failed: got %d items, want 3", len(aGroup))
	}
}

func TestRange(t *testing.T) {
	rangeSeq := Range(5, 3)
	rangeSlice := ToSlice(rangeSeq)
	expected := []int{5, 6, 7}
	if !slices.Equal(rangeSlice, expected) {
		t.Errorf("Range failed: got %v, want %v", rangeSlice, expected)
	}
}

func TestRepeat(t *testing.T) {
	repeatSeq := Repeat("hello", 3)
	repeatSlice := ToSlice(repeatSeq)
	expected := []string{"hello", "hello", "hello"}
	if !slices.Equal(repeatSlice, expected) {
		t.Errorf("Repeat failed: got %v, want %v", repeatSlice, expected)
	}
}

func TestZip(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]string{"a", "b", "c"})

	zipped := Zip(seq1, seq2, func(i int, s string) string {
		return s + string(rune('0'+i))
	})

	zippedSlice := ToSlice(zipped)
	expected := []string{"a1", "b2", "c3"}
	if !slices.Equal(zippedSlice, expected) {
		t.Errorf("Zip failed: got %v, want %v", zippedSlice, expected)
	}
}

// ===== New Order API Tests =====

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name       string
	Department string
	Level      int
	Salary     int
}

// TestOrderAPI tests the basic functionality of the new Order API
func TestOrderAPI(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 25},
		{"Eve", 35},
		{"Frank", 25},
	}

	// 使用新的 Order API：先按年龄排序，然后按姓名排序
	orderSpec := By(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b })
	orderSpec = Then(orderSpec, func(p Person) string { return p.Name }, func(a, b string) bool { return a < b })

	sorted := Order(slices.Values(people), orderSpec)
	result := ToSlice(sorted)

	expected := []Person{
		{"Bob", 25}, {"David", 25}, {"Frank", 25},
		{"Alice", 30}, {"Charlie", 30},
		{"Eve", 35},
	}

	if !slices.Equal(result, expected) {
		t.Errorf("Order API 失败: got %+v, want %+v", result, expected)
	}

	t.Logf("✅ 基础排序测试通过: %+v", result)
}

// TestOrderInlineSyntax tests the inline nested syntax style
func TestOrderInlineSyntax(t *testing.T) {
	people := []Person{
		{"Alice", 30}, {"Bob", 25}, {"Charlie", 30}, {"David", 25},
	}

	// 测试内联嵌套语法：Order(seq, Then(By(xxx), xxx))
	result := ToSlice(Order(
		slices.Values(people),
		Then(
			By(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b }),
			func(p Person) string { return p.Name }, func(a, b string) bool { return a < b },
		),
	))

	expected := []Person{
		{"Bob", 25}, {"David", 25}, {"Alice", 30}, {"Charlie", 30},
	}

	if !slices.Equal(result, expected) {
		t.Errorf("内联语法失败: got %+v, want %+v", result, expected)
	}

	t.Logf("✅ 内联语法测试通过: %+v", result)
}

// TestOrderComplexMultiLevel tests complex multi-level sorting
func TestOrderComplexMultiLevel(t *testing.T) {
	employees := []Employee{
		{"Alice", "Engineering", 3, 80000},
		{"Bob", "Engineering", 2, 70000},
		{"Charlie", "Sales", 3, 75000},
		{"David", "Engineering", 3, 85000},
		{"Eve", "Sales", 2, 65000},
		{"Frank", "Engineering", 2, 72000},
		{"Grace", "Sales", 3, 78000},
	}

	// 四级排序：部门 -> 级别(降序) -> 薪水(降序) -> 姓名
	result := ToSlice(Order(
		slices.Values(employees),
		Then(
			ThenDescending(
				ThenDescending(
					By(func(e Employee) string { return e.Department }, func(a, b string) bool { return a < b }),
					func(e Employee) int { return e.Level }, func(a, b int) bool { return a < b },
				),
				func(e Employee) int { return e.Salary }, func(a, b int) bool { return a < b },
			),
			func(e Employee) string { return e.Name }, func(a, b string) bool { return a < b },
		),
	))

	expected := []Employee{
		{"David", "Engineering", 3, 85000},
		{"Alice", "Engineering", 3, 80000},
		{"Frank", "Engineering", 2, 72000},
		{"Bob", "Engineering", 2, 70000},
		{"Grace", "Sales", 3, 78000},
		{"Charlie", "Sales", 3, 75000},
		{"Eve", "Sales", 2, 65000},
	}

	if !slices.Equal(result, expected) {
		t.Errorf("复杂排序失败")
		for i := 0; i < len(result) && i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Logf("位置 %d: got %+v, want %+v", i, result[i], expected[i])
			}
		}
	} else {
		t.Log("✅ 复杂四级排序测试通过!")
	}
}

// TestOrderDescending tests descending order functionality
func TestOrderDescending(t *testing.T) {
	people := []Person{
		{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}, {"David", 25},
	}

	// 按年龄降序，然后按姓名升序
	result := ToSlice(Order(
		slices.Values(people),
		Then(
			ByDescending(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b }),
			func(p Person) string { return p.Name }, func(a, b string) bool { return a < b },
		),
	))

	expected := []Person{
		{"Charlie", 35}, {"Alice", 30}, {"Bob", 25}, {"David", 25},
	}

	if !slices.Equal(result, expected) {
		t.Errorf("降序排序失败: got %+v, want %+v", result, expected)
	}

	t.Logf("✅ 降序排序测试通过: %+v", result)
}

// TestOrderAPIComparison compares new API with native Go sorting
func TestOrderAPIComparison(t *testing.T) {
	people := []Person{
		{"Alice", 30}, {"Bob", 25}, {"Charlie", 30}, {"David", 25}, {"Eve", 35},
	}

	// 新API
	newAPIResult := ToSlice(Order(
		slices.Values(people),
		Then(
			By(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b }),
			func(p Person) string { return p.Name }, func(a, b string) bool { return a < b },
		),
	))

	// 原生Go排序作为对比
	nativeResult := make([]Person, len(people))
	copy(nativeResult, people)
	slices.SortStableFunc(nativeResult, func(a, b Person) int {
		if a.Age != b.Age {
			if a.Age < b.Age {
				return -1
			}
			return 1
		}
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})

	if !slices.Equal(newAPIResult, nativeResult) {
		t.Errorf("新API结果与原生排序不一致")
		t.Logf("新API: %+v", newAPIResult)
		t.Logf("原生: %+v", nativeResult)
	} else {
		t.Log("✅ 新API与原生排序结果一致")
	}
}

// TestOrderEdgeCases tests edge cases
func TestOrderEdgeCases(t *testing.T) {
	// 空序列
	emptyResult := ToSlice(Order(
		slices.Values([]Person{}),
		By(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b }),
	))
	if len(emptyResult) != 0 {
		t.Errorf("空序列排序失败: got %v, want empty slice", emptyResult)
	}

	// 单元素
	singleResult := ToSlice(Order(
		slices.Values([]Person{{"Alice", 30}}),
		Then(
			By(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b }),
			func(p Person) string { return p.Name }, func(a, b string) bool { return a < b },
		),
	))
	expected := []Person{{"Alice", 30}}
	if !slices.Equal(singleResult, expected) {
		t.Errorf("单元素排序失败: got %+v, want %+v", singleResult, expected)
	}

	t.Log("✅ 边界情况测试通过")
}

func TestOrderSimple(t *testing.T) {
	people := []Person{
		{"Alice", 30}, {"Bob", 25}, {"Charlie", 30}, {"David", 25}, {"Eve", 35},
	}

	// 只用一个less函数，先按年龄升序，再按姓名升序
	sorted := ToSlice(OrderBy(slices.Values(people), func(a, b Person) bool {
		if a.Age != b.Age {
			return a.Age < b.Age
		}
		return a.Name < b.Name
	}))

	expected := []Person{
		{"Bob", 25}, {"David", 25}, {"Alice", 30}, {"Charlie", 30}, {"Eve", 35},
	}

	if !slices.Equal(sorted, expected) {
		t.Errorf("OrderSimple 排序失败: got %+v, want %+v", sorted, expected)
	} else {
		t.Logf("✅ OrderSimple 排序通过: %+v", sorted)
	}

	// 只用一个less函数，按姓名降序
	sortedDesc := ToSlice(OrderBy(slices.Values(people), func(a, b Person) bool {
		return a.Name > b.Name
	}))

	expectedDesc := []Person{
		{"Eve", 35}, {"David", 25}, {"Charlie", 30}, {"Bob", 25}, {"Alice", 30},
	}

	if !slices.Equal(sortedDesc, expectedDesc) {
		t.Errorf("OrderSimple 降序排序失败: got %+v, want %+v", sortedDesc, expectedDesc)
	} else {
		t.Logf("✅ OrderSimple 降序排序通过: %+v", sortedDesc)
	}
}
