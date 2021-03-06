// Copyright 2014 Vadim Kravcenko
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package gopencils

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type BasicAuth struct {
	Username string
	Password string
}

type ApiStruct struct {
	Base      string
	Methods   map[string]*Resource
	Options   map[string]bool
	BasicAuth *BasicAuth
	Client    *http.Client
	Cookies   *cookiejar.Jar
}

func Api(url string, options ...interface{}) *Resource {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	apiInstance := &ApiStruct{Base: url, Methods: make(map[string]*Resource), BasicAuth: nil}

	if len(options) > 0 {
		if auth, ok := options[0].(*BasicAuth); ok {
			apiInstance.BasicAuth = auth
		}
		if oauthClient, ok := options[0].(*http.Client); ok {
			apiInstance.Client = oauthClient
		}
	}
	if apiInstance.Client == nil {
		apiInstance.Cookies, _ = cookiejar.New(nil)

		// Skip verify by default?
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{
			Transport: tr,
			Jar:       apiInstance.Cookies,
		}
		apiInstance.Client = client
	}
	return &Resource{Url: "", Api: apiInstance}
}
