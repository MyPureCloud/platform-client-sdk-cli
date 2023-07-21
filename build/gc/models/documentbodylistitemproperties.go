package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodylistitempropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodylistitempropertiesDud struct { 
    


    


    


    


    


    


    


    

}

// Documentbodylistitemproperties
type Documentbodylistitemproperties struct { 
    // BackgroundColor - The background color for the list item. The valid values in hex color code representation. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`


    // Align - The align type for the list item.
    Align string `json:"align"`


    // Indentation - The indentation property for the list item. The valid values in 'em'.
    Indentation float32 `json:"indentation"`


    // FontSize - The font size for the list item. The valid values in 'em'.
    FontSize string `json:"fontSize"`


    // FontType - The font type for the list item.
    FontType string `json:"fontType"`


    // TextColor - The text color for the list item. The valid values in hex color code representation. For example black color - #000000
    TextColor string `json:"textColor"`


    // UnorderedType - The type of icon for the unordered list.
    UnorderedType string `json:"unorderedType"`


    // OrderedType - The type of icon for the ordered list.
    OrderedType string `json:"orderedType"`

}

// String returns a JSON representation of the model
func (o *Documentbodylistitemproperties) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodylistitemproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodylistitemproperties

    if DocumentbodylistitempropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodylistitempropertiesMarshalled = true

    return json.Marshal(&struct {
        
        BackgroundColor string `json:"backgroundColor"`
        
        Align string `json:"align"`
        
        Indentation float32 `json:"indentation"`
        
        FontSize string `json:"fontSize"`
        
        FontType string `json:"fontType"`
        
        TextColor string `json:"textColor"`
        
        UnorderedType string `json:"unorderedType"`
        
        OrderedType string `json:"orderedType"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

