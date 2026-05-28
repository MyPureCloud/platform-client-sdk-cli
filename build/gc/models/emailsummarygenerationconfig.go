package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailsummarygenerationconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailsummarygenerationconfigDud struct { 
    


    

}

// Emailsummarygenerationconfig
type Emailsummarygenerationconfig struct { 
    // Enabled - Email summary is enabled.
    Enabled bool `json:"enabled"`


    // SummarySetting - Configured summary setting object.
    SummarySetting Emailsummarysettingsentity `json:"summarySetting"`

}

// String returns a JSON representation of the model
func (o *Emailsummarygenerationconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailsummarygenerationconfig) MarshalJSON() ([]byte, error) {
    type Alias Emailsummarygenerationconfig

    if EmailsummarygenerationconfigMarshalled {
        return []byte("{}"), nil
    }
    EmailsummarygenerationconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        SummarySetting Emailsummarysettingsentity `json:"summarySetting"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

