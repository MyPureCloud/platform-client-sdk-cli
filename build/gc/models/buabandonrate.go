package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuabandonrateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuabandonrateDud struct { 
    


    

}

// Buabandonrate
type Buabandonrate struct { 
    // Include - Whether to include abandon rate in the associated configuration
    Include bool `json:"include"`


    // Percent - Abandon rate percent goal. Required if include == true
    Percent int `json:"percent"`

}

// String returns a JSON representation of the model
func (o *Buabandonrate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buabandonrate) MarshalJSON() ([]byte, error) {
    type Alias Buabandonrate

    if BuabandonrateMarshalled {
        return []byte("{}"), nil
    }
    BuabandonrateMarshalled = true

    return json.Marshal(&struct {
        
        Include bool `json:"include"`
        
        Percent int `json:"percent"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

