// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSNMPServerView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSNMPServerViewConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_snmp_server_view.test", "view_name", "VIEW12"),
					resource.TestCheckResourceAttr("iosxr_snmp_server_view.test", "mib_view_families.0.name", "iso"),
					resource.TestCheckResourceAttr("iosxr_snmp_server_view.test", "mib_view_families.0.included", "true"),
				),
			},
			{
				ResourceName:  "iosxr_snmp_server_view.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server/views/view[view-name=VIEW12]",
			},
		},
	})
}

func testAccIosxrSNMPServerViewConfig_minimum() string {
	return `
	resource "iosxr_snmp_server_view" "test" {
		view_name = "VIEW12"
	}
	`
}

func testAccIosxrSNMPServerViewConfig_all() string {
	return `
	resource "iosxr_snmp_server_view" "test" {
		view_name = "VIEW12"
		mib_view_families = [{
			name = "iso"
			included = true
		}]
	}
	`
}