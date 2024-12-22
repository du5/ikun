package main

import (
	"container/list"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/du5/ikun/mp3"
	hook "github.com/robotn/gohook"
)

type BoundedList struct {
	*list.List
	maxSize int
	sync.Mutex
}

type valWithTime struct {
	val  string
	time int64
}

func (v valWithTime) String() string {
	return strings.ToLower(string(v.val))
}

func (b *BoundedList) Push(val ...string) {
	b.Lock()
	defer b.Unlock()
	// 在尾部插入
	b.PushBack(valWithTime{
		val[0],
		time.Now().Unix(),
	})
	// 超过容量则移除头部（最旧）元素
	if b.Len() > b.maxSize {
		front := b.Front()
		if front != nil {
			b.Remove(front)
		}
	}
	if len(val) > 1 {
		mp3.Play(val[1])
	}
}

func (b *BoundedList) GetAll() (l []valWithTime, start int64) {
	for e := b.Front(); e != nil; e = e.Next() {
		l = append(l, e.Value.(valWithTime))
	}
	if len(l) > 0 {
		start = l[0].time
	}
	return
}

func (b *BoundedList) GetAllStr() (str string, start int64) {
	var l []valWithTime
	l, start = b.GetAll()
	for _, v := range l {
		str += v.String()
	}
	return
}

var b *BoundedList

func init() {
	b = &BoundedList{
		list.New(),
		3,
		sync.Mutex{},
	}
}

func hookCallBack(k string) func(e hook.Event) {
	return func(e hook.Event) {
		b.Push(k, k)
	}
}

func hookRegister(k string) {
	hook.Register(hook.KeyDown, []string{k}, hookCallBack(k))
}

func main() {
	runtime.LockOSThread()

	log.Println("开始监听键盘，连续按 3 下 ESC 退出")

	hookRegister("c")
	hookRegister("r")
	hookRegister("l")
	hookRegister("j")
	hookRegister("z")
	hookRegister("n")
	hookRegister("y")

	hook.Register(hook.KeyDown, []string{"m"}, func(e hook.Event) {
		m := "m_short"
		defer func() {
			b.Push("m", m)
		}()
		switch str, start := b.GetAllStr(); str {
		case "jnt", "zyt", "jyt", "znt":
			if time.Now().Unix()-start < 3 {
				m = "m_long"
			}
		}
	})
	hook.Register(hook.KeyDown, []string{"t"}, func(e hook.Event) {
		// 获取 BoundedList List 中最后一个元素
		m := "tiao"
		defer func() {
			b.Push("t", m)
		}()

		lastEle := b.Back()
		if lastEle == nil {
			return
		}

		_, ok := map[string]struct{}{"n": {}, "y": {}}[lastEle.Value.(valWithTime).String()]
		if ok {
			m = "tai"
		}

	})

	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		k := "esc"
		b.Push(k)
		if b.Len() == b.maxSize {
			esc := 0
			l, start := b.GetAll()
			for _, v := range l {
				if v.val == k {
					esc++
				}
			}
			if esc == b.maxSize && time.Now().Unix()-start < 2 {
				hook.End()
			}
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
