package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggroupmodelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggroupmodelDud struct { 
    

}

// Planninggroupmodel
type Planninggroupmodel struct { 
    // PlanningGroup - Planning group id
    PlanningGroup string `json:"planningGroup"`

}

// String returns a JSON representation of the model
func (o *Planninggroupmodel) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggroupmodel) MarshalJSON() ([]byte, error) {
    type Alias Planninggroupmodel

    if PlanninggroupmodelMarshalled {
        return []byte("{}"), nil
    }
    PlanninggroupmodelMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroup string `json:"planningGroup"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

