// 四：ringbuffer + CAS + 消息状态 + channel
package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// 消息结构体
type Message struct {
	ID      int
	Payload string
	Status  uint32 // 0 表示未处理/无消息; 1 表示处理完成； 2 表示已提交ACK
}

const (
	statusNew uint32 = iota
	statusFinished
	statusAcked
	bufferSize = 50 // 缓存数组大小
)

type ringBuffer struct {
	buffer                 [bufferSize]*Message      // 环形消息缓冲数组
	pullIndex, commitIndex uint32                    // 拉取和提交双指针
	handleChan             [bufferSize]chan *Message // 存储 Message 指针的双向通道
}

var (
	once sync.Once
	rb   *ringBuffer
)

func GetRingBuffer() *ringBuffer {
	once.Do(func() {
		rb = &ringBuffer{}
		for i := range rb.handleChan {
			rb.handleChan[i] = make(chan *Message, bufferSize)
			go HandleChanMsg(rb.handleChan[i])
		}
	})
	return rb
}

// Pull 拉取消息
func (rb *ringBuffer) Pull(msg *Message) bool {
	rb.pullIndex = rb.pullIndex % bufferSize
	for {
		// 当 pull 指针指向的位置为空 或 指向的旧消息已经提交，则可以新拉消息
		if rb.buffer[rb.pullIndex] == nil || atomic.LoadUint32(&rb.buffer[rb.pullIndex].Status) == statusAcked {
			// 新拉下的 Message 状态默认为 0，即 statusNew，因此无需 CAS 状态
			rb.buffer[rb.pullIndex] = msg
			rb.handleChan[msg.ID%bufferSize] <- msg
			log.Printf("%v 已加入缓冲圈\n", msg)
			rb.pullIndex++
			return true
		}
	}
}

// Commit 提交消息
func (rb *ringBuffer) Commit() {
	for {
		rb.commitIndex = rb.commitIndex % bufferSize
		// 当 commit 指针指向的位置为空 或者 指向的旧消息还没处理完成，则不能提交
		if rb.buffer[rb.commitIndex] == nil || atomic.LoadUint32(&rb.buffer[rb.commitIndex].Status) != statusFinished {
			continue
		}
		msg := rb.buffer[rb.commitIndex]
		log.Printf("%v 已提交\n", msg)
		// 省略 Kafka 手动提交commitIndex位置的消息，接下来移动 commit指针....
		time.Sleep(50 * time.Millisecond)
		atomic.CompareAndSwapUint32(&msg.Status, statusFinished, statusAcked)
		rb.commitIndex++
	}
}

func HandleChanMsg(c chan *Message) {
	for msg := range c {
		// 省略处理逻辑...
		time.Sleep(500 * time.Millisecond)
		atomic.CompareAndSwapUint32(&msg.Status, statusNew, statusFinished)
		log.Printf("%v 已成功处理\n", msg)
	}
}

func main() {
	// 初始化
	b := GetRingBuffer()

	go func() {
		b.Commit()
	}()

	// 模拟生成消息并入队
	const numMessages = 100
	for i := 0; i < numMessages; i++ {
		msg := &Message{
			ID:      i,
			Payload: "Message Contents",
			Status:  statusNew,
		}
		log.Printf("生成消息 %v\n", msg)
		b.Pull(msg)
	}
	fmt.Scanln()
}

// 三：ringbuffer + CAS + 消息状态
// package main

// import (
// 	"log"
// 	"sync"
// 	"sync/atomic"
// )

// type Message struct {
// 	ID         int
// 	Contents   string
// 	Processing int32
// }

// const bufferSize = 10

// var ringBuffer [bufferSize]*Message
// var readIndex, writeIndex int32

// func init() {
// 	log.Println("初始化")
// 	for i := range ringBuffer {
// 		ringBuffer[i] = &Message{}
// 		log.Printf("初始化 %d\n", i)
// 	}

// }

// // func enqueue(msg *Message) bool {
// // 	nextWriteIndex := atomic.AddInt32(&writeIndex, 1) % bufferSize
// // 	if atomic.LoadInt32(&ringBuffer[nextWriteIndex].Processing) == 0 {
// // 		if atomic.SwapInt32(&ringBuffer[nextWriteIndex].Processing, 1) == 0 {
// // 			ringBuffer[nextWriteIndex] = msg
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func dequeue() *Message {
// // 	nextReadIndex := atomic.AddInt32(&readIndex, 1) % bufferSize
// // 	msg := ringBuffer[nextReadIndex]
// // 	if atomic.LoadInt32(&msg.Processing) == 2 {
// // 		atomic.StoreInt32(&msg.Processing, 0)
// // 		return msg
// // 	}
// // 	return nil
// // }

// func enqueue(msg *Message) bool {
// 	for {
// 		currentWriteIndex := atomic.LoadInt32(&writeIndex)
// 		nextWriteIndex := (currentWriteIndex + 1) % bufferSize
// 		if atomic.LoadInt32(&ringBuffer[nextWriteIndex].Processing) == 0 {
// 			if atomic.CompareAndSwapInt32(&writeIndex, currentWriteIndex, nextWriteIndex) {
// 				if atomic.SwapInt32(&ringBuffer[nextWriteIndex].Processing, 1) == 0 {
// 					ringBuffer[nextWriteIndex] = msg
// 					return true
// 				}
// 				atomic.StoreInt32(&ringBuffer[nextWriteIndex].Processing, 0) // 重置处理状态
// 			}
// 		}
// 	}
// }

// func dequeue() *Message {
// 	for {
// 		currentReadIndex := atomic.LoadInt32(&readIndex)
// 		nextReadIndex := (currentReadIndex + 1) % bufferSize
// 		msg := ringBuffer[nextReadIndex]
// 		if atomic.LoadInt32(&msg.Processing) == 2 {
// 			if atomic.CompareAndSwapInt32(&readIndex, currentReadIndex, nextReadIndex) {
// 				atomic.StoreInt32(&msg.Processing, 0) // 重置处理状态
// 				return msg
// 			}
// 		} else {
// 			return nil
// 		}
// 	}
// }

// func processMessages() {
// 	for {
// 		msg := dequeue()
// 		if msg == nil {
// 			continue
// 		}
// 		// 处理消息
// 		processMessage(msg)
// 		// 提交 ACK
// 		ackMessage(msg)
// 	}
// }

// func processMessage(msg *Message) {
// 	// 处理消息逻辑
// }

// func ackMessage(msg *Message) {
// 	// 提交 ACK
// 	atomic.StoreInt32(&msg.Processing, 2)
// }

// func main() {
// 	// 启动消息处理协程
// 	var wg sync.WaitGroup
// 	const numWorkers = 10
// 	wg.Add(numWorkers)
// 	for i := 0; i < numWorkers; i++ {
// 		go func() {
// 			defer wg.Done()
// 			processMessages()
// 		}()
// 	}

// 	// 模拟生成消息并入队
// 	const numMessages = 100
// 	for i := 0; i < numMessages; i++ {
// 		msg := &Message{
// 			ID:         i,
// 			Contents:   "Message Contents",
// 			Processing: 0,
// 		}

// 		// 尝试将消息入队，直到成功
// 		for !enqueue(msg) {
// 			// 阻塞等待空位
// 		}
// 	}

// 	// 等待所有协程完成
// 	wg.Wait()
// }

// 二：这里用 channel 实现一个内存队列，保证消费
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"time"
// )

// const (
// 	BufferSize    = 10  // 环形缓冲区大小
// 	NumConsumers  = 5   // 消费者数量
// 	NumPartitions = 10  // 分区数量
// 	NumMessages   = 100 // 消息总数
// )

// type Message struct {
// 	ID       int
// 	Topic    string
// 	Contents string
// }

// func main() {
// 	buffer := make(chan Message, BufferSize)
// 	done := make(chan struct{})

// 	// 启动异步消息处理协程
// 	var wg sync.WaitGroup
// 	wg.Add(NumConsumers)
// 	for i := 0; i < NumConsumers; i++ {
// 		go consumeMessages(buffer, &wg)
// 	}

// 	// 生成测试消息并发送到缓冲区
// 	go produceMessages(buffer, done)

// 	// 等待中断信号
// 	waitForInterrupt()

// 	// 发送停止信号给消息生成器和消费者
// 	close(done)

// 	// 等待所有消费者完成
// 	wg.Wait()
// }

// func consumeMessages(buffer chan Message, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for msg := range buffer {
// 		// 处理消息逻辑
// 		log.Printf("Processing message: ID=%d, Topic=%s, Contents=%s\n", msg.ID, msg.Topic, msg.Contents)

// 		// 模拟耗时操作
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func produceMessages(buffer chan Message, done chan struct{}) {
// 	for i := 0; i < NumMessages; i++ {
// 		msg := Message{
// 			ID:       i,
// 			Topic:    getTopic(i),
// 			Contents: "Message Contents",
// 		}

// 		select {
// 		case buffer <- msg:
// 			// 消息成功发送到缓冲区
// 		case <-done:
// 			// 收到停止信号，退出生成消息循环
// 			return
// 		}
// 	}

// 	close(buffer)
// }

// func getTopic(messageID int) string {
// 	// 假设根据消息ID进行哈希计算得到Topic
// 	return "Topic-" + fmt.Sprint(messageID%NumPartitions)
// }

// func waitForInterrupt() {
// 	// 等待中断信号以优雅地关闭程序
// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, os.Interrupt)
// 	<-signals
// }

// 一：初始版本，ringbuffer 消费消息，但是这里的实现有问题，ringbuffer 满了直接丢消息而不会阻塞等待消息处理完再拉消息
// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"
// )

// const (
// 	BufferSize    = 10   // 环形缓冲区大小
// 	NumConsumers  = 5    // 消费者数量
// 	NumPartitions = 10   // 分区数量
// 	NumMessages   = 1000 // 消息总数
// )

// type Message struct {
// 	ID       int
// 	Topic    string
// 	Contents string
// }

// type RingBuffer struct {
// 	Buffer   []*Message
// 	ReadPos  int
// 	WritePos int
// 	Full     bool
// }

// func NewRingBuffer(size int) *RingBuffer {
// 	return &RingBuffer{
// 		Buffer: make([]*Message, size),
// 	}
// }

// func (rb *RingBuffer) Push(msg *Message) {
// 	if rb.Full {
// 		log.Printf("Buffer is full, dropping message: %v\n", msg)
// 		return
// 	}

// 	rb.Buffer[rb.WritePos] = msg
// 	rb.WritePos = (rb.WritePos + 1) % len(rb.Buffer)

// 	if rb.WritePos == rb.ReadPos {
// 		rb.Full = true
// 	}
// }

// func (rb *RingBuffer) Pop() *Message {
// 	if !rb.Full && rb.ReadPos == rb.WritePos {
// 		return &Message{}
// 	}

// 	msg := rb.Buffer[rb.ReadPos]
// 	rb.ReadPos = (rb.ReadPos + 1) % len(rb.Buffer)
// 	rb.Full = false

// 	return msg
// }

// func main() {
// 	buffer := NewRingBuffer(BufferSize)

// 	// 启动异步消息处理协程
// 	var wg sync.WaitGroup
// 	wg.Add(NumConsumers)
// 	for i := 0; i < NumConsumers; i++ {
// 		go consumeMessages(buffer, &wg)
// 	}

// 	// 生成测试消息并入队
// 	for i := 0; i < NumMessages; i++ {
// 		msg := &Message{
// 			ID:       i,
// 			Topic:    getTopic(i),
// 			Contents: "Message Contents",
// 		}
// 		buffer.Push(msg)
// 		if i >= 10 {
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}

// 	// 等待所有消费者完成
// 	wg.Wait()
// 	fmt.Scanln()
// }

// func consumeMessages(buffer *RingBuffer, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for {
// 		msg := buffer.Pop()
// 		// 处理消息逻辑
// 		log.Printf("Processing message: ID=%d, Topic=%s, Contents=%s\n", msg.ID, msg.Topic, msg.Contents)

// 		// 模拟耗时操作
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func getTopic(messageID int) string {
// 	// 假设根据消息ID进行哈希计算得到Topic
// 	return "Topic-" + string(messageID%NumPartitions)
// }
