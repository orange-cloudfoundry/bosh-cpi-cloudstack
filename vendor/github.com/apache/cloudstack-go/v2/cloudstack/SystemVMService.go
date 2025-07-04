//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SystemVMServiceIface interface {
	ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error)
	NewChangeServiceForSystemVmParams(id string, serviceofferingid string) *ChangeServiceForSystemVmParams
	DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error)
	NewDestroySystemVmParams(id string) *DestroySystemVmParams
	ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error)
	NewListSystemVmsParams() *ListSystemVmsParams
	GetSystemVmID(name string, opts ...OptionFunc) (string, int, error)
	GetSystemVmByName(name string, opts ...OptionFunc) (*SystemVm, int, error)
	GetSystemVmByID(id string, opts ...OptionFunc) (*SystemVm, int, error)
	ListSystemVmsUsageHistory(p *ListSystemVmsUsageHistoryParams) (*ListSystemVmsUsageHistoryResponse, error)
	NewListSystemVmsUsageHistoryParams() *ListSystemVmsUsageHistoryParams
	GetSystemVmsUsageHistoryID(name string, opts ...OptionFunc) (string, int, error)
	GetSystemVmsUsageHistoryByName(name string, opts ...OptionFunc) (*SystemVmsUsageHistory, int, error)
	GetSystemVmsUsageHistoryByID(id string, opts ...OptionFunc) (*SystemVmsUsageHistory, int, error)
	MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error)
	NewMigrateSystemVmParams(virtualmachineid string) *MigrateSystemVmParams
	RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error)
	NewRebootSystemVmParams(id string) *RebootSystemVmParams
	ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error)
	NewScaleSystemVmParams(id string, serviceofferingid string) *ScaleSystemVmParams
	StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error)
	NewStartSystemVmParams(id string) *StartSystemVmParams
	StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error)
	NewStopSystemVmParams(id string) *StopSystemVmParams
	PatchSystemVm(p *PatchSystemVmParams) (*PatchSystemVmResponse, error)
	NewPatchSystemVmParams() *PatchSystemVmParams
}

type ChangeServiceForSystemVmParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForSystemVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ChangeServiceForSystemVmParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ChangeServiceForSystemVmParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *ChangeServiceForSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeServiceForSystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ChangeServiceForSystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeServiceForSystemVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ChangeServiceForSystemVmParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ChangeServiceForSystemVmParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeServiceForSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewChangeServiceForSystemVmParams(id string, serviceofferingid string) *ChangeServiceForSystemVmParams {
	p := &ChangeServiceForSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Changes the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("changeServiceForSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type DestroySystemVmParams struct {
	p map[string]interface{}
}

func (p *DestroySystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroySystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DestroySystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DestroySystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DestroySystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewDestroySystemVmParams(id string) *DestroySystemVmParams {
	p := &DestroySystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a system virtual machine.
func (s *SystemVMService) DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("destroySystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroySystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DestroySystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type ListSystemVmsParams struct {
	p map[string]interface{}
}

func (p *ListSystemVmsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSystemVmsParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *ListSystemVmsParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *ListSystemVmsParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListSystemVmsParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListSystemVmsParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSystemVmsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSystemVmsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSystemVmsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSystemVmsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListSystemVmsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListSystemVmsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSystemVmsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSystemVmsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSystemVmsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSystemVmsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSystemVmsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSystemVmsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListSystemVmsParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListSystemVmsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListSystemVmsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListSystemVmsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListSystemVmsParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListSystemVmsParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
}

func (p *ListSystemVmsParams) ResetSystemvmtype() {
	if p.p != nil && p.p["systemvmtype"] != nil {
		delete(p.p, "systemvmtype")
	}
}

func (p *ListSystemVmsParams) GetSystemvmtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["systemvmtype"].(string)
	return value, ok
}

func (p *ListSystemVmsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListSystemVmsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListSystemVmsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSystemVmsParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewListSystemVmsParams() *ListSystemVmsParams {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSystemVms(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SystemVms[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SystemVms {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmByName(name string, opts ...OptionFunc) (*SystemVm, int, error) {
	id, count, err := s.GetSystemVmID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSystemVmByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmByID(id string, opts ...OptionFunc) (*SystemVm, int, error) {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSystemVms(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.SystemVms[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SystemVm UUID: %s!", id)
}

// List system virtual machines.
func (s *SystemVMService) ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error) {
	resp, err := s.cs.newRequest("listSystemVms", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSystemVmsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSystemVmsResponse struct {
	Count     int         `json:"count"`
	SystemVms []*SystemVm `json:"systemvm"`
}

type SystemVm struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type ListSystemVmsUsageHistoryParams struct {
	p map[string]interface{}
}

func (p *ListSystemVmsUsageHistoryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *ListSystemVmsUsageHistoryParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSystemVmsUsageHistoryParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ListSystemVmsUsageHistoryParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *ListSystemVmsUsageHistoryParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new ListSystemVmsUsageHistoryParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewListSystemVmsUsageHistoryParams() *ListSystemVmsUsageHistoryParams {
	p := &ListSystemVmsUsageHistoryParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmsUsageHistoryID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListSystemVmsUsageHistoryParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSystemVmsUsageHistory(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SystemVmsUsageHistory[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SystemVmsUsageHistory {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmsUsageHistoryByName(name string, opts ...OptionFunc) (*SystemVmsUsageHistory, int, error) {
	id, count, err := s.GetSystemVmsUsageHistoryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSystemVmsUsageHistoryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmsUsageHistoryByID(id string, opts ...OptionFunc) (*SystemVmsUsageHistory, int, error) {
	p := &ListSystemVmsUsageHistoryParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSystemVmsUsageHistory(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.SystemVmsUsageHistory[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SystemVmsUsageHistory UUID: %s!", id)
}

// Lists System VM stats
func (s *SystemVMService) ListSystemVmsUsageHistory(p *ListSystemVmsUsageHistoryParams) (*ListSystemVmsUsageHistoryResponse, error) {
	resp, err := s.cs.newRequest("listSystemVmsUsageHistory", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSystemVmsUsageHistoryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSystemVmsUsageHistoryResponse struct {
	Count                 int                      `json:"count"`
	SystemVmsUsageHistory []*SystemVmsUsageHistory `json:"systemvmsusagehistory"`
}

type SystemVmsUsageHistory struct {
	Displayname string   `json:"displayname"`
	Id          string   `json:"id"`
	JobID       string   `json:"jobid"`
	Jobstatus   int      `json:"jobstatus"`
	Name        string   `json:"name"`
	Stats       []string `json:"stats"`
}

type MigrateSystemVmParams struct {
	p map[string]interface{}
}

func (p *MigrateSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoselect"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("autoselect", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateSystemVmParams) SetAutoselect(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoselect"] = v
}

func (p *MigrateSystemVmParams) ResetAutoselect() {
	if p.p != nil && p.p["autoselect"] != nil {
		delete(p.p, "autoselect")
	}
}

func (p *MigrateSystemVmParams) GetAutoselect() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoselect"].(bool)
	return value, ok
}

func (p *MigrateSystemVmParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *MigrateSystemVmParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *MigrateSystemVmParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *MigrateSystemVmParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *MigrateSystemVmParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *MigrateSystemVmParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *MigrateSystemVmParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *MigrateSystemVmParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *MigrateSystemVmParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewMigrateSystemVmParams(virtualmachineid string) *MigrateSystemVmParams {
	p := &MigrateSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a system virtual machine to the host specified.
func (s *SystemVMService) MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("migrateSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type MigrateSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type RebootSystemVmParams struct {
	p map[string]interface{}
}

func (p *RebootSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootSystemVmParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *RebootSystemVmParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *RebootSystemVmParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *RebootSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RebootSystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RebootSystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RebootSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewRebootSystemVmParams(id string) *RebootSystemVmParams {
	p := &RebootSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reboots a system VM.
func (s *SystemVMService) RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("rebootSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RebootSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type ScaleSystemVmParams struct {
	p map[string]interface{}
}

func (p *ScaleSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ScaleSystemVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ScaleSystemVmParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ScaleSystemVmParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *ScaleSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ScaleSystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ScaleSystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ScaleSystemVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ScaleSystemVmParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ScaleSystemVmParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ScaleSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewScaleSystemVmParams(id string, serviceofferingid string) *ScaleSystemVmParams {
	p := &ScaleSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Scale the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("scaleSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ScaleSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type StartSystemVmParams struct {
	p map[string]interface{}
}

func (p *StartSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StartSystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StartSystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStartSystemVmParams(id string) *StartSystemVmParams {
	p := &StartSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a system virtual machine.
func (s *SystemVMService) StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("startSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type StartSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type StopSystemVmParams struct {
	p map[string]interface{}
}

func (p *StopSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopSystemVmParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *StopSystemVmParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *StopSystemVmParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *StopSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StopSystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StopSystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStopSystemVmParams(id string) *StopSystemVmParams {
	p := &StopSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a system VM.
func (s *SystemVMService) StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("stopSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type StopSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Arch                  string   `json:"arch"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostcontrolstate      string   `json:"hostcontrolstate"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	Serviceofferingid     string   `json:"serviceofferingid"`
	Serviceofferingname   string   `json:"serviceofferingname"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type PatchSystemVmParams struct {
	p map[string]interface{}
}

func (p *PatchSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *PatchSystemVmParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *PatchSystemVmParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *PatchSystemVmParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *PatchSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *PatchSystemVmParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *PatchSystemVmParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new PatchSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewPatchSystemVmParams() *PatchSystemVmParams {
	p := &PatchSystemVmParams{}
	p.p = make(map[string]interface{})
	return p
}

// Attempts to live patch systemVMs - CPVM, SSVM
func (s *SystemVMService) PatchSystemVm(p *PatchSystemVmParams) (*PatchSystemVmResponse, error) {
	resp, err := s.cs.newPostRequest("patchSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PatchSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type PatchSystemVmResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
