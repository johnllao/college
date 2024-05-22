package args

type PingReply struct {
	N       int
	Message string
}

type StatusReply struct {
	ProfileNames []string
}
