// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewVRFResource() resource.Resource {
	return &VRFResource{}
}

type VRFResource struct {
	client *client.Client
}

func (r *VRFResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf"
}

func (r *VRFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the VRF configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF name").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A description for the VRF").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 1024),
				},
			},
			"vpn_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`([0-9a-f]{1,8}):([0-9a-f]{1,8})`), ""),
				},
			},
			"address_family_ipv4_unicast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unicast sub address family").String,
				Optional:            true,
			},
			"address_family_ipv4_unicast_import_route_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use route-policy for import filtering").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"address_family_ipv4_unicast_export_route_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use route-policy for export").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"address_family_ipv4_multicast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Multicast topology").String,
				Optional:            true,
			},
			"address_family_ipv4_flowspec": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flowspec sub address family").String,
				Optional:            true,
			},
			"address_family_ipv6_unicast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unicast sub address family").String,
				Optional:            true,
			},
			"address_family_ipv6_unicast_import_route_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use route-policy for import filtering").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"address_family_ipv6_unicast_export_route_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use route-policy for export").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"address_family_ipv6_multicast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Multicast topology").String,
				Optional:            true,
			},
			"address_family_ipv6_flowspec": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flowspec sub address family").String,
				Optional:            true,
			},
			"rd_two_byte_as_as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Optional:            true,
			},
			"rd_two_byte_as_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"rd_four_byte_as_as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("4-byte AS number").String,
				Optional:            true,
			},
			"rd_four_byte_as_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"rd_ip_address_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("configure this node").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
				},
			},
			"rd_ip_address_index": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"address_family_ipv4_unicast_import_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_import_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_import_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_export_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_export_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv4_unicast_export_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_import_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_import_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_import_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_export_route_target_two_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Two Byte AS Number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_export_route_target_four_byte_as_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number Route Target").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Four Byte AS number").AddIntegerRangeDescription(65536, 4294967295).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(65536, 4294967295),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASN2:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
			"address_family_ipv6_unicast_export_route_target_ip_address_format": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"index": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4Address:index (hex or decimal format)").AddIntegerRangeDescription(0, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"stitching": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("These are stitching RTs").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *VRFResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *VRFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VRFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	import_ := false
	if state.Id.ValueString() == "" {
		import_ = true
		state.Id = types.StringValue(state.getPath())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	respBody := getResp.Notification[0].Update[0].Val.GetJsonIetfVal()
	if import_ {
		state.fromBody(ctx, respBody)
	} else {
		state.updateFromBody(ctx, respBody)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *VRFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	deletedListItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VRFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"
	if state.DeleteMode.ValueString() == "all" {
		deleteMode = "all"
	} else if state.DeleteMode.ValueString() == "attributes" {
		deleteMode = "attributes"
	}

	if deleteMode == "all" {
		ops = append(ops, client.SetOperation{Path: state.Id.ValueString(), Body: "", Operation: client.Delete})
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
		}
	}

	_, diags = r.client.Set(ctx, state.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *VRFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <vrf_name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_name"), idParts[0])...)
}
