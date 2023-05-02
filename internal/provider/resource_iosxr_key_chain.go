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

var _ resource.Resource = (*KeyChainResource)(nil)

func NewKeyChainResource() resource.Resource {
	return &KeyChainResource{}
}

type KeyChainResource struct {
	client *client.Client
}

func (r *KeyChainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_key_chain"
}

func (r *KeyChainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Key Chain configuration.",

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the key chain - maximum 32 characters").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"keys": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a Key").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure a Key").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 800),
								stringvalidator.RegexMatches(regexp.MustCompile(`[\w\-\.:,_@#%$\+=\|;]+`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]{1,15}`), ""),
							},
						},
						"key_string_password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Encrypted key string (even number of characters with first two as digits and sum less than 53, and rest of the characters should be hex digits)").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(!.+)|([^!].+)`), ""),
							},
						},
						"cryptographic_algorithm": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choose cryptographic algorithm").AddStringEnumDescription("aes-128-cmac-96", "hmac-md5", "hmac-sha-256", "hmac-sha1-12", "hmac-sha1-20", "hmac-sha1-96", "md5", "sha-1").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes-128-cmac-96", "hmac-md5", "hmac-sha-256", "hmac-sha1-12", "hmac-sha1-20", "hmac-sha1-96", "md5", "sha-1"),
							},
						},
						"accept_lifetime_start_time_hour": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start time hour").AddIntegerRangeDescription(0, 23).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 23),
							},
						},
						"accept_lifetime_start_time_minute": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start time minute").AddIntegerRangeDescription(0, 59).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
						},
						"accept_lifetime_start_time_second": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start time second").AddIntegerRangeDescription(0, 59).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
						},
						"accept_lifetime_start_time_day_of_month": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Day of the month").AddIntegerRangeDescription(1, 31).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 31),
							},
						},
						"accept_lifetime_start_time_month": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Month of the year").AddStringEnumDescription("april", "august", "december", "february", "january", "july", "june", "march", "may", "november", "october", "september").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("april", "august", "december", "february", "january", "july", "june", "march", "may", "november", "october", "september"),
							},
						},
						"accept_lifetime_start_time_year": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Year").AddIntegerRangeDescription(1993, 2035).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1993, 2035),
							},
						},
						"accept_lifetime_infinite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Never expires").String,
							Optional:            true,
						},
						"send_lifetime_start_time_hour": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start time hour").AddIntegerRangeDescription(0, 23).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 23),
							},
						},
						"send_lifetime_start_time_minute": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start time minute").AddIntegerRangeDescription(0, 59).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
						},
						"send_lifetime_start_time_second": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start time second").AddIntegerRangeDescription(0, 59).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
						},
						"send_lifetime_start_time_day_of_month": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Day of the month").AddIntegerRangeDescription(1, 31).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 31),
							},
						},
						"send_lifetime_start_time_month": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Month of the year").AddStringEnumDescription("april", "august", "december", "february", "january", "july", "june", "march", "may", "november", "october", "september").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("april", "august", "december", "february", "january", "july", "june", "march", "may", "november", "october", "september"),
							},
						},
						"send_lifetime_start_time_year": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Year").AddIntegerRangeDescription(1993, 2035).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1993, 2035),
							},
						},
						"send_lifetime_infinite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Never expires").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *KeyChainResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *KeyChainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan KeyChain

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *KeyChainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state KeyChain

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

func (r *KeyChainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state KeyChain

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
	body := plan.toBody(ctx)

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), plan.getPath(), body, client.Update)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		_, diags = r.client.Set(ctx, plan.Device.ValueString(), i, "", client.Delete)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
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

func (r *KeyChainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state KeyChain

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

func (r *KeyChainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
