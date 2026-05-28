package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryopportunityenrollmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryopportunityenrollmentsrequestDud struct { 
    


    

}

// Queryopportunityenrollmentsrequest
type Queryopportunityenrollmentsrequest struct { 
    // VarRange - The date range for the query. Exactly one of opportunityId or range must be set
    VarRange Requireddaterange `json:"range"`


    // OpportunityId - The ID of the specific opportunity by which to filter. Exactly one of opportunityId or range must be set
    OpportunityId string `json:"opportunityId"`

}

// String returns a JSON representation of the model
func (o *Queryopportunityenrollmentsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryopportunityenrollmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryopportunityenrollmentsrequest

    if QueryopportunityenrollmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryopportunityenrollmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarRange Requireddaterange `json:"range"`
        
        OpportunityId string `json:"opportunityId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

