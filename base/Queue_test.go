package base

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("队列新增3个元素",*queue)
	queue.Dequeue()
	fmt.Println("队列删除一个元素后",*queue)
	fmt.Println("获得队列的长度:",queue.Size())
}
