package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyactionDud struct { 
    


    

}

// Journeyaction
type Journeyaction struct { 
    // Id - The ID of an action from the Journey System (an action is spawned from an actionMap)
    Id string `json:"id"`


    // ActionMap - Details about the action map from the Journey System which triggered this action
    ActionMap Journeyactionmap `json:"actionMap"`

}

// String returns a JSON representation of the model
func (o *Journeyaction) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyaction) MarshalJSON() ([]byte, error) {
    type Alias Journeyaction

    if JourneyactionMarshalled {
        return []byte("{}"), nil
    }
    JourneyactionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        ActionMap Journeyactionmap `json:"actionMap"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

