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

var _ resource.Resource = (*RouterISISResource)(nil)

func NewRouterISISResource() resource.Resource {
	return &RouterISISResource{}
}

type RouterISISResource struct {
	client *client.Client
}

func (r *RouterISISResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_isis"
}

func (r *RouterISISResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Router ISIS configuration.",

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
			"process_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Process ID").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 36),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"is_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Area type (level)").AddStringEnumDescription("level-1", "level-1-2", "level-2-only").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("level-1", "level-1-2", "level-2-only"),
				},
			},
			"set_overload_bit_levels": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set overload-bit for one level only").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"level_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set overload-bit for one level only").AddIntegerRangeDescription(1, 2).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2),
							},
						},
						"on_startup_advertise_as_overloaded": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time in seconds to advertise ourself as overloaded after reboot").String,
							Optional:            true,
						},
						"on_startup_advertise_as_overloaded_time_to_advertise": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time in seconds to advertise ourself as overloaded after reboot").AddIntegerRangeDescription(5, 86400).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(5, 86400),
							},
						},
						"on_startup_wait_for_bgp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set overload bit on startup until BGP signals convergence, or timeout").String,
							Optional:            true,
						},
						"advertise_external": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If overload-bit set advertise IP prefixes learned from other protocols").String,
							Optional:            true,
						},
						"advertise_interlevel": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If overload-bit set advertise IP prefixes learned from another ISIS level").String,
							Optional:            true,
						},
					},
				},
			},
			"nsr": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable NSR").String,
				Optional:            true,
			},
			"nsf_cisco": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco Proprietary NSF restart").String,
				Optional:            true,
			},
			"nsf_ietf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IETF NSF restar").String,
				Optional:            true,
			},
			"nsf_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum route lifetime following restart (seconds)").AddIntegerRangeDescription(5, 300).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 300),
				},
			},
			"nsf_interface_timer": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timer used to wait for a restart ACK (seconds)").AddIntegerRangeDescription(1, 20).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"nsf_interface_expires": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("# of times T1 can expire waiting for the restart ACK").AddIntegerRangeDescription(1, 10).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"log_adjacency_changes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging adjacency state changes").String,
				Optional:            true,
			},
			"lsp_gen_interval_maximum_wait": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum delay before generating an LSP").AddIntegerRangeDescription(0, 120000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 120000),
				},
			},
			"lsp_gen_interval_initial_wait": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Initial delay before generating an LSP").AddIntegerRangeDescription(0, 120000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 120000),
				},
			},
			"lsp_gen_interval_secondary_wait": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary delay before generating an LSP").AddIntegerRangeDescription(0, 120000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 120000),
				},
			},
			"lsp_refresh_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set LSP refresh interval").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"max_lsp_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum LSP lifetime").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"lsp_password_keychain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies a Key Chain name will follow").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 1024),
				},
			},
			"distribute_link_state_instance_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distribution process instance identifier").AddIntegerRangeDescription(32, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(32, 4294967295),
				},
			},
			"distribute_link_state_throttle": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set throttle update in seconds").AddIntegerRangeDescription(1, 20).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"distribute_link_state_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distribution for one level only").AddIntegerRangeDescription(1, 2).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2),
				},
			},
			"affinity_maps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Affinity map configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Affinity map configuration").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
							},
						},
						"bit_position": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bit position for affinity attribute value").AddIntegerRangeDescription(0, 255).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
					},
				},
			},
			"flex_algos": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flex Algorithm definition").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"algorithm_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Flex Algorithm definition").AddIntegerRangeDescription(128, 255).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(128, 255),
							},
						},
						"advertise_definition": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Advertise the Flex-Algo Definition").String,
							Optional:            true,
						},
						"metric_type_delay": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use delay as metric").String,
							Optional:            true,
						},
					},
				},
			},
			"nets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A Network Entity Title (NET) for this process").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"net_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A Network Entity Title (NET) for this process").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 1024),
							},
						},
					},
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter the IS-IS interface configuration submode").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enter the IS-IS interface configuration submode").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9.:_/-]+`), ""),
							},
						},
						"circuit_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure circuit type for interface").AddStringEnumDescription("level-1", "level-1-2", "level-2-only").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("level-1", "level-1-2", "level-2-only"),
							},
						},
						"hello_padding_disable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Disable hello-padding").String,
							Optional:            true,
						},
						"hello_padding_sometimes": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable hello-padding during adjacency formation only").String,
							Optional:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set priority for Designated Router election").AddIntegerRangeDescription(0, 127).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 127),
							},
						},
						"point_to_point": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Treat active LAN interface as point-to-point").String,
							Optional:            true,
						},
						"passive": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not establish adjacencies over this interface").String,
							Optional:            true,
						},
						"suppressed": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not advertise connected prefixes of this interface").String,
							Optional:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Shutdown IS-IS on this interface").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *RouterISISResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RouterISISResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouterISIS

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

func (r *RouterISISResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouterISIS

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

	state.updateFromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *RouterISISResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouterISIS

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

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

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

func (r *RouterISISResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouterISIS

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	_, diags = r.client.Set(ctx, state.Device.ValueString(), client.SetOperation{Path: state.getPath(), Body: "", Operation: client.Delete})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *RouterISISResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
