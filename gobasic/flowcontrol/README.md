# 流程控制语句

## if 语句

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

## for 循环

```go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

- for 循环可以通过`break`、`goto`、`return`、`panic`语句强制退出循环
- 使用`for range`遍历数组、切片、字符串、map 及通道（channel），通过for range遍历的返回值有以下规律：
  1. 数组、切片、字
  2. 符串返回索引和值
  3. map返回键和值通道（channel）只返回通道内的值

## switch case

```go
func testSwitch() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
```

- 每个`switch`只能有一个`default`分支
- 一个分支可以有多个值，多个`case`值中间使用英文逗号分隔
- 分支还可以使用表达式，这时候`switch`语句后面不需要再跟判断变量
- `fallthrough`语法可以执行满足条件的case的下一个case

## goto 语句

- `goto`语句通过标签进行代码间的无条件跳转

## break 语句

- `break`语句可以结束`for`、`switch`和`select`的代码块
- `break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和 `select`的代码块上

## continue 语句

- `continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限在`for`循环内使用
- 在 `continue`语句后添加标签时，表示开始标签对应的循环

