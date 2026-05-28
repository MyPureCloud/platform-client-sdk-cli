package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryopportunitiesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryopportunitiesrequestDud struct { 
    

}

// Queryopportunitiesrequest
type Queryopportunitiesrequest struct { 
    // VarRange - The date range for the query
    VarRange Requireddaterange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Queryopportunitiesrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryopportunitiesrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryopportunitiesrequest

    if QueryopportunitiesrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryopportunitiesrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarRange Requireddaterange `json:"range"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

