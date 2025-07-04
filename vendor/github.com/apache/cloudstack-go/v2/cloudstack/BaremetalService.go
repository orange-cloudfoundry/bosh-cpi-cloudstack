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
	"net/url"
	"strconv"
)

type BaremetalServiceIface interface {
	AddBaremetalDhcp(p *AddBaremetalDhcpParams) (*AddBaremetalDhcpResponse, error)
	NewAddBaremetalDhcpParams(dhcpservertype string, password string, physicalnetworkid string, url string, username string) *AddBaremetalDhcpParams
	AddBaremetalPxeKickStartServer(p *AddBaremetalPxeKickStartServerParams) (*AddBaremetalPxeKickStartServerResponse, error)
	NewAddBaremetalPxeKickStartServerParams(password string, physicalnetworkid string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxeKickStartServerParams
	AddBaremetalPxePingServer(p *AddBaremetalPxePingServerParams) (*AddBaremetalPxePingServerResponse, error)
	NewAddBaremetalPxePingServerParams(password string, physicalnetworkid string, pingdir string, pingstorageserverip string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxePingServerParams
	AddBaremetalRct(p *AddBaremetalRctParams) (*AddBaremetalRctResponse, error)
	NewAddBaremetalRctParams(baremetalrcturl string) *AddBaremetalRctParams
	DeleteBaremetalRct(p *DeleteBaremetalRctParams) (*DeleteBaremetalRctResponse, error)
	NewDeleteBaremetalRctParams(id string) *DeleteBaremetalRctParams
	ListBaremetalDhcp(p *ListBaremetalDhcpParams) (*ListBaremetalDhcpResponse, error)
	NewListBaremetalDhcpParams(physicalnetworkid string) *ListBaremetalDhcpParams
	ListBaremetalPxeServers(p *ListBaremetalPxeServersParams) (*ListBaremetalPxeServersResponse, error)
	NewListBaremetalPxeServersParams(physicalnetworkid string) *ListBaremetalPxeServersParams
	ListBaremetalRct(p *ListBaremetalRctParams) (*ListBaremetalRctResponse, error)
	NewListBaremetalRctParams() *ListBaremetalRctParams
	NotifyBaremetalProvisionDone(p *NotifyBaremetalProvisionDoneParams) (*NotifyBaremetalProvisionDoneResponse, error)
	NewNotifyBaremetalProvisionDoneParams(mac string) *NotifyBaremetalProvisionDoneParams
}

type AddBaremetalDhcpParams struct {
	p map[string]interface{}
}

func (p *AddBaremetalDhcpParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["dhcpservertype"]; found {
		u.Set("dhcpservertype", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddBaremetalDhcpParams) SetDhcpservertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpservertype"] = v
}

func (p *AddBaremetalDhcpParams) ResetDhcpservertype() {
	if p.p != nil && p.p["dhcpservertype"] != nil {
		delete(p.p, "dhcpservertype")
	}
}

func (p *AddBaremetalDhcpParams) GetDhcpservertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpservertype"].(string)
	return value, ok
}

func (p *AddBaremetalDhcpParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddBaremetalDhcpParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddBaremetalDhcpParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddBaremetalDhcpParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddBaremetalDhcpParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddBaremetalDhcpParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddBaremetalDhcpParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddBaremetalDhcpParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddBaremetalDhcpParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddBaremetalDhcpParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddBaremetalDhcpParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddBaremetalDhcpParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalDhcpParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalDhcpParams(dhcpservertype string, password string, physicalnetworkid string, url string, username string) *AddBaremetalDhcpParams {
	p := &AddBaremetalDhcpParams{}
	p.p = make(map[string]interface{})
	p.p["dhcpservertype"] = dhcpservertype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// adds a baremetal dhcp server
func (s *BaremetalService) AddBaremetalDhcp(p *AddBaremetalDhcpParams) (*AddBaremetalDhcpResponse, error) {
	resp, err := s.cs.newPostRequest("addBaremetalDhcp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalDhcpResponse
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

type AddBaremetalDhcpResponse struct {
	Dhcpservertype    string `json:"dhcpservertype"`
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Url               string `json:"url"`
}

type AddBaremetalPxeKickStartServerParams struct {
	p map[string]interface{}
}

func (p *AddBaremetalPxeKickStartServerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["pxeservertype"]; found {
		u.Set("pxeservertype", v.(string))
	}
	if v, found := p.p["tftpdir"]; found {
		u.Set("tftpdir", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddBaremetalPxeKickStartServerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddBaremetalPxeKickStartServerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddBaremetalPxeKickStartServerParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *AddBaremetalPxeKickStartServerParams) SetPxeservertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pxeservertype"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetPxeservertype() {
	if p.p != nil && p.p["pxeservertype"] != nil {
		delete(p.p, "pxeservertype")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetPxeservertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pxeservertype"].(string)
	return value, ok
}

func (p *AddBaremetalPxeKickStartServerParams) SetTftpdir(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tftpdir"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetTftpdir() {
	if p.p != nil && p.p["tftpdir"] != nil {
		delete(p.p, "tftpdir")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetTftpdir() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tftpdir"].(string)
	return value, ok
}

func (p *AddBaremetalPxeKickStartServerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddBaremetalPxeKickStartServerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddBaremetalPxeKickStartServerParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddBaremetalPxeKickStartServerParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalPxeKickStartServerParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalPxeKickStartServerParams(password string, physicalnetworkid string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxeKickStartServerParams {
	p := &AddBaremetalPxeKickStartServerParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["pxeservertype"] = pxeservertype
	p.p["tftpdir"] = tftpdir
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// add a baremetal pxe server
func (s *BaremetalService) AddBaremetalPxeKickStartServer(p *AddBaremetalPxeKickStartServerParams) (*AddBaremetalPxeKickStartServerResponse, error) {
	resp, err := s.cs.newPostRequest("addBaremetalPxeKickStartServer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalPxeKickStartServerResponse
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

type AddBaremetalPxeKickStartServerResponse struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Tftpdir           string `json:"tftpdir"`
	Url               string `json:"url"`
}

type AddBaremetalPxePingServerParams struct {
	p map[string]interface{}
}

func (p *AddBaremetalPxePingServerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["pingcifspassword"]; found {
		u.Set("pingcifspassword", v.(string))
	}
	if v, found := p.p["pingcifsusername"]; found {
		u.Set("pingcifsusername", v.(string))
	}
	if v, found := p.p["pingdir"]; found {
		u.Set("pingdir", v.(string))
	}
	if v, found := p.p["pingstorageserverip"]; found {
		u.Set("pingstorageserverip", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["pxeservertype"]; found {
		u.Set("pxeservertype", v.(string))
	}
	if v, found := p.p["tftpdir"]; found {
		u.Set("tftpdir", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddBaremetalPxePingServerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPingcifspassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingcifspassword"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPingcifspassword() {
	if p.p != nil && p.p["pingcifspassword"] != nil {
		delete(p.p, "pingcifspassword")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPingcifspassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pingcifspassword"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPingcifsusername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingcifsusername"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPingcifsusername() {
	if p.p != nil && p.p["pingcifsusername"] != nil {
		delete(p.p, "pingcifsusername")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPingcifsusername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pingcifsusername"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPingdir(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingdir"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPingdir() {
	if p.p != nil && p.p["pingdir"] != nil {
		delete(p.p, "pingdir")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPingdir() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pingdir"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPingstorageserverip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingstorageserverip"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPingstorageserverip() {
	if p.p != nil && p.p["pingstorageserverip"] != nil {
		delete(p.p, "pingstorageserverip")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPingstorageserverip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pingstorageserverip"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetPxeservertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pxeservertype"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetPxeservertype() {
	if p.p != nil && p.p["pxeservertype"] != nil {
		delete(p.p, "pxeservertype")
	}
}

func (p *AddBaremetalPxePingServerParams) GetPxeservertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pxeservertype"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetTftpdir(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tftpdir"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetTftpdir() {
	if p.p != nil && p.p["tftpdir"] != nil {
		delete(p.p, "tftpdir")
	}
}

func (p *AddBaremetalPxePingServerParams) GetTftpdir() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tftpdir"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddBaremetalPxePingServerParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddBaremetalPxePingServerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddBaremetalPxePingServerParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddBaremetalPxePingServerParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalPxePingServerParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalPxePingServerParams(password string, physicalnetworkid string, pingdir string, pingstorageserverip string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxePingServerParams {
	p := &AddBaremetalPxePingServerParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["pingdir"] = pingdir
	p.p["pingstorageserverip"] = pingstorageserverip
	p.p["pxeservertype"] = pxeservertype
	p.p["tftpdir"] = tftpdir
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// add a baremetal ping pxe server
func (s *BaremetalService) AddBaremetalPxePingServer(p *AddBaremetalPxePingServerParams) (*AddBaremetalPxePingServerResponse, error) {
	resp, err := s.cs.newPostRequest("addBaremetalPxePingServer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalPxePingServerResponse
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

type AddBaremetalPxePingServerResponse struct {
	Id                  string `json:"id"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Pingdir             string `json:"pingdir"`
	Pingstorageserverip string `json:"pingstorageserverip"`
	Provider            string `json:"provider"`
	Tftpdir             string `json:"tftpdir"`
	Url                 string `json:"url"`
}

type AddBaremetalRctParams struct {
	p map[string]interface{}
}

func (p *AddBaremetalRctParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["baremetalrcturl"]; found {
		u.Set("baremetalrcturl", v.(string))
	}
	return u
}

func (p *AddBaremetalRctParams) SetBaremetalrcturl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["baremetalrcturl"] = v
}

func (p *AddBaremetalRctParams) ResetBaremetalrcturl() {
	if p.p != nil && p.p["baremetalrcturl"] != nil {
		delete(p.p, "baremetalrcturl")
	}
}

func (p *AddBaremetalRctParams) GetBaremetalrcturl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["baremetalrcturl"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalRctParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalRctParams(baremetalrcturl string) *AddBaremetalRctParams {
	p := &AddBaremetalRctParams{}
	p.p = make(map[string]interface{})
	p.p["baremetalrcturl"] = baremetalrcturl
	return p
}

// adds baremetal rack configuration text
func (s *BaremetalService) AddBaremetalRct(p *AddBaremetalRctParams) (*AddBaremetalRctResponse, error) {
	resp, err := s.cs.newPostRequest("addBaremetalRct", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalRctResponse
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

type AddBaremetalRctResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Url       string `json:"url"`
}

type DeleteBaremetalRctParams struct {
	p map[string]interface{}
}

func (p *DeleteBaremetalRctParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteBaremetalRctParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBaremetalRctParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBaremetalRctParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBaremetalRctParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewDeleteBaremetalRctParams(id string) *DeleteBaremetalRctParams {
	p := &DeleteBaremetalRctParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// deletes baremetal rack configuration text
func (s *BaremetalService) DeleteBaremetalRct(p *DeleteBaremetalRctParams) (*DeleteBaremetalRctResponse, error) {
	resp, err := s.cs.newPostRequest("deleteBaremetalRct", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBaremetalRctResponse
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

type DeleteBaremetalRctResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListBaremetalDhcpParams struct {
	p map[string]interface{}
}

func (p *ListBaremetalDhcpParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["dhcpservertype"]; found {
		u.Set("dhcpservertype", v.(string))
	}
	if v, found := p.p["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListBaremetalDhcpParams) SetDhcpservertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpservertype"] = v
}

func (p *ListBaremetalDhcpParams) ResetDhcpservertype() {
	if p.p != nil && p.p["dhcpservertype"] != nil {
		delete(p.p, "dhcpservertype")
	}
}

func (p *ListBaremetalDhcpParams) GetDhcpservertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpservertype"].(string)
	return value, ok
}

func (p *ListBaremetalDhcpParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBaremetalDhcpParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBaremetalDhcpParams) GetId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int64)
	return value, ok
}

func (p *ListBaremetalDhcpParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBaremetalDhcpParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBaremetalDhcpParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBaremetalDhcpParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBaremetalDhcpParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBaremetalDhcpParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBaremetalDhcpParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBaremetalDhcpParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBaremetalDhcpParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBaremetalDhcpParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListBaremetalDhcpParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ListBaremetalDhcpParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBaremetalDhcpParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewListBaremetalDhcpParams(physicalnetworkid string) *ListBaremetalDhcpParams {
	p := &ListBaremetalDhcpParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// list baremetal dhcp servers
func (s *BaremetalService) ListBaremetalDhcp(p *ListBaremetalDhcpParams) (*ListBaremetalDhcpResponse, error) {
	resp, err := s.cs.newRequest("listBaremetalDhcp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBaremetalDhcpResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBaremetalDhcpResponse struct {
	Count         int              `json:"count"`
	BaremetalDhcp []*BaremetalDhcp `json:"baremetaldhcp"`
}

type BaremetalDhcp struct {
	Dhcpservertype    string `json:"dhcpservertype"`
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Url               string `json:"url"`
}

type ListBaremetalPxeServersParams struct {
	p map[string]interface{}
}

func (p *ListBaremetalPxeServersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListBaremetalPxeServersParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBaremetalPxeServersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBaremetalPxeServersParams) GetId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int64)
	return value, ok
}

func (p *ListBaremetalPxeServersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBaremetalPxeServersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBaremetalPxeServersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBaremetalPxeServersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBaremetalPxeServersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBaremetalPxeServersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBaremetalPxeServersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBaremetalPxeServersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBaremetalPxeServersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBaremetalPxeServersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListBaremetalPxeServersParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ListBaremetalPxeServersParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBaremetalPxeServersParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewListBaremetalPxeServersParams(physicalnetworkid string) *ListBaremetalPxeServersParams {
	p := &ListBaremetalPxeServersParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// list baremetal pxe server
func (s *BaremetalService) ListBaremetalPxeServers(p *ListBaremetalPxeServersParams) (*ListBaremetalPxeServersResponse, error) {
	resp, err := s.cs.newRequest("listBaremetalPxeServers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBaremetalPxeServersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBaremetalPxeServersResponse struct {
	Count               int                   `json:"count"`
	BaremetalPxeServers []*BaremetalPxeServer `json:"baremetalpxeserver"`
}

type BaremetalPxeServer struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Url               string `json:"url"`
}

type ListBaremetalRctParams struct {
	p map[string]interface{}
}

func (p *ListBaremetalRctParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListBaremetalRctParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBaremetalRctParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBaremetalRctParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBaremetalRctParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBaremetalRctParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBaremetalRctParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBaremetalRctParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBaremetalRctParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBaremetalRctParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListBaremetalRctParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewListBaremetalRctParams() *ListBaremetalRctParams {
	p := &ListBaremetalRctParams{}
	p.p = make(map[string]interface{})
	return p
}

// list baremetal rack configuration
func (s *BaremetalService) ListBaremetalRct(p *ListBaremetalRctParams) (*ListBaremetalRctResponse, error) {
	resp, err := s.cs.newRequest("listBaremetalRct", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBaremetalRctResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBaremetalRctResponse struct {
	Count        int             `json:"count"`
	BaremetalRct []*BaremetalRct `json:"baremetalrct"`
}

type BaremetalRct struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Url       string `json:"url"`
}

type NotifyBaremetalProvisionDoneParams struct {
	p map[string]interface{}
}

func (p *NotifyBaremetalProvisionDoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["mac"]; found {
		u.Set("mac", v.(string))
	}
	return u
}

func (p *NotifyBaremetalProvisionDoneParams) SetMac(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mac"] = v
}

func (p *NotifyBaremetalProvisionDoneParams) ResetMac() {
	if p.p != nil && p.p["mac"] != nil {
		delete(p.p, "mac")
	}
}

func (p *NotifyBaremetalProvisionDoneParams) GetMac() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mac"].(string)
	return value, ok
}

// You should always use this function to get a new NotifyBaremetalProvisionDoneParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewNotifyBaremetalProvisionDoneParams(mac string) *NotifyBaremetalProvisionDoneParams {
	p := &NotifyBaremetalProvisionDoneParams{}
	p.p = make(map[string]interface{})
	p.p["mac"] = mac
	return p
}

// Notify provision has been done on a host. This api is for baremetal virtual router service, not for end user
func (s *BaremetalService) NotifyBaremetalProvisionDone(p *NotifyBaremetalProvisionDoneParams) (*NotifyBaremetalProvisionDoneResponse, error) {
	resp, err := s.cs.newPostRequest("notifyBaremetalProvisionDone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r NotifyBaremetalProvisionDoneResponse
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

type NotifyBaremetalProvisionDoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
