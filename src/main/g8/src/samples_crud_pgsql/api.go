package samples_crud_pgsql

import (
	"github.com/btnguyen2k/consu/reddo"
	"main/src/itineris"
	"strings"
)

/*
API handler "pgsqlGetListDepartments"
*/
func apiListDepartments(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	depts, err := daoDepartment.GetAll()
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	result := make([]map[string]interface{}, 0)
	for _, d := range depts {
		result = append(result, d.ToMap())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}

/*
API handler "pgsqlCreateDepartment"
*/
func apiCreateDepartment(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	name, err := params.GetParamAsType("name", reddo.TypeString)
	if err != nil || name == nil || strings.TrimSpace(name.(string)) == "" {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Department must have a name.")
	}
	bo := NewBoDepartment().SetName(strings.TrimSpace(name.(string)))
	dbResult, err := daoDepartment.Create(bo)
	if err != nil || !dbResult {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage("Error while creating department: " + err.Error())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(bo.ToMap())
}

/*
API handler "pgsqlGetDepartment"
*/
func apiGetDepartment(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	id, err := params.GetParamAsType("id", reddo.TypeString)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Cannot parse parameter [id].")
	}
	bo, err := daoDepartment.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if bo == nil {
		return itineris.ResultNotFound
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(bo.ToMap())
}

/*
API handler "pgsqlUpdateDepartment"
*/
func apiUpdateDepartment(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	id, err := params.GetParamAsType("id", reddo.TypeString)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Cannot parse parameter [id].")
	}
	bo, err := daoDepartment.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if bo == nil {
		return itineris.ResultNotFound
	}

	name, _ := params.GetParamAsType("name", reddo.TypeString)
	if name != nil && strings.TrimSpace(name.(string)) != "" {
		bo.SetName(strings.TrimSpace(name.(string)))
	}
	ok, err := daoDepartment.Update(bo)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(ok)
}

/*
API handler "pgsqlDeleteDepartment"
*/
func apiDeleteDepartment(_ *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	id, err := params.GetParamAsType("id", reddo.TypeString)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Cannot parse parameter [id].")
	}
	bo, err := daoDepartment.Get(id.(string))
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	if bo == nil {
		return itineris.ResultNotFound
	}
	ok, err := daoDepartment.Delete(bo)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(ok)
}
