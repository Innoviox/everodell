package everodell

type Action struct {
	gains Bundle
	nCards int

	special func() // todo
}