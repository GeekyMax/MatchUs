package entity

import (
	"errors"
	"github.com/geekymax/match_us/util"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

type Widget interface {
	Parse(param string) error                             // 解析value
	Check(param string) (pass bool, score int, err error) //
}

// 多选, 参数与需求中匹配上一个即可
// 参数为: ["option1", "option2"]
// 需求为: ["option1", "option2"]
type WidgetMulti struct {
	Options []string
}

func (w *WidgetMulti) Parse(param string) error {
	return jsoniter.UnmarshalFromString(param, &w.Options)
}

func (w *WidgetMulti) Check(param string) (pass bool, score int, err error) {
	filters := make([]string, 0)
	err = jsoniter.UnmarshalFromString(param, &filters)
	if err != nil {
		return
	}
	hits := 0
	for _, filter := range filters {
		for _, option := range w.Options {
			if option == filter {
				hits += 1
				break
			}
		}
	}
	return hits > 0, hits * 100, nil
}

// 单选, 参数与需求中匹配上一个即可
// 参数为: "option1"
// 需求为: ["option1", "option2"]
type WidgetSingle struct {
	Option string
}

func (w *WidgetSingle) Parse(param string) error {
	w.Option = param
	return nil
}

func (w *WidgetSingle) Check(param string) (pass bool, score int, err error) {
	filters := make([]string, 0)
	err = jsoniter.UnmarshalFromString(param, &filters)
	if err != nil {
		return
	}
	for _, filter := range filters {
		if filter == w.Option {
			pass = true
			score = 100
		}
	}
	return
}

// 数字, 区间匹配
// 参数为: 2
// 需求为: [1, 3]
type WidgetNumber struct {
	Num int64
}

func (w *WidgetNumber) Parse(param string) error {
	var err error
	w.Num, err = strconv.ParseInt(param, 10, 64)
	return err
}

func (w *WidgetNumber) Check(param string) (pass bool, score int, err error) {
	numberRange := make([]int64, 0)
	err = jsoniter.UnmarshalFromString(param, &numberRange)
	if err != nil {
		return
	}
	if len(numberRange) != 2 {
		err = errors.New("invalid params")
		return
	}
	pass = w.Num >= numberRange[0] && w.Num <= numberRange[1]
	if pass {
		score = 100
	}
	return
}

// todo ... Widget拓展点位
func GetWidgetById(id string) Widget {
	switch id {
	case WidgetTypeMulti:
		return &WidgetMulti{}
	case WidgetTypeSingle:
		return &WidgetSingle{}
	case WidgetTypeNumber:
		return &WidgetNumber{}
	default:
		util.Panicf("can't find Widget for %s", id)
		return nil
	}
}

const (
	WidgetTypeMulti  = "multi"
	WidgetTypeSingle = "single"
	WidgetTypeNumber = "number"
)
