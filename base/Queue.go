package base

type Queue struct {
	List []interface{} //元素
}

//队列添加元素的方法
func (queue *Queue) Enqueue(obj interface{}) {
	queue.List = append([]interface{}{obj}, queue.List[0:]...)
}

//删除最早的元素
func (queue *Queue) Dequeue() interface{}{
	q := queue.List[len(queue.List)-1]
	queue.List=queue.List[0:len(queue.List)-1]
	return q
}

//获得队列长度
func (queue *Queue) Size() int {
	return len(queue.List)
}

//获得队列是否是空
func (queue *Queue) IsEmpty() bool {
	return len(queue.List)==0
}


