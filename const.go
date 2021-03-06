package main

var (
	member = []string{
		"浩铭",
		"钊杰", // 1
		"方舟", //2
		"贵雄", //3
		"文星", //4
		"焕斑", //5
	}
	phone = []string{
		"13076886494",
		"15901092701",
		"15671568977",
		"13250507260",
		"18607588619",
		"15914548037",
	}
	//钉钉群昵称
	userIds = []string{
		"李浩铭 Haoming Li",  // 0
		"杨钊杰 Johnny Yang", // 1
		"方舟 Zhou Fang",    // 2
		"阮贵雄 Chase Ruan",  // 3
		"史文星 Wenxing Shi", // 4
		"蔡焕斑 Huanban Cai", // 5
	}
	// 每个[]代表当周值班同学id
	timeMap = [][]int{
		{1, 3},
		{0, 4},
		{3, 5},
		{2, 1},
		{5, 0},
		{4, 2},
	}
	cron_ = "0 0 10 ? * MON"
	// 钉钉webhook
	// 测试群链接：
	//url = "https://oapi.dingtalk.com/robot/send?access_token=0ac0f9469176242b763d3b1ab92c54fe10d93588ae997fe7aa433c38065a9cb1"
	url = "https://oapi.dingtalk.com/robot/send?access_token=6c0465483418b335583760f6b295627d11bbc6d64ce64b25c6e9a0c287b665c6"
)
