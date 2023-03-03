// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0

package interfaces

import (
	pluginsdkmsg "github.com/dvonthenen/enterprise-reference-implementation/pkg/interfaces"
)

/*
	Higher Level Application Message
*/
type Data struct {
	Type string `json:"type,omitempty"`
	Data string `json:"data,omitempty"`
}

type AppSpecificScaffold struct {
	*pluginsdkmsg.AppSpecificType

	Data Data `json:"historical,omitempty"`
}
