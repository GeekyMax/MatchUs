package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/geekymax/match_us/entity"
	"github.com/geekymax/match_us/util"
	jsoniter "github.com/json-iterator/go"
	"sort"
	"strconv"
	"strings"
)

// todo 加载用户,包括读文件,数据解析等
func LoadUserList(isGirl bool) []*entity.User {
	path := "file/dataset.xlsx"
	sheetName := "boy"
	if isGirl {
		sheetName = "girl"
	}
	table, err := loadExcel(path, sheetName)
	util.PanicIfError(err)
	users := make([]*entity.User, 0)
	for index, row := range table {
		user := &entity.User{
			Id:     fmt.Sprintf("%s_%d", sheetName, index),
			Number: index,
			BasicInfo: entity.BasicInfo{
				Name:   row["name"],
				Sex:    0,
				Phone:  "",
				Wechat: "",
			},
			Attributes:   map[string]string{},
			Requirements: map[string]entity.Requirement{},
		}
		for columnName, parser := range ColumnParser {
			value, ok := row[columnName]
			if ok {
				user.Attributes[columnName] = parser(value)
			}
		}
		for columnName, parser := range RequirementParser {
			value, ok := row[fmt.Sprintf("r_%s", columnName)]
			if ok {
				user.Requirements[columnName] = entity.Requirement{
					Value:  parser(value),
					Must:   false,
					Weight: 10,
				}
			}
		}
		users = append(users, user)
	}
	return users
}

// todo 拓展点位,数据解析工作
var ColumnParser = map[string]func(string) string{
	"age":    doNothing,
	"height": doNothing,
	"body":   doNothing,
}

// todo 拓展点位,要求数据解析工作
var RequirementParser = map[string]func(string) string{
	"age":    rangeParser,
	"height": rangeParser,
	"body":   listParser,
}

func doNothing(s string) string {
	return s
}

func listParser(s string) string {
	l := strings.Split(strings.Trim(s, " "), "┋")
	r, _ := jsoniter.MarshalToString(l)
	return r
}
func rangeParser(s string) string {
	l := strings.Split(strings.Trim(s, " "), ",")
	if len(l) != 2 {
		return "[0,100000]"
	}
	min, err := strconv.ParseInt(l[0], 10, 64)
	if err != nil {
		min = 0
	}
	max, err := strconv.ParseInt(l[1], 10, 64)
	if err != nil {
		max = 10000
	}
	r, _ := jsoniter.MarshalToString([]int64{min, max})
	return r
}

// 加载excel文件
func loadExcel(path string, sheetName string) (table []map[string]string, err error) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	rows := file.GetRows(sheetName)
	util.PanicIf(len(rows) < 2)
	header := rows[0]
	rows = rows[1:]
	table = make([]map[string]string, 0)
	for _, row := range rows {
		data := map[string]string{}
		for colNum, colName := range header {
			data[colName] = row[colNum]
		}
		table = append(table, data)
	}
	return table, nil
}

// 对用户进行排序,默认按照编号进行排序
func SortUserList(users []*entity.User) []*entity.User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Number < users[j].Number
	})
	return users
}
