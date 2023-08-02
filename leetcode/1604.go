package main

import (
	"fmt"
)

/*
1604. 警告一小时内使用相同员工卡大于等于三次的人
力扣公司的员工都使用员工卡来开办公室的门。每当一个员工使用一次他的员工卡，安保系统会记录下员工的名字和使用时间。如果一个员工在一小时时间内使用员工卡的次数大于等于三次，这个系统会自动发布一个 警告 。
给你字符串数组 keyName 和 keyTime ，其中 [keyName[i], keyTime[i]] 对应一个人的名字和他在 某一天 内使用员工卡的时间。
使用时间的格式是 24小时制 ，形如 "HH:MM" ，比方说 "23:51" 和 "09:49" 。
请你返回去重后的收到系统警告的员工名字，将它们按 字典序升序 排序后返回。
请注意 "10:00" - "11:00" 视为一个小时时间范围内，而 "23:51" - "00:10" 不被视为一小时内，因为系统记录的是某一天内的使用情况。
输入：keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"], keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
输出：["daniel"]
解释："daniel" 在一小时内使用了 3 次员工卡（"10:00"，"10:40"，"11:00"）
*/
func main() {
	keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	ans := alertNames(keyName, keyTime)
	fmt.Sprint("asdadasd : %s", ans)

}

func alertNames(keyName []string, keyTime []string) []string {
	nameTimesMap := map[string][]int{}
	for i, v := range keyName {
		time := keyTime[i]
		hour := int(time[0]-'0')*10 + int(time[1])
		minute := int(time[3]-'0')*10 + int(time[4])
		nameTimesMap[v] = append(nameTimesMap[v], hour*60+minute)
	}

	return []string{}
}
