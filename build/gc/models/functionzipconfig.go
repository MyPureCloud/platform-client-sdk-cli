package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionzipconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionzipconfigDud struct { 
    Status string `json:"status"`


    Id string `json:"id"`


    Name string `json:"name"`


    DateCreated time.Time `json:"dateCreated"`


    ErrorMessage string `json:"errorMessage"`


    RequestId string `json:"requestId"`

}

// Functionzipconfig - Action function zip file upload settings and state.
type Functionzipconfig struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Functionzipconfig) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Functionzipconfig) MarshalJSON() ([]byte, error) {
    type Alias Functionzipconfig

    if FunctionzipconfigMarshalled {
        return []byte("{}"), nil
    }
    FunctionzipconfigMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

