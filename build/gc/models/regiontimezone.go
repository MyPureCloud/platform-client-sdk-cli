package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegiontimezoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegiontimezoneDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Regiontimezone
type Regiontimezone struct { 
    


    // Name
    Name string `json:"name"`


    // Offset
    Offset int `json:"offset"`


    

}

// String returns a JSON representation of the model
func (o *Regiontimezone) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Regiontimezone) MarshalJSON() ([]byte, error) {
    type Alias Regiontimezone

    if RegiontimezoneMarshalled {
        return []byte("{}"), nil
    }
    RegiontimezoneMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Offset int `json:"offset"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

