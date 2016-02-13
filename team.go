package team

type teamInfo struct{
	name string //이름
	nick string //별칭
	stadium string //경기장
	nation string //국가
	city string // 연고지
	founded int //창단일
	league string // 소속리그
	manager string // 감독
}

type teamStatus struct{
	leagueRank map[string]int //참가리그별 순위

}
