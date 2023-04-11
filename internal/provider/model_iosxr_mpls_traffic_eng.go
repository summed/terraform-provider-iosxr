// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MPLSTrafficEng struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	TrafficEng types.Bool   `tfsdk:"traffic_eng"`
}

func (data MPLSTrafficEng) getPath() string {
	return "Cisco-IOS-XR-um-mpls-te-cfg:mpls"
}

func (data MPLSTrafficEng) toBody(ctx context.Context) string {
	body := "{}"
	if !data.TrafficEng.IsNull() && !data.TrafficEng.IsUnknown() {
		if data.TrafficEng.ValueBool() {
			body, _ = sjson.Set(body, "traffic-eng", map[string]string{})
		}
	}
	return body
}

func (data *MPLSTrafficEng) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "traffic-eng"); !data.TrafficEng.IsNull() {
		if value.Exists() {
			data.TrafficEng = types.BoolValue(true)
		} else {
			data.TrafficEng = types.BoolValue(false)
		}
	} else {
		data.TrafficEng = types.BoolNull()
	}
}

func (data *MPLSTrafficEng) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "traffic-eng"); value.Exists() {
		data.TrafficEng = types.BoolValue(true)
	} else {
		data.TrafficEng = types.BoolValue(false)
	}
}

func (data *MPLSTrafficEng) fromPlan(ctx context.Context, plan MPLSTrafficEng) {
	data.Device = plan.Device
}

func (data *MPLSTrafficEng) getDeletedListItems(ctx context.Context, state MPLSTrafficEng) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *MPLSTrafficEng) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
