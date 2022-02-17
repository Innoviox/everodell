package everodell

type Action struct {
	gains Bundle

	nUses int

	special func() // todo
}
