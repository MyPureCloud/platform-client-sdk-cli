package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentpositionpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentpositionpropertiesDud struct { 
    


    


    


    

}

// Contentpositionproperties
type Contentpositionproperties struct { 
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
func (o *Contentpositionproperties) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentpositionproperties) MarshalJSON() ([]byte, error) {
    type Alias Contentpositionproperties

    if ContentpositionpropertiesMarshalled {
        return []byte("{}"), nil
    }
    ContentpositionpropertiesMarshalled = true

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

