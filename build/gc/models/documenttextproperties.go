package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumenttextpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumenttextpropertiesDud struct { 
    


    


    

}

// Documenttextproperties
type Documenttextproperties struct { 
    // FontSize - The font size for the text. The valid values in 'em'.
    FontSize string `json:"fontSize"`


    // TextColor - The text color for the text. The valid values in hex color code representation. For example black color - #000000
    TextColor string `json:"textColor"`


    // BackgroundColor - The background color for the text. The valid values in hex color code representation. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`

}

// String returns a JSON representation of the model
func (o *Documenttextproperties) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documenttextproperties) MarshalJSON() ([]byte, error) {
    type Alias Documenttextproperties

    if DocumenttextpropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumenttextpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        FontSize string `json:"fontSize"`
        
        TextColor string `json:"textColor"`
        
        BackgroundColor string `json:"backgroundColor"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

