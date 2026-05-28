package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpportunityreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpportunityreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Opportunityreference
type Opportunityreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Opportunityreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opportunityreference) MarshalJSON() ([]byte, error) {
    type Alias Opportunityreference

    if OpportunityreferenceMarshalled {
        return []byte("{}"), nil
    }
    OpportunityreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

