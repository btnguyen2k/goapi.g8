package samples_crud_mongodb

import (
    "github.com/btnguyen2k/consu/reddo"
    "github.com/btnguyen2k/godal"
    "github.com/btnguyen2k/godal/mongo"
    "github.com/btnguyen2k/prom"
)

type DaoPetMongodb struct {
    *mongo.GenericDaoMongo
    collectionName string
}

func NewDaoPetMongodb(mc *prom.MongoConnect, collectionName string) IDaoPet {
    dao := &DaoPetMongodb{collectionName: collectionName}
    dao.GenericDaoMongo = mongo.NewGenericDaoMongo(mc, godal.NewAbstractGenericDao(dao))
    return dao
}

/*----------------------------------------------------------------------*/
// GdaoCreateFilter implements godal.IGenericDao.GdaoCreateFilter.
//
//  - DAO must implement GdaoCreateFilter
func (dao *DaoPetMongodb) GdaoCreateFilter(storageId string, bo godal.IGenericBo) interface{} {
    id, _ := bo.GboGetAttr("id", reddo.TypeString)
    return map[string]interface{}{"id": id}
}

// toBo transforms godal.IGenericBo to BoPet
func (dao *DaoPetMongodb) toBo(gbo godal.IGenericBo) *BoPet {
    if gbo == nil {
        return nil
    }
    return (&BoPet{}).fromGenericBo(gbo)
}

// Create implements IDaoPet.Create
func (dao *DaoPetMongodb) Create(bo *BoPet) (bool, error) {
    numRows, err := dao.GdaoCreate(dao.collectionName, bo.toGenericBo())
    return numRows > 0, err
}

// Get implements IDaoPet.Get
func (dao *DaoPetMongodb) Get(id string) (*BoPet, error) {
    filter := map[string]interface{}{"id": id}
    gbo, err := dao.GdaoFetchOne(dao.collectionName, filter)
    if err != nil || gbo == nil {
        return nil, err
    }
    return dao.toBo(gbo), nil
}

// GetAll implements IDaoPet.GetAll
func (dao *DaoPetMongodb) GetAll() ([]*BoPet, error) {
    sorting := map[string]int{"id": 1} // sort by "id" attribute, ascending
    rows, err := dao.GdaoFetchMany(dao.collectionName, nil, sorting, 0, 0)
    if err != nil {
        return nil, err
    }
    var result []*BoPet
    for _, e := range rows {
        bo := dao.toBo(e)
        if bo != nil {
            result = append(result, bo)
        }
    }
    return result, nil
}

// Update implements IDaoPet.Update
func (dao *DaoPetMongodb) Update(bo *BoPet) (bool, error) {
    numRows, err := dao.GdaoUpdate(dao.collectionName, bo.toGenericBo())
    return numRows > 0, err
}

// Delete implements IDaoPet.Delete
func (dao *DaoPetMongodb) Delete(bo *BoPet) (bool, error) {
    numRows, err := dao.GdaoDelete(dao.collectionName, bo.toGenericBo())
    return numRows > 0, err
}
