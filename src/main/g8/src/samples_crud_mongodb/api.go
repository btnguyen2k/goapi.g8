package samples_crud_mongodb

import (
    "github.com/btnguyen2k/consu/reddo"
    "main/src/itineris"
    "main/src/utils"
    "strings"
)

/*
API handler "mongoGetListPets"
*/
func apiListPets(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    pets, err := daoPet.GetAll()
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
    }
    result := make([]map[string]interface{}, 0)
    for _, p := range pets {
        result = append(result, p.ToMapHumanReadable())
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}

/*
API handler "mongoCreatePet"
*/
func apiCreatePet(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    name, err := params.GetParamAsType("name", reddo.TypeString)
    if err != nil || name == nil || strings.TrimSpace(name.(string)) == "" {
        return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Pet must have a name.")
    }

    pet := &BoPet{}
    pet.id = utils.UniqueId()
    pet.status = StatusAvailable
    pet.name = strings.TrimSpace(name.(string))

    dbResult, err := daoPet.Create(pet)
    if err != nil || !dbResult {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage("Error while creating pet: " + err.Error())
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(pet.ToMapHumanReadable())
}

/*
API handler "mongoGetPet"
*/
func apiGetPet(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    id, err := params.GetParamAsType("id", reddo.TypeString)
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Cannot parse parameter [id].")
    }
    pet, err := daoPet.Get(id.(string))
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
    }
    if pet == nil {
        return itineris.ResultNotFound
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(pet.ToMapHumanReadable())
}

/*
API handler "mongoUpdatePet"
*/
func apiUpdatePet(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    id, err := params.GetParamAsType("id", reddo.TypeString)
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Cannot parse parameter [id].")
    }
    pet, err := daoPet.Get(id.(string))
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
    }
    if pet == nil {
        return itineris.ResultNotFound
    }

    name, _ := params.GetParamAsType("name", reddo.TypeString)
    status, _ := params.GetParamAsType("status", reddo.TypeInt)

    if name != nil && strings.TrimSpace(name.(string)) != "" {
        pet.name = strings.TrimSpace(name.(string))
    }
    if status != nil && status.(int64) >= 0 {
        pet.status = PetStatus(status.(int64))
    }
    ok, err := daoPet.Update(pet)
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(ok)
}

/*
API handler "mongoDeletePete"
*/
func apiDeletePet(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    id, err := params.GetParamAsType("id", reddo.TypeString)
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("Cannot parse parameter [id].")
    }
    pet, err := daoPet.Get(id.(string))
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
    }
    if pet == nil {
        return itineris.ResultNotFound
    }
    ok, err := daoPet.Delete(pet)
    if err != nil {
        return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(ok)
}
