package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchctabuttonstylepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchctabuttonstylepropertiesDud struct { 
    


    


    


    


    

}

// Patchctabuttonstyleproperties
type Patchctabuttonstyleproperties struct { 
    // Color - Color of the text. (eg. #FFFFFF)
    Color string `json:"color"`


    // Font - Font of the text. (eg. Helvetica)
    Font string `json:"font"`


    // FontSize - Font size of the text. (eg. '12')
    FontSize string `json:"fontSize"`


    // TextAlign - Text alignment.
    TextAlign string `json:"textAlign"`


    // BackgroundColor - Background color of the CTA button. (eg. #A04033)
    BackgroundColor string `json:"backgroundColor"`

}

// String returns a JSON representation of the model
func (o *Patchctabuttonstyleproperties) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchctabuttonstyleproperties) MarshalJSON() ([]byte, error) {
    type Alias Patchctabuttonstyleproperties

    if PatchctabuttonstylepropertiesMarshalled {
        return []byte("{}"), nil
    }
    PatchctabuttonstylepropertiesMarshalled = true

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

