package main

type ChannelMap struct {
}

func (c *ChannelMap) Listen() {
}

func (c *ChannelMap) Stop() {
}

func (c *ChannelMap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
	_ = functor
	_ = accum_str
	_ = accum_int
	return "", 0
}

func (c *ChannelMap) AddWord(word string) {
	_ = word
}

func (c *ChannelMap) GetCount(word string) int {
	_ = word
	return 0
}

func NewChannelMap() *ChannelMap {
	return &ChannelMap{}
}
