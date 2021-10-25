package entity

type User struct {
	Id     string
	Number int // 性别内, 唯一编号
	BasicInfo
	Attributes   map[string]string      // attributeId -> value
	Requirements map[string]Requirement // attributeId -> requirement
}

// 基本信息, 不进行评分
type BasicInfo struct {
	Name   string
	Sex    int // 1为男生,2为女生
	Phone  string
	Wechat string
}

type Requirement struct {
	Value  string // 要求的值
	Must   bool   // 是否必须满足
	Weight int    // 权重
}
