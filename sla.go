package main

import (
	"fmt"
	"strconv"
)

/*
SELECT TOP (1000) [ID]
      ,[Name]
      ,[Description]
      ,[Value]
	  ,[icon] as [Icon]
  FROM [CM_Info_Test].[RMDM].[SLATypes]
*/

//Sla - resource of sla values
type SLA struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Icon        string `json:"icon"`
}

type Slas []SLA

func FindSla(id int) Slas {

	var cmd = "SELECT [ID],[Name],[Description],[Value],[icon] as [Icon] FROM [RMDM].[SLATypes] where ID=" + strconv.Itoa(id)
	sla, err := slaGet(cmd)
	if *debug {
		fmt.Printf("sla return:%v\n", sla)
	}
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	// return empty SLA if not found
	return sla
}

//ListSla - get list sla from DB
func ListSla() Slas {

	var cmd = "SELECT [ID],[Name],[Description],[Value],[icon] as [Icon] FROM [RMDM].[SLATypes]"
	sla, err := slaGet(cmd)
	if *debug {
		fmt.Printf("sla return:%v\n", sla)
	}
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	// return empty SLA if not found
	return sla
}
