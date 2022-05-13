package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentofferstylepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentofferstylepropertiesDud struct { 
    


    


    

}

// Contentofferstyleproperties
type Contentofferstyleproperties struct { 
    // Padding - Padding of the offer. (eg. 10px)
    Padding string `json:"padding"`


    // Color - Text color of the offer. (eg. #FF0000)
    Color string `json:"color"`


    // BackgroundColor - Background color of the offer. (eg. #000000)
    BackgroundColor string `json:"backgroundColor"`

}

// String returns a JSON representation of the model
func (o *Contentofferstyleproperties) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentofferstyleproperties) MarshalJSON() ([]byte, error) {
    type Alias Contentofferstyleproperties

    if ContentofferstylepropertiesMarshalled {
        return []byte("{}"), nil
    }
    ContentofferstylepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Padding string `json:"padding"`
        
        Color string `json:"color"`
        
        BackgroundColor string `json:"backgroundColor"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

