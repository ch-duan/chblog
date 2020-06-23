# Go map初识

在学习的过程中遇到一些问题总是记不住，写文章记录并帮助自己理解相关知识。如有错误，敬请指正。

## map基本使用

map是一种无序的键值对的集合，未初始化的值为nil，如果获取不存在的键返回的值为值相应类型的零值。

```Go
var q map[string]int
w:=make(map[string]int)
fmt.Println(q,w,q==nil,w==nil)
```

结果如下：

```Go
map[] map[] true false
```

这里打印出map[]根据go源码print.go中723行printValue函数看出go打印会给map自动加上map[]
如果想知道map中是否有key存在可以用下面的方式

```Go
    e := make(map[string]int, 10)
    e["qqq"] = 1
    value, ok := e["qqq"]
    value2, ok2 := e["www"]
    fmt.Println(value, ok, value2, ok2)
```

结果如下：不存在时返回相应value类型的零值

```Go
1 true 0 false
```

map是引用类型，只能和nil比较判断map是否为nil，另外有len()函数计算map长度，delete()函数来删除相应键值对

```Go
    e := make(map[string]int, 10)
    e["qqq"] = 1
    fmt.Println(len(e))
    delete(e, "qqq")
    value, ok = e["qqq"]
    fmt.Println(value, ok)
```

结果如下：

```Go
1
0 false
```
