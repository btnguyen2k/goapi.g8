package samples_crud_pgsql

import (
	"github.com/btnguyen2k/consu/reddo"
	"main/src/utils"
)

func NewBoDepartment() *BoDepartment {
	bo := &BoDepartment{id: utils.UniqueIdSmall(), data: make(map[string]interface{})}
	return bo
}

type BoDepartment struct {
	id   string
	data map[string]interface{}
}

func (bo *BoDepartment) ToMap() map[string]interface{} {
	// clone/copy data field
	var data = make(map[string]interface{})
	for k, v := range bo.data {
		data[k] = v
	}
	data["id"] = bo.id
	return data
}

func (bo *BoDepartment) GetName() string {
	name, err := reddo.ToString(bo.data["name"])
	if err != nil {
		return ""
	}
	return name
}

func (bo *BoDepartment) SetName(value string) *BoDepartment {
	bo.data["name"] = value
	return bo
}

// IDaoDepartment defines DAO APIs to access department database storage.
type IDaoDepartment interface {
	// Create persists a new department to database storage. If the department already existed, this function returns (false, nil)
	Create(bo *BoDepartment) (bool, error)

	// Get finds a department by id & fetches it from database storage.
	Get(id string) (*BoDepartment, error)

	// GetAll retrieves all available departments from database storage and returns them as a list.
	GetAll() ([]*BoDepartment, error)

	// Update modifies an existing department in the database storage. If the department does not exist in database, this function returns (false, nil).
	Update(bo *BoDepartment) (bool, error)

	// Delete removes a department from database storage.
	Delete(bo *BoDepartment) (bool, error)
}

/*-----------------------------------------------------------------------------*/

func NewBoEmployee() *BoEmployee {
	bo := &BoEmployee{id: utils.UniqueIdSmall(), data: make(map[string]interface{})}
	return bo
}

type BoEmployee struct {
	id   string
	data map[string]interface{}
}

func (bo *BoEmployee) ToMap() map[string]interface{} {
	// clone/copy data field
	var data = make(map[string]interface{})
	for k, v := range bo.data {
		data[k] = v
	}
	data["id"] = bo.id
	return data
}

// IDaoEmployee defines DAO APIs to access employee database storage.
type IDaoEmployee interface {
	// Create persists a new employee to database storage. If the department already existed, this function returns (false, nil)
	Create(bo *BoEmployee) (bool, error)

	// Get finds an employee by id & fetches it from database storage.
	Get(id string) (*BoEmployee, error)

	// GetAll retrieves all available employees from database storage and returns them as a list.
	GetAll() ([]*BoEmployee, error)

	// Update modifies an existing employee in the database storage. If the employee does not exist in database, this function returns (false, nil).
	Update(bo *BoEmployee) (bool, error)

	// Delete removes an employee from database storage.
	Delete(bo *BoEmployee) (bool, error)
}
