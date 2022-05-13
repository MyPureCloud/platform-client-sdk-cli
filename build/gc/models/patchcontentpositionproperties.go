package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcontentpositionpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcontentpositionpropertiesDud struct { 
    


    


    


    

}

// Patchcontentpositionproperties
type Patchcontentpositionproperties struct { 
    // Top - Top positioning offset.
    Top string `json:"top"`


    // Bottom - Bottom positioning offset.
    Bottom string `json:"bottom"`


    // Left - Left positioning offset.
    Left string `json:"left"`


    // Right - Right positioning offset.
    Right string `json:"right"`

}

// String returns a JSON representation of the model
func (o *Patchcontentpositionproperties) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcontentpositionproperties) MarshalJSON() ([]byte, error) {
    type Alias Patchcontentpositionproperties

    if PatchcontentpositionpropertiesMarshalled {
        return []byte("{}"), nil
    }
    PatchcontentpositionpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Top string `json:"top"`
        
        Bottom string `json:"bottom"`
        
        Left string `json:"left"`
        
        Right string `json:"right"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

