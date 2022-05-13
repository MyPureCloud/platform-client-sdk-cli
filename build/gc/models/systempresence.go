package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SystempresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SystempresenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Systempresence
type Systempresence struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Systempresence) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Systempresence) MarshalJSON() ([]byte, error) {
    type Alias Systempresence

    if SystempresenceMarshalled {
        return []byte("{}"), nil
    }
    SystempresenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

