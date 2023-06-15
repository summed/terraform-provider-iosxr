// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPNeighborGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPNeighborGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor_group.test", "remote_as", "65001"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor_group.test", "update_source", "Loopback0"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor_group.test", "address_families.0.af_name", "ipv4-labeled-unicast"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_neighbor_group.test", "address_families.0.soft_reconfiguration_inbound_always", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPNeighborGroupConfig = `

resource "iosxr_router_bgp_neighbor_group" "test" {
	as_number = "65001"
	name = "GROUP1"
	remote_as = "65001"
	update_source = "Loopback0"
	address_families = [{
		af_name = "ipv4-labeled-unicast"
		soft_reconfiguration_inbound_always = true
	}]
}

data "iosxr_router_bgp_neighbor_group" "test" {
	as_number = "65001"
	name = "GROUP1"
	depends_on = [iosxr_router_bgp_neighbor_group.test]
}
`