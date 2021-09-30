package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecastgenerationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecastgenerationresultDud struct { 
    

}

// Buforecastgenerationresult
type Buforecastgenerationresult struct { 
    // PlanningGroupResults - Generation results, broken down by planning group
    PlanningGroupResults []Buforecastgenerationplanninggroupresult `json:"planningGroupResults"`

}

// String returns a JSON representation of the model
func (o *Buforecastgenerationresult) String() string {
    
    
     o.PlanningGroupResults = []Buforecastgenerationplanninggroupresult{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecastgenerationresult) MarshalJSON() ([]byte, error) {
    type Alias Buforecastgenerationresult

    if BuforecastgenerationresultMarshalled {
        return []byte("{}"), nil
    }
    BuforecastgenerationresultMarshalled = true

    return json.Marshal(&struct { 
        PlanningGroupResults []Buforecastgenerationplanninggroupresult `json:"planningGroupResults"`
        
        *Alias
    }{
        

        
        PlanningGroupResults: []Buforecastgenerationplanninggroupresult{{}},
        

        
        Alias: (*Alias)(u),
    })
}

