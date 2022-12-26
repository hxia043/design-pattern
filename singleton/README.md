# 0. singleton introduction
singleton is to build a globally unique object. The singleton can implementation With the `sync.Once` of `sync` package in Go. 

# 1. example
The example of the singleton implementation as:
```
var oncelogger sync.Once

func NewLogger() *logger {
	if log == nil {
		oncelogger.Do(func() {
			log = &logger{timestamp: time.Now()}
		})

		return log
	}

	return log
}
```

Beside the singleton implement, the most useful case is to lazy initialization the env/config and parameters. The positive to use the `sync.Once` is:
1. To lazy initialization, is not heavy and quick than init function in package. Which can init the info till the info used.
2. Support concurrency with the more fine granularity of `lock`.

# 2. code analysis of sync.Once
Let's see if the `sync.Once` can support concurrency, then the `lock` is required. And to globally initialization, a flag to decide whether the action performed is needed.

With the though, check the struct of `sync.Once`:
```
type Once struct {
	done uint32
	m    Mutex
}
```

`done` is the flag, and `m` is the lock.

See the `sync.Do` function, with an atomic operation to keep the uniqueness. And then into the `doSlow` method, all goroutines can going to the method, so in the method the first is to add `lock`, then one goroutine to execute the `f` in critical area, after that make `done` with non-zero by atomic operation:
```
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

# 1. Reference
- [设计模式学习笔记](https://www.cnblogs.com/xingzheanan/p/16091910.html)
- [Go sync.Once](https://geektutu.com/post/hpg-sync-once.html)
- [sync.Once 惰性初始化](https://books.studygolang.com/gopl-zh/ch9/ch9-05.html)
