// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package utils contains methods to support security, cache, and webhook functions.
package utils

import (
	"net/http"

	"github.com/vmware/harbor/src/common/utils/log"
)

// VerifySecret verifies the UI_SECRET cookie in a http request.
func VerifySecret(r *http.Request, expectedSecret string) bool {
	c, err := r.Cookie("secret")
	if err != nil {
		log.Warningf("Failed to get secret cookie, error: %v", err)
	}
	return c != nil && c.Value == expectedSecret
}
