package main

import "fmt"

var Nodes = 6
var LineCounts = 9
var MAPS = [9][3]int{
	{2,4,11},
	{3,5,13},
	{4,6,3},
	{5,6,4},
	{2,3,6},
	{4,5,7},
	{1,2,1},
	{3,4,9},
	{1,3,2},
}

func QuickSort(left , right int)  {
	if left > right{
		return //已经有序了
	}

	i := left
	j := right //左右重合循环终止

	for i!=j{
		for edgeList[j].w >= edgeList[i].w && i < j{
			j--
		}
		for edgeList[i].w <= edgeList[j].w && i < j{
			i ++
		}
		if i<j{
			edgeList[i],edgeList[j] = edgeList[j],edgeList[i]
		}
	}

	QuickSort(left,i-1)
	QuickSort(i+1,right)
}

func getParent(v int)int  {
	if f[v] == v{
		return v
	}
	f[v] = getParent(f[v]) //寻找父节点
	return f[v]
}
//合并两个集合
func merge(s , e int) bool  {
	t1 := getParent(s)
	t2 := getParent(e)
	if t1 != t2{
		f[t2] = t1
		return true
	}
	return false
}

type edge struct {
	s int   //开始点

	e int  //结束点

	w int   //权重点

}

var edgeList [10]edge //边长

var f[7]int  //存储边长数量

var sum,count int //计算总长度 ，边长长度



//把边长排序  每次取最短 思路
func main()  {
	for k , line := range MAPS{
		edgeList[k+1].s = line[0]
		edgeList[k+1].e = line[1]
		edgeList[k+1].w = line[2]
	}
	fmt.Println(edgeList)  //边长排序
	QuickSort(1,LineCounts) //根据权重排序

	//排完序后  初始化集合

	for k , _ := range f{
		f[k] = k  //并集
	}
	for i:=1 ; i<=LineCounts ; i++{
		//这里是 最小生成树 克鲁斯卡尔算法核心
		if merge(edgeList[i].s,edgeList[i].e){
			count++  //计算count
			sum += edgeList[i].w  //计算边长 + 权重
			fmt.Println("way", i-1)
		}

		if count == Nodes -1 {
			break//已经是最短了
		}
	}
	fmt.Println(sum)  //打印最短路线
}