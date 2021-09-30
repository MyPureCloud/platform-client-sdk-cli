package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchactiontargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchactiontargetDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Patchactiontarget
type Patchactiontarget struct { 
    


    // Name
    Name string `json:"name"`


    // ServiceLevel - Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
    ServiceLevel Servicelevel `json:"serviceLevel"`


    // ShortAbandonThreshold - Indicates the non-default short abandon threshold
    ShortAbandonThreshold int `json:"shortAbandonThreshold"`


    

}

// String returns a JSON representation of the model
func (o *Patchactiontarget) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchactiontarget) MarshalJSON() ([]byte, error) {
    type Alias Patchactiontarget

    if PatchactiontargetMarshalled {
        return []byte("{}"), nil
    }
    PatchactiontargetMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ServiceLevel Servicelevel `json:"serviceLevel"`
        
        ShortAbandonThreshold int `json:"shortAbandonThreshold"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

