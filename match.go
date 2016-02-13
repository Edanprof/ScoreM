package match

import "team"

type MatchInfo struct{
	date int // 경기 시작날짜
	time int // 경기 시작시간
	home team.TeamInfo.name //home 팀이름
	away team.TeamInfo.name // away 팀이름
}

type MatchStatus struct{
	time int // 경기시작후 경과시간
	state string // 경기시작전, 경기중, 하프타임, 경기종료
	weather string // 날씨
	home_score int // 홈 스코어
	away_score int // 어웨이 스코어
	scoreman map[int]string //시간(int)에 해당득점자(string)
}
