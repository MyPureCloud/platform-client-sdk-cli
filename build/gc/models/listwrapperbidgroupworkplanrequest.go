package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperbidgroupworkplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperbidgroupworkplanrequestDud struct { 
    

}

// Listwrapperbidgroupworkplanrequest
type Listwrapperbidgroupworkplanrequest struct { 
    // Values
    Values []Bidgroupworkplanrequest `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperbidgroupworkplanrequest) String() string {
     o.Values = []Bidgroupworkplanrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperbidgroupworkplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperbidgroupworkplanrequest

    if ListwrapperbidgroupworkplanrequestMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperbidgroupworkplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Values []Bidgroupworkplanrequest `json:"values"`
        *Alias
    }{

        
        Values: []Bidgroupworkplanrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

