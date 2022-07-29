package respx

import (
	"errors"
	"luck-draw/common/errorx"
	"reflect"
)

/** 与前端约定的数据规范
interface ErrorInfoStructure {
	success: boolean; // 请求状态
	data?: any; // 返回数据
	total?: number; // 数据总个数
	pageSize: number //单页数量
	current?: number; // 当前页码
	page?: number; // 【自己加上去的】总共有多少页，根据前端的pageSize来计算
}
**/

/**
* 前端要求的返回格式规范
  const result = {
   data: dataSource,
   total: tableListDataSource.length,
   success: true,
   pageSize: finalPageSize,
   current: parseInt(`${params.currentPage}`, 10) || 1,
 };
*/

/** Respx 对接前端antpro框架的接口格式规范
 */
// type Respx struct {
// 	Status   bool   `json:"success"`  // 【必需】请求状态
// 	Message  string `json:"message"`  // 【必需】响应消息
// 	Data     any    `json:"data"`     // 【必需】返回数据
// 	Total    int64  `json:"total"`    // 【选填】数据总个数
// 	PageSize int64  `json:"pageSize"` // 【选填】单页数量
// 	Current  int64  `json:"current"`  // 【选填】当前页码
// 	Page     int64  `json:"page"`     // 【自定】总共有多少页，根据前端的pageSize来计算
// }

var (
	//respx只会返回成功的请求
	SUCCESS_MSG = "request success"
)

//StateData 状态响应结构体
type StateData struct {
	Status  bool   `json:"success"` // 请求状态。
	Message string `json:"message"` // 响应消息
}

//SingleData 单数据响应结构体
type SingleData struct {
	Status  bool   `json:"success"` // 请求状态。可有可无
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 返回数据
}

//ListData 列表响应数据结构体
type ListData struct {
	Status   bool   `json:"success"`  // 【必需】请求状态
	Message  string `json:"message"`  // 【必需】响应消息
	Data     any    `json:"data"`     // 【选填】返回数据
	Total    int64  `json:"total"`    // 【选填】数据总个数
	PageSize int64  `json:"pageSize"` // 【选填】单页数量
	Current  int64  `json:"current"`  // 【选填】当前页码
	Page     int64  `json:"page"`     // 【自定】总共有多少页，根据前端的pageSize来计算
}

//StateResp 返回响应状态
func StateResp(alertMsg string, status bool) *StateData {
	return &StateData{
		Status:  status,
		Message: alertMsg,
	}
}

//DefaultStateResp 默认返回响应状态，主要用于对接go-zero
func DefaultStateResp(err error) (*StateData, error) {
	if err != nil {
		return nil, err
	}
	return &StateData{
		Status:  true,
		Message: SUCCESS_MSG,
	}, nil
}

//SingleResp 返回单数据响应
func SingleResp(alertMsg string, data any) *SingleData {
	return &SingleData{
		Status:  true,
		Message: alertMsg,
		Data:    data,
	}
}

//DefaultSingleResp 默认返回单数据响应，主要用于对接go-zero
func DefaultSingleResp(data any, err error) (*SingleData, error) {
	if err != nil {
		return nil, err
	}
	return &SingleData{
		Status:  true,
		Message: SUCCESS_MSG,
		Data:    data,
	}, nil
}

/** ListResp 返回列表数据响应，用于分页查询数据
* alertMsg string 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func ListResp(alertMsg string, listData any, totalRecord, pageSize, Current int64) *ListData {
	data := listData
	if listData != nil {
		if reflect.TypeOf(listData).Kind().String() != "slice" {
			data = []any{listData}
		}
	} else {
		data = []any{}
	}
	if pageSize == 0 {
		pageSize = 1
	}
	return &ListData{
		Status:   true,
		Message:  alertMsg,
		Data:     data,
		Total:    totalRecord,
		PageSize: pageSize,
		Current:  Current,
		Page:     (totalRecord + pageSize - 1) / pageSize,
	}
}

/** DefaultListResp 默认返回列表数据响应，用于分页查询数据，主要用于对接go-zero
* err error 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func DefaultListResp(listData any, totalRecord, pageSize, Current int64, err error) (*ListData, error) {
	if err != nil {
		return nil, err
	}
	data := listData
	if listData != nil {
		if reflect.TypeOf(listData).Kind().String() != "slice" {
			data = []any{listData}
		}
	} else {
		data = []any{}
	}
	if pageSize == 0 {
		pageSize = 1
	}
	return &ListData{
		Status:   true,
		Message:  SUCCESS_MSG,
		Data:     data,
		Total:    totalRecord,
		PageSize: pageSize,
		Current:  Current,
		Page:     (totalRecord + pageSize - 1) / pageSize,
	}, nil
}

//Msg 根据error返回不同提示信息
func Msg(data any, err error, successMsg, failedMsg string) (any, error) {
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	return SingleResp(successMsg, data), nil
}

//DefaultBoolResp 默认返回响应状态，主要用于对接go-zero
func DefaultBoolResp(ok bool, err error) (*StateData, error) {
	if ok {
		return &StateData{
			Status:  true,
			Message: SUCCESS_MSG,
		}, nil
	}
	if err == nil {
		err = errors.New("Request failed! But the server has no error prompt")
	}
	return nil, err
}
