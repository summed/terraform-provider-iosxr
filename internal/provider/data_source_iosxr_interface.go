// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &InterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceDataSource{}
)

func NewInterfaceDataSource() datasource.DataSource {
	return &InterfaceDataSource{}
}

type InterfaceDataSource struct {
	client *client.Client
}

func (d *InterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface"
}

func (d *InterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface configuration subcommands",
				Required:            true,
			},
			"l2transport": schema.BoolAttribute{
				MarkdownDescription: "l2transport sub-interface",
				Computed:            true,
			},
			"point_to_point": schema.BoolAttribute{
				MarkdownDescription: "point-to-point sub-interface",
				Computed:            true,
			},
			"multipoint": schema.BoolAttribute{
				MarkdownDescription: "multipoint sub-interface",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "shutdown the given interface",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Set the MTU on an interface",
				Computed:            true,
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: "Set the bandwidth of an interface",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Set description for this interface",
				Computed:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "Set VRF in which the interface operates",
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
			},
			"ipv4_netmask": schema.StringAttribute{
				MarkdownDescription: "IP subnet mask",
				Computed:            true,
			},
			"unnumbered": schema.StringAttribute{
				MarkdownDescription: "Enable IP processing without an explicit address",
				Computed:            true,
			},
			"ipv6_link_local_address": schema.StringAttribute{
				MarkdownDescription: "IPv6 address",
				Computed:            true,
			},
			"ipv6_link_local_zone": schema.StringAttribute{
				MarkdownDescription: "IPv6 address zone",
				Computed:            true,
			},
			"ipv6_autoconfig": schema.BoolAttribute{
				MarkdownDescription: "Enable slaac on Mgmt interface",
				Computed:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable IPv6 on interface",
				Computed:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 address",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 name or address",
							Computed:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: "Prefix length in bits",
							Computed:            true,
						},
						"zone": schema.StringAttribute{
							MarkdownDescription: "IPv6 address zone",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *InterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *InterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Interface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
