package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModelmetadataresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModelmetadataresultDud struct { 
    


    

}

// Modelmetadataresult
type Modelmetadataresult struct { 
    // SessionInfo - Information about the continuous forecast session
    SessionInfo Sessioninfo `json:"sessionInfo"`


    // PlanningGroups - List of planning groups
    PlanningGroups []Planninggroupmodel `json:"planningGroups"`

}

// String returns a JSON representation of the model
func (o *Modelmetadataresult) String() string {
    
     o.PlanningGroups = []Planninggroupmodel{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modelmetadataresult) MarshalJSON() ([]byte, error) {
    type Alias Modelmetadataresult

    if ModelmetadataresultMarshalled {
        return []byte("{}"), nil
    }
    ModelmetadataresultMarshalled = true

    return json.Marshal(&struct {
        
        SessionInfo Sessioninfo `json:"sessionInfo"`
        
        PlanningGroups []Planninggroupmodel `json:"planningGroups"`
        *Alias
    }{

        


        
        PlanningGroups: []Planninggroupmodel{{}},
        

        Alias: (*Alias)(u),
    })
}

