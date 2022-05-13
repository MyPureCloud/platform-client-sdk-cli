package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateresponseassetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateresponseassetresponseDud struct { 
    Id string `json:"id"`


    Url string `json:"url"`


    Headers map[string]string `json:"headers"`

}

// Createresponseassetresponse
type Createresponseassetresponse struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Createresponseassetresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createresponseassetresponse) MarshalJSON() ([]byte, error) {
    type Alias Createresponseassetresponse

    if CreateresponseassetresponseMarshalled {
        return []byte("{}"), nil
    }
    CreateresponseassetresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

