package samples_crud_pgsql

import (
	"fmt"
	"github.com/btnguyen2k/prom"
	_ "github.com/lib/pq"
	"log"
	"main/src/goems"
	"main/src/itineris"
	"main/src/utils"
)

type MyBootstrapper struct {
	name string
}

var (
	Bootstrapper  = &MyBootstrapper{name: "samples_crud_pgsql"}
	sqlConnect    *prom.SqlConnect
	daoDepartment IDaoDepartment
	daoEmployee   IDaoEmployee
)

const (
	tblDepartment = "tbl_department"
	tblEmployee   = "tbl_employee"
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

// construct an 'prom.SqlConnect' instance
//
// - DO NOT FORGET to import sql driver for postgres, such as: import _ "github.com/lib/pq"
func createSqlConnect() *prom.SqlConnect {
	driver := goems.AppConfig.GetString("samples_crud_pgsql.driver", "postgres")
	dsn := goems.AppConfig.GetString("samples_crud_pgsql.postgres.url", "postgres://test:test@localhost:5432/test?client_encoding=UTF-8&application_name=gems")
	timeoutMs := goems.AppConfig.GetInt32("samples_crud_pgsql.postgres.timeout", 10000)
	sqlConnect, err := prom.NewSqlConnectWithFlavor(driver, dsn, int(timeoutMs), nil, prom.FlavorPgSql)
	if sqlConnect == nil || err != nil {
		if err != nil {
			log.Println(err)
		}
		panic("error creating [prom.SqlConnect] instance")
	}
	sqlConnect.SetLocation(utils.Location)
	return sqlConnect
}

func initDaos() {
	sqlConnect = createSqlConnect()
	sql := "CREATE TABLE IF NOT EXISTS %s (id VARCHAR(64), data JSONB, PRIMARY KEY (id))"

	_, err := sqlConnect.GetDB().Exec(fmt.Sprintf(sql, tblDepartment))
	if err != nil {
		log.Println(err)
		panic("error creating table [" + tblDepartment + "].")
	}

	_, err = sqlConnect.GetDB().Exec(fmt.Sprintf(sql, tblEmployee))
	if err != nil {
		log.Println(err)
		panic("error creating table [" + tblEmployee + "].")
	}

	daoDepartment = NewDaoDepartmentPgsql(sqlConnect, tblDepartment)
	// daoEmployee = NewDaoEmployeePgsql(sqlConnect, tblEmployee)
}

/*
Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)

    - api-handler function must has the following signature: func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
*/
func initApiHandlers(router *itineris.ApiRouter) {
	router.SetHandler("pgsqlListDepartments", apiListDepartments)
	router.SetHandler("pgsqlCreateDepartment", apiCreateDepartment)
	router.SetHandler("pgsqlGetDepartment", apiGetDepartment)
	router.SetHandler("pgsqlUpdateDepartment", apiUpdateDepartment)
	router.SetHandler("pgsqlDeleteDepartment", apiDeleteDepartment)
}
