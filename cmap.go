package main

type addOp struct {
  key string
  response chan int
}

type askOp struct {
  key string
  response chan int
}

type reduceOp struct {
  functr ReduceFunc
  acc_str string
  acc_int int
  completed chan int
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
          add_obj.response <- 1
      
      case ask_obj := <- c.ask_channel:
          //fmt.Printf("ask: %s\n", ask_obj)
          ask_obj.response <- c.words[ask_obj.key]
      
      case reduce_obj := <- c.reduce_channel:
          //fmt.Printf("reduce: %s\n", reduce_obj)
          for k, v := range c.words {
            reduce_obj.acc_str, reduce_obj.acc_int = reduce_obj.functr(reduce_obj.acc_str, reduce_obj.acc_int, k, v)
          }
          reduce_obj.completed <- 1
      case <- c.stop_channel:
          //fmt.Printf("stop_channel: %d\n", exit)
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
	//_ = functor
	//_ = accum_str
	//_ = accum_int
	//return "", 0
  //for k, v := range c.words {
    //accum_str, accum_int = functor(accum_str, accum_int, k, v)
  //}
  //return accum_str, accum_int

  reduce_obj := &reduceOp{
    functr: functor,
    acc_str: accum_str,
    acc_int: accum_int,
    completed: make(chan int)}
  
  c.reduce_channel <- reduce_obj
  <- reduce_obj.completed
  return reduce_obj.acc_str,reduce_obj.acc_int

}

func (c *ChannelMap) AddWord(word string) {
  add_obj := &addOp{
    key: word, 
    response: make(chan int)}
  c.add_channel <- add_obj
  <- add_obj.response
  //c.words[word]++
  //<- c.add_channel
}

func (c *ChannelMap) GetCount(word string) int {
  //res := c.words[word]
  //<- c.ask_channel
  //return res
  //return 0

  ask_obj := &askOp{
    key: word,
    response: make(chan int)}
  c.ask_channel <- ask_obj
  res := <- ask_obj.response
  return res
}

func NewChannelMap() *ChannelMap {
	//return &ChannelMap{}
  cm := new(ChannelMap)
  cm.words = make(map[string]int)
  cm.add_channel = make(chan *addOp, ADD_BUFFER_SIZE)
  cm.ask_channel = make(chan *askOp, ASK_BUFFER_SIZE)
  cm.reduce_channel = make(chan *reduceOp, REDUCE_BUFFER_SIZE)
  
  cm.stop_channel = make(chan int)
  return cm
}
