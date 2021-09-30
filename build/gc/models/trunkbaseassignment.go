package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkbaseassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkbaseassignmentDud struct { 
    


    

}

// Trunkbaseassignment
type Trunkbaseassignment struct { 
    // Family - The address family to use with the trunk base settings. 2=IPv4, 23=IPv6
    Family int `json:"family"`


    // TrunkBase - A trunk base settings reference.
    TrunkBase Trunkbase `json:"trunkBase"`

}

// String returns a JSON representation of the model
func (o *Trunkbaseassignment) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkbaseassignment) MarshalJSON() ([]byte, error) {
    type Alias Trunkbaseassignment

    if TrunkbaseassignmentMarshalled {
        return []byte("{}"), nil
    }
    TrunkbaseassignmentMarshalled = true

    return json.Marshal(&struct { 
        Family int `json:"family"`
        
        TrunkBase Trunkbase `json:"trunkBase"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

