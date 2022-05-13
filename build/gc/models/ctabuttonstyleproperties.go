package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CtabuttonstylepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CtabuttonstylepropertiesDud struct { 
    


    


    


    


    

}

// Ctabuttonstyleproperties
type Ctabuttonstyleproperties struct { 
    // Color - Color of the text. (eg. #FFFFFF)
    Color string `json:"color"`


    // Font - Font of the text. (eg. Helvetica)
    Font string `json:"font"`


    // FontSize - Font size of the text. (eg. '12')
    FontSize string `json:"fontSize"`


    // TextAlign - Text alignment.
    TextAlign string `json:"textAlign"`


    // BackgroundColor - Background color of the CTA button. (eg. #FF0000)
    BackgroundColor string `json:"backgroundColor"`

}

// String returns a JSON representation of the model
func (o *Ctabuttonstyleproperties) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ctabuttonstyleproperties) MarshalJSON() ([]byte, error) {
    type Alias Ctabuttonstyleproperties

    if CtabuttonstylepropertiesMarshalled {
        return []byte("{}"), nil
    }
    CtabuttonstylepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Color string `json:"color"`
        
        Font string `json:"font"`
        
        FontSize string `json:"fontSize"`
        
        TextAlign string `json:"textAlign"`
        
        BackgroundColor string `json:"backgroundColor"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

