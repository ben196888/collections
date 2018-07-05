package collections

type Queue struct {
    list []interface{}
}

func NewQueue() *Queue {
    return &Queue{make([]interface{}, 0)}
}

func (this *Queue) Len() int {
    return len(this.list)
}

func (this *Queue) IsEmpty() bool {
    return this.Len() == 0
}

func (this *Queue) Peek() interface{} {
    if this.IsEmpty() {
        return nil
    }
    return this.list[this.Len() - 1]
}

func (this *Queue) Enquque(i interface{}) {
    this.list = append(this.list, i)
}

func (this *Queue) Dequeue() interface{} {
    if this.IsEmpty() {
        return nil
    }
    result := this.Peek()
    this.list = this.list[1:]
    return result
}
