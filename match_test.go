package main

import (
	"github.com/geekymax/match_us/entity"
	"github.com/geekymax/match_us/util"
	"testing"
)

func TestDoMatch(t *testing.T) {
	girls := []*entity.User{
		{
			Id:        "g1",
			BasicInfo: entity.BasicInfo{},
			Attributes: map[string]string{
				"age":     "20",
				"grade":   "3", // 从大一开始算的定量
				"college": "cs",
				"height":  "160",
				"weight":  "50",
			},
			Requirements: map[string]entity.Requirement{
				"age": {
					Value:  "[18,23]",
					Must:   true,
					Weight: 10,
				},
				"height": {
					Value:  "[178,190]",
					Must:   true,
					Weight: 10,
				},
				"college": {
					Value:  "[\"cs\",\"ee\"]",
					Must:   false,
					Weight: 3,
				},
			},
		},
	}
	boys := []*entity.User{
		{
			Id:        "b1",
			BasicInfo: entity.BasicInfo{},
			Attributes: map[string]string{
				"age":     "24",
				"grade":   "3", // 从大一开始算的定量
				"college": "cs",
				"height":  "170",
				"weight":  "68",
			},
			Requirements: map[string]entity.Requirement{
				"age": {
					Value:  "[18,23]",
					Must:   true,
					Weight: 10,
				},
			},
		},
		{
			Id:        "b2",
			BasicInfo: entity.BasicInfo{},
			Attributes: map[string]string{
				"age":     "22",
				"grade":   "3",
				"college": "unknown",
				"height":  "180",
				"weight":  "75",
			},
			Requirements: map[string]entity.Requirement{
				"age": {
					Value:  "[18,23]",
					Must:   true,
					Weight: 10,
				},
			},
		},
		{
			Id:        "b3",
			BasicInfo: entity.BasicInfo{},
			Attributes: map[string]string{
				"age":     "22",
				"grade":   "3",
				"college": "cs",
				"height":  "180",
				"weight":  "75",
			},
			Requirements: map[string]entity.Requirement{
				"age": {
					Value:  "[18,23]",
					Must:   true,
					Weight: 10,
				},
			},
		},
	}
	result := DoMatch(girls, boys)
	util.Infof("match result is %s", util.MustToString(result))
}
