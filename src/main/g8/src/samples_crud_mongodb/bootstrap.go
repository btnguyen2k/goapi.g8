package samples_crud_mongodb

import (
	"github.com/btnguyen2k/prom"
	"log"
	"main/src/goems"
	"main/src/itineris"
)

type MyBootstrapper struct {
	name string
}

var (
	Bootstrapper = &MyBootstrapper{name: "samples_crud_mongodb"}
	mongoConnect *prom.MongoConnect
	daoPet       IDaoPet
)

const (
	collectionPet = "pets"
)

/*
Bootstrap implements goems.IBootstrapper.Bootstrap

Bootstrapper usually does:
- register api-handlers with the global ApiRouter
- other initializing work (e.g. creating DAO, initializing database, etc)
*/
func (b *MyBootstrapper) Bootstrap() error {
	initDaos()
	initApiHandlers(goems.ApiRouter)
	return nil
}

// construct an 'prom.MongoConnect' instance
func createMongoConnect() *prom.MongoConnect {
	url := goems.AppConfig.GetString("samples_crud_mongodb.mongodb.url", "mongodb://test:test@localhost:27017/test")
	db := goems.AppConfig.GetString("samples_crud_mongodb.mongodb.db", "test")
	timeoutMs := goems.AppConfig.GetInt32("samples_crud_mongodb.mongodb.timeout", 10000)
	mongoConnect, err := prom.NewMongoConnect(url, db, int(timeoutMs))
	if mongoConnect == nil || err != nil {
		if err != nil {
			log.Println(err)
		}
		panic("error creating [prom.MongoConnect] instance")
	}
	return mongoConnect
}

func initDaos() {
	mongoConnect = createMongoConnect()
	if ok, _ := mongoConnect.HasCollection(collectionPet); !ok {
		log.Printf("Creating collection [%s]...", collectionPet)
		result, err := mongoConnect.CreateCollection(collectionPet)
		log.Println("Result %#V %#V", result, err)
	}
	daoPet = NewDaoPetMongodb(mongoConnect, collectionPet)
}

/*
Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)

    - api-handler function must has the following signature: func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
*/
func initApiHandlers(router *itineris.ApiRouter) {
	router.SetHandler("mongoGetListPets", apiListPets)
	router.SetHandler("mongoCreatePet", apiCreatePet)
	router.SetHandler("mongoGetPet", apiGetPet)
	router.SetHandler("mongoUpdatePet", apiUpdatePet)
	router.SetHandler("mongoDeletePet", apiDeletePet)
}
