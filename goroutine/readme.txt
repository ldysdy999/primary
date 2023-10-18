让出调度：
1、sleep.time()
2、runtime.Gosched()

等待协程：
1、sync.WaitGroup  Add()、Done()、Wait()

加锁：
多个并发协程对同一资源进行修改，加锁，并且尽量小范围加锁
1、sync.Mutex  Lock() Unlock()

管道：
struct{} 表示一个空的结构体，不包含任何字段，因此它的大小为 0。不会占用任何内存。这使得 chan struct{} 成为一种用于同步和信号传递的轻量级方法
要声明一个空结构体类型的通道，您需要使用 chan struct{}，而不是 chan struct，这是因为 struct 是一个关键字，用于定义结构体类型。