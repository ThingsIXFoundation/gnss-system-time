// Copyright 2023 Stichting ThingsIX Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package gnsssystemtime

import (
	"math"
	"time"
)

var GST_START = time.Date(1999, 8, 22, 0, 0, 0, 0, time.UTC)

// GalileoToTime returns the current time in UTC, given a Galileo week-number (wno), time-of-week (tow) and the current leap-seconds
func GalileoToTime(wno uint16, tow uint32, leapSeconds int) time.Time {
	return GST_START.Add(time.Duration(wno) * 7 * 24 * time.Hour).Add(time.Duration(tow) * time.Second).Add(-1 * time.Duration(leapSeconds) * time.Second)
}

// TimeToGalileo returns the Galileo week-number and time-of-week given the current time and current leap-seconds
func TimeToGalileo(t time.Time, leapSeconds int) (uint16, uint32) {
	afterStart := t.UTC().Add(time.Duration(leapSeconds) * time.Second).Sub(GST_START)
	wno := uint16(math.Floor(afterStart.Hours() / (7 * 24)))
	tow := uint32((afterStart - (time.Duration(wno) * 7 * 24 * time.Hour)) / time.Second)
	return wno, tow
}

// GalileoTowToTime determines the first matching time based on a Galileo time-of-week (tow) before a given timestamp and
// given the current leap seconds. It searches non-inclusive, the returned time will always be before before but never
// before itself.
func GalileoTowToTime(tow uint32, before time.Time, leapSeconds int) time.Time {
	beforeWno, beforeTow := TimeToGalileo(before, leapSeconds)
	// We search for the matching galileo time before time
	// if the beforeTow is bigger than the tow, the wno is the same
	if beforeTow > tow {
		return GalileoToTime(beforeWno, tow, leapSeconds)
		// if the beforeTwo is smaller than the two, it must be from the previous week
	} else {
		return GalileoToTime(beforeWno-1, tow, leapSeconds)
	}
}
