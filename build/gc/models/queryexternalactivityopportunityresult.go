package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryexternalactivityopportunityresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryexternalactivityopportunityresultDud struct { 
    Id string `json:"id"`


    

}

// Queryexternalactivityopportunityresult
type Queryexternalactivityopportunityresult struct { 
    


    // OpportunityName - The name of the opportunity associated with this external activity
    OpportunityName string `json:"opportunityName"`

}

// String returns a JSON representation of the model
func (o *Queryexternalactivityopportunityresult) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryexternalactivityopportunityresult) MarshalJSON() ([]byte, error) {
    type Alias Queryexternalactivityopportunityresult

    if QueryexternalactivityopportunityresultMarshalled {
        return []byte("{}"), nil
    }
    QueryexternalactivityopportunityresultMarshalled = true

    return json.Marshal(&struct {
        
        OpportunityName string `json:"opportunityName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

