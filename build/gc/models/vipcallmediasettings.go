package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VipcallmediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VipcallmediasettingsDud struct { 
    


    

}

// Vipcallmediasettings
type Vipcallmediasettings struct { 
    // Enabled - Toggle that enables VIP experience for this feature.
    Enabled bool `json:"enabled"`


    // SkipOwnershipTime - Toggle that enables this media type to fallback immediately to the configured VIP Backup.
    SkipOwnershipTime bool `json:"skipOwnershipTime"`

}

// String returns a JSON representation of the model
func (o *Vipcallmediasettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Vipcallmediasettings) MarshalJSON() ([]byte, error) {
    type Alias Vipcallmediasettings

    if VipcallmediasettingsMarshalled {
        return []byte("{}"), nil
    }
    VipcallmediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        SkipOwnershipTime bool `json:"skipOwnershipTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

