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

type NetworkACLServiceIface interface {
	CreateNetworkACL(p *CreateNetworkACLParams) (*CreateNetworkACLResponse, error)
	NewCreateNetworkACLParams(protocol string) *CreateNetworkACLParams
	CreateNetworkACLList(p *CreateNetworkACLListParams) (*CreateNetworkACLListResponse, error)
	NewCreateNetworkACLListParams(name string, vpcid string) *CreateNetworkACLListParams
	DeleteNetworkACL(p *DeleteNetworkACLParams) (*DeleteNetworkACLResponse, error)
	NewDeleteNetworkACLParams(id string) *DeleteNetworkACLParams
	DeleteNetworkACLList(p *DeleteNetworkACLListParams) (*DeleteNetworkACLListResponse, error)
	NewDeleteNetworkACLListParams(id string) *DeleteNetworkACLListParams
	ListNetworkACLLists(p *ListNetworkACLListsParams) (*ListNetworkACLListsResponse, error)
	NewListNetworkACLListsParams() *ListNetworkACLListsParams
	GetNetworkACLListID(name string, opts ...OptionFunc) (string, int, error)
	GetNetworkACLListByName(name string, opts ...OptionFunc) (*NetworkACLList, int, error)
	GetNetworkACLListByID(id string, opts ...OptionFunc) (*NetworkACLList, int, error)
	ListNetworkACLs(p *ListNetworkACLsParams) (*ListNetworkACLsResponse, error)
	NewListNetworkACLsParams() *ListNetworkACLsParams
	GetNetworkACLByID(id string, opts ...OptionFunc) (*NetworkACL, int, error)
	MoveNetworkAclItem(p *MoveNetworkAclItemParams) (*MoveNetworkAclItemResponse, error)
	NewMoveNetworkAclItemParams(id string) *MoveNetworkAclItemParams
	ReplaceNetworkACLList(p *ReplaceNetworkACLListParams) (*ReplaceNetworkACLListResponse, error)
	NewReplaceNetworkACLListParams(aclid string) *ReplaceNetworkACLListParams
	UpdateNetworkACLItem(p *UpdateNetworkACLItemParams) (*UpdateNetworkACLItemResponse, error)
	NewUpdateNetworkACLItemParams(id string) *UpdateNetworkACLItemParams
	UpdateNetworkACLList(p *UpdateNetworkACLListParams) (*UpdateNetworkACLListResponse, error)
	NewUpdateNetworkACLListParams(id string) *UpdateNetworkACLListParams
}

type CreateNetworkACLParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkACLParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["number"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("number", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["reason"]; found {
		u.Set("reason", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *CreateNetworkACLParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
}

func (p *CreateNetworkACLParams) ResetAclid() {
	if p.p != nil && p.p["aclid"] != nil {
		delete(p.p, "aclid")
	}
}

func (p *CreateNetworkACLParams) GetAclid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["aclid"].(string)
	return value, ok
}

func (p *CreateNetworkACLParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *CreateNetworkACLParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *CreateNetworkACLParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *CreateNetworkACLParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateNetworkACLParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreateNetworkACLParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *CreateNetworkACLParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *CreateNetworkACLParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *CreateNetworkACLParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *CreateNetworkACLParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateNetworkACLParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateNetworkACLParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateNetworkACLParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *CreateNetworkACLParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *CreateNetworkACLParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *CreateNetworkACLParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *CreateNetworkACLParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *CreateNetworkACLParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *CreateNetworkACLParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateNetworkACLParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateNetworkACLParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateNetworkACLParams) SetNumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["number"] = v
}

func (p *CreateNetworkACLParams) ResetNumber() {
	if p.p != nil && p.p["number"] != nil {
		delete(p.p, "number")
	}
}

func (p *CreateNetworkACLParams) GetNumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["number"].(int)
	return value, ok
}

func (p *CreateNetworkACLParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateNetworkACLParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreateNetworkACLParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreateNetworkACLParams) SetReason(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["reason"] = v
}

func (p *CreateNetworkACLParams) ResetReason() {
	if p.p != nil && p.p["reason"] != nil {
		delete(p.p, "reason")
	}
}

func (p *CreateNetworkACLParams) GetReason() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["reason"].(string)
	return value, ok
}

func (p *CreateNetworkACLParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *CreateNetworkACLParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *CreateNetworkACLParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *CreateNetworkACLParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *CreateNetworkACLParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *CreateNetworkACLParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkACLParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewCreateNetworkACLParams(protocol string) *CreateNetworkACLParams {
	p := &CreateNetworkACLParams{}
	p.p = make(map[string]interface{})
	p.p["protocol"] = protocol
	return p
}

// Creates a ACL rule in the given network (the network has to belong to VPC)
func (s *NetworkACLService) CreateNetworkACL(p *CreateNetworkACLParams) (*CreateNetworkACLResponse, error) {
	resp, err := s.cs.newPostRequest("createNetworkACL", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkACLResponse
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

type CreateNetworkACLResponse struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type CreateNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreateNetworkACLListParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateNetworkACLListParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateNetworkACLListParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateNetworkACLListParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateNetworkACLListParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateNetworkACLListParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateNetworkACLListParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateNetworkACLListParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateNetworkACLListParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateNetworkACLListParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *CreateNetworkACLListParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *CreateNetworkACLListParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewCreateNetworkACLListParams(name string, vpcid string) *CreateNetworkACLListParams {
	p := &CreateNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["vpcid"] = vpcid
	return p
}

// Creates a network ACL. If no VPC is given, then it creates a global ACL that can be used by everyone.
func (s *NetworkACLService) CreateNetworkACLList(p *CreateNetworkACLListParams) (*CreateNetworkACLListResponse, error) {
	resp, err := s.cs.newPostRequest("createNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkACLListResponse
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

type CreateNetworkACLListResponse struct {
	Description string `json:"description"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Vpcid       string `json:"vpcid"`
	Vpcname     string `json:"vpcname"`
}

type DeleteNetworkACLParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkACLParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkACLParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetworkACLParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteNetworkACLParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkACLParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewDeleteNetworkACLParams(id string) *DeleteNetworkACLParams {
	p := &DeleteNetworkACLParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a network ACL
func (s *NetworkACLService) DeleteNetworkACL(p *DeleteNetworkACLParams) (*DeleteNetworkACLResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNetworkACL", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkACLResponse
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

type DeleteNetworkACLResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkACLListParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetworkACLListParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteNetworkACLListParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewDeleteNetworkACLListParams(id string) *DeleteNetworkACLListParams {
	p := &DeleteNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a network ACL
func (s *NetworkACLService) DeleteNetworkACLList(p *DeleteNetworkACLListParams) (*DeleteNetworkACLListResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkACLListResponse
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

type DeleteNetworkACLListResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNetworkACLListsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkACLListsParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
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
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListNetworkACLListsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListNetworkACLListsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListNetworkACLListsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListNetworkACLListsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListNetworkACLListsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListNetworkACLListsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListNetworkACLListsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListNetworkACLListsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListNetworkACLListsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListNetworkACLListsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListNetworkACLListsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworkACLListsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetworkACLListsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListNetworkACLListsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListNetworkACLListsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListNetworkACLListsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListNetworkACLListsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListNetworkACLListsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListNetworkACLListsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworkACLListsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetworkACLListsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworkACLListsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetworkACLListsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListNetworkACLListsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListNetworkACLListsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListNetworkACLListsParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListNetworkACLListsParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListNetworkACLListsParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkACLListsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewListNetworkACLListsParams() *ListNetworkACLListsParams {
	p := &ListNetworkACLListsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListNetworkACLListsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworkACLLists(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.NetworkACLLists[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetworkACLLists {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListByName(name string, opts ...OptionFunc) (*NetworkACLList, int, error) {
	id, count, err := s.GetNetworkACLListID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetNetworkACLListByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListByID(id string, opts ...OptionFunc) (*NetworkACLList, int, error) {
	p := &ListNetworkACLListsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworkACLLists(p)
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
		return l.NetworkACLLists[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for NetworkACLList UUID: %s!", id)
}

// Lists all network ACLs
func (s *NetworkACLService) ListNetworkACLLists(p *ListNetworkACLListsParams) (*ListNetworkACLListsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkACLLists", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkACLListsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkACLListsResponse struct {
	Count           int               `json:"count"`
	NetworkACLLists []*NetworkACLList `json:"networkacllist"`
}

type NetworkACLList struct {
	Description string `json:"description"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Vpcid       string `json:"vpcid"`
	Vpcname     string `json:"vpcname"`
}

type ListNetworkACLsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkACLsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
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
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *ListNetworkACLsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListNetworkACLsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListNetworkACLsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
}

func (p *ListNetworkACLsParams) ResetAclid() {
	if p.p != nil && p.p["aclid"] != nil {
		delete(p.p, "aclid")
	}
}

func (p *ListNetworkACLsParams) GetAclid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["aclid"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *ListNetworkACLsParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *ListNetworkACLsParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListNetworkACLsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListNetworkACLsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListNetworkACLsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListNetworkACLsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListNetworkACLsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListNetworkACLsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListNetworkACLsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListNetworkACLsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListNetworkACLsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListNetworkACLsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworkACLsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetworkACLsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListNetworkACLsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListNetworkACLsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListNetworkACLsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListNetworkACLsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListNetworkACLsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworkACLsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetworkACLsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworkACLsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworkACLsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetworkACLsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetworkACLsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListNetworkACLsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListNetworkACLsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *ListNetworkACLsParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *ListNetworkACLsParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListNetworkACLsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListNetworkACLsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListNetworkACLsParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *ListNetworkACLsParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *ListNetworkACLsParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkACLsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewListNetworkACLsParams() *ListNetworkACLsParams {
	p := &ListNetworkACLsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLByID(id string, opts ...OptionFunc) (*NetworkACL, int, error) {
	p := &ListNetworkACLsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworkACLs(p)
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
		return l.NetworkACLs[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for NetworkACL UUID: %s!", id)
}

// Lists all network ACL items
func (s *NetworkACLService) ListNetworkACLs(p *ListNetworkACLsParams) (*ListNetworkACLsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkACLs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkACLsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkACLsResponse struct {
	Count       int           `json:"count"`
	NetworkACLs []*NetworkACL `json:"networkacl"`
}

type NetworkACL struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type MoveNetworkAclItemParams struct {
	p map[string]interface{}
}

func (p *MoveNetworkAclItemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclconsistencyhash"]; found {
		u.Set("aclconsistencyhash", v.(string))
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["nextaclruleid"]; found {
		u.Set("nextaclruleid", v.(string))
	}
	if v, found := p.p["previousaclruleid"]; found {
		u.Set("previousaclruleid", v.(string))
	}
	return u
}

func (p *MoveNetworkAclItemParams) SetAclconsistencyhash(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclconsistencyhash"] = v
}

func (p *MoveNetworkAclItemParams) ResetAclconsistencyhash() {
	if p.p != nil && p.p["aclconsistencyhash"] != nil {
		delete(p.p, "aclconsistencyhash")
	}
}

func (p *MoveNetworkAclItemParams) GetAclconsistencyhash() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["aclconsistencyhash"].(string)
	return value, ok
}

func (p *MoveNetworkAclItemParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *MoveNetworkAclItemParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *MoveNetworkAclItemParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *MoveNetworkAclItemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *MoveNetworkAclItemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *MoveNetworkAclItemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *MoveNetworkAclItemParams) SetNextaclruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nextaclruleid"] = v
}

func (p *MoveNetworkAclItemParams) ResetNextaclruleid() {
	if p.p != nil && p.p["nextaclruleid"] != nil {
		delete(p.p, "nextaclruleid")
	}
}

func (p *MoveNetworkAclItemParams) GetNextaclruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nextaclruleid"].(string)
	return value, ok
}

func (p *MoveNetworkAclItemParams) SetPreviousaclruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["previousaclruleid"] = v
}

func (p *MoveNetworkAclItemParams) ResetPreviousaclruleid() {
	if p.p != nil && p.p["previousaclruleid"] != nil {
		delete(p.p, "previousaclruleid")
	}
}

func (p *MoveNetworkAclItemParams) GetPreviousaclruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["previousaclruleid"].(string)
	return value, ok
}

// You should always use this function to get a new MoveNetworkAclItemParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewMoveNetworkAclItemParams(id string) *MoveNetworkAclItemParams {
	p := &MoveNetworkAclItemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Move an ACL rule to a position bettwen two other ACL rules of the same ACL network list
func (s *NetworkACLService) MoveNetworkAclItem(p *MoveNetworkAclItemParams) (*MoveNetworkAclItemResponse, error) {
	resp, err := s.cs.newPostRequest("moveNetworkAclItem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MoveNetworkAclItemResponse
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

type MoveNetworkAclItemResponse struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type ReplaceNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *ReplaceNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (p *ReplaceNetworkACLListParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
}

func (p *ReplaceNetworkACLListParams) ResetAclid() {
	if p.p != nil && p.p["aclid"] != nil {
		delete(p.p, "aclid")
	}
}

func (p *ReplaceNetworkACLListParams) GetAclid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["aclid"].(string)
	return value, ok
}

func (p *ReplaceNetworkACLListParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
}

func (p *ReplaceNetworkACLListParams) ResetGatewayid() {
	if p.p != nil && p.p["gatewayid"] != nil {
		delete(p.p, "gatewayid")
	}
}

func (p *ReplaceNetworkACLListParams) GetGatewayid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gatewayid"].(string)
	return value, ok
}

func (p *ReplaceNetworkACLListParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ReplaceNetworkACLListParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ReplaceNetworkACLListParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

// You should always use this function to get a new ReplaceNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewReplaceNetworkACLListParams(aclid string) *ReplaceNetworkACLListParams {
	p := &ReplaceNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["aclid"] = aclid
	return p
}

// Replaces ACL associated with a network or private gateway
func (s *NetworkACLService) ReplaceNetworkACLList(p *ReplaceNetworkACLListParams) (*ReplaceNetworkACLListResponse, error) {
	resp, err := s.cs.newPostRequest("replaceNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReplaceNetworkACLListResponse
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

type ReplaceNetworkACLListResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateNetworkACLItemParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkACLItemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["number"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("number", vv)
	}
	if v, found := p.p["partialupgrade"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("partialupgrade", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["reason"]; found {
		u.Set("reason", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *UpdateNetworkACLItemParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *UpdateNetworkACLItemParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *UpdateNetworkACLItemParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *UpdateNetworkACLItemParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *UpdateNetworkACLItemParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateNetworkACLItemParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateNetworkACLItemParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *UpdateNetworkACLItemParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *UpdateNetworkACLItemParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateNetworkACLItemParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateNetworkACLItemParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *UpdateNetworkACLItemParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *UpdateNetworkACLItemParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *UpdateNetworkACLItemParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *UpdateNetworkACLItemParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateNetworkACLItemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateNetworkACLItemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetNumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["number"] = v
}

func (p *UpdateNetworkACLItemParams) ResetNumber() {
	if p.p != nil && p.p["number"] != nil {
		delete(p.p, "number")
	}
}

func (p *UpdateNetworkACLItemParams) GetNumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["number"].(int)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetPartialupgrade(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["partialupgrade"] = v
}

func (p *UpdateNetworkACLItemParams) ResetPartialupgrade() {
	if p.p != nil && p.p["partialupgrade"] != nil {
		delete(p.p, "partialupgrade")
	}
}

func (p *UpdateNetworkACLItemParams) GetPartialupgrade() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["partialupgrade"].(bool)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *UpdateNetworkACLItemParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *UpdateNetworkACLItemParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetReason(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["reason"] = v
}

func (p *UpdateNetworkACLItemParams) ResetReason() {
	if p.p != nil && p.p["reason"] != nil {
		delete(p.p, "reason")
	}
}

func (p *UpdateNetworkACLItemParams) GetReason() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["reason"].(string)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *UpdateNetworkACLItemParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *UpdateNetworkACLItemParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *UpdateNetworkACLItemParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *UpdateNetworkACLItemParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *UpdateNetworkACLItemParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkACLItemParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewUpdateNetworkACLItemParams(id string) *UpdateNetworkACLItemParams {
	p := &UpdateNetworkACLItemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates ACL item with specified ID
func (s *NetworkACLService) UpdateNetworkACLItem(p *UpdateNetworkACLItemParams) (*UpdateNetworkACLItemResponse, error) {
	resp, err := s.cs.newPostRequest("updateNetworkACLItem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkACLItemResponse
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

type UpdateNetworkACLItemResponse struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type UpdateNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateNetworkACLListParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateNetworkACLListParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateNetworkACLListParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateNetworkACLListParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateNetworkACLListParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateNetworkACLListParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateNetworkACLListParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateNetworkACLListParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateNetworkACLListParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateNetworkACLListParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateNetworkACLListParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateNetworkACLListParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateNetworkACLListParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateNetworkACLListParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateNetworkACLListParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewUpdateNetworkACLListParams(id string) *UpdateNetworkACLListParams {
	p := &UpdateNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates network ACL list
func (s *NetworkACLService) UpdateNetworkACLList(p *UpdateNetworkACLListParams) (*UpdateNetworkACLListResponse, error) {
	resp, err := s.cs.newPostRequest("updateNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkACLListResponse
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

type UpdateNetworkACLListResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
