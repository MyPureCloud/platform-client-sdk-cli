package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModelingprocessingerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModelingprocessingerrorDud struct { 
    InternalErrorCode string `json:"internalErrorCode"`


    Description string `json:"description"`

}

// Modelingprocessingerror
type Modelingprocessingerror struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Modelingprocessingerror) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modelingprocessingerror) MarshalJSON() ([]byte, error) {
    type Alias Modelingprocessingerror

    if ModelingprocessingerrorMarshalled {
        return []byte("{}"), nil
    }
    ModelingprocessingerrorMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

