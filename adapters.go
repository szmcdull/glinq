package glinq

import (
	"bufio"
	"io"
	"reflect"
	"sync/atomic"
)

type (
	Slice[T any]           []T
	SliceEnumerator[T any] struct {
		s   Slice[T]
		pos int
	}

	LineEnumerator struct {
		reader  *bufio.Reader
		scanner *bufio.Scanner
		taken   int32
	}

	Pair[K comparable, V any] struct {
		Key   K
		Value V
	}
	Map[K comparable, V any]           map[K]V
	MapEnumerator[K comparable, V any] struct {
		m    Map[K, V]
		iter *reflect.MapIter
	}
)

func (me Slice[T]) Count() int {
	return len(me)
}

func (me Slice[T]) Any() bool {
	return len(me) > 0
}

// func From[T any](v any) IEnumerable[T] {
// 	switch vv := v.(type) {
// 	case []T:
// 		return Slice[T](vv)
// 	default:
// 		return nil
// 	}
// 	t := reflect.TypeOf(v)
// 	if t.Kind() == reflect.Slice {
// 		reflect.SliceOf(t.Elem())
// 	}
// }

func FromSlice[T any](v []T) IEnumerable[T] {
	return Slice[T](v)
}

// 先复制出一个 []Pair[K,V]，再转成 IEnumerable，不适用于map很大的情况
func FromMapCopy[K comparable, V any](v map[K]V) IEnumerable[Pair[K, V]] {
	values := make([]Pair[K, V], 0, len(v))
	for kk, vv := range v {
		values = append(values, Pair[K, V]{kk, vv})
	}
	return FromSlice(values)
}

// 通过 reflect.MapRange 遍历 map，线性读取快，随机访问慢
func FromMapReflect[K comparable, V any](v map[K]V) IEnumerable[Pair[K, V]] {
	return Map[K, V](v)
}

func FromMap[K comparable, V any](v map[K]V) IEnumerable[Pair[K, V]] {
	if len(v) < 1024 {
		return FromMapCopy(v)
	} else {
		return FromMapReflect(v)
	}
}

func (me Slice[T]) GetAt(i int) (v T, err error) {
	if i < len(me) {
		v = me[i]
	} else {
		err = io.ErrUnexpectedEOF
	}
	return
}

func (me Slice[T]) GetEnumerator() IEnumerator[T] {
	return &SliceEnumerator[T]{s: me, pos: -1}
}

func (me *SliceEnumerator[T]) Current() T {
	return me.s[me.pos]
}

func (me *SliceEnumerator[T]) GetAt(pos int) (v T, err error) {
	return me.s.GetAt(pos)
}

func (me *SliceEnumerator[T]) SeekOnce(pos int) error {
	if pos < len(me.s) {
		me.pos = pos
		return nil
	}
	return io.ErrUnexpectedEOF
}

func (me *SliceEnumerator[T]) MoveNext() error {
	pos := me.pos + 1
	if pos >= len(me.s) {
		return io.EOF
	}
	me.pos = pos
	return nil
}

func (me *SliceEnumerator[T]) Reset() error {
	me.pos = -1
	return nil
}

func ReadLines(reader io.Reader) *LineEnumerator {
	var r *bufio.Reader
	var ok bool

	if r, ok = reader.(*bufio.Reader); !ok {
		r = bufio.NewReader(reader)
	}
	return &LineEnumerator{
		reader:  r,
		scanner: bufio.NewScanner(reader),
	}
}

func (me *LineEnumerator) GetEnumerator() IEnumerator[string] {
	if atomic.CompareAndSwapInt32(&me.taken, 0, 1) {
		return me
	}
	panic(`read iterators cannot be read more than once`)
}

func (me *LineEnumerator) MoveNext() error {
	if me.scanner.Scan() {
		return nil
	} else {
		err := me.scanner.Err()
		if err == nil {
			err = io.EOF
		}
		return err
	}
}

func (me *LineEnumerator) Current() string {
	return me.scanner.Text()
}

func (me *LineEnumerator) Any() bool {
	_, err := me.reader.Peek(1)
	return err != nil
}

func (me *LineEnumerator) Count() int {
	return -1
}

func (me Map[K, V]) Count() int {
	return len(me)
}

func (me Map[K, V]) Any() bool {
	return len(me) > 0
}

func (me Map[K, V]) GetEnumerator() IEnumerator[Pair[K, V]] {
	v := reflect.ValueOf(me)
	return &MapEnumerator[K, V]{me, v.MapRange()}
}

func (me *MapEnumerator[K, V]) Current() Pair[K, V] {
	return Pair[K, V]{me.iter.Key().Interface().(K), me.iter.Value().Interface().(V)}
}

func (me *MapEnumerator[K, V]) MoveNext() error {
	if me.iter.Next() {
		return nil
	}
	return io.EOF
}

func (me *MapEnumerator[K, V]) Reset() error {
	me.iter = reflect.ValueOf(me.m).MapRange()
	return nil
}
