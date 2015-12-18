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
	"time"
)

func wait(configuration Configuration) {
	time.Sleep(time.Duration(configuration.Period) * time.Second)
}

func main() {
	InitLog()
	log.Info("Starting graphite-squid...")

	var configuration = config()
	for {
		var page, err = getPage(configuration.Host, configuration.Port)
		if err != nil {
			continue
		}
		var status = parse(page)

		sendMetrics(status, configuration)

		wait(configuration)
	}
}
