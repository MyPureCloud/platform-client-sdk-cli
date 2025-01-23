package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationserviceasyncresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationserviceasyncresponseDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Validationserviceasyncresponse
type Validationserviceasyncresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Validationserviceasyncresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationserviceasyncresponse) MarshalJSON() ([]byte, error) {
    type Alias Validationserviceasyncresponse

    if ValidationserviceasyncresponseMarshalled {
        return []byte("{}"), nil
    }
    ValidationserviceasyncresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

