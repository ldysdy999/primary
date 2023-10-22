例程数量：runtime.NumGoroutine()

让出调度：
1、sleep.time()
2、runtime.Gosched()

等待协程：
1、sync.WaitGroup  Add()、Done()、Wait()
2、chan struct{}

加锁：
多个并发协程对同一资源进行修改，加锁，并且尽量小范围加锁
1、sync.Mutex  Lock() Unlock()

原子操作：
sync/atomic

管道阻塞：
1、struct{} 表示一个空的结构体，不包含任何字段，因此它的大小为 0。不会占用任何内存。这使得 chan struct{} 成为一种用于同步和信号传递的轻量级方法
2、要声明一个空结构体类型的通道，您需要使用 chan struct{}，而不是 chan struct，这是因为 struct 是一个关键字，用于定义结构体类型。

3、context.Context用于在不同的协程之间传递请求范围的元数据、超时、取消信号等信息，实现跨协程的数据共享和协作
  context.WithTimeout 、 优雅退出：context.WithCancel --> "cancle()" --> "<-ctx.Done()"
  详情：https://www.yuque.com/kubesre/golang/zc3lholhxgy4zti4#%E3%80%8A30.Golang%E4%B8%8A%E4%B8%8B%E6%96%87context%E3%80%8B
4、signal.Notify(c, syscall.SIGINT|syscall.SIGTERM) 当接受到终止信号，c管道就有信息了
5、ticker:=time.NewTicker(1 * time.Second) 每一秒给ticker.C管道信息
   <-time.After(3 * time.Second) 3秒后给管道信息


