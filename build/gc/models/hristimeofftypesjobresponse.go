package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HristimeofftypesjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HristimeofftypesjobresponseDud struct { 
    


    

}

// Hristimeofftypesjobresponse
type Hristimeofftypesjobresponse struct { 
    // Entities
    Entities []Hristimeofftyperesponse `json:"entities"`


    // Status - The status of the time off types job
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Hristimeofftypesjobresponse) String() string {
     o.Entities = []Hristimeofftyperesponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Hristimeofftypesjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Hristimeofftypesjobresponse

    if HristimeofftypesjobresponseMarshalled {
        return []byte("{}"), nil
    }
    HristimeofftypesjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Hristimeofftyperesponse `json:"entities"`
        
        Status string `json:"status"`
        *Alias
    }{

        
        Entities: []Hristimeofftyperesponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

