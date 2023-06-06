package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyimagepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyimagepropertiesDud struct { 
    


    


    

}

// Documentbodyimageproperties
type Documentbodyimageproperties struct { 
    // BackgroundColor - The background color property for the image. The valid values in hex color code representation. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`


    // Align - The align property for the image.
    Align string `json:"align"`


    // Indentation - The indentation property for the image. The valid values in 'em'.
    Indentation float32 `json:"indentation"`

}

// String returns a JSON representation of the model
func (o *Documentbodyimageproperties) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyimageproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyimageproperties

    if DocumentbodyimagepropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyimagepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        BackgroundColor string `json:"backgroundColor"`
        
        Align string `json:"align"`
        
        Indentation float32 `json:"indentation"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

