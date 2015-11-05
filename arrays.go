package main

import "fmt"

func main() {

    //这里我们创建了一个长度为5的数组. 这一组数组的初值是zero-valued。整型就是0
    var a [5]int
    fmt.Println("emp:", a)

    //可以通过array[index] = value语法赋值
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    //内置的len函数会返回数组长度
    fmt.Println("len:", len(a))

    //通过这个语法声明数组的默认初值
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    //数组类型是一维的，但是你可以通过组合创建多维数组结构
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
