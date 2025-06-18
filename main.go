package main

import (
	"flag"
	"github.com/isaacgr/nmos-node-generator/capabilities/is04"
)

func main() {
	flag.Parse()
	rc := is04.NewResourceCore(
		"evNode",
		"NMOS Test Node",
		false,
	)

	rn := is04.NewNode(
		"127.0.0.1",
		63001,
		2,
		nil,
		rc,
	)

	is04.NewDevice(
		rn,
		false,
	)

	s := is04.IS04Server{
		Node: rn,
	}
	s.Start()

}
