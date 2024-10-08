// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
)

// HTTPConfigApplyConfiguration represents a declarative configuration of the HTTPConfig type for use
// with apply.
type HTTPConfigApplyConfiguration struct {
	Authorization     *v1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	BasicAuth         *v1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	OAuth2            *v1.OAuth2ApplyConfiguration            `json:"oauth2,omitempty"`
	BearerTokenSecret *SecretKeySelectorApplyConfiguration    `json:"bearerTokenSecret,omitempty"`
	TLSConfig         *v1.SafeTLSConfigApplyConfiguration     `json:"tlsConfig,omitempty"`
	ProxyURL          *string                                 `json:"proxyURL,omitempty"`
	FollowRedirects   *bool                                   `json:"followRedirects,omitempty"`
}

// HTTPConfigApplyConfiguration constructs a declarative configuration of the HTTPConfig type for use with
// apply.
func HTTPConfig() *HTTPConfigApplyConfiguration {
	return &HTTPConfigApplyConfiguration{}
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithAuthorization(value *v1.SafeAuthorizationApplyConfiguration) *HTTPConfigApplyConfiguration {
	b.Authorization = value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithBasicAuth(value *v1.BasicAuthApplyConfiguration) *HTTPConfigApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithOAuth2 sets the OAuth2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth2 field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithOAuth2(value *v1.OAuth2ApplyConfiguration) *HTTPConfigApplyConfiguration {
	b.OAuth2 = value
	return b
}

// WithBearerTokenSecret sets the BearerTokenSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BearerTokenSecret field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithBearerTokenSecret(value *SecretKeySelectorApplyConfiguration) *HTTPConfigApplyConfiguration {
	b.BearerTokenSecret = value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithTLSConfig(value *v1.SafeTLSConfigApplyConfiguration) *HTTPConfigApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithProxyURL sets the ProxyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyURL field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithProxyURL(value string) *HTTPConfigApplyConfiguration {
	b.ProxyURL = &value
	return b
}

// WithFollowRedirects sets the FollowRedirects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FollowRedirects field is set to the value of the last call.
func (b *HTTPConfigApplyConfiguration) WithFollowRedirects(value bool) *HTTPConfigApplyConfiguration {
	b.FollowRedirects = &value
	return b
}
