package prototype

//注：本例子采用的是浅拷贝
//定义抽象可复制接口
type Prototype interface {
	Clone()
}

//具体实现
type Resume struct {
	Name string
	Sex  string
}

func (r *Resume) Clone() Resume {
	return Resume{
		Name: r.Name,
		Sex:  r.Sex,
	}
}
