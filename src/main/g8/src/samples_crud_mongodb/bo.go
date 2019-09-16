package samples_crud_mongodb

import (
    "github.com/btnguyen2k/consu/reddo"
    "github.com/btnguyen2k/godal"
)

const (
    StatusUnknown PetStatus = iota
    StatusAvailable
    StatusSold
    StatusPending
)

type PetStatus int

func (s PetStatus) String() string {
    names := []string{
        "Unknown",
        "Available",
        "Sold",
        "Pending",
    }
    if s < StatusUnknown || s > StatusPending {
        return "Undefined"
    }
    return names[s]
}

type BoPet struct {
    id     string
    name   string
    status PetStatus
}

func (pet *BoPet) toGenericBo() godal.IGenericBo {
    gbo := godal.NewGenericBo()
    gbo.GboSetAttr("id", pet.id)
    gbo.GboSetAttr("name", pet.name)
    gbo.GboSetAttr("status", pet.status)
    return gbo
}

func (pet *BoPet) fromGenericBo(gbo godal.IGenericBo) *BoPet {
    pet.id = gbo.GboGetAttrUnsafe("id", reddo.TypeString).(string)
    pet.name = gbo.GboGetAttrUnsafe("name", reddo.TypeString).(string)
    pet.status = PetStatus(gbo.GboGetAttrUnsafe("status", reddo.TypeInt).(int64))
    return pet
}

func (pet *BoPet) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":     pet.id,
        "name":   pet.name,
        "status": pet.status,
    }
}

func (pet *BoPet) ToMapHumanReadable() map[string]interface{} {
    return map[string]interface{}{
        "id":     pet.id,
        "name":   pet.name,
        "status": pet.status.String(),
    }
}

// IDaoPet defines DAO APIs to access pet database storage.
type IDaoPet interface {
    // Create persists a new pet to database storage. If the pet already existed, this function returns (false, nil)
    Create(bo *BoPet) (bool, error)

    // Get finds a pet by id & fetches it from database storage.
    Get(id string) (*BoPet, error)

    // GetAll retrieves all available pets from database storage and returns them as a list.
    GetAll() ([]*BoPet, error)

    // Update modifies an existing pet in the database storage. If the pet does not exist in database, this function returns (false, nil).
    Update(bo *BoPet) (bool, error)

    // Delete removes a pet from database storage.
    Delete(bo *BoPet) (bool, error)
}
