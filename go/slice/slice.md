# Go slice

## 基本使用

slice底层是一个指针和一个长度和容量的变量

```
//runtime/slice.go

type slice struct {
 array unsafe.Pointer
 len   int
 cap   int
}
```

![1664267937369](image/slice/1664267937369.png)
