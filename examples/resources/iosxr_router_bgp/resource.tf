resource "iosxr_router_bgp" "example" {
  as_number                                  = "65001"
  default_information_originate              = true
  default_metric                             = 125
  nsr_disable                                = false
  bgp_redistribute_internal                  = true
  segment_routing_srv6_locator               = "locator11"
  timers_bgp_keepalive_interval              = 5
  timers_bgp_holdtime                        = "20"
  timers_bgp_minimum_acceptable_holdtime     = "10"
  bgp_router_id                              = "22.22.22.22"
  bgp_graceful_restart_graceful_reset        = true
  ibgp_policy_out_enforce_modifications      = true
  bgp_log_neighbor_changes_detail            = true
  bfd_minimum_interval                       = 10
  bfd_multiplier                             = 4
  nexthop_validation_color_extcomm_sr_policy = true
  nexthop_validation_color_extcomm_disable   = true
  bgp_bestpath_as_path_ignore                = true
  bgp_bestpath_as_path_multipath_relax       = true
  bgp_bestpath_cost_community_ignore         = true
  bgp_bestpath_compare_routerid              = true
  bgp_bestpath_aigp_ignore                   = true
  bgp_bestpath_igp_metric_ignore             = true
  bgp_bestpath_igp_metric_sr_policy          = true
  bgp_bestpath_med_always                    = true
  bgp_bestpath_med_confed                    = true
  bgp_bestpath_med_missing_as_worst          = true
  bgp_bestpath_origin_as_use_validity        = true
  bgp_bestpath_origin_as_allow_invalid       = true
  bgp_bestpath_sr_policy_prefer              = false
  bgp_bestpath_sr_policy_force               = true
  neighbors = [
    {
      neighbor_address                   = "10.1.1.2"
      remote_as                          = "65002"
      description                        = "My Neighbor Description"
      use_neighbor_group                 = "GROUP1"
      ignore_connected_check             = true
      ebgp_multihop_maximum_hop_count    = 10
      bfd_minimum_interval               = 10
      bfd_multiplier                     = 4
      local_as                           = "65003"
      local_as_no_prepend                = true
      local_as_replace_as                = true
      local_as_dual_as                   = true
      password                           = "12341C2713181F13253920"
      shutdown                           = false
      timers_keepalive_interval          = 5
      timers_holdtime                    = "20"
      timers_minimum_acceptable_holdtime = "10"
      update_source                      = "GigabitEthernet0/0/0/1"
      ttl_security                       = false
    }
  ]
  neighbor_groups = [
    {
      name                 = "GROUP1"
      remote_as            = "65001"
      update_source        = "Loopback0"
      bfd_minimum_interval = 3
    }
  ]
}
