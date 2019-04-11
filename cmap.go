package main

type addOp struct {
  key string
  completed_channel chan int
}

type askOp struct {
  key string
  response_channel chan int
}

type result struct {
  acc_str string
  acc_int int
}

type reduceOp struct {
  functr ReduceFunc
  acc_str string
  acc_int int
  response_channel chan *result
}


type ChannelMap struct {
  words map[string]int // resource

  stop_channel chan int
  add_channel chan *addOp
  ask_channel chan *askOp
  reduce_channel chan *reduceOp
}

func (c *ChannelMap) Listen() {
  for {
    select {
      
      case add_obj := <- c.add_channel:
          //fmt.Printf("add: %s\n", add_obj)
          c.words[add_obj.key]++
          add_obj.completed_channel <- 1
      
      case ask_obj := <- c.ask_channel:
          //fmt.Printf("ask: %s\n", ask_obj)
          ask_obj.response_channel <- c.words[ask_obj.key]
      
      case reduce_obj := <- c.reduce_channel:
          //fmt.Printf("reduce: %s\n", reduce_obj)
          for k, v := range c.words {
            reduce_obj.acc_str, reduce_obj.acc_int = reduce_obj.functr(reduce_obj.acc_str, reduce_obj.acc_int, k, v)
          }
          result_obj := &result {
            acc_str: reduce_obj.acc_str,
            acc_int: reduce_obj.acc_int}
          reduce_obj.response_channel <- result_obj
      
      case <- c.stop_channel:
          close(c.add_channel)
          close(c.ask_channel)
          close(c.reduce_channel)
          return
      }
   } 
}

func (c *ChannelMap) Stop() {
  c.stop_channel <- 1
}

func (c *ChannelMap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
  reduce_obj := &reduceOp{
    functr: functor,
    acc_str: accum_str,
    acc_int: accum_int,
    response_channel: make(chan *result)}
  //fmt.Printf("Before: acc_int: %d\n", accum_int)
  //fmt.Printf("Before: acc_str: %s\n", accum_str)
  c.reduce_channel <- reduce_obj
  result_obj := <- reduce_obj.response_channel
  //fmt.Printf("After: acc_int: %d\n", result_obj.acc_int)
  //fmt.Printf("After: acc_str: %s\n", result_obj.acc_str)
  return result_obj.acc_str,result_obj.acc_int

}

func (c *ChannelMap) AddWord(word string) {
  add_obj := &addOp{
    key: word, 
    completed_channel: make(chan int)}
  c.add_channel <- add_obj
  <- add_obj.completed_channel
}

func (c *ChannelMap) GetCount(word string) int {
  ask_obj := &askOp{
    key: word,
    response_channel: make(chan int)}
  c.ask_channel <- ask_obj
  res := <- ask_obj.response_channel
  return res
}

func NewChannelMap() *ChannelMap {
  cm := new(ChannelMap)
  cm.words = make(map[string]int)
  cm.add_channel = make(chan *addOp, ADD_BUFFER_SIZE)
  cm.ask_channel = make(chan *askOp, ASK_BUFFER_SIZE)
  cm.reduce_channel = make(chan *reduceOp, REDUCE_BUFFER_SIZE)
  cm.stop_channel = make(chan int)
  return cm
}
