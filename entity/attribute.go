package entity

import (
	"github.com/geekymax/match_us/util"
)

type AttributeDefinition struct {
	WidgetType string
	Bias       string
}

var AttributeDefinitionMap = map[string]AttributeDefinition{
	"age": {
		WidgetType: WidgetTypeNumber,
	},
	"grade": {
		WidgetType: WidgetTypeNumber,
	},
	"college": {
		WidgetType: WidgetTypeSingle,
	},
	"height": {
		WidgetType: WidgetTypeNumber,
	},
	"weight": {
		WidgetType: WidgetTypeNumber,
	},
	"body": {
		WidgetType: WidgetTypeSingle,
	},
	// ... 拓展点位
}

func GetAttributeDefinition(id string) AttributeDefinition {
	def, ok := AttributeDefinitionMap[id]
	if !ok {
		util.Panicf("can't find AttributeDefinition for %s", id)
	}
	return def
}
