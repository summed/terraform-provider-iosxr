// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrCDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrCDPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_cdp.test", "enable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_cdp.test", "holdtime", "12"),
					resource.TestCheckResourceAttr("data.iosxr_cdp.test", "timer", "34"),
					resource.TestCheckResourceAttr("data.iosxr_cdp.test", "advertise_v1", "true"),
					resource.TestCheckResourceAttr("data.iosxr_cdp.test", "log_adjacency_changes", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrCDPConfig = `

resource "iosxr_cdp" "test" {
	enable = true
	holdtime = 12
	timer = 34
	advertise_v1 = true
	log_adjacency_changes = true
}

data "iosxr_cdp" "test" {
	depends_on = [iosxr_cdp.test]
}
`