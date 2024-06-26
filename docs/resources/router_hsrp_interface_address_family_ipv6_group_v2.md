---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_hsrp_interface_address_family_ipv6_group_v2 Resource - terraform-provider-iosxr"
subcategory: "HSRP"
description: |-
  This resource can manage the Router HSRP Interface Address Family IPv6 Group V2 configuration.
---

# iosxr_router_hsrp_interface_address_family_ipv6_group_v2 (Resource)

This resource can manage the Router HSRP Interface Address Family IPv6 Group V2 configuration.

## Example Usage

```terraform
resource "iosxr_router_hsrp_interface_address_family_ipv6_group_v2" "example" {
  interface_name                 = "GigabitEthernet0/0/0/2"
  group_id                       = 4055
  name                           = "gp2"
  mac_address                    = "00:01:00:02:00:02"
  timers_hold_time               = 10
  timers_hold_time2              = 20
  timers_msec                    = 100
  timers_msec2                   = 300
  preempt_delay                  = 256
  priority                       = 244
  bfd_fast_detect_peer_ipv6      = "fe80::240:d0ff:fe48:4672"
  bfd_fast_detect_peer_interface = "GigabitEthernet0/0/0/3"
  track_objects = [
    {
      object_name        = "TOBJ2"
      priority_decrement = 10
    }
  ]
  track_interfaces = [
    {
      track_name         = "GigabitEthernet0/0/0/4"
      priority_decrement = 244
    }
  ]
  addresses = [
    {
      address = "2001:db8:cafe:2100::bad1:1010"
    }
  ]
  address_link_local_autoconfig_legacy_compatible = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (Number) group number version 2
  - Range: `0`-`4095`
- `interface_name` (String) HSRP interface configuration subcommands

### Optional

- `address_link_local_autoconfig_legacy_compatible` (Boolean) Autoconfigure for Legacy compatibility (with IOS/NX-OS)
- `address_link_local_ipv6_address` (String) HSRP IPv6 linklocal address
- `addresses` (Attributes List) Global HSRP IPv6 address (see [below for nested schema](#nestedatt--addresses))
- `bfd_fast_detect_peer_interface` (String) Select an interface over which to run BFD
- `bfd_fast_detect_peer_ipv6` (String) BFD peer interface IPv6 address
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `mac_address` (String) Use specified mac address for the virtual router
- `name` (String) MGO session name
- `preempt_delay` (Number) Wait before preempting
  - Range: `0`-`3600`
- `priority` (Number) Priority level
  - Range: `0`-`255`
- `timers_hold_time` (Number) Hold time in seconds
  - Range: `1`-`255`
- `timers_hold_time2` (Number) Hold time in seconds
  - Range: `1`-`255`
- `timers_msec` (Number) Specify hellotime in milliseconds
  - Range: `100`-`3000`
- `timers_msec2` (Number) Specify hold time in milliseconds
  - Range: `100`-`3000`
- `track_interfaces` (Attributes List) Configure tracking (see [below for nested schema](#nestedatt--track_interfaces))
- `track_objects` (Attributes List) Object tracking (see [below for nested schema](#nestedatt--track_objects))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--addresses"></a>
### Nested Schema for `addresses`

Required:

- `address` (String) Set Global HSRP IPv6 address


<a id="nestedatt--track_interfaces"></a>
### Nested Schema for `track_interfaces`

Required:

- `priority_decrement` (Number) Priority decrement
  - Range: `1`-`255`
- `track_name` (String) Configure tracking


<a id="nestedatt--track_objects"></a>
### Nested Schema for `track_objects`

Required:

- `object_name` (String) Object tracking
- `priority_decrement` (Number) Priority decrement
  - Range: `1`-`255`

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_router_hsrp_interface_address_family_ipv6_group_v2.example "<interface_name>,<group_id>"
```
