//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azfile

import "net/http"

type ResponseError interface {
	Error() string
	Unwrap() error
	RawResponse() *http.Response
	NonRetriable()
}
