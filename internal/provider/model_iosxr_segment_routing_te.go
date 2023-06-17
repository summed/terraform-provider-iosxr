// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SegmentRoutingTE struct {
	Device                types.String                     `tfsdk:"device"`
	Id                    types.String                     `tfsdk:"id"`
	LoggingPcepPeerStatus types.Bool                       `tfsdk:"logging_pcep_peer_status"`
	LoggingPolicyStatus   types.Bool                       `tfsdk:"logging_policy_status"`
	PccReportAll          types.Bool                       `tfsdk:"pcc_report_all"`
	PccSourceAddress      types.String                     `tfsdk:"pcc_source_address"`
	PccDelegationTimeout  types.Int64                      `tfsdk:"pcc_delegation_timeout"`
	PccDeadTimer          types.Int64                      `tfsdk:"pcc_dead_timer"`
	PccInitiatedState     types.Int64                      `tfsdk:"pcc_initiated_state"`
	PccInitiatedOrphan    types.Int64                      `tfsdk:"pcc_initiated_orphan"`
	PcePeers              []SegmentRoutingTEPcePeers       `tfsdk:"pce_peers"`
	OnDemandColors        []SegmentRoutingTEOnDemandColors `tfsdk:"on_demand_colors"`
	Policies              []SegmentRoutingTEPolicies       `tfsdk:"policies"`
}
type SegmentRoutingTEPcePeers struct {
	PceAddress types.String `tfsdk:"pce_address"`
	Precedence types.Int64  `tfsdk:"precedence"`
}
type SegmentRoutingTEOnDemandColors struct {
	DynamicAnycastSidInclusion       types.Bool   `tfsdk:"dynamic_anycast_sid_inclusion"`
	DynamicMetricType                types.String `tfsdk:"dynamic_metric_type"`
	Color                            types.Int64  `tfsdk:"color"`
	Srv6Enable                       types.Bool   `tfsdk:"srv6_enable"`
	Srv6LocatorName                  types.String `tfsdk:"srv6_locator_name"`
	Srv6LocatorBehavior              types.String `tfsdk:"srv6_locator_behavior"`
	Srv6LocatorBindingSidType        types.String `tfsdk:"srv6_locator_binding_sid_type"`
	SourceAddress                    types.String `tfsdk:"source_address"`
	SourceAddressType                types.String `tfsdk:"source_address_type"`
	EffectiveMetricEnable            types.Bool   `tfsdk:"effective_metric_enable"`
	EffectiveMetricValue             types.Int64  `tfsdk:"effective_metric_value"`
	EffectiveMetricType              types.String `tfsdk:"effective_metric_type"`
	ConstraintSegmentsProtectionType types.String `tfsdk:"constraint_segments_protection_type"`
	ConstraintSegmentsSidAlgorithm   types.Int64  `tfsdk:"constraint_segments_sid_algorithm"`
}
type SegmentRoutingTEPolicies struct {
	PolicyName                 types.String `tfsdk:"policy_name"`
	Srv6Enable                 types.Bool   `tfsdk:"srv6_enable"`
	Srv6LocatorName            types.String `tfsdk:"srv6_locator_name"`
	Srv6LocatorBindingSidType  types.String `tfsdk:"srv6_locator_binding_sid_type"`
	Srv6LocatorBehavior        types.String `tfsdk:"srv6_locator_behavior"`
	SourceAddress              types.String `tfsdk:"source_address"`
	SourceAddressType          types.String `tfsdk:"source_address_type"`
	PolicyColorEndpointColor   types.Int64  `tfsdk:"policy_color_endpoint_color"`
	PolicyColorEndpointType    types.String `tfsdk:"policy_color_endpoint_type"`
	PolicyColorEndpointAddress types.String `tfsdk:"policy_color_endpoint_address"`
}

func (data SegmentRoutingTE) getPath() string {
	return "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering"
}

func (data SegmentRoutingTE) toBody(ctx context.Context) string {
	body := "{}"
	if !data.LoggingPcepPeerStatus.IsNull() && !data.LoggingPcepPeerStatus.IsUnknown() {
		if data.LoggingPcepPeerStatus.ValueBool() {
			body, _ = sjson.Set(body, "logging.pcep-peer-status", map[string]string{})
		}
	}
	if !data.LoggingPolicyStatus.IsNull() && !data.LoggingPolicyStatus.IsUnknown() {
		if data.LoggingPolicyStatus.ValueBool() {
			body, _ = sjson.Set(body, "logging.policy-status", map[string]string{})
		}
	}
	if !data.PccReportAll.IsNull() && !data.PccReportAll.IsUnknown() {
		if data.PccReportAll.ValueBool() {
			body, _ = sjson.Set(body, "pcc.report-all", map[string]string{})
		}
	}
	if !data.PccSourceAddress.IsNull() && !data.PccSourceAddress.IsUnknown() {
		body, _ = sjson.Set(body, "pcc.source-address", data.PccSourceAddress.ValueString())
	}
	if !data.PccDelegationTimeout.IsNull() && !data.PccDelegationTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "pcc.delegation-timeout", strconv.FormatInt(data.PccDelegationTimeout.ValueInt64(), 10))
	}
	if !data.PccDeadTimer.IsNull() && !data.PccDeadTimer.IsUnknown() {
		body, _ = sjson.Set(body, "pcc.dead-timer-interval", strconv.FormatInt(data.PccDeadTimer.ValueInt64(), 10))
	}
	if !data.PccInitiatedState.IsNull() && !data.PccInitiatedState.IsUnknown() {
		body, _ = sjson.Set(body, "pcc.initiated-state-interval", strconv.FormatInt(data.PccInitiatedState.ValueInt64(), 10))
	}
	if !data.PccInitiatedOrphan.IsNull() && !data.PccInitiatedOrphan.IsUnknown() {
		body, _ = sjson.Set(body, "pcc.initiated-orphan-interval", strconv.FormatInt(data.PccInitiatedOrphan.ValueInt64(), 10))
	}
	if len(data.PcePeers) > 0 {
		body, _ = sjson.Set(body, "pcc.pce-peers.pce-peer", []interface{}{})
		for index, item := range data.PcePeers {
			if !item.PceAddress.IsNull() && !item.PceAddress.IsUnknown() {
				body, _ = sjson.Set(body, "pcc.pce-peers.pce-peer"+"."+strconv.Itoa(index)+"."+"pce-address", item.PceAddress.ValueString())
			}
			if !item.Precedence.IsNull() && !item.Precedence.IsUnknown() {
				body, _ = sjson.Set(body, "pcc.pce-peers.pce-peer"+"."+strconv.Itoa(index)+"."+"precedence", strconv.FormatInt(item.Precedence.ValueInt64(), 10))
			}
		}
	}
	if len(data.OnDemandColors) > 0 {
		body, _ = sjson.Set(body, "on-demand-colors.on-demand-color", []interface{}{})
		for index, item := range data.OnDemandColors {
			if !item.DynamicAnycastSidInclusion.IsNull() && !item.DynamicAnycastSidInclusion.IsUnknown() {
				if item.DynamicAnycastSidInclusion.ValueBool() {
					body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"on-demand-color-dyn-mpls.on-demand-color-dyn-mpls-anycast", map[string]string{})
				}
			}
			if !item.DynamicMetricType.IsNull() && !item.DynamicMetricType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"on-demand-color-dyn-mpls.on-demand-color-dyn-mpls-metric.metric-type", item.DynamicMetricType.ValueString())
			}
			if !item.Color.IsNull() && !item.Color.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"color", strconv.FormatInt(item.Color.ValueInt64(), 10))
			}
			if !item.Srv6Enable.IsNull() && !item.Srv6Enable.IsUnknown() {
				if item.Srv6Enable.ValueBool() {
					body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.enable", map[string]string{})
				}
			}
			if !item.Srv6LocatorName.IsNull() && !item.Srv6LocatorName.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.locator.locator-name", item.Srv6LocatorName.ValueString())
			}
			if !item.Srv6LocatorBehavior.IsNull() && !item.Srv6LocatorBehavior.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.locator.behavior", item.Srv6LocatorBehavior.ValueString())
			}
			if !item.Srv6LocatorBindingSidType.IsNull() && !item.Srv6LocatorBindingSidType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"srv6.locator.binding-sid-type", item.Srv6LocatorBindingSidType.ValueString())
			}
			if !item.SourceAddress.IsNull() && !item.SourceAddress.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"source-address.source-address", item.SourceAddress.ValueString())
			}
			if !item.SourceAddressType.IsNull() && !item.SourceAddressType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"source-address.ip-address-type", item.SourceAddressType.ValueString())
			}
			if !item.EffectiveMetricEnable.IsNull() && !item.EffectiveMetricEnable.IsUnknown() {
				if item.EffectiveMetricEnable.ValueBool() {
					body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"effective-metric.enable", map[string]string{})
				}
			}
			if !item.EffectiveMetricValue.IsNull() && !item.EffectiveMetricValue.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"effective-metric.metric-value-type.metric-value", strconv.FormatInt(item.EffectiveMetricValue.ValueInt64(), 10))
			}
			if !item.EffectiveMetricType.IsNull() && !item.EffectiveMetricType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"effective-metric.metric-value-type.metric-type", item.EffectiveMetricType.ValueString())
			}
			if !item.ConstraintSegmentsProtectionType.IsNull() && !item.ConstraintSegmentsProtectionType.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"constraint.segments.protection-type", item.ConstraintSegmentsProtectionType.ValueString())
			}
			if !item.ConstraintSegmentsSidAlgorithm.IsNull() && !item.ConstraintSegmentsSidAlgorithm.IsUnknown() {
				body, _ = sjson.Set(body, "on-demand-colors.on-demand-color"+"."+strconv.Itoa(index)+"."+"constraint.segments.sid-algorithm", strconv.FormatInt(item.ConstraintSegmentsSidAlgorithm.ValueInt64(), 10))
			}
		}
	}
	if len(data.Policies) > 0 {
		body, _ = sjson.Set(body, "policies.policy", []interface{}{})
		for index, item := range data.Policies {
			if !item.PolicyName.IsNull() && !item.PolicyName.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-name", item.PolicyName.ValueString())
			}
			if !item.Srv6Enable.IsNull() && !item.Srv6Enable.IsUnknown() {
				if item.Srv6Enable.ValueBool() {
					body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.enable", map[string]string{})
				}
			}
			if !item.Srv6LocatorName.IsNull() && !item.Srv6LocatorName.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.locator.locator-name", item.Srv6LocatorName.ValueString())
			}
			if !item.Srv6LocatorBindingSidType.IsNull() && !item.Srv6LocatorBindingSidType.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.locator.binding-sid-type", item.Srv6LocatorBindingSidType.ValueString())
			}
			if !item.Srv6LocatorBehavior.IsNull() && !item.Srv6LocatorBehavior.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"srv6.locator.behavior", item.Srv6LocatorBehavior.ValueString())
			}
			if !item.SourceAddress.IsNull() && !item.SourceAddress.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"source-address.source-address", item.SourceAddress.ValueString())
			}
			if !item.SourceAddressType.IsNull() && !item.SourceAddressType.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"source-address.ip-address-type", item.SourceAddressType.ValueString())
			}
			if !item.PolicyColorEndpointColor.IsNull() && !item.PolicyColorEndpointColor.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-color-endpoint.color", strconv.FormatInt(item.PolicyColorEndpointColor.ValueInt64(), 10))
			}
			if !item.PolicyColorEndpointType.IsNull() && !item.PolicyColorEndpointType.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-color-endpoint.end-point-type", item.PolicyColorEndpointType.ValueString())
			}
			if !item.PolicyColorEndpointAddress.IsNull() && !item.PolicyColorEndpointAddress.IsUnknown() {
				body, _ = sjson.Set(body, "policies.policy"+"."+strconv.Itoa(index)+"."+"policy-color-endpoint.end-point-address", item.PolicyColorEndpointAddress.ValueString())
			}
		}
	}
	return body
}

func (data *SegmentRoutingTE) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "logging.pcep-peer-status"); !data.LoggingPcepPeerStatus.IsNull() {
		if value.Exists() {
			data.LoggingPcepPeerStatus = types.BoolValue(true)
		} else {
			data.LoggingPcepPeerStatus = types.BoolValue(false)
		}
	} else {
		data.LoggingPcepPeerStatus = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "logging.policy-status"); !data.LoggingPolicyStatus.IsNull() {
		if value.Exists() {
			data.LoggingPolicyStatus = types.BoolValue(true)
		} else {
			data.LoggingPolicyStatus = types.BoolValue(false)
		}
	} else {
		data.LoggingPolicyStatus = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "pcc.report-all"); !data.PccReportAll.IsNull() {
		if value.Exists() {
			data.PccReportAll = types.BoolValue(true)
		} else {
			data.PccReportAll = types.BoolValue(false)
		}
	} else {
		data.PccReportAll = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "pcc.source-address"); value.Exists() && !data.PccSourceAddress.IsNull() {
		data.PccSourceAddress = types.StringValue(value.String())
	} else {
		data.PccSourceAddress = types.StringNull()
	}
	if value := gjson.GetBytes(res, "pcc.delegation-timeout"); value.Exists() && !data.PccDelegationTimeout.IsNull() {
		data.PccDelegationTimeout = types.Int64Value(value.Int())
	} else {
		data.PccDelegationTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "pcc.dead-timer-interval"); value.Exists() && !data.PccDeadTimer.IsNull() {
		data.PccDeadTimer = types.Int64Value(value.Int())
	} else {
		data.PccDeadTimer = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "pcc.initiated-state-interval"); value.Exists() && !data.PccInitiatedState.IsNull() {
		data.PccInitiatedState = types.Int64Value(value.Int())
	} else {
		data.PccInitiatedState = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "pcc.initiated-orphan-interval"); value.Exists() && !data.PccInitiatedOrphan.IsNull() {
		data.PccInitiatedOrphan = types.Int64Value(value.Int())
	} else {
		data.PccInitiatedOrphan = types.Int64Null()
	}
	for i := range data.PcePeers {
		keys := [...]string{"pce-address"}
		keyValues := [...]string{data.PcePeers[i].PceAddress.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "pcc.pce-peers.pce-peer").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("pce-address"); value.Exists() && !data.PcePeers[i].PceAddress.IsNull() {
			data.PcePeers[i].PceAddress = types.StringValue(value.String())
		} else {
			data.PcePeers[i].PceAddress = types.StringNull()
		}
		if value := r.Get("precedence"); value.Exists() && !data.PcePeers[i].Precedence.IsNull() {
			data.PcePeers[i].Precedence = types.Int64Value(value.Int())
		} else {
			data.PcePeers[i].Precedence = types.Int64Null()
		}
	}
	for i := range data.OnDemandColors {
		keys := [...]string{"color"}
		keyValues := [...]string{strconv.FormatInt(data.OnDemandColors[i].Color.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "on-demand-colors.on-demand-color").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("on-demand-color-dyn-mpls.on-demand-color-dyn-mpls-anycast"); !data.OnDemandColors[i].DynamicAnycastSidInclusion.IsNull() {
			if value.Exists() {
				data.OnDemandColors[i].DynamicAnycastSidInclusion = types.BoolValue(true)
			} else {
				data.OnDemandColors[i].DynamicAnycastSidInclusion = types.BoolValue(false)
			}
		} else {
			data.OnDemandColors[i].DynamicAnycastSidInclusion = types.BoolNull()
		}
		if value := r.Get("on-demand-color-dyn-mpls.on-demand-color-dyn-mpls-metric.metric-type"); value.Exists() && !data.OnDemandColors[i].DynamicMetricType.IsNull() {
			data.OnDemandColors[i].DynamicMetricType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].DynamicMetricType = types.StringNull()
		}
		if value := r.Get("color"); value.Exists() && !data.OnDemandColors[i].Color.IsNull() {
			data.OnDemandColors[i].Color = types.Int64Value(value.Int())
		} else {
			data.OnDemandColors[i].Color = types.Int64Null()
		}
		if value := r.Get("srv6.enable"); !data.OnDemandColors[i].Srv6Enable.IsNull() {
			if value.Exists() {
				data.OnDemandColors[i].Srv6Enable = types.BoolValue(true)
			} else {
				data.OnDemandColors[i].Srv6Enable = types.BoolValue(false)
			}
		} else {
			data.OnDemandColors[i].Srv6Enable = types.BoolNull()
		}
		if value := r.Get("srv6.locator.locator-name"); value.Exists() && !data.OnDemandColors[i].Srv6LocatorName.IsNull() {
			data.OnDemandColors[i].Srv6LocatorName = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].Srv6LocatorName = types.StringNull()
		}
		if value := r.Get("srv6.locator.behavior"); value.Exists() && !data.OnDemandColors[i].Srv6LocatorBehavior.IsNull() {
			data.OnDemandColors[i].Srv6LocatorBehavior = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].Srv6LocatorBehavior = types.StringNull()
		}
		if value := r.Get("srv6.locator.binding-sid-type"); value.Exists() && !data.OnDemandColors[i].Srv6LocatorBindingSidType.IsNull() {
			data.OnDemandColors[i].Srv6LocatorBindingSidType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].Srv6LocatorBindingSidType = types.StringNull()
		}
		if value := r.Get("source-address.source-address"); value.Exists() && !data.OnDemandColors[i].SourceAddress.IsNull() {
			data.OnDemandColors[i].SourceAddress = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].SourceAddress = types.StringNull()
		}
		if value := r.Get("source-address.ip-address-type"); value.Exists() && !data.OnDemandColors[i].SourceAddressType.IsNull() {
			data.OnDemandColors[i].SourceAddressType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].SourceAddressType = types.StringNull()
		}
		if value := r.Get("effective-metric.enable"); !data.OnDemandColors[i].EffectiveMetricEnable.IsNull() {
			if value.Exists() {
				data.OnDemandColors[i].EffectiveMetricEnable = types.BoolValue(true)
			} else {
				data.OnDemandColors[i].EffectiveMetricEnable = types.BoolValue(false)
			}
		} else {
			data.OnDemandColors[i].EffectiveMetricEnable = types.BoolNull()
		}
		if value := r.Get("effective-metric.metric-value-type.metric-value"); value.Exists() && !data.OnDemandColors[i].EffectiveMetricValue.IsNull() {
			data.OnDemandColors[i].EffectiveMetricValue = types.Int64Value(value.Int())
		} else {
			data.OnDemandColors[i].EffectiveMetricValue = types.Int64Null()
		}
		if value := r.Get("effective-metric.metric-value-type.metric-type"); value.Exists() && !data.OnDemandColors[i].EffectiveMetricType.IsNull() {
			data.OnDemandColors[i].EffectiveMetricType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].EffectiveMetricType = types.StringNull()
		}
		if value := r.Get("constraint.segments.protection-type"); value.Exists() && !data.OnDemandColors[i].ConstraintSegmentsProtectionType.IsNull() {
			data.OnDemandColors[i].ConstraintSegmentsProtectionType = types.StringValue(value.String())
		} else {
			data.OnDemandColors[i].ConstraintSegmentsProtectionType = types.StringNull()
		}
		if value := r.Get("constraint.segments.sid-algorithm"); value.Exists() && !data.OnDemandColors[i].ConstraintSegmentsSidAlgorithm.IsNull() {
			data.OnDemandColors[i].ConstraintSegmentsSidAlgorithm = types.Int64Value(value.Int())
		} else {
			data.OnDemandColors[i].ConstraintSegmentsSidAlgorithm = types.Int64Null()
		}
	}
	for i := range data.Policies {
		keys := [...]string{"policy-name"}
		keyValues := [...]string{data.Policies[i].PolicyName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "policies.policy").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("policy-name"); value.Exists() && !data.Policies[i].PolicyName.IsNull() {
			data.Policies[i].PolicyName = types.StringValue(value.String())
		} else {
			data.Policies[i].PolicyName = types.StringNull()
		}
		if value := r.Get("srv6.enable"); !data.Policies[i].Srv6Enable.IsNull() {
			if value.Exists() {
				data.Policies[i].Srv6Enable = types.BoolValue(true)
			} else {
				data.Policies[i].Srv6Enable = types.BoolValue(false)
			}
		} else {
			data.Policies[i].Srv6Enable = types.BoolNull()
		}
		if value := r.Get("srv6.locator.locator-name"); value.Exists() && !data.Policies[i].Srv6LocatorName.IsNull() {
			data.Policies[i].Srv6LocatorName = types.StringValue(value.String())
		} else {
			data.Policies[i].Srv6LocatorName = types.StringNull()
		}
		if value := r.Get("srv6.locator.binding-sid-type"); value.Exists() && !data.Policies[i].Srv6LocatorBindingSidType.IsNull() {
			data.Policies[i].Srv6LocatorBindingSidType = types.StringValue(value.String())
		} else {
			data.Policies[i].Srv6LocatorBindingSidType = types.StringNull()
		}
		if value := r.Get("srv6.locator.behavior"); value.Exists() && !data.Policies[i].Srv6LocatorBehavior.IsNull() {
			data.Policies[i].Srv6LocatorBehavior = types.StringValue(value.String())
		} else {
			data.Policies[i].Srv6LocatorBehavior = types.StringNull()
		}
		if value := r.Get("source-address.source-address"); value.Exists() && !data.Policies[i].SourceAddress.IsNull() {
			data.Policies[i].SourceAddress = types.StringValue(value.String())
		} else {
			data.Policies[i].SourceAddress = types.StringNull()
		}
		if value := r.Get("source-address.ip-address-type"); value.Exists() && !data.Policies[i].SourceAddressType.IsNull() {
			data.Policies[i].SourceAddressType = types.StringValue(value.String())
		} else {
			data.Policies[i].SourceAddressType = types.StringNull()
		}
		if value := r.Get("policy-color-endpoint.color"); value.Exists() && !data.Policies[i].PolicyColorEndpointColor.IsNull() {
			data.Policies[i].PolicyColorEndpointColor = types.Int64Value(value.Int())
		} else {
			data.Policies[i].PolicyColorEndpointColor = types.Int64Null()
		}
		if value := r.Get("policy-color-endpoint.end-point-type"); value.Exists() && !data.Policies[i].PolicyColorEndpointType.IsNull() {
			data.Policies[i].PolicyColorEndpointType = types.StringValue(value.String())
		} else {
			data.Policies[i].PolicyColorEndpointType = types.StringNull()
		}
		if value := r.Get("policy-color-endpoint.end-point-address"); value.Exists() && !data.Policies[i].PolicyColorEndpointAddress.IsNull() {
			data.Policies[i].PolicyColorEndpointAddress = types.StringValue(value.String())
		} else {
			data.Policies[i].PolicyColorEndpointAddress = types.StringNull()
		}
	}
}

func (data *SegmentRoutingTE) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "logging.pcep-peer-status"); value.Exists() {
		data.LoggingPcepPeerStatus = types.BoolValue(true)
	} else {
		data.LoggingPcepPeerStatus = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "logging.policy-status"); value.Exists() {
		data.LoggingPolicyStatus = types.BoolValue(true)
	} else {
		data.LoggingPolicyStatus = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "pcc.report-all"); value.Exists() {
		data.PccReportAll = types.BoolValue(true)
	} else {
		data.PccReportAll = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "pcc.source-address"); value.Exists() {
		data.PccSourceAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "pcc.delegation-timeout"); value.Exists() {
		data.PccDelegationTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "pcc.dead-timer-interval"); value.Exists() {
		data.PccDeadTimer = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "pcc.initiated-state-interval"); value.Exists() {
		data.PccInitiatedState = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "pcc.initiated-orphan-interval"); value.Exists() {
		data.PccInitiatedOrphan = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "pcc.pce-peers.pce-peer"); value.Exists() {
		data.PcePeers = make([]SegmentRoutingTEPcePeers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingTEPcePeers{}
			if cValue := v.Get("pce-address"); cValue.Exists() {
				item.PceAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("precedence"); cValue.Exists() {
				item.Precedence = types.Int64Value(cValue.Int())
			}
			data.PcePeers = append(data.PcePeers, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "on-demand-colors.on-demand-color"); value.Exists() {
		data.OnDemandColors = make([]SegmentRoutingTEOnDemandColors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingTEOnDemandColors{}
			if cValue := v.Get("on-demand-color-dyn-mpls.on-demand-color-dyn-mpls-anycast"); cValue.Exists() {
				item.DynamicAnycastSidInclusion = types.BoolValue(true)
			} else {
				item.DynamicAnycastSidInclusion = types.BoolValue(false)
			}
			if cValue := v.Get("on-demand-color-dyn-mpls.on-demand-color-dyn-mpls-metric.metric-type"); cValue.Exists() {
				item.DynamicMetricType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("color"); cValue.Exists() {
				item.Color = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("srv6.enable"); cValue.Exists() {
				item.Srv6Enable = types.BoolValue(true)
			} else {
				item.Srv6Enable = types.BoolValue(false)
			}
			if cValue := v.Get("srv6.locator.locator-name"); cValue.Exists() {
				item.Srv6LocatorName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.behavior"); cValue.Exists() {
				item.Srv6LocatorBehavior = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.binding-sid-type"); cValue.Exists() {
				item.Srv6LocatorBindingSidType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.source-address"); cValue.Exists() {
				item.SourceAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.ip-address-type"); cValue.Exists() {
				item.SourceAddressType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("effective-metric.enable"); cValue.Exists() {
				item.EffectiveMetricEnable = types.BoolValue(true)
			} else {
				item.EffectiveMetricEnable = types.BoolValue(false)
			}
			if cValue := v.Get("effective-metric.metric-value-type.metric-value"); cValue.Exists() {
				item.EffectiveMetricValue = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("effective-metric.metric-value-type.metric-type"); cValue.Exists() {
				item.EffectiveMetricType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("constraint.segments.protection-type"); cValue.Exists() {
				item.ConstraintSegmentsProtectionType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("constraint.segments.sid-algorithm"); cValue.Exists() {
				item.ConstraintSegmentsSidAlgorithm = types.Int64Value(cValue.Int())
			}
			data.OnDemandColors = append(data.OnDemandColors, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "policies.policy"); value.Exists() {
		data.Policies = make([]SegmentRoutingTEPolicies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SegmentRoutingTEPolicies{}
			if cValue := v.Get("policy-name"); cValue.Exists() {
				item.PolicyName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.enable"); cValue.Exists() {
				item.Srv6Enable = types.BoolValue(true)
			} else {
				item.Srv6Enable = types.BoolValue(false)
			}
			if cValue := v.Get("srv6.locator.locator-name"); cValue.Exists() {
				item.Srv6LocatorName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.binding-sid-type"); cValue.Exists() {
				item.Srv6LocatorBindingSidType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("srv6.locator.behavior"); cValue.Exists() {
				item.Srv6LocatorBehavior = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.source-address"); cValue.Exists() {
				item.SourceAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("source-address.ip-address-type"); cValue.Exists() {
				item.SourceAddressType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("policy-color-endpoint.color"); cValue.Exists() {
				item.PolicyColorEndpointColor = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("policy-color-endpoint.end-point-type"); cValue.Exists() {
				item.PolicyColorEndpointType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("policy-color-endpoint.end-point-address"); cValue.Exists() {
				item.PolicyColorEndpointAddress = types.StringValue(cValue.String())
			}
			data.Policies = append(data.Policies, item)
			return true
		})
	}
}

func (data *SegmentRoutingTE) getDeletedListItems(ctx context.Context, state SegmentRoutingTE) []string {
	deletedListItems := make([]string, 0)
	for i := range state.PcePeers {
		keys := [...]string{"pce-address"}
		stateKeyValues := [...]string{state.PcePeers[i].PceAddress.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.PcePeers[i].PceAddress.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.PcePeers {
			found = true
			if state.PcePeers[i].PceAddress.ValueString() != data.PcePeers[j].PceAddress.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if found {
		} else {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/pcc/pce-peers/pce-peer%v", state.getPath(), keyString))
		}
	}
	for i := range state.OnDemandColors {
		keys := [...]string{"color"}
		stateKeyValues := [...]string{strconv.FormatInt(state.OnDemandColors[i].Color.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.OnDemandColors[i].Color.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.OnDemandColors {
			found = true
			if state.OnDemandColors[i].Color.ValueInt64() != data.OnDemandColors[j].Color.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if found {
		} else {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v", state.getPath(), keyString))
		}
	}
	for i := range state.Policies {
		keys := [...]string{"policy-name"}
		stateKeyValues := [...]string{state.Policies[i].PolicyName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Policies[i].PolicyName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Policies {
			found = true
			if state.Policies[i].PolicyName.ValueString() != data.Policies[j].PolicyName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if found {
		} else {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/policies/policy%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *SegmentRoutingTE) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.LoggingPcepPeerStatus.IsNull() && !data.LoggingPcepPeerStatus.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/logging/pcep-peer-status", data.getPath()))
	}
	if !data.LoggingPolicyStatus.IsNull() && !data.LoggingPolicyStatus.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/logging/policy-status", data.getPath()))
	}
	if !data.PccReportAll.IsNull() && !data.PccReportAll.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/pcc/report-all", data.getPath()))
	}
	for i := range data.PcePeers {
		keys := [...]string{"pce-address"}
		keyValues := [...]string{data.PcePeers[i].PceAddress.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.OnDemandColors {
		keys := [...]string{"color"}
		keyValues := [...]string{strconv.FormatInt(data.OnDemandColors[i].Color.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.OnDemandColors[i].DynamicAnycastSidInclusion.IsNull() && !data.OnDemandColors[i].DynamicAnycastSidInclusion.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v/on-demand-color-dyn-mpls/on-demand-color-dyn-mpls-anycast", data.getPath(), keyString))
		}
		if !data.OnDemandColors[i].Srv6Enable.IsNull() && !data.OnDemandColors[i].Srv6Enable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v/srv6/enable", data.getPath(), keyString))
		}
		if !data.OnDemandColors[i].EffectiveMetricEnable.IsNull() && !data.OnDemandColors[i].EffectiveMetricEnable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/on-demand-colors/on-demand-color%v/effective-metric/enable", data.getPath(), keyString))
		}
	}
	for i := range data.Policies {
		keys := [...]string{"policy-name"}
		keyValues := [...]string{data.Policies[i].PolicyName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.Policies[i].Srv6Enable.IsNull() && !data.Policies[i].Srv6Enable.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/policies/policy%v/srv6/enable", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}
