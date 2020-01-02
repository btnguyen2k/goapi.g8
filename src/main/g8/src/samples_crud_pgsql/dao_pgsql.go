package samples_crud_pgsql

import (
	"encoding/json"
	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/godal/sql"
	"github.com/btnguyen2k/prom"
)

var (
	// our table has only 2 columns: `id` and `data`
	tableColumns = []string{"id", "data"}
)

/*----------------------------------------------------------------------*/

type DaoDepartmentPgsql struct {
	*sql.GenericDaoSql
	tableName string
}

func NewDaoDepartmentPgsql(sqlc *prom.SqlConnect, tableName string) IDaoDepartment {
	dao := &DaoDepartmentPgsql{tableName: tableName}
	dao.GenericDaoSql = sql.NewGenericDaoSql(sqlc, godal.NewAbstractGenericDao(dao))
	dao.SetSqlFlavor(prom.FlavorPgSql)
	dao.SetRowMapper(&sql.GenericRowMapperSql{NameTransformation: sql.NameTransfLowerCase, ColumnsListMap: map[string][]string{tableName: tableColumns}})
	return dao
}

/*----------------------------------------------------------------------*/
// GdaoCreateFilter implements godal.IGenericDao.GdaoCreateFilter.
//
//  - DAO must implement GdaoCreateFilter
func (dao *DaoDepartmentPgsql) GdaoCreateFilter(storageId string, bo godal.IGenericBo) interface{} {
	id, _ := bo.GboGetAttr("id", reddo.TypeString)
	return map[string]interface{}{"id": id}
}

// toBo transforms godal.IGenericBo to BoDepartment
func (dao *DaoDepartmentPgsql) toBo(gbo godal.IGenericBo) *BoDepartment {
	if gbo == nil {
		return nil
	}
	id := gbo.GboGetAttrUnsafe("id", reddo.TypeString).(string)
	dataRaw := gbo.GboGetAttrUnsafe("data", reddo.TypeString).(string)

	bo := NewBoDepartment()
	bo.id = id
	json.Unmarshal([]byte(dataRaw), &bo.data)
	return bo
}

// toBo transforms godal.IGenericBo to BoDepartment
func (dao *DaoDepartmentPgsql) toGenericBo(bo *BoDepartment) godal.IGenericBo {
	if bo == nil {
		return nil
	}
	gbo := godal.NewGenericBo()
	gbo.GboSetAttr("id", bo.id)
	data, _ := json.Marshal(bo.data)
	gbo.GboSetAttr("data", string(data))
	return gbo
}

// Create implements IDaoPet.Create
func (dao *DaoDepartmentPgsql) Create(bo *BoDepartment) (bool, error) {
	numRows, err := dao.GdaoCreate(dao.tableName, dao.toGenericBo(bo))
	return numRows > 0, err
}

// Get implements IDaoPet.Get
func (dao *DaoDepartmentPgsql) Get(id string) (*BoDepartment, error) {
	filter := map[string]interface{}{"id": id} // alternative: filter := sql.FilterFieldValue{"id", "=", id}
	gbo, err := dao.GdaoFetchOne(dao.tableName, filter)
	if err != nil || gbo == nil {
		return nil, err
	}
	return dao.toBo(gbo), nil
}

// GetAll implements IDaoPet.GetAll
func (dao *DaoDepartmentPgsql) GetAll() ([]*BoDepartment, error) {
	sorting := map[string]int{"id": 1} // sort by "id" attribute, ascending
	rows, err := dao.GdaoFetchMany(dao.tableName, nil, sorting, 0, 0)
	if err != nil {
		return nil, err
	}
	var result []*BoDepartment
	for _, row := range rows {
		bo := dao.toBo(row)
		if bo != nil {
			result = append(result, bo)
		}
	}
	return result, nil
}

// Update implements IDaoPet.Update
func (dao *DaoDepartmentPgsql) Update(bo *BoDepartment) (bool, error) {
	numRows, err := dao.GdaoUpdate(dao.tableName, dao.toGenericBo(bo))
	return numRows > 0, err
}

// Delete implements IDaoPet.Delete
func (dao *DaoDepartmentPgsql) Delete(bo *BoDepartment) (bool, error) {
	numRows, err := dao.GdaoDelete(dao.tableName, dao.toGenericBo(bo))
	return numRows > 0, err
}
