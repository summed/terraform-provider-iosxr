// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SNMPServerMIB struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	DeleteMode       types.String `tfsdk:"delete_mode"`
	IfmibIfaliasLong types.Bool   `tfsdk:"ifmib_ifalias_long"`
	IfindexPersist   types.Bool   `tfsdk:"ifindex_persist"`
}

type SNMPServerMIBData struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	IfmibIfaliasLong types.Bool   `tfsdk:"ifmib_ifalias_long"`
	IfindexPersist   types.Bool   `tfsdk:"ifindex_persist"`
}

func (data SNMPServerMIB) getPath() string {
	return "Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server-mibs"
}

func (data SNMPServerMIBData) getPath() string {
	return "Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server-mibs"
}

func (data SNMPServerMIB) toBody(ctx context.Context) string {
	body := "{}"
	if !data.IfmibIfaliasLong.IsNull() && !data.IfmibIfaliasLong.IsUnknown() {
		if data.IfmibIfaliasLong.ValueBool() {
			body, _ = sjson.Set(body, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib.ifalias.long", map[string]string{})
		}
	}
	if !data.IfindexPersist.IsNull() && !data.IfindexPersist.IsUnknown() {
		if data.IfindexPersist.ValueBool() {
			body, _ = sjson.Set(body, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex.persist", map[string]string{})
		}
	}
	return body
}

func (data *SNMPServerMIB) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib.ifalias.long"); !data.IfmibIfaliasLong.IsNull() {
		if value.Exists() {
			data.IfmibIfaliasLong = types.BoolValue(true)
		} else {
			data.IfmibIfaliasLong = types.BoolValue(false)
		}
	} else {
		data.IfmibIfaliasLong = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex.persist"); !data.IfindexPersist.IsNull() {
		if value.Exists() {
			data.IfindexPersist = types.BoolValue(true)
		} else {
			data.IfindexPersist = types.BoolValue(false)
		}
	} else {
		data.IfindexPersist = types.BoolNull()
	}
}

func (data *SNMPServerMIB) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib.ifalias.long"); value.Exists() {
		data.IfmibIfaliasLong = types.BoolValue(true)
	} else {
		data.IfmibIfaliasLong = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex.persist"); value.Exists() {
		data.IfindexPersist = types.BoolValue(true)
	} else {
		data.IfindexPersist = types.BoolValue(false)
	}
}

func (data *SNMPServerMIBData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib.ifalias.long"); value.Exists() {
		data.IfmibIfaliasLong = types.BoolValue(true)
	} else {
		data.IfmibIfaliasLong = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex.persist"); value.Exists() {
		data.IfindexPersist = types.BoolValue(true)
	} else {
		data.IfindexPersist = types.BoolValue(false)
	}
}

func (data *SNMPServerMIB) getDeletedItems(ctx context.Context, state SNMPServerMIB) []string {
	deletedItems := make([]string, 0)
	if !state.IfmibIfaliasLong.IsNull() && data.IfmibIfaliasLong.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib/ifalias/long", state.getPath()))
	}
	if !state.IfindexPersist.IsNull() && data.IfindexPersist.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex/persist", state.getPath()))
	}
	return deletedItems
}

func (data *SNMPServerMIB) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.IfmibIfaliasLong.IsNull() && !data.IfmibIfaliasLong.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib/ifalias/long", data.getPath()))
	}
	if !data.IfindexPersist.IsNull() && !data.IfindexPersist.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex/persist", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *SNMPServerMIB) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.IfmibIfaliasLong.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XR-um-mibs-ifmib-cfg:ifmib/ifalias/long", data.getPath()))
	}
	if !data.IfindexPersist.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XR-um-mibs-ifmib-cfg:ifindex/persist", data.getPath()))
	}
	return deletePaths
}
