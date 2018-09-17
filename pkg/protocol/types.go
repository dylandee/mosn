/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package protocol

import (
	"github.com/alipay/sofa-mosn/pkg/types"
)

// Protocol type definition
const (
	SofaRPC   types.Protocol = "SofaRpc"
	HTTP1     types.Protocol = "Http1"
	HTTP2     types.Protocol = "Http2"
	Xprotocol types.Protocol = "X"
)

// Host key for routing in MOSN Header
const (
	MosnHeaderHostKey        = "host"
	MosnHeaderPathKey        = "path"
	MosnHeaderQueryStringKey = "querystring"
	MosnHeaderMethod         = "method"
)

// wrapper for map[string]string
type CommonHeader map[string]string

func (h CommonHeader) Get(key string) (value string, ok bool) {
	value, ok = h[key]
	return
}

func (h CommonHeader) Set(key string, value string) {
	h[key] = value
}

func (h CommonHeader) Del(key string) {
	delete(h, key)
}

func (h CommonHeader) Range(f func(key, value string) bool) {
	for k, v := range h {
		// stop if f return false
		if !f(k, v) {
			break
		}
	}
}
