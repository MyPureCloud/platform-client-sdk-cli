package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperplanninggroupminimumsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperplanninggroupminimumsrequestDud struct { 
    

}

// Listwrapperplanninggroupminimumsrequest
type Listwrapperplanninggroupminimumsrequest struct { 
    // Values
    Values []Planninggroupminimumsrequest `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperplanninggroupminimumsrequest) String() string {
     o.Values = []Planninggroupminimumsrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperplanninggroupminimumsrequest) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperplanninggroupminimumsrequest

    if ListwrapperplanninggroupminimumsrequestMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperplanninggroupminimumsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Values []Planninggroupminimumsrequest `json:"values"`
        *Alias
    }{

        
        Values: []Planninggroupminimumsrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

