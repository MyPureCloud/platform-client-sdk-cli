package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectrateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectrateDud struct { 
    Attempts int `json:"attempts"`


    Connects int `json:"connects"`


    ConnectRatio float64 `json:"connectRatio"`

}

// Connectrate
type Connectrate struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Connectrate) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectrate) MarshalJSON() ([]byte, error) {
    type Alias Connectrate

    if ConnectrateMarshalled {
        return []byte("{}"), nil
    }
    ConnectrateMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

