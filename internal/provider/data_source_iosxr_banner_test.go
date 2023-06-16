// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrBanner(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrBannerConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_banner.test", "line", " Hello user !"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrBannerConfig = `

resource "iosxr_banner" "test" {
	banner_type = "login"
	line = " Hello user !"
}

data "iosxr_banner" "test" {
	banner_type = "login"
	depends_on = [iosxr_banner.test]
}
`
