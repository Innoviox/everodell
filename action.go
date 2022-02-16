package everodell

type Action struct {
	gains Bundle
	nCards int

	nUses int

	special func() // todo
}