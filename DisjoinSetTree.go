package main

import "fmt"

type DisjoinSetTree struct {


	p *DisjoinSetTree
	rank int
	Value interface{}
}

const offlineminExtract  = -1

func OffLineMinum(seq []int) []int  {
	//归并两个集合
	if len(seq)==0 || seq[0] == offlineminExtract{
		panic("seq类型错误")
	}//构造两个并查集 对数组进行处理

	insertSet := make([]*DisjoinSetTree,0,0)
	values := make([]*DisjoinSetTree,len(seq),cap(seq)) //开辟内存
	n, m := 1,0
	insertSet = append(insertSet,NewDisjoinSetTree(0))  //不断加入节点
	for i:=range seq{
		if seq[i]==offlineminExtract{
			m++

			insertSet= append(insertSet,NewDisjoinSetTree(m))
		}else {
			values[seq[i]] = NewDisjoinSetTree(seq[i])
			insertSet[m] = Union(insertSet[m],values[seq[i]])
			n++
		}
	}

	extractSeq := make([]int,m,m)  //获取最短的
	for i:=1;i<n;i++{
		j := FindSet(values[i]).Value.(int)  //实例化value
		if j!=m{
			extractSeq[j]=i

			for l := j + 1 ; l <= m ; l++{
				if insertSet[l]!=nil{
					//去重
					insertSet[l] = Union(insertSet[l],insertSet[j])
					insertSet[l].Value=l
					insertSet[j]=nil
					break
				}
			}
		}
	}
	return extractSeq
}


func NewDisjoinSetTree(value interface{}) *DisjoinSetTree  {
	t := new(DisjoinSetTree)
	t.Value = value
	t.p = t
	t.rank = 0
	return t
}

func FindSet(dst *DisjoinSetTree) *DisjoinSetTree  {
	if dst.p!= dst{
		dst.p = FindSet(dst.p)
	}

	return dst.p
}

func Union(dst1,dst2 *DisjoinSetTree) *DisjoinSetTree  {
	return Link(FindSet(dst1),FindSet(dst2))
}

func Link(dst1,dst2 *DisjoinSetTree) *DisjoinSetTree  {
	if dst1 !=dst2{
		if dst1.rank < dst2.rank{
			dst1.p = dst2
			return dst2
		}
		dst2.p = dst1  //链接
		if dst1.rank == dst2.rank{
			dst1.rank ++
		}
		return dst1
	}
	return dst1
}

func main()  {
	seq:=[]int{4,8,offlineminExtract,3,offlineminExtract, 9,2,6,offlineminExtract,offlineminExtract,offlineminExtract,
	1,7,offlineminExtract,5}
	exp:=[]int{4,3,2,6,8,1,1}
	last:=OffLineMinum(seq)
	fmt.Println(seq)
	fmt.Println(exp)
	fmt.Println(last)
	fmt.Println(len(seq))
}