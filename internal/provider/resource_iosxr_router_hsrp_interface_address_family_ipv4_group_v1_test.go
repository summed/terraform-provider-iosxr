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

func TestAccIosxrRouterHSRPInterfaceAddressFamilyIPv4GroupV1(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "group_id", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "address", "22.22.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "address_learn", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "priority", "124"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "name", "NAME11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "preempt_delay", "3200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "bfd_fast_detect_peer_ipv4", "44.44.4.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "bfd_fast_detect_peer_interface", "GigabitEthernet0/0/0/7"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "track_interfaces.0.track_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "track_interfaces.0.priority_decrement", "166"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "track_objects.0.object_name", "OBJECT1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test", "track_objects.0.priority_decrement", "177"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterHSRPInterfaceAddressFamilyIPv4GroupV1Config_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterHSRPInterfaceAddressFamilyIPv4GroupV1Config_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_hsrp_interface_address_family_ipv4_group_v1.test",
		ImportState:   true,
		ImportStateId: "GigabitEthernet0/0/0/1,123",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrRouterHSRPInterfaceAddressFamilyIPv4GroupV1Config_minimum() string {
	config := `resource "iosxr_router_hsrp_interface_address_family_ipv4_group_v1" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	group_id = 123` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterHSRPInterfaceAddressFamilyIPv4GroupV1Config_all() string {
	config := `resource "iosxr_router_hsrp_interface_address_family_ipv4_group_v1" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	group_id = 123` + "\n"
	config += `	address = "22.22.1.1"` + "\n"
	config += `	address_learn = false` + "\n"
	config += `	priority = 124` + "\n"
	config += `	name = "NAME11"` + "\n"
	config += `	preempt_delay = 3200` + "\n"
	config += `	bfd_fast_detect_peer_ipv4 = "44.44.4.4"` + "\n"
	config += `	bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/7"` + "\n"
	config += `	track_interfaces = [{` + "\n"
	config += `		track_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		priority_decrement = 166` + "\n"
	config += `		}]` + "\n"
	config += `	track_objects = [{` + "\n"
	config += `		object_name = "OBJECT1"` + "\n"
	config += `		priority_decrement = 177` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
