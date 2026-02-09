package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemcasereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemcasereferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Workitemcasereference
type Workitemcasereference struct { 
    


    // Reference - The reference identifier of the Case.
    Reference string `json:"reference"`


    

}

// String returns a JSON representation of the model
func (o *Workitemcasereference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemcasereference) MarshalJSON() ([]byte, error) {
    type Alias Workitemcasereference

    if WorkitemcasereferenceMarshalled {
        return []byte("{}"), nil
    }
    WorkitemcasereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Reference string `json:"reference"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

