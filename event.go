package everodell

type Event struct {
	points int
	reqs []string
	achievable func() //todo
	onAchieve func() //todo
}