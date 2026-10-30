package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrgsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrgsettingsDud struct { 
    


    

}

// Orgsettings
type Orgsettings struct { 
    // ContentSearchEnabled - If true, enables full-text search of interaction transcripts
    ContentSearchEnabled bool `json:"contentSearchEnabled"`


    // UseDivisionedPermissions - If true, the results of searches for users are restricted to the divisions the user has directory:user:view permissions in
    UseDivisionedPermissions bool `json:"useDivisionedPermissions"`

}

// String returns a JSON representation of the model
func (o *Orgsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Orgsettings) MarshalJSON() ([]byte, error) {
    type Alias Orgsettings

    if OrgsettingsMarshalled {
        return []byte("{}"), nil
    }
    OrgsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ContentSearchEnabled bool `json:"contentSearchEnabled"`
        
        UseDivisionedPermissions bool `json:"useDivisionedPermissions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

