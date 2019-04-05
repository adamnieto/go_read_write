package main

type ChannelMap struct {
  words map[string]int
}

func (c *ChannelMap) Listen() {
}

func (c *ChannelMap) Stop() {
}

func (c *ChannelMap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
	//_ = functor
	//_ = accum_str
	//_ = accum_int
	//return "", 0
  for k, v := range c.words {
    accum_str, accum_int = functor(accum_str, accum_int, k, v)
  }
  return accum_str, accum_int
}

func (c *ChannelMap) AddWord(word string) {
	//_ = word
  c.words[word]++
}

func (c *ChannelMap) GetCount(word string) int {
	//_ = word
	//return 0
  return c.words[word] 
}

func NewChannelMap() *ChannelMap {
	//return &ChannelMap{}
  cm := new(ChannelMap)
  cm.words = make(map[string]int)
  return cm
}
