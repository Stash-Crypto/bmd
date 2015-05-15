// Copyright (c) 2015 Monetas
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.package main

package main

import (
	"github.com/monetas/bmd/bmpeer"
	"github.com/monetas/bmd/database"
)

// TstNewServer allows for the creation of a sever with extra parameters for testing
// purposes, such as a nonstandard set of default peers.
func (s *server) TstStart(startPeers []*DefaultPeer) {
	s.start(startPeers)
}

func TstNewServer(listenAddrs []string, db database.Db, listen func(string, string) (bmpeer.Listener, error)) (*server, error) {
	return newServer(listenAddrs, db, listen)
}
