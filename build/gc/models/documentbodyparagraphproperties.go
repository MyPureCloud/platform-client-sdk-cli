package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyparagraphpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyparagraphpropertiesDud struct { 
    


    


    


    


    


    

}

// Documentbodyparagraphproperties
type Documentbodyparagraphproperties struct { 
    // FontSize - The font size for the paragraph. The valid values in 'em'.
    FontSize string `json:"fontSize"`


    // FontType - The font type for the paragraph.
    FontType string `json:"fontType"`


    // TextColor - The text color for the paragraph. The valid values in hex color code representation. For example black color - #000000
    TextColor string `json:"textColor"`


    // BackgroundColor - The background color for the paragraph. The valid values in hex color code representation. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`


    // Align - The align type for the paragraph.
    Align string `json:"align"`


    // Indentation - The indentation color for the paragraph. The valid values in 'em'.
    Indentation float32 `json:"indentation"`

}

// String returns a JSON representation of the model
func (o *Documentbodyparagraphproperties) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyparagraphproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyparagraphproperties

    if DocumentbodyparagraphpropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyparagraphpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        FontSize string `json:"fontSize"`
        
        FontType string `json:"fontType"`
        
        TextColor string `json:"textColor"`
        
        BackgroundColor string `json:"backgroundColor"`
        
        Align string `json:"align"`
        
        Indentation float32 `json:"indentation"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

