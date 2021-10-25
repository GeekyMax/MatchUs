package main

import (
	"github.com/geekymax/match_us/entity"
	"github.com/geekymax/match_us/util"
)

// 核心函数
func DoMatch(girls []*entity.User, boys []*entity.User) (result [][]string) {
	result = make([][]string, 0)
	// 开始进行筛选
	for _, girl := range girls {
		// 获取匹配的男孩
		matchedBoy := GetMatchedBoy(girl, boys)
		if matchedBoy != nil {
			// 如果有匹配到的男孩,则从列表中删除
			util.Infof("[%s]match boy[%s]for girl", girl.Id, matchedBoy.Id)
			boys = deleteUserFromList(boys, matchedBoy)
			result = append(result, []string{girl.Id, matchedBoy.Id})
		} else {
			// 没有匹配到,打个日志
			util.Infof("[%s]can't match boy for girl", girl.Id)
			result = append(result, []string{girl.Id, "NONE"})

		}
	}
	return
}

func deleteUserFromList(list []*entity.User, user *entity.User) []*entity.User {
	result := make([]*entity.User, 0)
	for _, e := range list {
		if e.Id != user.Id {
			result = append(result, e)
		}
	}
	return result
}

// 获取与女孩最匹配的男孩:
// 1. 必须双向匹配
// 2. 匹配中分数最高的男生
func GetMatchedBoy(girl *entity.User, boys []*entity.User) (matchedBoy *entity.User) {
	maxScore := 0
	for _, boy := range boys {
		girlLikeBoy, score := Match(boy, girl)
		boyLikeGirl, _ := Match(girl, boy)
		if girlLikeBoy && boyLikeGirl && score > maxScore {
			maxScore = score
			matchedBoy = boy
		}
	}
	return
}

// 用户A是否符合用户B的条件,以及评分多少分
func Match(userA *entity.User, userB *entity.User) (match bool, score int) {
	match = true
	for attributeId, requirement := range userB.Requirements {
		attributeDefinition := entity.GetAttributeDefinition(attributeId)
		data := userA.Attributes[attributeId]
		widget := entity.GetWidgetById(attributeDefinition.WidgetType)
		err := widget.Parse(data)
		if err != nil {
			util.Errorf("Parse widget[%s] for data[%s] error[%v]", attributeDefinition.WidgetType, data, err)
			continue
		}
		thisMatch, thisScore, err := widget.Check(requirement.Value)
		if err != nil {
			util.Errorf("Check widget[%s] for data[%s] filter[%s] error[%v]", attributeDefinition.WidgetType, requirement.Value, data, err)
			continue
		}
		if requirement.Must {
			match = thisMatch && match
		}
		score += thisScore * requirement.Weight
	}
	return
}
