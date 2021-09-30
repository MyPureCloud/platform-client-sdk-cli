package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivealertcountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivealertcountDud struct { 
    

}

// Activealertcount
type Activealertcount struct { 
    // Count - The count of active alerts for a user.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Activealertcount) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activealertcount) MarshalJSON() ([]byte, error) {
    type Alias Activealertcount

    if ActivealertcountMarshalled {
        return []byte("{}"), nil
    }
    ActivealertcountMarshalled = true

    return json.Marshal(&struct { 
        Count int `json:"count"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

