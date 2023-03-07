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
	"reflect"
	"testing"
	"time"
)

func TestGalileoToTime(t *testing.T) {
	type args struct {
		wno         uint16
		tow         uint32
		leapSeconds int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"Tue, 07 Mar 2023 13:43:30 UTC",
			args{
				1228,
				222228,
				18,
			},
			time.Date(2023, 03, 7, 13, 43, 30, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GalileoToTime(tt.args.wno, tt.args.tow, tt.args.leapSeconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GalileoToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToGalileo(t *testing.T) {
	type args struct {
		t           time.Time
		leapSeconds int
	}
	tests := []struct {
		name string
		args args
		wno  uint16
		tow  uint32
	}{
		{

			"Tue, 07 Mar 2023 13:43:30 UTC",
			args{
				time.Date(2023, 03, 7, 13, 43, 30, 0, time.UTC),
				18,
			},
			1228,
			222228,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWno, gotTow := TimeToGalileo(tt.args.t, tt.args.leapSeconds)
			if gotWno != tt.wno {
				t.Errorf("TimeToGalileo() wno got = %v, want %v", gotWno, tt.wno)
			}
			if gotTow != tt.tow {
				t.Errorf("TimeToGalileo() tow got = %v, want %v", gotTow, tt.tow)
			}
		})
	}
}

func TestGalileoTowToTime(t *testing.T) {
	type args struct {
		tow         uint32
		before      time.Time
		leapSeconds int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"Tue, 07 Mar 2023 13:43:30 UTC",
			args{
				222228,
				time.Date(2023, 03, 7, 13, 43, 31, 0, time.UTC),
				18,
			},
			time.Date(2023, 03, 7, 13, 43, 30, 0, time.UTC),
		},
		{
			"Tue, 07 Mar 2023 13:43:30 UTC",
			args{
				222229,
				time.Date(2023, 03, 7, 13, 43, 30, 0, time.UTC),
				18,
			},
			time.Date(2023, 02, 28, 13, 43, 31, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GalileoTowToTime(tt.args.tow, tt.args.before, tt.args.leapSeconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GalileoTowToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
