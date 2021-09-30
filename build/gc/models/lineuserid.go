package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LineuseridMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LineuseridDud struct { 
    

}

// Lineuserid - Channel-specific User ID for Line accounts
type Lineuserid struct { 
    // UserId - The unique channel-specific userId for the user
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Lineuserid) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lineuserid) MarshalJSON() ([]byte, error) {
    type Alias Lineuserid

    if LineuseridMarshalled {
        return []byte("{}"), nil
    }
    LineuseridMarshalled = true

    return json.Marshal(&struct { 
        UserId string `json:"userId"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

