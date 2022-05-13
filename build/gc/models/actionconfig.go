package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionconfigDud struct { 
    


    

}

// Actionconfig - Defines components of the Action Config.
type Actionconfig struct { 
    // Request - Configuration of outbound request.
    Request Requestconfig `json:"request"`


    // Response - Configuration of response processing.
    Response Responseconfig `json:"response"`

}

// String returns a JSON representation of the model
func (o *Actionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionconfig) MarshalJSON() ([]byte, error) {
    type Alias Actionconfig

    if ActionconfigMarshalled {
        return []byte("{}"), nil
    }
    ActionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Request Requestconfig `json:"request"`
        
        Response Responseconfig `json:"response"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

