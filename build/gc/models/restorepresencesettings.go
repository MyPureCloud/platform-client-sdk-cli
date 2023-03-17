package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RestorepresencesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RestorepresencesettingsDud struct { 
    


    


    

}

// Restorepresencesettings
type Restorepresencesettings struct { 
    // Enabled - Whether the restore presence feature is enabled
    Enabled bool `json:"enabled"`


    // RestoreTimeMilliseconds - How many milliseconds the presence will be restored within
    RestoreTimeMilliseconds int `json:"restoreTimeMilliseconds"`


    // RestoreOnQueueEnabled - Whether the ON_QUEUE presence will be restored
    RestoreOnQueueEnabled bool `json:"restoreOnQueueEnabled"`

}

// String returns a JSON representation of the model
func (o *Restorepresencesettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Restorepresencesettings) MarshalJSON() ([]byte, error) {
    type Alias Restorepresencesettings

    if RestorepresencesettingsMarshalled {
        return []byte("{}"), nil
    }
    RestorepresencesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        RestoreTimeMilliseconds int `json:"restoreTimeMilliseconds"`
        
        RestoreOnQueueEnabled bool `json:"restoreOnQueueEnabled"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

