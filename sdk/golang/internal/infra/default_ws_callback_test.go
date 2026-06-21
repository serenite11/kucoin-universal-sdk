package infra

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"unsafe"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestCallbackManager_Add(t *testing.T) {
	cm := NewCallbackManager("/mock")
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: newEmptyCallback()}))

	assert.Equal(t, 1, len(cm.topicCallbackMapping))
	assert.Equal(t, 1, len(cm.idTopicMapping))

	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"b", "c"}, Callback: newEmptyCallback()}))

	assert.Equal(t, 3, len(cm.topicCallbackMapping))
	assert.Equal(t, 2, len(cm.idTopicMapping))
}

func TestCallbackManager_Add2(t *testing.T) {
	cm := NewCallbackManager("/mock")

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{fmt.Sprintf("%d", i)}, Callback: newEmptyCallback()}))
		}(i)
	}

	wg.Wait()

	assert.Equal(t, 10, len(cm.topicCallbackMapping))
	assert.Equal(t, 10, len(cm.idTopicMapping))
}

func TestCallbackManager_Add_Dup1(t *testing.T) {
	cm := NewCallbackManager("/mock")

	c1 := newEmptyCallback()
	c2 := newEmptyCallback()

	sub1 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}
	sub2 := &util.SubInfo{Prefix: "/mock", Args: []string{"a", "c"}, Callback: c2}
	assert.True(t, cm.Add(sub1))
	assert.True(t, cm.Add(sub2))

	assert.Equal(t, 2, len(cm.topicCallbackMapping))
	assert.Equal(t, 2, len(cm.idTopicMapping))

	assert.Equal(t, uintptr(unsafe.Pointer(c1)), uintptr(unsafe.Pointer(cm.topicCallbackMapping["/mock:a"].callback.(*cbObj))))
	assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.topicCallbackMapping["/mock:c"].callback.(*cbObj))))

	_, ok := cm.idTopicMapping[sub1.ToId()]["/mock:a"]
	assert.True(t, ok)
	_, ok = cm.idTopicMapping[sub2.ToId()]["/mock:c"]
	assert.True(t, ok)
}

func TestCallbackManager_Add_Dup2(t *testing.T) {
	cm := NewCallbackManager("/mock")

	c1 := newEmptyCallback()
	c2 := newEmptyCallback()

	sub1 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}
	sub2 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c2}
	assert.True(t, cm.Add(sub1))
	assert.False(t, cm.Add(sub2))

	assert.Equal(t, 1, len(cm.topicCallbackMapping))
	assert.Equal(t, 1, len(cm.idTopicMapping))

	assert.Equal(t, uintptr(unsafe.Pointer(c1)), uintptr(unsafe.Pointer(cm.topicCallbackMapping["/mock:a"].callback.(*cbObj))))

	_, ok := cm.idTopicMapping[sub1.ToId()]["/mock:a"]
	assert.True(t, ok)
}

func TestCallbackManager_Remove(t *testing.T) {
	cm := NewCallbackManager("/mock")

	c1 := newEmptyCallback()
	c2 := newEmptyCallback()

	sub1 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}
	sub2 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c2}
	assert.True(t, cm.Add(sub1))
	assert.False(t, cm.Add(sub2))

	cm.Remove(sub1.ToId())

	assert.Equal(t, 0, len(cm.topicCallbackMapping))
	assert.Equal(t, 0, len(cm.idTopicMapping))
}

func TestCallbackManager_Remove_2(t *testing.T) {
	cm := NewCallbackManager("/mock")

	c1 := newEmptyCallback()
	c2 := newEmptyCallback()

	sub1 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}
	sub2 := &util.SubInfo{Prefix: "/mock", Args: []string{"a", "c"}, Callback: c2}
	assert.True(t, cm.Add(sub1))
	assert.True(t, cm.Add(sub2))

	cm.Remove(sub1.ToId())

	assert.Equal(t, 1, len(cm.topicCallbackMapping))
	assert.Equal(t, 1, len(cm.idTopicMapping))

	assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.topicCallbackMapping["/mock:c"].callback.(*cbObj))))
	_, ok := cm.idTopicMapping[sub2.ToId()]["/mock:c"]
	assert.True(t, ok)
}

func TestCallbackManager_Remove_3(t *testing.T) {
	cm := NewCallbackManager("/mock")

	c1 := newEmptyCallback()
	c2 := newEmptyCallback()

	sub1 := &util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}
	sub2 := &util.SubInfo{Prefix: "/mock", Args: []string{"b", "c"}, Callback: c2}
	assert.True(t, cm.Add(sub1))
	assert.True(t, cm.Add(sub2))

	cm.Remove(sub1.ToId())

	assert.Equal(t, 2, len(cm.topicCallbackMapping))
	assert.Equal(t, 1, len(cm.idTopicMapping))

	assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.topicCallbackMapping["/mock:b"].callback.(*cbObj))))
	assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.topicCallbackMapping["/mock:c"].callback.(*cbObj))))
	_, ok := cm.idTopicMapping[sub2.ToId()]["/mock:b"]
	assert.True(t, ok)
	_, ok = cm.idTopicMapping[sub2.ToId()]["/mock:c"]
	assert.True(t, ok)
}

func TestCallbackManager_Add_Remove(t *testing.T) {
	cm := NewCallbackManager("/mock")

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sub := &util.SubInfo{Prefix: "/mock", Args: []string{fmt.Sprintf("%d", i)}, Callback: newEmptyCallback()}
			cm.Add(sub)
			cm.Remove(sub.ToId())
		}(i)
	}
	wg.Wait()

	assert.Equal(t, 0, len(cm.topicCallbackMapping))
	assert.Equal(t, 0, len(cm.idTopicMapping))

	wg = sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sub := &util.SubInfo{Prefix: "/mock", Args: []string{fmt.Sprintf("%d", i%3)}, Callback: newEmptyCallback()}
			cm.Add(sub)
			cm.Remove(sub.ToId())
		}(i)
	}
	wg.Wait()

	assert.Equal(t, 0, len(cm.topicCallbackMapping))
	assert.Equal(t, 0, len(cm.idTopicMapping))
}

func TestCallbackManager_Get(t *testing.T) {
	cm := NewCallbackManager("/mock")
	c1 := newEmptyCallback()
	c2 := newEmptyCallback()
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}))
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"b", "c"}, Callback: c2}))

	assert.Equal(t, uintptr(unsafe.Pointer(c1)), uintptr(unsafe.Pointer(cm.Get("/mock:a").(*cbObj))))
	assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.Get("/mock:b").(*cbObj))))
	assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.Get("/mock:c").(*cbObj))))

}

func TestCallbackManager_Get2(t *testing.T) {
	cm := NewCallbackManager("/mock")
	c1 := newEmptyCallback()
	c2 := newEmptyCallback()
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}))
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"b", "c"}, Callback: c2}))

	for i := 0; i < 10; i++ {
		go func() {
			assert.Equal(t, uintptr(unsafe.Pointer(c1)), uintptr(unsafe.Pointer(cm.Get("/mock:a").(*cbObj))))
			assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.Get("/mock:b").(*cbObj))))
			assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(cm.Get("/mock:c").(*cbObj))))
		}()
	}

}

func TestCallbackManager_GetSubInfo(t *testing.T) {
	cm := NewCallbackManager("/mock")
	c1 := newEmptyCallback()
	c2 := newEmptyCallback()
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"a"}, Callback: c1}))
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{"b", "c"}, Callback: c2}))

	info := cm.GetSubInfo()
	for _, i := range info {
		assert.Equal(t, i.Prefix, "/mock")
		if len(i.Args) == 1 {
			assert.Equal(t, i.Args[0], "a")
			assert.Equal(t, uintptr(unsafe.Pointer(c1)), uintptr(unsafe.Pointer(i.Callback.(*cbObj))))
		} else {
			assert.Equal(t, uintptr(unsafe.Pointer(c2)), uintptr(unsafe.Pointer(i.Callback.(*cbObj))))
			sort.Strings(i.Args)
			assert.Equal(t, i.Args[0], "b")
			assert.Equal(t, i.Args[1], "c")
		}
	}
}

func TestCallbackManager_GetSubInfo2(t *testing.T) {
	cm := NewCallbackManager("/mock")
	c1 := newEmptyCallback()
	assert.True(t, cm.Add(&util.SubInfo{Prefix: "/mock", Args: []string{}, Callback: c1}))

	info := cm.GetSubInfo()
	assert.Equal(t, len(info[0].Args), 0)
	assert.Equal(t, uintptr(unsafe.Pointer(c1)), uintptr(unsafe.Pointer(info[0].Callback.(*cbObj))))
	assert.Equal(t, "/mock", info[0].Prefix)
}
