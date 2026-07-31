package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailrefinementconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailrefinementconfigDud struct { 
    


    

}

// Emailrefinementconfig
type Emailrefinementconfig struct { 
    // Enabled - Email refinement is enabled.
    Enabled bool `json:"enabled"`


    // RefinementSetting - Configured refinement setting object.
    RefinementSetting Refinementsettingentity `json:"refinementSetting"`

}

// String returns a JSON representation of the model
func (o *Emailrefinementconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailrefinementconfig) MarshalJSON() ([]byte, error) {
    type Alias Emailrefinementconfig

    if EmailrefinementconfigMarshalled {
        return []byte("{}"), nil
    }
    EmailrefinementconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        RefinementSetting Refinementsettingentity `json:"refinementSetting"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

