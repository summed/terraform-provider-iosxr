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

type LoggingSourceInterface struct {
	Device types.String                 `tfsdk:"device"`
	Id     types.String                 `tfsdk:"id"`
	Name   types.String                 `tfsdk:"name"`
	Vrfs   []LoggingSourceInterfaceVrfs `tfsdk:"vrfs"`
}
type LoggingSourceInterfaceVrfs struct {
	Name types.String `tfsdk:"name"`
}

func (data LoggingSourceInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-logging-cfg:/logging/source-interfaces/source-interface[source-interface-name=%v]", data.Name.ValueString())
}

func (data LoggingSourceInterface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "source-interface-name", data.Name.ValueString())
	}
	if len(data.Vrfs) > 0 {
		body, _ = sjson.Set(body, "vrfs.vrf", []interface{}{})
		for index, item := range data.Vrfs {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf"+"."+strconv.Itoa(index)+"."+"vrf-name", item.Name.ValueString())
			}
		}
	}
	return body
}

func (data *LoggingSourceInterface) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.Vrfs {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.Vrfs[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "vrfs.vrf").ForEach(
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
		if value := r.Get("vrf-name"); value.Exists() && !data.Vrfs[i].Name.IsNull() {
			data.Vrfs[i].Name = types.StringValue(value.String())
		} else {
			data.Vrfs[i].Name = types.StringNull()
		}
	}
}

func (data *LoggingSourceInterface) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "vrfs.vrf"); value.Exists() {
		data.Vrfs = make([]LoggingSourceInterfaceVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LoggingSourceInterfaceVrfs{}
			if cValue := v.Get("vrf-name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.Vrfs = append(data.Vrfs, item)
			return true
		})
	}
}

func (data *LoggingSourceInterface) getDeletedListItems(ctx context.Context, state LoggingSourceInterface) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Vrfs {
		keys := [...]string{"vrf-name"}
		stateKeyValues := [...]string{state.Vrfs[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Vrfs[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Vrfs {
			found = true
			if state.Vrfs[i].Name.ValueString() != data.Vrfs[j].Name.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/vrfs/vrf%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *LoggingSourceInterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.Vrfs {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.Vrfs[i].Name.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}
