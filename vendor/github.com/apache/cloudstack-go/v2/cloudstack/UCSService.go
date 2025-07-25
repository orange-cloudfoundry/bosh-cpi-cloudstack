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

type UCSServiceIface interface {
	AddUcsManager(p *AddUcsManagerParams) (*AddUcsManagerResponse, error)
	NewAddUcsManagerParams(password string, url string, username string, zoneid string) *AddUcsManagerParams
	AssociateUcsProfileToBlade(p *AssociateUcsProfileToBladeParams) (*AssociateUcsProfileToBladeResponse, error)
	NewAssociateUcsProfileToBladeParams(bladeid string, profiledn string, ucsmanagerid string) *AssociateUcsProfileToBladeParams
	DeleteUcsManager(p *DeleteUcsManagerParams) (*DeleteUcsManagerResponse, error)
	NewDeleteUcsManagerParams(ucsmanagerid string) *DeleteUcsManagerParams
	ListUcsBlades(p *ListUcsBladesParams) (*ListUcsBladesResponse, error)
	NewListUcsBladesParams(ucsmanagerid string) *ListUcsBladesParams
	ListUcsManagers(p *ListUcsManagersParams) (*ListUcsManagersResponse, error)
	NewListUcsManagersParams() *ListUcsManagersParams
	GetUcsManagerID(keyword string, opts ...OptionFunc) (string, int, error)
	GetUcsManagerByName(name string, opts ...OptionFunc) (*UcsManager, int, error)
	GetUcsManagerByID(id string, opts ...OptionFunc) (*UcsManager, int, error)
	ListUcsProfiles(p *ListUcsProfilesParams) (*ListUcsProfilesResponse, error)
	NewListUcsProfilesParams(ucsmanagerid string) *ListUcsProfilesParams
}

type AddUcsManagerParams struct {
	p map[string]interface{}
}

func (p *AddUcsManagerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddUcsManagerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddUcsManagerParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddUcsManagerParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddUcsManagerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddUcsManagerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddUcsManagerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddUcsManagerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddUcsManagerParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddUcsManagerParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddUcsManagerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddUcsManagerParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddUcsManagerParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *AddUcsManagerParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddUcsManagerParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddUcsManagerParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddUcsManagerParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewAddUcsManagerParams(password string, url string, username string, zoneid string) *AddUcsManagerParams {
	p := &AddUcsManagerParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["url"] = url
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// Adds a Ucs manager
func (s *UCSService) AddUcsManager(p *AddUcsManagerParams) (*AddUcsManagerResponse, error) {
	resp, err := s.cs.newPostRequest("addUcsManager", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddUcsManagerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddUcsManagerResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Zoneid    string `json:"zoneid"`
}

type AssociateUcsProfileToBladeParams struct {
	p map[string]interface{}
}

func (p *AssociateUcsProfileToBladeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bladeid"]; found {
		u.Set("bladeid", v.(string))
	}
	if v, found := p.p["profiledn"]; found {
		u.Set("profiledn", v.(string))
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *AssociateUcsProfileToBladeParams) SetBladeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bladeid"] = v
}

func (p *AssociateUcsProfileToBladeParams) ResetBladeid() {
	if p.p != nil && p.p["bladeid"] != nil {
		delete(p.p, "bladeid")
	}
}

func (p *AssociateUcsProfileToBladeParams) GetBladeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bladeid"].(string)
	return value, ok
}

func (p *AssociateUcsProfileToBladeParams) SetProfiledn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["profiledn"] = v
}

func (p *AssociateUcsProfileToBladeParams) ResetProfiledn() {
	if p.p != nil && p.p["profiledn"] != nil {
		delete(p.p, "profiledn")
	}
}

func (p *AssociateUcsProfileToBladeParams) GetProfiledn() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["profiledn"].(string)
	return value, ok
}

func (p *AssociateUcsProfileToBladeParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
}

func (p *AssociateUcsProfileToBladeParams) ResetUcsmanagerid() {
	if p.p != nil && p.p["ucsmanagerid"] != nil {
		delete(p.p, "ucsmanagerid")
	}
}

func (p *AssociateUcsProfileToBladeParams) GetUcsmanagerid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new AssociateUcsProfileToBladeParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewAssociateUcsProfileToBladeParams(bladeid string, profiledn string, ucsmanagerid string) *AssociateUcsProfileToBladeParams {
	p := &AssociateUcsProfileToBladeParams{}
	p.p = make(map[string]interface{})
	p.p["bladeid"] = bladeid
	p.p["profiledn"] = profiledn
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// associate a profile to a blade
func (s *UCSService) AssociateUcsProfileToBlade(p *AssociateUcsProfileToBladeParams) (*AssociateUcsProfileToBladeResponse, error) {
	resp, err := s.cs.newPostRequest("associateUcsProfileToBlade", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssociateUcsProfileToBladeResponse
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

type AssociateUcsProfileToBladeResponse struct {
	Bladedn      string `json:"bladedn"`
	Hostid       string `json:"hostid"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Profiledn    string `json:"profiledn"`
	Ucsmanagerid string `json:"ucsmanagerid"`
}

type DeleteUcsManagerParams struct {
	p map[string]interface{}
}

func (p *DeleteUcsManagerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *DeleteUcsManagerParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
}

func (p *DeleteUcsManagerParams) ResetUcsmanagerid() {
	if p.p != nil && p.p["ucsmanagerid"] != nil {
		delete(p.p, "ucsmanagerid")
	}
}

func (p *DeleteUcsManagerParams) GetUcsmanagerid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUcsManagerParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewDeleteUcsManagerParams(ucsmanagerid string) *DeleteUcsManagerParams {
	p := &DeleteUcsManagerParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// Delete a Ucs manager
func (s *UCSService) DeleteUcsManager(p *DeleteUcsManagerParams) (*DeleteUcsManagerResponse, error) {
	resp, err := s.cs.newPostRequest("deleteUcsManager", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUcsManagerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteUcsManagerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteUcsManagerResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteUcsManagerResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListUcsBladesParams struct {
	p map[string]interface{}
}

func (p *ListUcsBladesParams) toURLValues() url.Values {
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
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *ListUcsBladesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUcsBladesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUcsBladesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUcsBladesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUcsBladesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUcsBladesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUcsBladesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUcsBladesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUcsBladesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListUcsBladesParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
}

func (p *ListUcsBladesParams) ResetUcsmanagerid() {
	if p.p != nil && p.p["ucsmanagerid"] != nil {
		delete(p.p, "ucsmanagerid")
	}
}

func (p *ListUcsBladesParams) GetUcsmanagerid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUcsBladesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsBladesParams(ucsmanagerid string) *ListUcsBladesParams {
	p := &ListUcsBladesParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// List ucs blades
func (s *UCSService) ListUcsBlades(p *ListUcsBladesParams) (*ListUcsBladesResponse, error) {
	resp, err := s.cs.newRequest("listUcsBlades", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsBladesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUcsBladesResponse struct {
	Count     int         `json:"count"`
	UcsBlades []*UcsBlade `json:"ucsblade"`
}

type UcsBlade struct {
	Bladedn      string `json:"bladedn"`
	Hostid       string `json:"hostid"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Profiledn    string `json:"profiledn"`
	Ucsmanagerid string `json:"ucsmanagerid"`
}

type ListUcsManagersParams struct {
	p map[string]interface{}
}

func (p *ListUcsManagersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListUcsManagersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListUcsManagersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListUcsManagersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListUcsManagersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUcsManagersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUcsManagersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUcsManagersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUcsManagersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUcsManagersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUcsManagersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUcsManagersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUcsManagersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListUcsManagersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListUcsManagersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListUcsManagersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUcsManagersParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsManagersParams() *ListUcsManagersParams {
	p := &ListUcsManagersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListUcsManagersParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListUcsManagers(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.UcsManagers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.UcsManagers {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerByName(name string, opts ...OptionFunc) (*UcsManager, int, error) {
	id, count, err := s.GetUcsManagerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetUcsManagerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerByID(id string, opts ...OptionFunc) (*UcsManager, int, error) {
	p := &ListUcsManagersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUcsManagers(p)
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
		return l.UcsManagers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for UcsManager UUID: %s!", id)
}

// List ucs manager
func (s *UCSService) ListUcsManagers(p *ListUcsManagersParams) (*ListUcsManagersResponse, error) {
	resp, err := s.cs.newRequest("listUcsManagers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsManagersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUcsManagersResponse struct {
	Count       int           `json:"count"`
	UcsManagers []*UcsManager `json:"ucsmanager"`
}

type UcsManager struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Zoneid    string `json:"zoneid"`
}

type ListUcsProfilesParams struct {
	p map[string]interface{}
}

func (p *ListUcsProfilesParams) toURLValues() url.Values {
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
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *ListUcsProfilesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUcsProfilesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUcsProfilesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUcsProfilesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUcsProfilesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUcsProfilesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUcsProfilesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUcsProfilesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUcsProfilesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListUcsProfilesParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
}

func (p *ListUcsProfilesParams) ResetUcsmanagerid() {
	if p.p != nil && p.p["ucsmanagerid"] != nil {
		delete(p.p, "ucsmanagerid")
	}
}

func (p *ListUcsProfilesParams) GetUcsmanagerid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUcsProfilesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsProfilesParams(ucsmanagerid string) *ListUcsProfilesParams {
	p := &ListUcsProfilesParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// List profile in ucs manager
func (s *UCSService) ListUcsProfiles(p *ListUcsProfilesParams) (*ListUcsProfilesResponse, error) {
	resp, err := s.cs.newRequest("listUcsProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUcsProfilesResponse struct {
	Count       int           `json:"count"`
	UcsProfiles []*UcsProfile `json:"ucsprofile"`
}

type UcsProfile struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Ucsdn     string `json:"ucsdn"`
}
