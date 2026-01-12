package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateagentunavailabletimesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateagentunavailabletimesresponseDud struct { 
    Id string `json:"id"`


    


    

}

// Validateagentunavailabletimesresponse
type Validateagentunavailabletimesresponse struct { 
    


    // Status - The current status of the job
    Status string `json:"status"`


    // Result - Validation results if status == 'Complete'
    Result Unavailabletimesvalidationresult `json:"result"`

}

// String returns a JSON representation of the model
func (o *Validateagentunavailabletimesresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateagentunavailabletimesresponse) MarshalJSON() ([]byte, error) {
    type Alias Validateagentunavailabletimesresponse

    if ValidateagentunavailabletimesresponseMarshalled {
        return []byte("{}"), nil
    }
    ValidateagentunavailabletimesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Result Unavailabletimesvalidationresult `json:"result"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

