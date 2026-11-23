package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VideorecordingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VideorecordingsettingsDud struct { 
    


    

}

// Videorecordingsettings - The settings for video chat recording and customer consent
type Videorecordingsettings struct { 
    // Allowed - whether or not video chat recording is allowed. Defaults to false when not set
    Allowed bool `json:"allowed"`


    // ConsentMode - customer consent behaviour when recording is allowed. Required when allowed is true
    ConsentMode string `json:"consentMode"`

}

// String returns a JSON representation of the model
func (o *Videorecordingsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Videorecordingsettings) MarshalJSON() ([]byte, error) {
    type Alias Videorecordingsettings

    if VideorecordingsettingsMarshalled {
        return []byte("{}"), nil
    }
    VideorecordingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Allowed bool `json:"allowed"`
        
        ConsentMode string `json:"consentMode"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

