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

var _ resource.Resource = (*L2VPNXconnectGroupP2PResource)(nil)

func NewL2VPNXconnectGroupP2PResource() resource.Resource {
	return &L2VPNXconnectGroupP2PResource{}
}

type L2VPNXconnectGroupP2PResource struct {
	client *client.Client
}

func (r *L2VPNXconnectGroupP2PResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn_xconnect_group_p2p"
}

func (r *L2VPNXconnectGroupP2PResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the L2VPN Xconnect Group P2P configuration.",

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
			"group_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the group the cross connects belong to").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"p2p_xconnect_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure point to point cross connect commands").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 38),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description for cross connect").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify (sub-)interface name to cross connect").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify (sub-)interface name to cross connect").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
					},
				},
			},
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9\.]*`), ""),
							},
						},
						"pw_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the pseudowire id").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"pw_class": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("PW class template name to use for this XC").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
					},
				},
			},
			"ipv6_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F:\.]*`), ""),
							},
						},
						"pw_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the pseudowire id").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"pw_class": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("PW class template name to use for this XC").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
					},
				},
			},
		},
	}
}

func (r *L2VPNXconnectGroupP2PResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *L2VPNXconnectGroupP2PResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan L2VPNXconnectGroupP2P

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

func (r *L2VPNXconnectGroupP2PResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state L2VPNXconnectGroupP2P

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

func (r *L2VPNXconnectGroupP2PResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state L2VPNXconnectGroupP2P

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

func (r *L2VPNXconnectGroupP2PResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state L2VPNXconnectGroupP2P

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

func (r *L2VPNXconnectGroupP2PResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
