package main

import "github.com/geekymax/match_us/util"

func main() {
	girls := LoadUserList(true)
	girls = SortUserList(girls)
	boys := LoadUserList(false)
	boys = SortUserList(boys)
	result := DoMatch(girls, boys)
	util.Infof("match result is %s", util.MustToString(result))
}
