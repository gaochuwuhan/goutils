package cafe


//求交集
func SliceIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 { //对slice1的元素计数
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times > 0 { //slice1中出现过大于一次的
			nn = append(nn, v)
		}
	}
	return nn
}

//两个切片的差集，s1本身减交集
func SliceDiff(s1,s2 []string) []string{
	intersect:=SliceIntersect(s1,s2)
	m := make(map[string]int)
	nn := make([]string,0)
	for _,v :=range intersect{
		m[v]++
	}
	for _, v:= range s1{
		times,_:=m[v]
		if times == 0{
			nn=append(nn, v)
		}
	}
	return nn
}

//判断某元素是否在 string 切片中
func ElementInSlice(element string,s1 []string) (isIn bool){
	for _,v :=range s1{
		if element==v{
			isIn=true
			return
		}
	}
	return
}

//切片元素去重
func SliceRemoveDup(s1 []string) []string{
	m:=make(map[string]bool)
	for _,v:=range s1{
		m[v]=true
	}
	s2:=make([]string,0)
	for key:=range m{
		s2=append(s2,key)
	}
	return s2
}
