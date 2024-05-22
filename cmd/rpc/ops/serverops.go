package ops

import (
	"runtime"
	"runtime/pprof"

	"github.com/johnllao/college/cmd/rpc/args"
)

type ServerOps struct {
	count int
}

func (o *ServerOps) Ping(a *int, reply *args.PingReply) error {
	o.count++
	reply.Message = "PONG"
	reply.N = o.count
	return nil
}

func (o *ServerOps) Status(a *int, reply *args.StatusReply) error {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	reply.ProfileNames = make([]string, len(pprof.Profiles()))
	var profiles = pprof.Profiles()
	for i := 0; i < len(profiles); i++ {
		reply.ProfileNames[i] = profiles[i].Name()
	}

	return nil
}
