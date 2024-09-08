package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionconfigDud struct { 
    Id string `json:"id"`


    Function Function `json:"function"`


    Zip Functionzipconfig `json:"zip"`


    UploadExceptionHistory []Functionzipconfig `json:"uploadExceptionHistory"`


    SelfUri string `json:"selfUri"`

}

// Functionconfig - Current action function configuration and zip upload configuration.
type Functionconfig struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Functionconfig) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Functionconfig) MarshalJSON() ([]byte, error) {
    type Alias Functionconfig

    if FunctionconfigMarshalled {
        return []byte("{}"), nil
    }
    FunctionconfigMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

