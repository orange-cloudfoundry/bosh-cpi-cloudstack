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

type VPNServiceIface interface {
	AddVpnUser(p *AddVpnUserParams) (*AddVpnUserResponse, error)
	NewAddVpnUserParams(password string, username string) *AddVpnUserParams
	CreateRemoteAccessVpn(p *CreateRemoteAccessVpnParams) (*CreateRemoteAccessVpnResponse, error)
	NewCreateRemoteAccessVpnParams(publicipid string) *CreateRemoteAccessVpnParams
	CreateVpnConnection(p *CreateVpnConnectionParams) (*CreateVpnConnectionResponse, error)
	NewCreateVpnConnectionParams(s2scustomergatewayid string, s2svpngatewayid string) *CreateVpnConnectionParams
	CreateVpnCustomerGateway(p *CreateVpnCustomerGatewayParams) (*CreateVpnCustomerGatewayResponse, error)
	NewCreateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, ikepolicy string, ipsecpsk string) *CreateVpnCustomerGatewayParams
	CreateVpnGateway(p *CreateVpnGatewayParams) (*CreateVpnGatewayResponse, error)
	NewCreateVpnGatewayParams(vpcid string) *CreateVpnGatewayParams
	DeleteRemoteAccessVpn(p *DeleteRemoteAccessVpnParams) (*DeleteRemoteAccessVpnResponse, error)
	NewDeleteRemoteAccessVpnParams(publicipid string) *DeleteRemoteAccessVpnParams
	DeleteVpnConnection(p *DeleteVpnConnectionParams) (*DeleteVpnConnectionResponse, error)
	NewDeleteVpnConnectionParams(id string) *DeleteVpnConnectionParams
	DeleteVpnCustomerGateway(p *DeleteVpnCustomerGatewayParams) (*DeleteVpnCustomerGatewayResponse, error)
	NewDeleteVpnCustomerGatewayParams(id string) *DeleteVpnCustomerGatewayParams
	DeleteVpnGateway(p *DeleteVpnGatewayParams) (*DeleteVpnGatewayResponse, error)
	NewDeleteVpnGatewayParams(id string) *DeleteVpnGatewayParams
	ListRemoteAccessVpns(p *ListRemoteAccessVpnsParams) (*ListRemoteAccessVpnsResponse, error)
	NewListRemoteAccessVpnsParams() *ListRemoteAccessVpnsParams
	GetRemoteAccessVpnByID(id string, opts ...OptionFunc) (*RemoteAccessVpn, int, error)
	ListVpnConnections(p *ListVpnConnectionsParams) (*ListVpnConnectionsResponse, error)
	NewListVpnConnectionsParams() *ListVpnConnectionsParams
	GetVpnConnectionByID(id string, opts ...OptionFunc) (*VpnConnection, int, error)
	ListVpnCustomerGateways(p *ListVpnCustomerGatewaysParams) (*ListVpnCustomerGatewaysResponse, error)
	NewListVpnCustomerGatewaysParams() *ListVpnCustomerGatewaysParams
	GetVpnCustomerGatewayID(keyword string, opts ...OptionFunc) (string, int, error)
	GetVpnCustomerGatewayByName(name string, opts ...OptionFunc) (*VpnCustomerGateway, int, error)
	GetVpnCustomerGatewayByID(id string, opts ...OptionFunc) (*VpnCustomerGateway, int, error)
	ListVpnGateways(p *ListVpnGatewaysParams) (*ListVpnGatewaysResponse, error)
	NewListVpnGatewaysParams() *ListVpnGatewaysParams
	GetVpnGatewayByID(id string, opts ...OptionFunc) (*VpnGateway, int, error)
	ListVpnUsers(p *ListVpnUsersParams) (*ListVpnUsersResponse, error)
	NewListVpnUsersParams() *ListVpnUsersParams
	GetVpnUserByID(id string, opts ...OptionFunc) (*VpnUser, int, error)
	RemoveVpnUser(p *RemoveVpnUserParams) (*RemoveVpnUserResponse, error)
	NewRemoveVpnUserParams(username string) *RemoveVpnUserParams
	ResetVpnConnection(p *ResetVpnConnectionParams) (*ResetVpnConnectionResponse, error)
	NewResetVpnConnectionParams(id string) *ResetVpnConnectionParams
	UpdateRemoteAccessVpn(p *UpdateRemoteAccessVpnParams) (*UpdateRemoteAccessVpnResponse, error)
	NewUpdateRemoteAccessVpnParams(id string) *UpdateRemoteAccessVpnParams
	UpdateVpnConnection(p *UpdateVpnConnectionParams) (*UpdateVpnConnectionResponse, error)
	NewUpdateVpnConnectionParams(id string) *UpdateVpnConnectionParams
	UpdateVpnCustomerGateway(p *UpdateVpnCustomerGatewayParams) (*UpdateVpnCustomerGatewayResponse, error)
	NewUpdateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, id string, ikepolicy string, ipsecpsk string) *UpdateVpnCustomerGatewayParams
	UpdateVpnGateway(p *UpdateVpnGatewayParams) (*UpdateVpnGatewayResponse, error)
	NewUpdateVpnGatewayParams(id string) *UpdateVpnGatewayParams
}

type AddVpnUserParams struct {
	p map[string]interface{}
}

func (p *AddVpnUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddVpnUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AddVpnUserParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *AddVpnUserParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AddVpnUserParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AddVpnUserParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *AddVpnUserParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *AddVpnUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddVpnUserParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddVpnUserParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddVpnUserParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AddVpnUserParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *AddVpnUserParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AddVpnUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddVpnUserParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddVpnUserParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddVpnUserParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewAddVpnUserParams(password string, username string) *AddVpnUserParams {
	p := &AddVpnUserParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Adds vpn users
func (s *VPNService) AddVpnUser(p *AddVpnUserParams) (*AddVpnUserResponse, error) {
	resp, err := s.cs.newPostRequest("addVpnUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddVpnUserResponse
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

type AddVpnUserResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	State      string `json:"state"`
	Username   string `json:"username"`
}

type CreateRemoteAccessVpnParams struct {
	p map[string]interface{}
}

func (p *CreateRemoteAccessVpnParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["iprange"]; found {
		u.Set("iprange", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (p *CreateRemoteAccessVpnParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateRemoteAccessVpnParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateRemoteAccessVpnParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateRemoteAccessVpnParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateRemoteAccessVpnParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateRemoteAccessVpnParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateRemoteAccessVpnParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateRemoteAccessVpnParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateRemoteAccessVpnParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateRemoteAccessVpnParams) SetIprange(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iprange"] = v
}

func (p *CreateRemoteAccessVpnParams) ResetIprange() {
	if p.p != nil && p.p["iprange"] != nil {
		delete(p.p, "iprange")
	}
}

func (p *CreateRemoteAccessVpnParams) GetIprange() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iprange"].(string)
	return value, ok
}

func (p *CreateRemoteAccessVpnParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
}

func (p *CreateRemoteAccessVpnParams) ResetOpenfirewall() {
	if p.p != nil && p.p["openfirewall"] != nil {
		delete(p.p, "openfirewall")
	}
}

func (p *CreateRemoteAccessVpnParams) GetOpenfirewall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["openfirewall"].(bool)
	return value, ok
}

func (p *CreateRemoteAccessVpnParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
}

func (p *CreateRemoteAccessVpnParams) ResetPublicipid() {
	if p.p != nil && p.p["publicipid"] != nil {
		delete(p.p, "publicipid")
	}
}

func (p *CreateRemoteAccessVpnParams) GetPublicipid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicipid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateRemoteAccessVpnParams(publicipid string) *CreateRemoteAccessVpnParams {
	p := &CreateRemoteAccessVpnParams{}
	p.p = make(map[string]interface{})
	p.p["publicipid"] = publicipid
	return p
}

// Creates a l2tp/ipsec remote access vpn
func (s *VPNService) CreateRemoteAccessVpn(p *CreateRemoteAccessVpnParams) (*CreateRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newPostRequest("createRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateRemoteAccessVpnResponse
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

type CreateRemoteAccessVpnResponse struct {
	Account      string `json:"account"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Domainpath   string `json:"domainpath"`
	Fordisplay   bool   `json:"fordisplay"`
	Id           string `json:"id"`
	Iprange      string `json:"iprange"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Presharedkey string `json:"presharedkey"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Publicip     string `json:"publicip"`
	Publicipid   string `json:"publicipid"`
	State        string `json:"state"`
}

type CreateVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *CreateVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["passive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passive", vv)
	}
	if v, found := p.p["s2scustomergatewayid"]; found {
		u.Set("s2scustomergatewayid", v.(string))
	}
	if v, found := p.p["s2svpngatewayid"]; found {
		u.Set("s2svpngatewayid", v.(string))
	}
	return u
}

func (p *CreateVpnConnectionParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateVpnConnectionParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateVpnConnectionParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateVpnConnectionParams) SetPassive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passive"] = v
}

func (p *CreateVpnConnectionParams) ResetPassive() {
	if p.p != nil && p.p["passive"] != nil {
		delete(p.p, "passive")
	}
}

func (p *CreateVpnConnectionParams) GetPassive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passive"].(bool)
	return value, ok
}

func (p *CreateVpnConnectionParams) SetS2scustomergatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["s2scustomergatewayid"] = v
}

func (p *CreateVpnConnectionParams) ResetS2scustomergatewayid() {
	if p.p != nil && p.p["s2scustomergatewayid"] != nil {
		delete(p.p, "s2scustomergatewayid")
	}
}

func (p *CreateVpnConnectionParams) GetS2scustomergatewayid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["s2scustomergatewayid"].(string)
	return value, ok
}

func (p *CreateVpnConnectionParams) SetS2svpngatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["s2svpngatewayid"] = v
}

func (p *CreateVpnConnectionParams) ResetS2svpngatewayid() {
	if p.p != nil && p.p["s2svpngatewayid"] != nil {
		delete(p.p, "s2svpngatewayid")
	}
}

func (p *CreateVpnConnectionParams) GetS2svpngatewayid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["s2svpngatewayid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnConnectionParams(s2scustomergatewayid string, s2svpngatewayid string) *CreateVpnConnectionParams {
	p := &CreateVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["s2scustomergatewayid"] = s2scustomergatewayid
	p.p["s2svpngatewayid"] = s2svpngatewayid
	return p
}

// Create site to site vpn connection
func (s *VPNService) CreateVpnConnection(p *CreateVpnConnectionParams) (*CreateVpnConnectionResponse, error) {
	resp, err := s.cs.newPostRequest("createVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnConnectionResponse
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

type CreateVpnConnectionResponse struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Domainpath           string `json:"domainpath"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type CreateVpnCustomerGatewayParams struct {
	p map[string]interface{}
}

func (p *CreateVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		u.Set("cidrlist", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["dpd"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dpd", vv)
	}
	if v, found := p.p["esplifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("esplifetime", vv)
	}
	if v, found := p.p["esppolicy"]; found {
		u.Set("esppolicy", v.(string))
	}
	if v, found := p.p["forceencap"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forceencap", vv)
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["ikelifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("ikelifetime", vv)
	}
	if v, found := p.p["ikepolicy"]; found {
		u.Set("ikepolicy", v.(string))
	}
	if v, found := p.p["ikeversion"]; found {
		u.Set("ikeversion", v.(string))
	}
	if v, found := p.p["ipsecpsk"]; found {
		u.Set("ipsecpsk", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["splitconnections"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("splitconnections", vv)
	}
	return u
}

func (p *CreateVpnCustomerGatewayParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetCidrlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetCidrlist() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetDpd(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dpd"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetDpd() {
	if p.p != nil && p.p["dpd"] != nil {
		delete(p.p, "dpd")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetDpd() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dpd"].(bool)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetEsplifetime(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esplifetime"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetEsplifetime() {
	if p.p != nil && p.p["esplifetime"] != nil {
		delete(p.p, "esplifetime")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetEsplifetime() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["esplifetime"].(int64)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetEsppolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esppolicy"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetEsppolicy() {
	if p.p != nil && p.p["esppolicy"] != nil {
		delete(p.p, "esppolicy")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetEsppolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["esppolicy"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetForceencap(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forceencap"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetForceencap() {
	if p.p != nil && p.p["forceencap"] != nil {
		delete(p.p, "forceencap")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetForceencap() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forceencap"].(bool)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetIkelifetime(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikelifetime"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetIkelifetime() {
	if p.p != nil && p.p["ikelifetime"] != nil {
		delete(p.p, "ikelifetime")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetIkelifetime() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ikelifetime"].(int64)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetIkepolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikepolicy"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetIkepolicy() {
	if p.p != nil && p.p["ikepolicy"] != nil {
		delete(p.p, "ikepolicy")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetIkepolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ikepolicy"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetIkeversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikeversion"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetIkeversion() {
	if p.p != nil && p.p["ikeversion"] != nil {
		delete(p.p, "ikeversion")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetIkeversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ikeversion"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetIpsecpsk(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipsecpsk"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetIpsecpsk() {
	if p.p != nil && p.p["ipsecpsk"] != nil {
		delete(p.p, "ipsecpsk")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetIpsecpsk() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipsecpsk"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateVpnCustomerGatewayParams) SetSplitconnections(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["splitconnections"] = v
}

func (p *CreateVpnCustomerGatewayParams) ResetSplitconnections() {
	if p.p != nil && p.p["splitconnections"] != nil {
		delete(p.p, "splitconnections")
	}
}

func (p *CreateVpnCustomerGatewayParams) GetSplitconnections() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["splitconnections"].(bool)
	return value, ok
}

// You should always use this function to get a new CreateVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, ikepolicy string, ipsecpsk string) *CreateVpnCustomerGatewayParams {
	p := &CreateVpnCustomerGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["cidrlist"] = cidrlist
	p.p["esppolicy"] = esppolicy
	p.p["gateway"] = gateway
	p.p["ikepolicy"] = ikepolicy
	p.p["ipsecpsk"] = ipsecpsk
	return p
}

// Creates site to site vpn customer gateway
func (s *VPNService) CreateVpnCustomerGateway(p *CreateVpnCustomerGatewayParams) (*CreateVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("createVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnCustomerGatewayResponse
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

type CreateVpnCustomerGatewayResponse struct {
	Account          string `json:"account"`
	Cidrlist         string `json:"cidrlist"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	Dpd              bool   `json:"dpd"`
	Esplifetime      int64  `json:"esplifetime"`
	Esppolicy        string `json:"esppolicy"`
	Forceencap       bool   `json:"forceencap"`
	Gateway          string `json:"gateway"`
	Hasannotations   bool   `json:"hasannotations"`
	Id               string `json:"id"`
	Ikelifetime      int64  `json:"ikelifetime"`
	Ikepolicy        string `json:"ikepolicy"`
	Ikeversion       string `json:"ikeversion"`
	Ipaddress        string `json:"ipaddress"`
	Ipsecpsk         string `json:"ipsecpsk"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Removed          string `json:"removed"`
	Splitconnections bool   `json:"splitconnections"`
}

type CreateVpnGatewayParams struct {
	p map[string]interface{}
}

func (p *CreateVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreateVpnGatewayParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateVpnGatewayParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateVpnGatewayParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateVpnGatewayParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *CreateVpnGatewayParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *CreateVpnGatewayParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnGatewayParams(vpcid string) *CreateVpnGatewayParams {
	p := &CreateVpnGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["vpcid"] = vpcid
	return p
}

// Creates site to site vpn local gateway
func (s *VPNService) CreateVpnGateway(p *CreateVpnGatewayParams) (*CreateVpnGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("createVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnGatewayResponse
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

type CreateVpnGatewayResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Fordisplay bool   `json:"fordisplay"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	Publicip   string `json:"publicip"`
	Removed    string `json:"removed"`
	Vpcid      string `json:"vpcid"`
	Vpcname    string `json:"vpcname"`
}

type DeleteRemoteAccessVpnParams struct {
	p map[string]interface{}
}

func (p *DeleteRemoteAccessVpnParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (p *DeleteRemoteAccessVpnParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
}

func (p *DeleteRemoteAccessVpnParams) ResetPublicipid() {
	if p.p != nil && p.p["publicipid"] != nil {
		delete(p.p, "publicipid")
	}
}

func (p *DeleteRemoteAccessVpnParams) GetPublicipid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicipid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteRemoteAccessVpnParams(publicipid string) *DeleteRemoteAccessVpnParams {
	p := &DeleteRemoteAccessVpnParams{}
	p.p = make(map[string]interface{})
	p.p["publicipid"] = publicipid
	return p
}

// Destroys a l2tp/ipsec remote access vpn
func (s *VPNService) DeleteRemoteAccessVpn(p *DeleteRemoteAccessVpnParams) (*DeleteRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newPostRequest("deleteRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRemoteAccessVpnResponse
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

type DeleteRemoteAccessVpnResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *DeleteVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVpnConnectionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVpnConnectionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVpnConnectionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnConnectionParams(id string) *DeleteVpnConnectionParams {
	p := &DeleteVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete site to site vpn connection
func (s *VPNService) DeleteVpnConnection(p *DeleteVpnConnectionParams) (*DeleteVpnConnectionResponse, error) {
	resp, err := s.cs.newPostRequest("deleteVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnConnectionResponse
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

type DeleteVpnConnectionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVpnCustomerGatewayParams struct {
	p map[string]interface{}
}

func (p *DeleteVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVpnCustomerGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVpnCustomerGatewayParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVpnCustomerGatewayParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnCustomerGatewayParams(id string) *DeleteVpnCustomerGatewayParams {
	p := &DeleteVpnCustomerGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete site to site vpn customer gateway
func (s *VPNService) DeleteVpnCustomerGateway(p *DeleteVpnCustomerGatewayParams) (*DeleteVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("deleteVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnCustomerGatewayResponse
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

type DeleteVpnCustomerGatewayResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVpnGatewayParams struct {
	p map[string]interface{}
}

func (p *DeleteVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVpnGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVpnGatewayParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVpnGatewayParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnGatewayParams(id string) *DeleteVpnGatewayParams {
	p := &DeleteVpnGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete site to site vpn gateway
func (s *VPNService) DeleteVpnGateway(p *DeleteVpnGatewayParams) (*DeleteVpnGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("deleteVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnGatewayResponse
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

type DeleteVpnGatewayResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListRemoteAccessVpnsParams struct {
	p map[string]interface{}
}

func (p *ListRemoteAccessVpnsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (p *ListRemoteAccessVpnsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListRemoteAccessVpnsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListRemoteAccessVpnsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListRemoteAccessVpnsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListRemoteAccessVpnsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListRemoteAccessVpnsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListRemoteAccessVpnsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListRemoteAccessVpnsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListRemoteAccessVpnsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListRemoteAccessVpnsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListRemoteAccessVpnsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListRemoteAccessVpnsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListRemoteAccessVpnsParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
}

func (p *ListRemoteAccessVpnsParams) ResetPublicipid() {
	if p.p != nil && p.p["publicipid"] != nil {
		delete(p.p, "publicipid")
	}
}

func (p *ListRemoteAccessVpnsParams) GetPublicipid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicipid"].(string)
	return value, ok
}

// You should always use this function to get a new ListRemoteAccessVpnsParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListRemoteAccessVpnsParams() *ListRemoteAccessVpnsParams {
	p := &ListRemoteAccessVpnsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetRemoteAccessVpnByID(id string, opts ...OptionFunc) (*RemoteAccessVpn, int, error) {
	p := &ListRemoteAccessVpnsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListRemoteAccessVpns(p)
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
		return l.RemoteAccessVpns[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for RemoteAccessVpn UUID: %s!", id)
}

// Lists remote access vpns
func (s *VPNService) ListRemoteAccessVpns(p *ListRemoteAccessVpnsParams) (*ListRemoteAccessVpnsResponse, error) {
	resp, err := s.cs.newRequest("listRemoteAccessVpns", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRemoteAccessVpnsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRemoteAccessVpnsResponse struct {
	Count            int                `json:"count"`
	RemoteAccessVpns []*RemoteAccessVpn `json:"remoteaccessvpn"`
}

type RemoteAccessVpn struct {
	Account      string `json:"account"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Domainpath   string `json:"domainpath"`
	Fordisplay   bool   `json:"fordisplay"`
	Id           string `json:"id"`
	Iprange      string `json:"iprange"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Presharedkey string `json:"presharedkey"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Publicip     string `json:"publicip"`
	Publicipid   string `json:"publicipid"`
	State        string `json:"state"`
}

type ListVpnConnectionsParams struct {
	p map[string]interface{}
}

func (p *ListVpnConnectionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListVpnConnectionsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVpnConnectionsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVpnConnectionsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVpnConnectionsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVpnConnectionsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListVpnConnectionsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListVpnConnectionsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVpnConnectionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVpnConnectionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVpnConnectionsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVpnConnectionsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVpnConnectionsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVpnConnectionsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVpnConnectionsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVpnConnectionsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVpnConnectionsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVpnConnectionsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVpnConnectionsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVpnConnectionsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVpnConnectionsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVpnConnectionsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVpnConnectionsParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListVpnConnectionsParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListVpnConnectionsParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnConnectionsParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnConnectionsParams() *ListVpnConnectionsParams {
	p := &ListVpnConnectionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnConnectionByID(id string, opts ...OptionFunc) (*VpnConnection, int, error) {
	p := &ListVpnConnectionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnConnections(p)
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
		return l.VpnConnections[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnConnection UUID: %s!", id)
}

// Lists site to site vpn connection gateways
func (s *VPNService) ListVpnConnections(p *ListVpnConnectionsParams) (*ListVpnConnectionsResponse, error) {
	resp, err := s.cs.newRequest("listVpnConnections", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnConnectionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnConnectionsResponse struct {
	Count          int              `json:"count"`
	VpnConnections []*VpnConnection `json:"vpnconnection"`
}

type VpnConnection struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Domainpath           string `json:"domainpath"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type ListVpnCustomerGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListVpnCustomerGatewaysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ListVpnCustomerGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVpnCustomerGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVpnCustomerGatewaysParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVpnCustomerGatewaysParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnCustomerGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnCustomerGatewaysParams() *ListVpnCustomerGatewaysParams {
	p := &ListVpnCustomerGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListVpnCustomerGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVpnCustomerGateways(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.VpnCustomerGateways[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VpnCustomerGateways {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayByName(name string, opts ...OptionFunc) (*VpnCustomerGateway, int, error) {
	id, count, err := s.GetVpnCustomerGatewayID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVpnCustomerGatewayByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayByID(id string, opts ...OptionFunc) (*VpnCustomerGateway, int, error) {
	p := &ListVpnCustomerGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnCustomerGateways(p)
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
		return l.VpnCustomerGateways[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnCustomerGateway UUID: %s!", id)
}

// Lists site to site vpn customer gateways
func (s *VPNService) ListVpnCustomerGateways(p *ListVpnCustomerGatewaysParams) (*ListVpnCustomerGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listVpnCustomerGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnCustomerGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnCustomerGatewaysResponse struct {
	Count               int                   `json:"count"`
	VpnCustomerGateways []*VpnCustomerGateway `json:"vpncustomergateway"`
}

type VpnCustomerGateway struct {
	Account          string `json:"account"`
	Cidrlist         string `json:"cidrlist"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	Dpd              bool   `json:"dpd"`
	Esplifetime      int64  `json:"esplifetime"`
	Esppolicy        string `json:"esppolicy"`
	Forceencap       bool   `json:"forceencap"`
	Gateway          string `json:"gateway"`
	Hasannotations   bool   `json:"hasannotations"`
	Id               string `json:"id"`
	Ikelifetime      int64  `json:"ikelifetime"`
	Ikepolicy        string `json:"ikepolicy"`
	Ikeversion       string `json:"ikeversion"`
	Ipaddress        string `json:"ipaddress"`
	Ipsecpsk         string `json:"ipsecpsk"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Removed          string `json:"removed"`
	Splitconnections bool   `json:"splitconnections"`
}

type ListVpnGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListVpnGatewaysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListVpnGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVpnGatewaysParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVpnGatewaysParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVpnGatewaysParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVpnGatewaysParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListVpnGatewaysParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListVpnGatewaysParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVpnGatewaysParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVpnGatewaysParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVpnGatewaysParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVpnGatewaysParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVpnGatewaysParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVpnGatewaysParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVpnGatewaysParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVpnGatewaysParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVpnGatewaysParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVpnGatewaysParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVpnGatewaysParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVpnGatewaysParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVpnGatewaysParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVpnGatewaysParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVpnGatewaysParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListVpnGatewaysParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListVpnGatewaysParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnGatewaysParams() *ListVpnGatewaysParams {
	p := &ListVpnGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnGatewayByID(id string, opts ...OptionFunc) (*VpnGateway, int, error) {
	p := &ListVpnGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnGateways(p)
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
		return l.VpnGateways[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnGateway UUID: %s!", id)
}

// Lists site 2 site vpn gateways
func (s *VPNService) ListVpnGateways(p *ListVpnGatewaysParams) (*ListVpnGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listVpnGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnGatewaysResponse struct {
	Count       int           `json:"count"`
	VpnGateways []*VpnGateway `json:"vpngateway"`
}

type VpnGateway struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Fordisplay bool   `json:"fordisplay"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	Publicip   string `json:"publicip"`
	Removed    string `json:"removed"`
	Vpcid      string `json:"vpcid"`
	Vpcname    string `json:"vpcname"`
}

type ListVpnUsersParams struct {
	p map[string]interface{}
}

func (p *ListVpnUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListVpnUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVpnUsersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVpnUsersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVpnUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVpnUsersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVpnUsersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVpnUsersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVpnUsersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVpnUsersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVpnUsersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVpnUsersParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVpnUsersParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVpnUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVpnUsersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVpnUsersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVpnUsersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVpnUsersParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVpnUsersParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVpnUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVpnUsersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVpnUsersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVpnUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVpnUsersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVpnUsersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVpnUsersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVpnUsersParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVpnUsersParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVpnUsersParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *ListVpnUsersParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *ListVpnUsersParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnUsersParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnUsersParams() *ListVpnUsersParams {
	p := &ListVpnUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnUserByID(id string, opts ...OptionFunc) (*VpnUser, int, error) {
	p := &ListVpnUsersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnUsers(p)
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
		return l.VpnUsers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnUser UUID: %s!", id)
}

// Lists vpn users
func (s *VPNService) ListVpnUsers(p *ListVpnUsersParams) (*ListVpnUsersResponse, error) {
	resp, err := s.cs.newRequest("listVpnUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnUsersResponse struct {
	Count    int        `json:"count"`
	VpnUsers []*VpnUser `json:"vpnuser"`
}

type VpnUser struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	State      string `json:"state"`
	Username   string `json:"username"`
}

type RemoveVpnUserParams struct {
	p map[string]interface{}
}

func (p *RemoveVpnUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *RemoveVpnUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *RemoveVpnUserParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *RemoveVpnUserParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *RemoveVpnUserParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *RemoveVpnUserParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *RemoveVpnUserParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *RemoveVpnUserParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *RemoveVpnUserParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *RemoveVpnUserParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *RemoveVpnUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *RemoveVpnUserParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *RemoveVpnUserParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveVpnUserParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewRemoveVpnUserParams(username string) *RemoveVpnUserParams {
	p := &RemoveVpnUserParams{}
	p.p = make(map[string]interface{})
	p.p["username"] = username
	return p
}

// Removes vpn user
func (s *VPNService) RemoveVpnUser(p *RemoveVpnUserParams) (*RemoveVpnUserResponse, error) {
	resp, err := s.cs.newPostRequest("removeVpnUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveVpnUserResponse
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

type RemoveVpnUserResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ResetVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *ResetVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ResetVpnConnectionParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ResetVpnConnectionParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ResetVpnConnectionParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ResetVpnConnectionParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ResetVpnConnectionParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ResetVpnConnectionParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ResetVpnConnectionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ResetVpnConnectionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ResetVpnConnectionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ResetVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewResetVpnConnectionParams(id string) *ResetVpnConnectionParams {
	p := &ResetVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reset site to site vpn connection
func (s *VPNService) ResetVpnConnection(p *ResetVpnConnectionParams) (*ResetVpnConnectionResponse, error) {
	resp, err := s.cs.newPostRequest("resetVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetVpnConnectionResponse
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

type ResetVpnConnectionResponse struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Domainpath           string `json:"domainpath"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type UpdateRemoteAccessVpnParams struct {
	p map[string]interface{}
}

func (p *UpdateRemoteAccessVpnParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateRemoteAccessVpnParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateRemoteAccessVpnParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateRemoteAccessVpnParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateRemoteAccessVpnParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateRemoteAccessVpnParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateRemoteAccessVpnParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateRemoteAccessVpnParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateRemoteAccessVpnParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateRemoteAccessVpnParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateRemoteAccessVpnParams(id string) *UpdateRemoteAccessVpnParams {
	p := &UpdateRemoteAccessVpnParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates remote access vpn
func (s *VPNService) UpdateRemoteAccessVpn(p *UpdateRemoteAccessVpnParams) (*UpdateRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newPostRequest("updateRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRemoteAccessVpnResponse
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

type UpdateRemoteAccessVpnResponse struct {
	Account      string `json:"account"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Domainpath   string `json:"domainpath"`
	Fordisplay   bool   `json:"fordisplay"`
	Id           string `json:"id"`
	Iprange      string `json:"iprange"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Presharedkey string `json:"presharedkey"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Publicip     string `json:"publicip"`
	Publicipid   string `json:"publicipid"`
	State        string `json:"state"`
}

type UpdateVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *UpdateVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateVpnConnectionParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateVpnConnectionParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateVpnConnectionParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateVpnConnectionParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateVpnConnectionParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateVpnConnectionParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateVpnConnectionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVpnConnectionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVpnConnectionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnConnectionParams(id string) *UpdateVpnConnectionParams {
	p := &UpdateVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates site to site vpn connection
func (s *VPNService) UpdateVpnConnection(p *UpdateVpnConnectionParams) (*UpdateVpnConnectionResponse, error) {
	resp, err := s.cs.newPostRequest("updateVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnConnectionResponse
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

type UpdateVpnConnectionResponse struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Domainpath           string `json:"domainpath"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type UpdateVpnCustomerGatewayParams struct {
	p map[string]interface{}
}

func (p *UpdateVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		u.Set("cidrlist", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["dpd"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dpd", vv)
	}
	if v, found := p.p["esplifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("esplifetime", vv)
	}
	if v, found := p.p["esppolicy"]; found {
		u.Set("esppolicy", v.(string))
	}
	if v, found := p.p["forceencap"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forceencap", vv)
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ikelifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("ikelifetime", vv)
	}
	if v, found := p.p["ikepolicy"]; found {
		u.Set("ikepolicy", v.(string))
	}
	if v, found := p.p["ikeversion"]; found {
		u.Set("ikeversion", v.(string))
	}
	if v, found := p.p["ipsecpsk"]; found {
		u.Set("ipsecpsk", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["splitconnections"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("splitconnections", vv)
	}
	return u
}

func (p *UpdateVpnCustomerGatewayParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetCidrlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetCidrlist() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetDpd(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dpd"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetDpd() {
	if p.p != nil && p.p["dpd"] != nil {
		delete(p.p, "dpd")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetDpd() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dpd"].(bool)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetEsplifetime(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esplifetime"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetEsplifetime() {
	if p.p != nil && p.p["esplifetime"] != nil {
		delete(p.p, "esplifetime")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetEsplifetime() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["esplifetime"].(int64)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetEsppolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esppolicy"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetEsppolicy() {
	if p.p != nil && p.p["esppolicy"] != nil {
		delete(p.p, "esppolicy")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetEsppolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["esppolicy"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetForceencap(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forceencap"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetForceencap() {
	if p.p != nil && p.p["forceencap"] != nil {
		delete(p.p, "forceencap")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetForceencap() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forceencap"].(bool)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetIkelifetime(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikelifetime"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetIkelifetime() {
	if p.p != nil && p.p["ikelifetime"] != nil {
		delete(p.p, "ikelifetime")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetIkelifetime() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ikelifetime"].(int64)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetIkepolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikepolicy"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetIkepolicy() {
	if p.p != nil && p.p["ikepolicy"] != nil {
		delete(p.p, "ikepolicy")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetIkepolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ikepolicy"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetIkeversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikeversion"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetIkeversion() {
	if p.p != nil && p.p["ikeversion"] != nil {
		delete(p.p, "ikeversion")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetIkeversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ikeversion"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetIpsecpsk(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipsecpsk"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetIpsecpsk() {
	if p.p != nil && p.p["ipsecpsk"] != nil {
		delete(p.p, "ipsecpsk")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetIpsecpsk() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipsecpsk"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVpnCustomerGatewayParams) SetSplitconnections(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["splitconnections"] = v
}

func (p *UpdateVpnCustomerGatewayParams) ResetSplitconnections() {
	if p.p != nil && p.p["splitconnections"] != nil {
		delete(p.p, "splitconnections")
	}
}

func (p *UpdateVpnCustomerGatewayParams) GetSplitconnections() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["splitconnections"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, id string, ikepolicy string, ipsecpsk string) *UpdateVpnCustomerGatewayParams {
	p := &UpdateVpnCustomerGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["cidrlist"] = cidrlist
	p.p["esppolicy"] = esppolicy
	p.p["gateway"] = gateway
	p.p["id"] = id
	p.p["ikepolicy"] = ikepolicy
	p.p["ipsecpsk"] = ipsecpsk
	return p
}

// Update site to site vpn customer gateway
func (s *VPNService) UpdateVpnCustomerGateway(p *UpdateVpnCustomerGatewayParams) (*UpdateVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("updateVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnCustomerGatewayResponse
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

type UpdateVpnCustomerGatewayResponse struct {
	Account          string `json:"account"`
	Cidrlist         string `json:"cidrlist"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	Dpd              bool   `json:"dpd"`
	Esplifetime      int64  `json:"esplifetime"`
	Esppolicy        string `json:"esppolicy"`
	Forceencap       bool   `json:"forceencap"`
	Gateway          string `json:"gateway"`
	Hasannotations   bool   `json:"hasannotations"`
	Id               string `json:"id"`
	Ikelifetime      int64  `json:"ikelifetime"`
	Ikepolicy        string `json:"ikepolicy"`
	Ikeversion       string `json:"ikeversion"`
	Ipaddress        string `json:"ipaddress"`
	Ipsecpsk         string `json:"ipsecpsk"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Removed          string `json:"removed"`
	Splitconnections bool   `json:"splitconnections"`
}

type UpdateVpnGatewayParams struct {
	p map[string]interface{}
}

func (p *UpdateVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateVpnGatewayParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateVpnGatewayParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateVpnGatewayParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateVpnGatewayParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateVpnGatewayParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateVpnGatewayParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateVpnGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVpnGatewayParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVpnGatewayParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnGatewayParams(id string) *UpdateVpnGatewayParams {
	p := &UpdateVpnGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates site to site vpn local gateway
func (s *VPNService) UpdateVpnGateway(p *UpdateVpnGatewayParams) (*UpdateVpnGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("updateVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnGatewayResponse
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

type UpdateVpnGatewayResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Fordisplay bool   `json:"fordisplay"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	Publicip   string `json:"publicip"`
	Removed    string `json:"removed"`
	Vpcid      string `json:"vpcid"`
	Vpcname    string `json:"vpcname"`
}
