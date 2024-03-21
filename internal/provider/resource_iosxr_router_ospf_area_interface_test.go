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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterOSPFAreaInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_non_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_point_to_point", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "network_point_to_multipoint", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "cost", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "passive_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "passive_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "fast_reroute_per_prefix_ti_lfa", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "fast_reroute_per_prefix_tiebreaker_srlg_disjoint", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_ospf_area_interface.test", "fast_reroute_per_prefix_tiebreaker_node_protecting", "33"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterOSPFAreaInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterOSPFAreaInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_ospf_area_interface.test",
		ImportState:   true,
		ImportStateId: "OSPF1,0,GigabitEthernet0/0/0/1",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrRouterOSPFAreaInterfaceConfig_minimum() string {
	config := `resource "iosxr_router_ospf_area_interface" "test" {` + "\n"
	config += `	process_name = "OSPF1"` + "\n"
	config += `	area_id = "0"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterOSPFAreaInterfaceConfig_all() string {
	config := `resource "iosxr_router_ospf_area_interface" "test" {` + "\n"
	config += `	process_name = "OSPF1"` + "\n"
	config += `	area_id = "0"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	network_broadcast = false` + "\n"
	config += `	network_non_broadcast = false` + "\n"
	config += `	network_point_to_point = true` + "\n"
	config += `	network_point_to_multipoint = false` + "\n"
	config += `	cost = 20` + "\n"
	config += `	priority = 100` + "\n"
	config += `	passive_enable = false` + "\n"
	config += `	passive_disable = true` + "\n"
	config += `	fast_reroute_per_prefix_ti_lfa = true` + "\n"
	config += `	fast_reroute_per_prefix_tiebreaker_srlg_disjoint = 22` + "\n"
	config += `	fast_reroute_per_prefix_tiebreaker_node_protecting = 33` + "\n"
	config += `}` + "\n"
	return config
}
