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
	"regexp"
	"strings"
)

type Counter struct {
	Requests  string
	Hits      string
	Errors    string
	KBytesIn  string
	KBytesOut string
}

type Status struct {
	Client      Counter
	Server      Counter
	ServerHttp  Counter
	ServerFtp   Counter
	ServerOther Counter
}

/* I'm going to do it in most simple and stupid ways.
 * Call me Kumar :)
 */
var ClientRequestRX, _ = regexp.Compile(`client_http.requests\s+=\s+(\d+)`)
var ClientHitsRX, _ = regexp.Compile(`client_http.hits\s+=\s+(\d+)`)
var ClientErrorsRX, _ = regexp.Compile(`client_http.errors\s+=\s+(\d+)`)
var ClientKBytesInRX, _ = regexp.Compile(`client_http.kbytes_in\s+=\s+(\d+)`)
var ClientKBytesOutRX, _ = regexp.Compile(`client_http.kbytes_out\s+=\s+(\d+)`)

var ServerRequestRX, _ = regexp.Compile(`server.all.requests\s+=\s+(\d+)`)
var ServerErrorsRX, _ = regexp.Compile(`server.all.errors\s+=\s+(\d+)`)
var ServerKBytesInRX, _ = regexp.Compile(`server.all.kbytes_in\s+=\s+(\d+)`)
var ServerKBytesOutRX, _ = regexp.Compile(`server.all.kbytes_out\s+=\s+(\d+)`)

var ServerHttpRequestRX, _ = regexp.Compile(`server.http.requests\s+=\s+(\d+)`)
var ServerHttpErrorsRX, _ = regexp.Compile(`server.http.errors\s+=\s+(\d+)`)
var ServerHttpKBytesInRX, _ = regexp.Compile(`server.http.kbytes_in\s+=\s+(\d+)`)
var ServerHttpKBytesOutRX, _ = regexp.Compile(`server.http.kbytes_out\s+=\s+(\d+)`)

var ServerFtpRequestRX, _ = regexp.Compile(`server.ftp.requests\s+=\s+(\d+)`)
var ServerFtpErrorsRX, _ = regexp.Compile(`server.ftp.errors\s+=\s+(\d+)`)
var ServerFtpKBytesInRX, _ = regexp.Compile(`server.ftp.kbytes_in\s+=\s+(\d+)`)
var ServerFtpKBytesOutRX, _ = regexp.Compile(`server.ftp.kbytes_out\s+=\s+(\d+)`)

var ServerOtherRequestRX, _ = regexp.Compile(`server.other.requests\s+=\s+(\d+)`)
var ServerOtherErrorsRX, _ = regexp.Compile(`server.other.errors\s+=\s+(\d+)`)
var ServerOtherKBytesInRX, _ = regexp.Compile(`server.other.kbytes_in\s+=\s+(\d+)`)
var ServerOtherKBytesOutRX, _ = regexp.Compile(`server.other.kbytes_out\s+=\s+(\d+)`)

func parse(page string) Status {
	var result = Status{Client: Counter{}, Server: Counter{}, ServerHttp: Counter{}, ServerFtp: Counter{}, ServerOther: Counter{}}

	var statusData = strings.Split(page, "\n")
	for _, line := range statusData {
		if ClientHitsRX.MatchString(line) {
			result.Client.Hits = ClientHitsRX.FindStringSubmatch(line)[1]
		}
		if ClientErrorsRX.MatchString(line) {
			result.Client.Errors = ClientErrorsRX.FindStringSubmatch(line)[1]
		}
		if ClientKBytesInRX.MatchString(line) {
			result.Client.KBytesIn = ClientKBytesInRX.FindStringSubmatch(line)[1]
		}
		if ClientKBytesOutRX.MatchString(line) {
			result.Client.KBytesOut = ClientKBytesOutRX.FindStringSubmatch(line)[1]
		}
		if ClientRequestRX.MatchString(line) {
			result.Client.Requests = ClientRequestRX.FindStringSubmatch(line)[1]
		}

		if ServerErrorsRX.MatchString(line) {
			result.Server.Errors = ServerErrorsRX.FindStringSubmatch(line)[1]
		}
		if ServerKBytesInRX.MatchString(line) {
			result.Server.KBytesIn = ServerKBytesInRX.FindStringSubmatch(line)[1]
		}
		if ServerKBytesOutRX.MatchString(line) {
			result.Server.KBytesOut = ServerKBytesOutRX.FindStringSubmatch(line)[1]
		}
		if ServerRequestRX.MatchString(line) {
			result.Server.Requests = ServerRequestRX.FindStringSubmatch(line)[1]
		}

		if ServerHttpErrorsRX.MatchString(line) {
			result.ServerHttp.Errors = ServerHttpErrorsRX.FindStringSubmatch(line)[1]
		}
		if ServerHttpKBytesInRX.MatchString(line) {
			result.ServerHttp.KBytesIn = ServerHttpKBytesInRX.FindStringSubmatch(line)[1]
		}
		if ServerHttpKBytesOutRX.MatchString(line) {
			result.ServerHttp.KBytesOut = ServerHttpKBytesOutRX.FindStringSubmatch(line)[1]
		}
		if ServerHttpRequestRX.MatchString(line) {
			result.ServerHttp.Requests = ServerHttpRequestRX.FindStringSubmatch(line)[1]
		}

		if ServerFtpErrorsRX.MatchString(line) {
			result.ServerFtp.Errors = ServerFtpErrorsRX.FindStringSubmatch(line)[1]
		}
		if ServerFtpKBytesInRX.MatchString(line) {
			result.ServerFtp.KBytesIn = ServerFtpKBytesInRX.FindStringSubmatch(line)[1]
		}
		if ServerFtpKBytesOutRX.MatchString(line) {
			result.ServerFtp.KBytesOut = ServerFtpKBytesOutRX.FindStringSubmatch(line)[1]
		}
		if ServerFtpRequestRX.MatchString(line) {
			result.ServerFtp.Requests = ServerFtpRequestRX.FindStringSubmatch(line)[1]
		}

		if ServerOtherErrorsRX.MatchString(line) {
			result.ServerOther.Errors = ServerOtherErrorsRX.FindStringSubmatch(line)[1]
		}
		if ServerOtherKBytesInRX.MatchString(line) {
			result.ServerOther.KBytesIn = ServerOtherKBytesInRX.FindStringSubmatch(line)[1]
		}
		if ServerOtherKBytesOutRX.MatchString(line) {
			result.ServerOther.KBytesOut = ServerOtherKBytesOutRX.FindStringSubmatch(line)[1]
		}
		if ServerOtherRequestRX.MatchString(line) {
			result.ServerOther.Requests = ServerOtherRequestRX.FindStringSubmatch(line)[1]
		}
	}

	return result
}
