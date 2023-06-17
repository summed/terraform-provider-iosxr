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

type Interface struct {
	Device                                   types.String                   `tfsdk:"device"`
	Id                                       types.String                   `tfsdk:"id"`
	InterfaceName                            types.String                   `tfsdk:"interface_name"`
	L2transport                              types.Bool                     `tfsdk:"l2transport"`
	PointToPoint                             types.Bool                     `tfsdk:"point_to_point"`
	Multipoint                               types.Bool                     `tfsdk:"multipoint"`
	DampeningDecayHalfLifeValue              types.Int64                    `tfsdk:"dampening_decay_half_life_value"`
	Ipv4PointToPoint                         types.Bool                     `tfsdk:"ipv4_point_to_point"`
	ServicePolicyInput                       []InterfaceServicePolicyInput  `tfsdk:"service_policy_input"`
	ServicePolicyOutput                      []InterfaceServicePolicyOutput `tfsdk:"service_policy_output"`
	BfdModeIetf                              types.Bool                     `tfsdk:"bfd_mode_ietf"`
	EncapsulationDot1qVlanId                 types.Int64                    `tfsdk:"encapsulation_dot1q_vlan_id"`
	L2transportEncapsulationDot1qVlanId      types.String                   `tfsdk:"l2transport_encapsulation_dot1q_vlan_id"`
	L2transportEncapsulationDot1qSecondDot1q types.String                   `tfsdk:"l2transport_encapsulation_dot1q_second_dot1q"`
	RewriteIngressTagPopOne                  types.Bool                     `tfsdk:"rewrite_ingress_tag_pop_one"`
	RewriteIngressTagPopTwo                  types.Bool                     `tfsdk:"rewrite_ingress_tag_pop_two"`
	Shutdown                                 types.Bool                     `tfsdk:"shutdown"`
	Mtu                                      types.Int64                    `tfsdk:"mtu"`
	Bandwidth                                types.Int64                    `tfsdk:"bandwidth"`
	Description                              types.String                   `tfsdk:"description"`
	LoadInterval                             types.Int64                    `tfsdk:"load_interval"`
	Vrf                                      types.String                   `tfsdk:"vrf"`
	Ipv4Address                              types.String                   `tfsdk:"ipv4_address"`
	Ipv4Netmask                              types.String                   `tfsdk:"ipv4_netmask"`
	Unnumbered                               types.String                   `tfsdk:"unnumbered"`
	Ipv6LinkLocalAddress                     types.String                   `tfsdk:"ipv6_link_local_address"`
	Ipv6LinkLocalZone                        types.String                   `tfsdk:"ipv6_link_local_zone"`
	Ipv6Autoconfig                           types.Bool                     `tfsdk:"ipv6_autoconfig"`
	Ipv6Enable                               types.Bool                     `tfsdk:"ipv6_enable"`
	Ipv6Addresses                            []InterfaceIpv6Addresses       `tfsdk:"ipv6_addresses"`
}
type InterfaceServicePolicyInput struct {
	Name types.String `tfsdk:"name"`
}
type InterfaceServicePolicyOutput struct {
	Name types.String `tfsdk:"name"`
}
type InterfaceIpv6Addresses struct {
	Address      types.String `tfsdk:"address"`
	PrefixLength types.Int64  `tfsdk:"prefix_length"`
	Zone         types.String `tfsdk:"zone"`
}

func (data Interface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=%s]", data.InterfaceName.ValueString())
}

func (data Interface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interface-name", data.InterfaceName.ValueString())
	}
	if !data.L2transport.IsNull() && !data.L2transport.IsUnknown() {
		if data.L2transport.ValueBool() {
			body, _ = sjson.Set(body, "sub-interface-type.l2transport", map[string]string{})
		}
	}
	if !data.PointToPoint.IsNull() && !data.PointToPoint.IsUnknown() {
		if data.PointToPoint.ValueBool() {
			body, _ = sjson.Set(body, "sub-interface-type.point-to-point", map[string]string{})
		}
	}
	if !data.Multipoint.IsNull() && !data.Multipoint.IsUnknown() {
		if data.Multipoint.ValueBool() {
			body, _ = sjson.Set(body, "sub-interface-type.multipoint", map[string]string{})
		}
	}
	if !data.DampeningDecayHalfLifeValue.IsNull() && !data.DampeningDecayHalfLifeValue.IsUnknown() {
		body, _ = sjson.Set(body, "dampening.decay-half-life.value", strconv.FormatInt(data.DampeningDecayHalfLifeValue.ValueInt64(), 10))
	}
	if !data.Ipv4PointToPoint.IsNull() && !data.Ipv4PointToPoint.IsUnknown() {
		if data.Ipv4PointToPoint.ValueBool() {
			body, _ = sjson.Set(body, "ipv4.Cisco-IOS-XR-um-if-ipv4-cfg:point-to-point", map[string]string{})
		}
	}
	if !data.BfdModeIetf.IsNull() && !data.BfdModeIetf.IsUnknown() {
		if data.BfdModeIetf.ValueBool() {
			body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-bundle-cfg:bfd.mode.ietf", map[string]string{})
		}
	}
	if !data.EncapsulationDot1qVlanId.IsNull() && !data.EncapsulationDot1qVlanId.IsUnknown() {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-l2-ethernet-cfg:encapsulation.dot1q.vlan-id", strconv.FormatInt(data.EncapsulationDot1qVlanId.ValueInt64(), 10))
	}
	if !data.L2transportEncapsulationDot1qVlanId.IsNull() && !data.L2transportEncapsulationDot1qVlanId.IsUnknown() {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-l2-ethernet-cfg:l2transport-encapsulation.dot1q.vlan-id", data.L2transportEncapsulationDot1qVlanId.ValueString())
	}
	if !data.L2transportEncapsulationDot1qSecondDot1q.IsNull() && !data.L2transportEncapsulationDot1qSecondDot1q.IsUnknown() {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-l2-ethernet-cfg:l2transport-encapsulation.dot1q.second-dot1q", data.L2transportEncapsulationDot1qSecondDot1q.ValueString())
	}
	if !data.RewriteIngressTagPopOne.IsNull() && !data.RewriteIngressTagPopOne.IsUnknown() {
		if data.RewriteIngressTagPopOne.ValueBool() {
			body, _ = sjson.Set(body, "Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite.ingress.tag.pop.one", map[string]string{})
		}
	}
	if !data.RewriteIngressTagPopTwo.IsNull() && !data.RewriteIngressTagPopTwo.IsUnknown() {
		if data.RewriteIngressTagPopTwo.ValueBool() {
			body, _ = sjson.Set(body, "Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite.ingress.tag.pop.two", map[string]string{})
		}
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, "shutdown", map[string]string{})
		}
	}
	if !data.Mtu.IsNull() && !data.Mtu.IsUnknown() {
		body, _ = sjson.Set(body, "mtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if !data.Bandwidth.IsNull() && !data.Bandwidth.IsUnknown() {
		body, _ = sjson.Set(body, "bandwidth", strconv.FormatInt(data.Bandwidth.ValueInt64(), 10))
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.LoadInterval.IsNull() && !data.LoadInterval.IsUnknown() {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-statistics-cfg:load-interval", strconv.FormatInt(data.LoadInterval.ValueInt64(), 10))
	}
	if !data.Vrf.IsNull() && !data.Vrf.IsUnknown() {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-vrf-cfg:vrf", data.Vrf.ValueString())
	}
	if !data.Ipv4Address.IsNull() && !data.Ipv4Address.IsUnknown() {
		body, _ = sjson.Set(body, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address", data.Ipv4Address.ValueString())
	}
	if !data.Ipv4Netmask.IsNull() && !data.Ipv4Netmask.IsUnknown() {
		body, _ = sjson.Set(body, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask", data.Ipv4Netmask.ValueString())
	}
	if !data.Unnumbered.IsNull() && !data.Unnumbered.IsUnknown() {
		body, _ = sjson.Set(body, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered", data.Unnumbered.ValueString())
	}
	if !data.Ipv6LinkLocalAddress.IsNull() && !data.Ipv6LinkLocalAddress.IsUnknown() {
		body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.link-local-address.address", data.Ipv6LinkLocalAddress.ValueString())
	}
	if !data.Ipv6LinkLocalZone.IsNull() && !data.Ipv6LinkLocalZone.IsUnknown() {
		body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.link-local-address.zone", data.Ipv6LinkLocalZone.ValueString())
	}
	if !data.Ipv6Autoconfig.IsNull() && !data.Ipv6Autoconfig.IsUnknown() {
		if data.Ipv6Autoconfig.ValueBool() {
			body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.autoconfig", map[string]string{})
		}
	}
	if !data.Ipv6Enable.IsNull() && !data.Ipv6Enable.IsUnknown() {
		if data.Ipv6Enable.ValueBool() {
			body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:enable", map[string]string{})
		}
	}
	if len(data.ServicePolicyInput) > 0 {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.input", []interface{}{})
		for index, item := range data.ServicePolicyInput {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.input"+"."+strconv.Itoa(index)+"."+"service-policy-name", item.Name.ValueString())
			}
		}
	}
	if len(data.ServicePolicyOutput) > 0 {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.output", []interface{}{})
		for index, item := range data.ServicePolicyOutput {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.output"+"."+strconv.Itoa(index)+"."+"service-policy-name", item.Name.ValueString())
			}
		}
	}
	if len(data.Ipv6Addresses) > 0 {
		body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.ipv6-address", []interface{}{})
		for index, item := range data.Ipv6Addresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.ipv6-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.PrefixLength.IsNull() && !item.PrefixLength.IsUnknown() {
				body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.ipv6-address"+"."+strconv.Itoa(index)+"."+"prefix-length", strconv.FormatInt(item.PrefixLength.ValueInt64(), 10))
			}
			if !item.Zone.IsNull() && !item.Zone.IsUnknown() {
				body, _ = sjson.Set(body, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.ipv6-address"+"."+strconv.Itoa(index)+"."+"zone", item.Zone.ValueString())
			}
		}
	}
	return body
}

func (data *Interface) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "sub-interface-type.l2transport"); !data.L2transport.IsNull() {
		if value.Exists() {
			data.L2transport = types.BoolValue(true)
		} else {
			data.L2transport = types.BoolValue(false)
		}
	} else {
		data.L2transport = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "sub-interface-type.point-to-point"); !data.PointToPoint.IsNull() {
		if value.Exists() {
			data.PointToPoint = types.BoolValue(true)
		} else {
			data.PointToPoint = types.BoolValue(false)
		}
	} else {
		data.PointToPoint = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "sub-interface-type.multipoint"); !data.Multipoint.IsNull() {
		if value.Exists() {
			data.Multipoint = types.BoolValue(true)
		} else {
			data.Multipoint = types.BoolValue(false)
		}
	} else {
		data.Multipoint = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "dampening.decay-half-life.value"); value.Exists() && !data.DampeningDecayHalfLifeValue.IsNull() {
		data.DampeningDecayHalfLifeValue = types.Int64Value(value.Int())
	} else {
		data.DampeningDecayHalfLifeValue = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ipv4-cfg:point-to-point"); !data.Ipv4PointToPoint.IsNull() {
		if value.Exists() {
			data.Ipv4PointToPoint = types.BoolValue(true)
		} else {
			data.Ipv4PointToPoint = types.BoolValue(false)
		}
	} else {
		data.Ipv4PointToPoint = types.BoolNull()
	}
	for i := range data.ServicePolicyInput {
		keys := [...]string{"service-policy-name"}
		keyValues := [...]string{data.ServicePolicyInput[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.input").ForEach(
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
		if value := r.Get("service-policy-name"); value.Exists() && !data.ServicePolicyInput[i].Name.IsNull() {
			data.ServicePolicyInput[i].Name = types.StringValue(value.String())
		} else {
			data.ServicePolicyInput[i].Name = types.StringNull()
		}
	}
	for i := range data.ServicePolicyOutput {
		keys := [...]string{"service-policy-name"}
		keyValues := [...]string{data.ServicePolicyOutput[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.output").ForEach(
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
		if value := r.Get("service-policy-name"); value.Exists() && !data.ServicePolicyOutput[i].Name.IsNull() {
			data.ServicePolicyOutput[i].Name = types.StringValue(value.String())
		} else {
			data.ServicePolicyOutput[i].Name = types.StringNull()
		}
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-bundle-cfg:bfd.mode.ietf"); !data.BfdModeIetf.IsNull() {
		if value.Exists() {
			data.BfdModeIetf = types.BoolValue(true)
		} else {
			data.BfdModeIetf = types.BoolValue(false)
		}
	} else {
		data.BfdModeIetf = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:encapsulation.dot1q.vlan-id"); value.Exists() && !data.EncapsulationDot1qVlanId.IsNull() {
		data.EncapsulationDot1qVlanId = types.Int64Value(value.Int())
	} else {
		data.EncapsulationDot1qVlanId = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:l2transport-encapsulation.dot1q.vlan-id"); value.Exists() && !data.L2transportEncapsulationDot1qVlanId.IsNull() {
		data.L2transportEncapsulationDot1qVlanId = types.StringValue(value.String())
	} else {
		data.L2transportEncapsulationDot1qVlanId = types.StringNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:l2transport-encapsulation.dot1q.second-dot1q"); value.Exists() && !data.L2transportEncapsulationDot1qSecondDot1q.IsNull() {
		data.L2transportEncapsulationDot1qSecondDot1q = types.StringValue(value.String())
	} else {
		data.L2transportEncapsulationDot1qSecondDot1q = types.StringNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite.ingress.tag.pop.one"); !data.RewriteIngressTagPopOne.IsNull() {
		if value.Exists() {
			data.RewriteIngressTagPopOne = types.BoolValue(true)
		} else {
			data.RewriteIngressTagPopOne = types.BoolValue(false)
		}
	} else {
		data.RewriteIngressTagPopOne = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite.ingress.tag.pop.two"); !data.RewriteIngressTagPopTwo.IsNull() {
		if value.Exists() {
			data.RewriteIngressTagPopTwo = types.BoolValue(true)
		} else {
			data.RewriteIngressTagPopTwo = types.BoolValue(false)
		}
	} else {
		data.RewriteIngressTagPopTwo = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "shutdown"); !data.Shutdown.IsNull() {
		if value.Exists() {
			data.Shutdown = types.BoolValue(true)
		} else {
			data.Shutdown = types.BoolValue(false)
		}
	} else {
		data.Shutdown = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "mtu"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bandwidth"); value.Exists() && !data.Bandwidth.IsNull() {
		data.Bandwidth = types.Int64Value(value.Int())
	} else {
		data.Bandwidth = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-statistics-cfg:load-interval"); value.Exists() && !data.LoadInterval.IsNull() {
		data.LoadInterval = types.Int64Value(value.Int())
	} else {
		data.LoadInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-vrf-cfg:vrf"); value.Exists() && !data.Vrf.IsNull() {
		data.Vrf = types.StringValue(value.String())
	} else {
		data.Vrf = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask"); value.Exists() && !data.Ipv4Netmask.IsNull() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered"); value.Exists() && !data.Unnumbered.IsNull() {
		data.Unnumbered = types.StringValue(value.String())
	} else {
		data.Unnumbered = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.link-local-address.address"); value.Exists() && !data.Ipv6LinkLocalAddress.IsNull() {
		data.Ipv6LinkLocalAddress = types.StringValue(value.String())
	} else {
		data.Ipv6LinkLocalAddress = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.link-local-address.zone"); value.Exists() && !data.Ipv6LinkLocalZone.IsNull() {
		data.Ipv6LinkLocalZone = types.StringValue(value.String())
	} else {
		data.Ipv6LinkLocalZone = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.autoconfig"); !data.Ipv6Autoconfig.IsNull() {
		if value.Exists() {
			data.Ipv6Autoconfig = types.BoolValue(true)
		} else {
			data.Ipv6Autoconfig = types.BoolValue(false)
		}
	} else {
		data.Ipv6Autoconfig = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:enable"); !data.Ipv6Enable.IsNull() {
		if value.Exists() {
			data.Ipv6Enable = types.BoolValue(true)
		} else {
			data.Ipv6Enable = types.BoolValue(false)
		}
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	for i := range data.Ipv6Addresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6Addresses[i].Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.ipv6-address").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.Ipv6Addresses[i].Address.IsNull() {
			data.Ipv6Addresses[i].Address = types.StringValue(value.String())
		} else {
			data.Ipv6Addresses[i].Address = types.StringNull()
		}
		if value := r.Get("prefix-length"); value.Exists() && !data.Ipv6Addresses[i].PrefixLength.IsNull() {
			data.Ipv6Addresses[i].PrefixLength = types.Int64Value(value.Int())
		} else {
			data.Ipv6Addresses[i].PrefixLength = types.Int64Null()
		}
		if value := r.Get("zone"); value.Exists() && !data.Ipv6Addresses[i].Zone.IsNull() {
			data.Ipv6Addresses[i].Zone = types.StringValue(value.String())
		} else {
			data.Ipv6Addresses[i].Zone = types.StringNull()
		}
	}
}

func (data *Interface) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "sub-interface-type.l2transport"); value.Exists() {
		data.L2transport = types.BoolValue(true)
	} else {
		data.L2transport = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "sub-interface-type.point-to-point"); value.Exists() {
		data.PointToPoint = types.BoolValue(true)
	} else {
		data.PointToPoint = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "sub-interface-type.multipoint"); value.Exists() {
		data.Multipoint = types.BoolValue(true)
	} else {
		data.Multipoint = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "dampening.decay-half-life.value"); value.Exists() {
		data.DampeningDecayHalfLifeValue = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ipv4-cfg:point-to-point"); value.Exists() {
		data.Ipv4PointToPoint = types.BoolValue(true)
	} else {
		data.Ipv4PointToPoint = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.input"); value.Exists() {
		data.ServicePolicyInput = make([]InterfaceServicePolicyInput, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceServicePolicyInput{}
			if cValue := v.Get("service-policy-name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.ServicePolicyInput = append(data.ServicePolicyInput, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy.output"); value.Exists() {
		data.ServicePolicyOutput = make([]InterfaceServicePolicyOutput, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceServicePolicyOutput{}
			if cValue := v.Get("service-policy-name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.ServicePolicyOutput = append(data.ServicePolicyOutput, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-bundle-cfg:bfd.mode.ietf"); value.Exists() {
		data.BfdModeIetf = types.BoolValue(true)
	} else {
		data.BfdModeIetf = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:encapsulation.dot1q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:l2transport-encapsulation.dot1q.vlan-id"); value.Exists() {
		data.L2transportEncapsulationDot1qVlanId = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:l2transport-encapsulation.dot1q.second-dot1q"); value.Exists() {
		data.L2transportEncapsulationDot1qSecondDot1q = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite.ingress.tag.pop.one"); value.Exists() {
		data.RewriteIngressTagPopOne = types.BoolValue(true)
	} else {
		data.RewriteIngressTagPopOne = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite.ingress.tag.pop.two"); value.Exists() {
		data.RewriteIngressTagPopTwo = types.BoolValue(true)
	} else {
		data.RewriteIngressTagPopTwo = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "mtu"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bandwidth"); value.Exists() {
		data.Bandwidth = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-statistics-cfg:load-interval"); value.Exists() {
		data.LoadInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-vrf-cfg:vrf"); value.Exists() {
		data.Vrf = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address"); value.Exists() {
		data.Ipv4Address = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask"); value.Exists() {
		data.Ipv4Netmask = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv4.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered"); value.Exists() {
		data.Unnumbered = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.link-local-address.address"); value.Exists() {
		data.Ipv6LinkLocalAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.link-local-address.zone"); value.Exists() {
		data.Ipv6LinkLocalZone = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.autoconfig"); value.Exists() {
		data.Ipv6Autoconfig = types.BoolValue(true)
	} else {
		data.Ipv6Autoconfig = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:enable"); value.Exists() {
		data.Ipv6Enable = types.BoolValue(true)
	} else {
		data.Ipv6Enable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "ipv6.Cisco-IOS-XR-um-if-ip-address-cfg:addresses.ipv6-address"); value.Exists() {
		data.Ipv6Addresses = make([]InterfaceIpv6Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceIpv6Addresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("prefix-length"); cValue.Exists() {
				item.PrefixLength = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("zone"); cValue.Exists() {
				item.Zone = types.StringValue(cValue.String())
			}
			data.Ipv6Addresses = append(data.Ipv6Addresses, item)
			return true
		})
	}
}

func (data *Interface) getDeletedListItems(ctx context.Context, state Interface) []string {
	deletedListItems := make([]string, 0)
	for i := range state.ServicePolicyInput {
		keys := [...]string{"service-policy-name"}
		stateKeyValues := [...]string{state.ServicePolicyInput[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.ServicePolicyInput[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.ServicePolicyInput {
			found = true
			if state.ServicePolicyInput[i].Name.ValueString() != data.ServicePolicyInput[j].Name.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy/input%v", state.getPath(), keyString))
		}
	}
	for i := range state.ServicePolicyOutput {
		keys := [...]string{"service-policy-name"}
		stateKeyValues := [...]string{state.ServicePolicyOutput[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.ServicePolicyOutput[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.ServicePolicyOutput {
			found = true
			if state.ServicePolicyOutput[i].Name.ValueString() != data.ServicePolicyOutput[j].Name.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XR-um-if-service-policy-qos-cfg:service-policy/output%v", state.getPath(), keyString))
		}
	}
	for i := range state.Ipv6Addresses {
		keys := [...]string{"address"}
		stateKeyValues := [...]string{state.Ipv6Addresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6Addresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6Addresses {
			found = true
			if state.Ipv6Addresses[i].Address.ValueString() != data.Ipv6Addresses[j].Address.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ipv6/Cisco-IOS-XR-um-if-ip-address-cfg:addresses/ipv6-address%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *Interface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.L2transport.IsNull() && !data.L2transport.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/sub-interface-type/l2transport", data.getPath()))
	}
	if !data.PointToPoint.IsNull() && !data.PointToPoint.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/sub-interface-type/point-to-point", data.getPath()))
	}
	if !data.Multipoint.IsNull() && !data.Multipoint.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/sub-interface-type/multipoint", data.getPath()))
	}
	if !data.Ipv4PointToPoint.IsNull() && !data.Ipv4PointToPoint.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv4/Cisco-IOS-XR-um-if-ipv4-cfg:point-to-point", data.getPath()))
	}
	for i := range data.ServicePolicyInput {
		keys := [...]string{"service-policy-name"}
		keyValues := [...]string{data.ServicePolicyInput[i].Name.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.ServicePolicyOutput {
		keys := [...]string{"service-policy-name"}
		keyValues := [...]string{data.ServicePolicyOutput[i].Name.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	if !data.BfdModeIetf.IsNull() && !data.BfdModeIetf.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XR-um-if-bundle-cfg:bfd/mode/ietf", data.getPath()))
	}
	if !data.RewriteIngressTagPopOne.IsNull() && !data.RewriteIngressTagPopOne.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite/ingress/tag/pop/one", data.getPath()))
	}
	if !data.RewriteIngressTagPopTwo.IsNull() && !data.RewriteIngressTagPopTwo.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XR-um-l2-ethernet-cfg:rewrite/ingress/tag/pop/two", data.getPath()))
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.Ipv6Autoconfig.IsNull() && !data.Ipv6Autoconfig.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/Cisco-IOS-XR-um-if-ip-address-cfg:addresses/autoconfig", data.getPath()))
	}
	if !data.Ipv6Enable.IsNull() && !data.Ipv6Enable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/Cisco-IOS-XR-um-if-ip-address-cfg:enable", data.getPath()))
	}
	for i := range data.Ipv6Addresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6Addresses[i].Address.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}
