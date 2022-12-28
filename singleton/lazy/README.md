# Note
```
if log == nil {
    mu.Lock()
    defer mu.Unlock()

    if log == nil {
        log = &logger{Timestamp: time.Now()}
        return log
    }
}
```

这里做了两次 `log == nil` 判断。是因为，如果没有第一次判断，进入临界区的每个 goroutine 都会等待，拿锁，判断，降低性能。  

由于 log 是每个 goroutine 共享的，当有 goroutine 改动 log 时，会反映到其它 goroutine 上。因此，并不需要所有 goroutine 等待拿锁。通过第一次判断过滤不符合条件的 goroutine，从而提升性能。
