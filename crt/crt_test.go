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

package crt

import (
	"testing"
)

func TestCRTInit(t *testing.T) {
	configStr := `{
                    "docker" : {
                        "socket" : "unix:///var/run/docker.sock"
                    },
                    "crt" : {
                        "type": "docker"
                    }
                  }`
	crt := CRT{}
	err := crt.Init(configStr)
	if err != nil {
		t.Fatalf("crt init failed: Error: %s", err)
	}
	defer func() { crt.Deinit() }()
}

func TestCRTInitInvalidConfigMissingCRT(t *testing.T) {
	configStr := `{
                    "docker" : {
                        "socket" : "unix:///var/run/docker.sock"
                    }
                  }`
	crt := CRT{}
	err := crt.Init(configStr)
	if err == nil {
		t.Fatalf("crt init succeeded!")
	}
}

func TestCRTInitInvalidConfigMissingCRTIf(t *testing.T) {
	configStr := `{
                    "crt" : {
                        "type": "docker"
                    }
                  }`
	crt := CRT{}
	err := crt.Init(configStr)
	if err == nil {
		t.Fatalf("crt init succeeded!")
	}
}

func TestCRTInitInvalidConfigInvalidCRT(t *testing.T) {
	configStr := `{
                    "crt" : {
                        "type": "rocket"
                    }
                  }`
	crt := CRT{}
	err := crt.Init(configStr)
	if err == nil {
		t.Fatalf("crt init succeeded!")
	}
}
