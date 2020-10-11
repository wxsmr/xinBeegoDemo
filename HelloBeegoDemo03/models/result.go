package models

//公用的用于返回的结构体的定义，Data可以表示任意类型
type Result struct {
	Code int //接口返回状态类型 1表示成功 ， 0表示失败  2表示其他
	Message string //接口返回状态对应的描述信息
	Data interface{}//返回的数据
}