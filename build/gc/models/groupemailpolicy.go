package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupemailpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupemailpolicyDud struct { 
    


    

}

// Groupemailpolicy
type Groupemailpolicy struct { 
    // EmailMembers
    EmailMembers bool `json:"emailMembers"`


    // EmailGroup
    EmailGroup bool `json:"emailGroup"`

}

// String returns a JSON representation of the model
func (o *Groupemailpolicy) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupemailpolicy) MarshalJSON() ([]byte, error) {
    type Alias Groupemailpolicy

    if GroupemailpolicyMarshalled {
        return []byte("{}"), nil
    }
    GroupemailpolicyMarshalled = true

    return json.Marshal(&struct {
        
        EmailMembers bool `json:"emailMembers"`
        
        EmailGroup bool `json:"emailGroup"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

