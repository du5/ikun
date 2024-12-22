package main

import (
	"container/list"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/du5/ikun/mp3"
	hook "github.com/robotn/gohook"
)

type BoundedList struct {
	*list.List
	maxSize int
}

type valWithTime struct {
	val  rune
	time int64
}

func (v valWithTime) String() string {
	return strings.ToLower(string(v.val))
}

func (b *BoundedList) Push(val rune) {
	// 在尾部插入
	b.PushBack(valWithTime{
		val,
		time.Now().Unix(),
	})
	// 超过容量则移除头部（最旧）元素
	if b.Len() > b.maxSize {
		front := b.Front()
		if front != nil {
			b.Remove(front)
		}
	}
}

func (b *BoundedList) GetAll() []valWithTime {
	arr := make([]valWithTime, 0, b.Len())
	for e := b.Front(); e != nil; e = e.Next() {
		arr = append(arr, e.Value.(valWithTime))
	}
	return arr
}

func (b *BoundedList) GetAllStrWithTT() (str string, start int64) {
	temp := b.GetAll()
	for _, v := range temp {
		str += v.String()
	}
	if len(temp) > 0 {
		start = temp[0].time
	}
	return
}

var b *BoundedList

func init() {
	b = &BoundedList{
		list.New(),
		3,
	}
}

func main() {
	// 必须要在主线程上（main goroutine）执行
	runtime.LockOSThread()

	log.Println("开始监听键盘，连续按 3 下 ESC 退出")

	done := make(chan struct{})
	// 开启一个协程监听事件
	go func() {
		events := hook.Start() // 返回一个 channel，里面不断产生键盘/鼠标事件
		defer func() {
			hook.End()
			done <- struct{}{}
		}()

		for e := range events {
			// 判断是否是键盘事件
			if e.Kind == hook.KeyDown {

				k := strings.ToLower(string(e.Keychar))
				switch k {
				case "z", "j":
					mp3.PlayAcc(mp3.Ji)
				case "n", "y":
					mp3.PlayAcc(mp3.Ni)
				case "t":
					lastEle := b.Back()
					if lastEle != nil {
						switch lastEle.Value.(valWithTime).String() {
						case "n", "y":
							mp3.PlayAcc(mp3.Tai)
						default:
							mp3.PlayAcc(mp3.Tiao)
						}
					} else {
						mp3.PlayAcc(mp3.Tiao)
					}
				case "m":
					switch str, start := b.GetAllStrWithTT(); str {
					case "jnt", "zyn", "jyn", "znt":
						if time.Now().Unix()-start < 3 {
							mp3.PlayAcc(mp3.Mei2)
							continue
						}
					}
					mp3.PlayAcc(mp3.Mei)

				case "c":
					mp3.PlayAcc(mp3.Chang)
				case "r":
					mp3.PlayAcc(mp3.Rap)
				case "l":
					mp3.PlayAcc(mp3.Lanqiu)
				case "a":
					// acc.PlayAcc(acc.Aiyou)
				case "g":
					// acc.PlayAcc(acc.Nigangma)
				}
				b.Push(e.Keychar)
				// 连续按 3 次 ESC 退出
				if b.Len() == b.maxSize {
					eixt := 0
					for _, v := range b.GetAll() {
						if v.val == rune(27) {
							eixt++
						}
					}
					if eixt == b.maxSize {
						return
					}
				}
			}
		}
	}()

	// 阻塞主线程
	<-done
}
