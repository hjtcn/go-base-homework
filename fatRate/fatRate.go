package fatRate

import "fmt"

/**
计算多人体脂率
*/

func getSexWeight(sex string) int {
	if sex == "男" {
		return 1
	}

	return 0
}

func getBmi(weight, tall float64) float64 {
	return weight / (tall * tall)
}

func getFatRate(bmi float64, age, sexWeight int) float64 {
	return (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
}

func suggest(fatRateMsg string) string {
	switch fatRateMsg {
	case "体脂率有误":
		return "请检验个人信息是否输入正确"
	case "偏瘦":
		return "偏瘦, 多吃蛋白"
	case "标准":
		return "标准, 注意保持哟"
	case "偏重":
		return "偏重, 需要进行适当运动哟"
	case "肥胖":
		return "肥胖, 你要多加运动哟"
	case "严重肥胖":
		return "严重肥胖, 需要看下医生了"
	default:
		return fatRateMsg
	}
}

func getFatRateResult(fatRate float64, age, sexWeight int) string {
	fatRateMsg := "没什么建议"
	if sexWeight == 1 {
		if age >= 18 && age <= 39 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.1 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.16 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.21 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.26 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age <= 59 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.11 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.17 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.22 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.27 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age >= 60 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.13 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.19 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.24 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.29 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		}
	} else {
		if age >= 18 && age <= 39 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.2 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.27 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.34 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.39 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age <= 59 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.21 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.28 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.35 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.4 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age >= 60 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.22 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.29 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.36 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.41 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		}
	}

	return fatRateMsg
}

func MainFatRateBody() {
	person_num := 3
	fmt.Println("支持计算", person_num, "个人的体脂率哟")

	names, ages, sexs, talls, weights := getPersonInfo(person_num)

	sexWeights := make([]int, person_num)
	bmis := make([]float64, person_num)
	fatRates := make([]float64, person_num)
	fatRateMsgs := make([]string, person_num)
	suggestInfos := make([]string, person_num)

	for i := 0; i < person_num; i++ {
		sexWeights[i] = getSexWeight(sexs[i])
		bmis[i] = getBmi(weights[i], talls[i])
		fatRates[i] = getFatRate(bmis[i], ages[i], sexWeights[i])

		fatRateMsgs[i] = getFatRateResult(fatRates[i], ages[i], sexWeights[i])
		suggestInfos[i] = suggest(fatRateMsgs[i])
	}

	fmt.Println("")
	for i := 0; i < person_num; i++ {
		fmt.Printf("%s, 你的体脂率是%.2f, %s\n", names[i], fatRates[i], suggestInfos[i])
	}

	fmt.Printf("平均体脂率为是%.2f", (fatRates[0]+fatRates[1]+fatRates[2])/3)
}

func getPersonInfo(person_num int) ([]string, []int, []string, []float64, []float64) {
	names := make([]string, person_num)
	ages := make([]int, person_num)
	sexs := make([]string, person_num)
	talls := make([]float64, person_num)
	weights := make([]float64, person_num)

	fmt.Print("请输入姓名：")
	fmt.Scan(&names[0], &names[1], &names[2])

	fmt.Print("请输入年龄(岁)：")
	fmt.Scan(&ages[0], &ages[1], &ages[2])

	fmt.Print("请输入性别：")
	fmt.Scan(&sexs[0], &sexs[1], &sexs[2])

	fmt.Print("请输入体重(公斤): ")
	fmt.Scan(&weights[0], &weights[1], &weights[2])

	fmt.Print("请输入身高(米): ")
	fmt.Scan(&talls[0], &talls[1], &talls[2])
	return names, ages, sexs, talls, weights
}

func WhetherContinue() bool {
	var goOn string
	fmt.Println("")
	fmt.Print("请问您是否还要继续计算体脂率(yes/no): ")
	fmt.Scanf("%s", &goOn)

	err_num := 1
	for goOn != "yes" && goOn != "no" {
		if err_num >= 3 {
			return false
		}
		fmt.Print("请问您是否还要继续计算体脂率(yes/no): ")
		fmt.Scanf("%s", &goOn)
		err_num++
	}

	if goOn == "no" {
		return false
	}

	return true
}

/**
目前只能计算固定的三人体脂，用户体验不是太好，入参没有做严格的校验
*/
