package main

import (
	_ "database/sql"
	"fmt"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

//Sla - resource of sla values
type OBJECT struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ID_City        int    `json:"city_id"`
	Address        string `json:"address"`
	ID_ObjectType  int    `json:"object_type_id"`
	DT_Open        string `json:"dt_open"`
	DT_Close       string `json:"dt_close"`
	ID_Focus       string `json:"focus_id"`
	ID_Datamanager string `json:"datamanager_id"`
	IsEnabled      bool   `json:"isenabled"`
}

type Objects []OBJECT

func FindObject(id int) Objects {

	var cmd = "set dateformat ymd; " +
		"SELECT oo.[ID] " +
		",oo.[Name] " +
		",oo.[ID_City] " +
		",isnull(oo.[Address],'не определен') [Address] " +
		",oo.[ID_ObjectType] " +
		",cast(isnull(oo.[DT_ObjectOpen],'2000-01-01 00:00:00') as date) [DT_Open] " +
		",cast(isnull(oo.[DT_ObjectClose],dateadd(yyyy,10,getdate())) as date) [DT_Close] " +
		",convert(char(36),isnull(oo.[ID_Focus],'00000000-0000-0000-0000-000000000000')) [ID_Focus] " +
		",isnull(oc.ID_Datamanager,0) [ID_Datamanager] " +
		",oo.[IsEnabled] " +
		"FROM [Object].[Objects] oo left join RMDM.UD_ObjectConfigs oc on oo.ID=oc.ID_Object " +
		"where oo.ID=" + strconv.Itoa(id)
	obj, err := ObjectGet(cmd)
	if *debug {
		fmt.Printf("object return:%v\n", obj)
	}
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	// return empty Object if not found
	return obj
}

//ListObject - get list object from DB
func ListObject() Objects {

	var cmd = "set dateformat ymd; " +
		"SELECT oo.[ID] " +
		",oo.[Name] " +
		",oo.[ID_City] " +
		",isnull(oo.[Address],'не определен') [Address] " +
		",oo.[ID_ObjectType] " +
		",cast(isnull(oo.[DT_ObjectOpen],'2000-01-01 00:00:00') as date) [DT_Open] " +
		",cast(isnull(oo.[DT_ObjectClose],dateadd(yyyy,10,getdate())) as date) [DT_Close] " +
		",convert(char(36),isnull(oo.[ID_Focus],'00000000-0000-0000-0000-000000000000')) [ID_Focus] " +
		",isnull(oc.ID_Datamanager,0) [ID_Datamanager]" +
		",oo.[IsEnabled] " +
		"FROM [Object].[Objects] oo left join RMDM.UD_ObjectConfigs oc on oo.ID=oc.ID_Object " +
		"where oo.ID_ObjectType in (1,2) " +
		"order by oo.ID_ObjectType, oo.Name"
	obj, err := ObjectGet(cmd)
	if *debug {
		fmt.Printf("sla return:%v\n", obj)
	}
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	// return empty OBJ if not found
	return obj
}
