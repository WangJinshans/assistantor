package core

// 本地延时队列
// zrange score为过期时间戳
// 不停遍历 取出score小于当前时间的任务

type DelayQueue struct {
	BatchMode  bool // 异步批量  同步此参数无意义
	AsyncQueue chan Message
}

func NewDelayQueue(conf *Config) {
	// 连接redis
}

type Topic struct {
	Name string
}

type Message struct {
	MessageId string
	Message   interface{}
	Expire    int64
}

type Config struct {
	BatchSize int64
	LinerMs   int
}

func (queue *DelayQueue) OnProduce(messageId string, err error) {

}

func (queue *DelayQueue) PushMessage(message Message) {

}

func (queue *DelayQueue) SyncPushMessage(message Message) {
	// 依赖于参数

}
