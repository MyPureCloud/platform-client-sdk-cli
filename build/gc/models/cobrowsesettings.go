package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CobrowsesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CobrowsesettingsDud struct { 
    


    


    

}

// Cobrowsesettings - Settings concerning cobrowse
type Cobrowsesettings struct { 
    // Enabled - Whether or not cobrowse is enabled
    Enabled bool `json:"enabled"`


    // AllowAgentControl - Whether the viewer should have option to request control
    AllowAgentControl bool `json:"allowAgentControl"`


    // MaskSelectors - Mask patterns that will apply to pages being shared
    MaskSelectors []string `json:"maskSelectors"`

}

// String returns a JSON representation of the model
func (o *Cobrowsesettings) String() string {
    
    
    
    
    
    
    
    
    
    
     o.MaskSelectors = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cobrowsesettings) MarshalJSON() ([]byte, error) {
    type Alias Cobrowsesettings

    if CobrowsesettingsMarshalled {
        return []byte("{}"), nil
    }
    CobrowsesettingsMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        AllowAgentControl bool `json:"allowAgentControl"`
        
        MaskSelectors []string `json:"maskSelectors"`
        
        *Alias
    }{
        

        

        

        

        

        
        MaskSelectors: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

