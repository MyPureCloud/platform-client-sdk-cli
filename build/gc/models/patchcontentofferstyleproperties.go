package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcontentofferstylepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcontentofferstylepropertiesDud struct { 
    


    


    

}

// Patchcontentofferstyleproperties
type Patchcontentofferstyleproperties struct { 
    // Padding - Padding of the offer. (eg. 10px)
    Padding string `json:"padding"`


    // Color - Text color of the offer. (eg. #FF0000)
    Color string `json:"color"`


    // BackgroundColor - Background color of the offer. (eg. #000000)
    BackgroundColor string `json:"backgroundColor"`

}

// String returns a JSON representation of the model
func (o *Patchcontentofferstyleproperties) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcontentofferstyleproperties) MarshalJSON() ([]byte, error) {
    type Alias Patchcontentofferstyleproperties

    if PatchcontentofferstylepropertiesMarshalled {
        return []byte("{}"), nil
    }
    PatchcontentofferstylepropertiesMarshalled = true

    return json.Marshal(&struct { 
        Padding string `json:"padding"`
        
        Color string `json:"color"`
        
        BackgroundColor string `json:"backgroundColor"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

