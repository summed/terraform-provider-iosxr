---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_vrrp_interface_address_family_ipv4 Data Source - terraform-provider-iosxr"
subcategory: "VRRP"
description: |-
  This data source can read the Router VRRP Interface Address Family IPv4 configuration.
---

# iosxr_router_vrrp_interface_address_family_ipv4 (Data Source)

This data source can read the Router VRRP Interface Address Family IPv4 configuration.

## Example Usage

```terraform
data "iosxr_router_vrrp_interface_address_family_ipv4" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
  vrrp_id        = 123
  version        = 2
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_name` (String) VRRP interface configuration subcommands
- `version` (Number) VRRP version
- `vrrp_id` (Number) VRRP configuration

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `accept_mode_disable` (Boolean) Disable accept mode
- `address` (String) Enable VRRP and specify IP address(es)
- `bfd_fast_detect_peer_ipv4` (String) BFD peer interface IPv4 address
- `id` (String) The path of the retrieved object.
- `name` (String) Configure VRRP Session name
- `preempt_delay` (Number) Wait before preempting
- `preempt_disable` (Boolean) Disable preemption
- `priority` (Number) Set priority level
- `secondary_addresses` (Attributes List) VRRP IPv4 address (see [below for nested schema](#nestedatt--secondary_addresses))
- `text_authentication` (String) Set plain text authentication string
- `timer_advertisement_milliseconds` (Number) Configure in milliseconds
- `timer_advertisement_seconds` (Number) Advertisement time in seconds
- `timer_force` (Boolean) Force the configured values to be used
- `track_interfaces` (Attributes List) Track an interface (see [below for nested schema](#nestedatt--track_interfaces))
- `track_objects` (Attributes List) Object Tracking (see [below for nested schema](#nestedatt--track_objects))

<a id="nestedatt--secondary_addresses"></a>
### Nested Schema for `secondary_addresses`

Read-Only:

- `address` (String) VRRP IPv4 address


<a id="nestedatt--track_interfaces"></a>
### Nested Schema for `track_interfaces`

Read-Only:

- `interface_name` (String) Track an interface
- `priority_decrement` (Number) Priority decrement


<a id="nestedatt--track_objects"></a>
### Nested Schema for `track_objects`

Read-Only:

- `object_name` (String) Object to be tracked
- `priority_decrement` (Number) Priority decrement
