// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

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
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
)

var _ resource.Resource = (*RouterBGPVRFAddressFamilyResource)(nil)

func NewRouterBGPVRFAddressFamilyResource() resource.Resource {
	return &RouterBGPVRFAddressFamilyResource{}
}

type RouterBGPVRFAddressFamilyResource struct {
	client *client.Client
}

func (r *RouterBGPVRFAddressFamilyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_bgp_vrf_address_family"
}

func (r *RouterBGPVRFAddressFamilyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router BGP VRF Address Family configuration.",

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
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bgp as-number").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a vrf name").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"af_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter Address Family command mode").AddStringEnumDescription("all-address-family", "ipv4-flowspec", "ipv4-labeled-unicast", "ipv4-mdt", "ipv4-multicast", "ipv4-mvpn", "ipv4-rt-filter", "ipv4-sr-policy", "ipv4-tunnel", "ipv4-unicast", "ipv6-flowspec", "ipv6-labeled-unicast", "ipv6-multicast", "ipv6-mvpn", "ipv6-sr-policy", "ipv6-unicast", "l2vpn-evpn", "l2vpn-mspw", "l2vpn-vpls-vpws", "link-state-link-state", "vpnv4-flowspec", "vpnv4-multicast", "vpnv4-unicast", "vpnv6-flowspec", "vpnv6-multicast", "vpnv6-unicast").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all-address-family", "ipv4-flowspec", "ipv4-labeled-unicast", "ipv4-mdt", "ipv4-multicast", "ipv4-mvpn", "ipv4-rt-filter", "ipv4-sr-policy", "ipv4-tunnel", "ipv4-unicast", "ipv6-flowspec", "ipv6-labeled-unicast", "ipv6-multicast", "ipv6-mvpn", "ipv6-sr-policy", "ipv6-unicast", "l2vpn-evpn", "l2vpn-mspw", "l2vpn-vpls-vpws", "link-state-link-state", "vpnv4-flowspec", "vpnv4-multicast", "vpnv4-unicast", "vpnv6-flowspec", "vpnv6-multicast", "vpnv6-unicast"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"maximum_paths_ebgp_multipath": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("eBGP-multipath").AddIntegerRangeDescription(2, 128).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 128),
				},
			},
			"maximum_paths_eibgp_multipath": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("eiBGP-multipath").AddIntegerRangeDescription(2, 128).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 128),
				},
			},
			"maximum_paths_ibgp_multipath": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("iBGP-multipath").AddIntegerRangeDescription(2, 128).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 128),
				},
			},
			"label_mode_per_ce": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set per CE label mode").String,
				Optional:            true,
				Computed:            true,
			},
			"label_mode_per_vrf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set per VRF label mode").String,
				Optional:            true,
				Computed:            true,
			},
			"redistribute_connected": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Connected routes").String,
				Optional:            true,
				Computed:            true,
			},
			"redistribute_connected_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Metric for redistributed routes").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"redistribute_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static routes").String,
				Optional:            true,
				Computed:            true,
			},
			"redistribute_static_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Metric for redistributed routes").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Aggregate address and mask or masklength").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Aggregate address and mask or masklength").String,
							Optional:            true,
							Computed:            true,
						},
						"masklength": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network in prefix/length format (prefix part)").AddIntegerRangeDescription(0, 128).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 128),
							},
						},
						"as_set": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Generate AS set path information").String,
							Optional:            true,
							Computed:            true,
						},
						"as_confed_set": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Generate AS confed set path information").String,
							Optional:            true,
							Computed:            true,
						},
						"summary_only": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter more specific routes from updates").String,
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 network and mask or masklength").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 network and mask or masklength").String,
							Optional:            true,
							Computed:            true,
						},
						"masklength": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network in prefix/length format (prefix part)").AddIntegerRangeDescription(0, 128).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 128),
							},
						},
					},
				},
			},
			"redistribute_ospf": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Open Shortest Path First (OSPF/OSPFv3)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"router_tag": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Open Shortest Path First (OSPF)").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF internal routes").String,
							Optional:            true,
							Computed:            true,
						},
						"match_internal_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF external routes").String,
							Optional:            true,
							Computed:            true,
						},
						"match_internal_nssa_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF NSSA external routes").String,
							Optional:            true,
							Computed:            true,
						},
						"match_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF external routes").String,
							Optional:            true,
							Computed:            true,
						},
						"match_external_nssa_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF NSSA external routes").String,
							Optional:            true,
							Computed:            true,
						},
						"match_nssa_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute OSPF NSSA external routes").String,
							Optional:            true,
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric for redistributed routes").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
					},
				},
			},
		},
	}
}

func (r *RouterBGPVRFAddressFamilyResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterBGPVRFAddressFamilyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterBGPVRFAddressFamily

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody()

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete()
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.setUnknownValues()

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterBGPVRFAddressFamilyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterBGPVRFAddressFamily

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.updateFromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterBGPVRFAddressFamilyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterBGPVRFAddressFamily

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody()

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.setUnknownValues()

	deletedListItems := plan.getDeletedListItems(state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete()
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterBGPVRFAddressFamilyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterBGPVRFAddressFamily

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), state.getPath(), "", client.Delete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *RouterBGPVRFAddressFamilyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
