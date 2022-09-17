package common

import (
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
)

func ConvertTreeOptions(data []*admin_pb.TreeOptions) (res []types.TreeOptions) {
	res = make([]types.TreeOptions, len(data))
	for _, item := range data {
		res = append(res, types.TreeOptions{
			Title:    item.Title,
			Value:    item.Value,
			Children: ConvertTreeOptions(item.Children),
		})
	}
	return
}

func ConvertEnumList(data []*admin_pb.Enum) (res []types.Enum) {
	res = make([]types.Enum, len(data))
	for _, item := range data {
		res = append(res, types.Enum{
			Label: item.Label,
			Value: item.Value,
		})
	}
	return
}
