package collections

type Stack struct {
    list []interface{}
}

func NewStack() *Stack {
    return &Stack{make([]interface{}, 0)}
}

func (this *Stack) Len() int {
    return len(this.list)
}

func (this *Stack) IsEmpty() bool {
    return this.Len() == 0
}

func (this *Stack) Peek() interface{} {
    if this.IsEmpty() {
        return nil
    }
    return this.list[this.Len() - 1]
}

func (this *Stack) Push(i interface{}) {
    this.list = append(this.list, i)
}

func (this *Stack) Pop() interface{} {
    if this.IsEmpty() {
        return nil
    }
    result := this.Peek()
    this.list = this.list[:this.Len() - 1]
    return result
}
