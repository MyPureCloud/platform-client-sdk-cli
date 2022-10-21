package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffbalancesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffbalancesresponseDud struct { 
    


    

}

// Timeoffbalancesresponse
type Timeoffbalancesresponse struct { 
    // Job - The asynchronous job handling the query
    Job Timeoffbalancejobreference `json:"job"`


    // Entities - The list of time off balances. May come via notification
    Entities []Timeoffbalanceresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Timeoffbalancesresponse) String() string {
    
     o.Entities = []Timeoffbalanceresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffbalancesresponse) MarshalJSON() ([]byte, error) {
    type Alias Timeoffbalancesresponse

    if TimeoffbalancesresponseMarshalled {
        return []byte("{}"), nil
    }
    TimeoffbalancesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Timeoffbalancejobreference `json:"job"`
        
        Entities []Timeoffbalanceresponse `json:"entities"`
        *Alias
    }{

        


        
        Entities: []Timeoffbalanceresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

