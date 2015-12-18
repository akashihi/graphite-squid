/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"github.com/marpaia/graphite-golang"
)

func sendMetrics(status Status, config Configuration) {
	var Graphite, err = graphite.NewGraphite(config.MetricsHost, config.MetricsPort)
	if err != nil {
		log.Error("Can't connect to graphite collector: %v", err)
		return
	}

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.client.requests"), status.Client.Requests)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.client.hits"), status.Client.Hits)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.client.errors"), status.Client.Errors)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.client.kbytes.in"), status.Client.KBytesIn)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.client.kbytes.out"), status.Client.KBytesOut)

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.requests"), status.Server.Requests)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.errors"), status.Server.Errors)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.kbytes.in"), status.Server.KBytesIn)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.kbytes.out"), status.Server.KBytesOut)

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.http.requests"), status.ServerHttp.Requests)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.http.errors"), status.ServerHttp.Errors)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.http.kbytes.in"), status.ServerHttp.KBytesIn)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.http.kbytes.out"), status.ServerHttp.KBytesOut)

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.ftp.requests"), status.ServerFtp.Requests)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.ftp.errors"), status.ServerFtp.Errors)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.ftp.kbytes.in"), status.ServerFtp.KBytesIn)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.ftp.kbytes.out"), status.ServerFtp.KBytesOut)

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.other.requests"), status.ServerOther.Requests)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.other.errors"), status.ServerOther.Errors)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.other.kbytes.in"), status.ServerOther.KBytesIn)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".squid.server.other.kbytes.out"), status.ServerOther.KBytesOut)
}
