package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityadherencepresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityadherencepresenceDud struct { 
    


    


    

}

// Useractivityadherencepresence
type Useractivityadherencepresence struct { 
    // PresenceDefinition - The current presence definition for the user
    PresenceDefinition Useractivitypresencedefinition `json:"presenceDefinition"`


    // PresenceMessage - The free-form presence message the user has set, if any
    PresenceMessage string `json:"presenceMessage"`


    // ModifiedDate - The date the presence was last modified, in ISO-8601 format
    ModifiedDate time.Time `json:"modifiedDate"`

}

// String returns a JSON representation of the model
func (o *Useractivityadherencepresence) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityadherencepresence) MarshalJSON() ([]byte, error) {
    type Alias Useractivityadherencepresence

    if UseractivityadherencepresenceMarshalled {
        return []byte("{}"), nil
    }
    UseractivityadherencepresenceMarshalled = true

    return json.Marshal(&struct {
        
        PresenceDefinition Useractivitypresencedefinition `json:"presenceDefinition"`
        
        PresenceMessage string `json:"presenceMessage"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

