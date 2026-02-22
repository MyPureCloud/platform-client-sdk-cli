package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributesidresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributesidresponseDud struct { 
    

}

// Customattributesidresponse
type Customattributesidresponse struct { 
    // Id - The id of the newly created or updated Custom Attributes record.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Customattributesidresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributesidresponse) MarshalJSON() ([]byte, error) {
    type Alias Customattributesidresponse

    if CustomattributesidresponseMarshalled {
        return []byte("{}"), nil
    }
    CustomattributesidresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

