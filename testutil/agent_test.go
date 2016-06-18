//
// Copyright (c) 2016 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package testutil_test

import (
	"errors"
	"testing"

	"github.com/01org/ciao/ssntp"
	. "github.com/01org/ciao/testutil"
)

func TestNewSsntpTestClientonnectionArgs(t *testing.T) {
	_, err := NewSsntpTestClientConnection("AGENT Client", ssntp.UNKNOWN, AgentUUID)
	if err == nil {
		t.Fatalf("NewSsntpTestClientConnection incorrectly accepted unknown role")
	}

	_, err = NewSsntpTestClientConnection("AGENT Client", ssntp.AGENT, "")
	if err == nil {
		t.Fatalf("NewSsntpTestClientConnection incorrectly accepted empty uuid")
	}
}

func TestAgentErrorChan(t *testing.T) {
	agentCh := agent.AddErrorChan(ssntp.StopFailure)

	var result Result
	result.Err = errors.New("foo")
	go agent.SendResultAndDelErrorChan(ssntp.StopFailure, result)

	r, err := agent.GetErrorChanResult(agentCh, ssntp.StopFailure)
	if err == nil {
		t.Fatal(err)
	}
	if r.Err != result.Err {
		t.Fatalf("channel returned wrong result: expected \"%s\", got \"%s\"\n", result.Err, r.Err)
	}
}

func TestAgentErrorChanTimeout(t *testing.T) {
	agentCh := agent.AddErrorChan(ssntp.StopFailure)

	_, err := agent.GetErrorChanResult(agentCh, ssntp.StopFailure)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAgentEventChan(t *testing.T) {
	agentCh := agent.AddEventChan(ssntp.TraceReport)

	var result Result
	result.Err = errors.New("foo")
	go agent.SendResultAndDelEventChan(ssntp.TraceReport, result)

	r, err := agent.GetEventChanResult(agentCh, ssntp.TraceReport)
	if err == nil {
		t.Fatal(err)
	}
	if r.Err != result.Err {
		t.Fatalf("channel returned wrong result: expected \"%s\", got \"%s\"\n", result.Err, r.Err)
	}
}

func TestAgentEventChanTimeout(t *testing.T) {
	agentCh := agent.AddEventChan(ssntp.TraceReport)

	_, err := agent.GetEventChanResult(agentCh, ssntp.TraceReport)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAgentCmdChan(t *testing.T) {
	agentCh := agent.AddCmdChan(ssntp.START)

	var result Result
	result.Err = errors.New("foo")
	go agent.SendResultAndDelCmdChan(ssntp.START, result)

	r, err := agent.GetCmdChanResult(agentCh, ssntp.START)
	if err == nil {
		t.Fatal(err)
	}
	if r.Err != result.Err {
		t.Fatalf("channel returned wrong result: expected \"%s\", got \"%s\"\n", result.Err, r.Err)
	}
}

func TestAgentCmdChanTimeout(t *testing.T) {
	agentCh := agent.AddCmdChan(ssntp.START)

	_, err := agent.GetCmdChanResult(agentCh, ssntp.START)
	if err == nil {
		t.Fatal(err)
	}
}
