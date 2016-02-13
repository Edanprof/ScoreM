package player

type PlayerInfo struct{
	name string // 이름
	nick string // 별명
	brith int // 생년월일
	height int // 키cm
	weight int // 몸무게 kg
	position map[string]int{} // 포지션, 포지션별 숙련도

	nation string // 소속국가
	team string // 소속팀
}

type PlayerStatus struct{
	status string
	first boolean
	goal int
	asis int
	foul int
	inj string
}
