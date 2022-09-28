package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VipmediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VipmediasettingsDud struct { 
    


    

}

// Vipmediasettings
type Vipmediasettings struct { 
    // Enabled - Toggle that enables VIP experience for this feature.
    Enabled bool `json:"enabled"`


    // SkipOwnershipTime - Toggle that enables this media type to fallback immediately to the configured VIP Backup.
    SkipOwnershipTime bool `json:"skipOwnershipTime"`

}

// String returns a JSON representation of the model
func (o *Vipmediasettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Vipmediasettings) MarshalJSON() ([]byte, error) {
    type Alias Vipmediasettings

    if VipmediasettingsMarshalled {
        return []byte("{}"), nil
    }
    VipmediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        SkipOwnershipTime bool `json:"skipOwnershipTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

