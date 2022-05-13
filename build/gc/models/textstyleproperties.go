package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextstylepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextstylepropertiesDud struct { 
    


    


    


    

}

// Textstyleproperties
type Textstyleproperties struct { 
    // Color - Color of the text. (eg. #FFFFFF)
    Color string `json:"color"`


    // Font - Font of the text. (eg. Helvetica)
    Font string `json:"font"`


    // FontSize - Font size of the text. (eg. '12')
    FontSize string `json:"fontSize"`


    // TextAlign - Text alignment.
    TextAlign string `json:"textAlign"`

}

// String returns a JSON representation of the model
func (o *Textstyleproperties) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textstyleproperties) MarshalJSON() ([]byte, error) {
    type Alias Textstyleproperties

    if TextstylepropertiesMarshalled {
        return []byte("{}"), nil
    }
    TextstylepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Color string `json:"color"`
        
        Font string `json:"font"`
        
        FontSize string `json:"fontSize"`
        
        TextAlign string `json:"textAlign"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

