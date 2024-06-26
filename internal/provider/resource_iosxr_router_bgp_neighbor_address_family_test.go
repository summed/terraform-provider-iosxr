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

func TestAccIosxrRouterBGPNeighborAddressFamily(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "af_name", "vpnv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "import_stitching_rt_re_originate_stitching_rt", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "route_reflector_client", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "route_reflector_client_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "advertise_vpnv4_unicast_enable_re_originated_stitching_rt", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "next_hop_self", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "next_hop_self_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "encapsulation_type_srv6", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "route_policy_in", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "route_policy_out", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "soft_reconfiguration_inbound_always", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "send_community_ebgp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "send_community_ebgp_inheritance_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "maximum_prefix_limit", "1248576"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "maximum_prefix_threshold", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "maximum_prefix_restart", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "maximum_prefix_discard_extra_paths", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "maximum_prefix_warning_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "default_originate_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_neighbor_address_family.test", "default_originate_inheritance_disable", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPNeighborAddressFamilyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPNeighborAddressFamilyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_bgp_neighbor_address_family.test",
		ImportState:   true,
		ImportStateId: "65001,10.1.1.2,vpnv4-unicast",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccIosxrRouterBGPNeighborAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
	lists = [
		{
			name = "address-families/address-family"
			key = "af-name"
			items = [
				{
					"af-name" = "vpnv4-unicast"
				},
			]
		},
		{
			name = "neighbors/neighbor"
			key = "neighbor-address"
			items = [
				{
					"neighbor-address" = "10.1.1.2"
					"remote-as" = "65001"
				},
			]
		},
	]
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
	attributes = {
		"route-policy-name" = "ROUTE_POLICY_1"
		"rpl-route-policy" = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
	}
}

`

func testAccIosxrRouterBGPNeighborAddressFamilyConfig_minimum() string {
	config := `resource "iosxr_router_bgp_neighbor_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	neighbor_address = "10.1.1.2"` + "\n"
	config += `	af_name = "vpnv4-unicast"` + "\n"
	config += `	maximum_prefix_limit = 1248576` + "\n"
	config += `	maximum_prefix_threshold = 80` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterBGPNeighborAddressFamilyConfig_all() string {
	config := `resource "iosxr_router_bgp_neighbor_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	neighbor_address = "10.1.1.2"` + "\n"
	config += `	af_name = "vpnv4-unicast"` + "\n"
	config += `	import_stitching_rt_re_originate_stitching_rt = true` + "\n"
	config += `	route_reflector_client = true` + "\n"
	config += `	route_reflector_client_inheritance_disable = true` + "\n"
	config += `	advertise_vpnv4_unicast_enable_re_originated_stitching_rt = true` + "\n"
	config += `	next_hop_self = true` + "\n"
	config += `	next_hop_self_inheritance_disable = true` + "\n"
	config += `	encapsulation_type_srv6 = true` + "\n"
	config += `	route_policy_in = "ROUTE_POLICY_1"` + "\n"
	config += `	route_policy_out = "ROUTE_POLICY_1"` + "\n"
	config += `	soft_reconfiguration_inbound_always = true` + "\n"
	config += `	send_community_ebgp = true` + "\n"
	config += `	send_community_ebgp_inheritance_disable = true` + "\n"
	config += `	maximum_prefix_limit = 1248576` + "\n"
	config += `	maximum_prefix_threshold = 80` + "\n"
	config += `	maximum_prefix_restart = 5` + "\n"
	config += `	maximum_prefix_discard_extra_paths = true` + "\n"
	config += `	maximum_prefix_warning_only = true` + "\n"
	config += `	default_originate_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	default_originate_inheritance_disable = true` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
