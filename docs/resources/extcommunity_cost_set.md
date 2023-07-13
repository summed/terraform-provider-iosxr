---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_extcommunity_cost_set Resource - terraform-provider-iosxr"
subcategory: "Route Policy"
description: |-
  This resource can manage the Extcommunity Cost Set configuration.
---

# iosxr_extcommunity_cost_set (Resource)

This resource can manage the Extcommunity Cost Set configuration.

## Example Usage

```terraform
resource "iosxr_extcommunity_cost_set" "example" {
  set_name = "COST2"
  rpl      = "extcommunity-set cost COST2\nend-set\n"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `rpl` (String) Extended Community Cost Set
- `set_name` (String) Set name

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_extcommunity_cost_set.example "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/extended-community-cost-sets/extended-community-cost-set[set-name=COST2]"
```