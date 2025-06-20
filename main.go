package main

import (
	"flag"
	"github.com/isaacgr/nmos-node-generator/capabilities/is04"
	"github.com/isaacgr/nmos-node-generator/constants"
)

func main() {
	flag.Parse()
	nc := is04.NewResourceCore(
		"evNode",
		"NMOS Test Node",
		false,
	)
	dc := is04.NewResourceCore(
		"evDevice",
		"NMOS Test Device",
		false,
	)
	sc := is04.NewResourceCore(
		"evSource",
		"NMOS Test Source",
		false,
	)

	rn := is04.NewNode(
		"127.0.0.1",
		63001,
		2,
		nil,
		nc,
	)

	rd := is04.NewDevice(
		rn,
		dc,
	)

	is04.NewSource(
		rd,
		constants.VideoFormat,
		sc,
	)

	s := is04.IS04Server{
		Node: rn,
	}
	s.Start()

}
