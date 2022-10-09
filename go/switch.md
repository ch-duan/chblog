switch中默认不需要break，默认不执行下一个分支，只有加上fullthrough才会执行下一个分支

```
switch(x){
  case 1:
    fmt.Println("11")   //x=1时，不执行后面的分支
  case 2:  //默认不执行
    fmt.Println("22")
    fulltrough         //x=2时，会执行case 3
  case 3:
    fmt.Println("33")
}
```
