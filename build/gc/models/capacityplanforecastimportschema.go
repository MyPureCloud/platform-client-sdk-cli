package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastimportschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastimportschemaDud struct { 
    


    


    

}

// Capacityplanforecastimportschema
type Capacityplanforecastimportschema struct { 
    // Granularity - The granularity at which forecast values are provided for each planning group
    Granularity string `json:"granularity"`


    // BusinessUnitStartDate - The first day of the forecast relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BusinessUnitStartDate time.Time `json:"businessUnitStartDate"`


    // PlanningGroups - The forecast data for the planning groups included in this forecast
    PlanningGroups []Capacityplanforecastplanninggroup `json:"planningGroups"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastimportschema) String() string {
    
    
     o.PlanningGroups = []Capacityplanforecastplanninggroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastimportschema) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastimportschema

    if CapacityplanforecastimportschemaMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastimportschemaMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        
        BusinessUnitStartDate time.Time `json:"businessUnitStartDate"`
        
        PlanningGroups []Capacityplanforecastplanninggroup `json:"planningGroups"`
        *Alias
    }{

        


        


        
        PlanningGroups: []Capacityplanforecastplanninggroup{{}},
        

        Alias: (*Alias)(u),
    })
}

