package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutlierresultbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutlierresultbodyDud struct { 
    


    

}

// Outlierresultbody
type Outlierresultbody struct { 
    // PlanningGroupId - The ID of the planning group for which outliers are present
    PlanningGroupId string `json:"planningGroupId"`


    // Outliers - Outliers detected in the forecast data
    Outliers []Outlier `json:"outliers"`

}

// String returns a JSON representation of the model
func (o *Outlierresultbody) String() string {
    
     o.Outliers = []Outlier{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outlierresultbody) MarshalJSON() ([]byte, error) {
    type Alias Outlierresultbody

    if OutlierresultbodyMarshalled {
        return []byte("{}"), nil
    }
    OutlierresultbodyMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        Outliers []Outlier `json:"outliers"`
        *Alias
    }{

        


        
        Outliers: []Outlier{{}},
        

        Alias: (*Alias)(u),
    })
}

