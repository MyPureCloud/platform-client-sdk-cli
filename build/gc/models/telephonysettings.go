package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TelephonysettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TelephonysettingsDud struct { 
    


    

}

// Telephonysettings
type Telephonysettings struct { 
    // PersistentConnectionRequired - Determines if a persistent connection will be established before routing calls to agents
    PersistentConnectionRequired bool `json:"persistentConnectionRequired"`


    // BlockCallerIdAccessCode - Configurable code that should be exempt from caller ID manipulation (e.g. *67 for anonymized calling)
    BlockCallerIdAccessCode string `json:"blockCallerIdAccessCode"`

}

// String returns a JSON representation of the model
func (o *Telephonysettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Telephonysettings) MarshalJSON() ([]byte, error) {
    type Alias Telephonysettings

    if TelephonysettingsMarshalled {
        return []byte("{}"), nil
    }
    TelephonysettingsMarshalled = true

    return json.Marshal(&struct {
        
        PersistentConnectionRequired bool `json:"persistentConnectionRequired"`
        
        BlockCallerIdAccessCode string `json:"blockCallerIdAccessCode"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

