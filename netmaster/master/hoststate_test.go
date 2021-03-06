/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package master

import (
	"testing"

	"github.com/contiv/netplugin/core"
)

const (
	testHostID = "testHost"
	hostCfgKey = hostConfigPathPrefix + testHostID
)

var hostState = &testHostStateDriver{}

type testHostStateDriver struct {
}

func (d *testHostStateDriver) Init(config *core.Config) error {
	return core.Errorf("Shouldn't be called!")
}

func (d *testHostStateDriver) Deinit() {
}

func (d *testHostStateDriver) Write(key string, value []byte) error {
	return core.Errorf("Shouldn't be called!")
}

func (d *testHostStateDriver) Read(key string) ([]byte, error) {
	return []byte{}, core.Errorf("Shouldn't be called!")
}

func (d *testHostStateDriver) ReadAll(baseKey string) ([][]byte, error) {
	return [][]byte{}, core.Errorf("Shouldn't be called!")
}

func (d *testHostStateDriver) WatchAll(baseKey string, rsps chan [2][]byte) error {
	return core.Errorf("not supported")
}

func (d *testHostStateDriver) validateKey(key string) error {
	if key != hostCfgKey {
		return core.Errorf("Unexpected key. recvd: %s expected: %s",
			key, hostCfgKey)
	}

	return nil
}

func (d *testHostStateDriver) ClearState(key string) error {
	return d.validateKey(key)
}

func (d *testHostStateDriver) ReadState(key string, value core.State,
	unmarshal func([]byte, interface{}) error) error {
	return d.validateKey(key)
}

func (d *testHostStateDriver) ReadAllState(key string, value core.State,
	unmarshal func([]byte, interface{}) error) ([]core.State, error) {
	return nil, core.Errorf("Shouldn't be called!")
}

func (d *testHostStateDriver) WatchAllState(baseKey string, sType core.State,
	unmarshal func([]byte, interface{}) error, rsps chan core.WatchState) error {
	return core.Errorf("not supported")
}

func (d *testHostStateDriver) WriteState(key string, value core.State,
	marshal func(interface{}) ([]byte, error)) error {
	return d.validateKey(key)
}

func TestHostConfigRead(t *testing.T) {
	hostCfg := &HostConfig{}
	hostCfg.StateDriver = hostState

	err := hostCfg.Read(testHostID)
	if err != nil {
		t.Fatalf("read config state failed. Error: %s", err)
	}
}

func TestHostConfigWrite(t *testing.T) {
	hostCfg := &HostConfig{}
	hostCfg.StateDriver = hostState
	hostCfg.Name = testHostID

	err := hostCfg.Write()
	if err != nil {
		t.Fatalf("write config state failed. Error: %s", err)
	}
}

func TestHostConfigClear(t *testing.T) {
	hostCfg := &HostConfig{}
	hostCfg.StateDriver = hostState
	hostCfg.Name = testHostID

	err := hostCfg.Clear()
	if err != nil {
		t.Fatalf("clear config state failed. Error: %s", err)
	}
}
