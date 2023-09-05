package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HristimeofftyperesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HristimeofftyperesponseDud struct { 
    


    


    

}

// Hristimeofftyperesponse
type Hristimeofftyperesponse struct { 
    // Id - ID of the time off type configured in integration
    Id string `json:"id"`


    // Name - Name of the time off type configured in integration
    Name string `json:"name"`


    // SecondaryId - Secondary ID of the time off type, if configured in integration
    SecondaryId string `json:"secondaryId"`

}

// String returns a JSON representation of the model
func (o *Hristimeofftyperesponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Hristimeofftyperesponse) MarshalJSON() ([]byte, error) {
    type Alias Hristimeofftyperesponse

    if HristimeofftyperesponseMarshalled {
        return []byte("{}"), nil
    }
    HristimeofftyperesponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SecondaryId string `json:"secondaryId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

