package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitDud struct { 
    Key string `json:"key"`


    Value float64 `json:"value"`

}

// Limit
type Limit struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Limit) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limit) MarshalJSON() ([]byte, error) {
    type Alias Limit

    if LimitMarshalled {
        return []byte("{}"), nil
    }
    LimitMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

