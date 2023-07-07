package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablepropertiesDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Documentbodytableproperties
type Documentbodytableproperties struct { 
    // Width - The width for the table. The valid values in 'em'.
    Width float32 `json:"width"`


    // Height - The height for the table.
    Height float32 `json:"height"`


    // CellSpacing - The cell spacing for the table. The valid values in 'em'.
    CellSpacing float32 `json:"cellSpacing"`


    // CellPadding - The cell padding for the table. The valid values in 'em'.
    CellPadding float32 `json:"cellPadding"`


    // BorderWidth - The border width for the table. The valid values in 'em'
    BorderWidth float32 `json:"borderWidth"`


    // Alignment - The alignment for the table.
    Alignment string `json:"alignment"`


    // BorderStyle - The border style for the table.
    BorderStyle string `json:"borderStyle"`


    // BorderColor - The border color for the table. The valid values in hex color code representation. For example black color - #000000
    BorderColor string `json:"borderColor"`


    // BackgroundColor - The background color for the table. The valid values in hex color code representation. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`


    // Caption - The caption for the table. The valid values in hex color code representation. For example black color - #000000
    Caption Documentbodytablecaptionblock `json:"caption"`

}

// String returns a JSON representation of the model
func (o *Documentbodytableproperties) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytableproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytableproperties

    if DocumentbodytablepropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Width float32 `json:"width"`
        
        Height float32 `json:"height"`
        
        CellSpacing float32 `json:"cellSpacing"`
        
        CellPadding float32 `json:"cellPadding"`
        
        BorderWidth float32 `json:"borderWidth"`
        
        Alignment string `json:"alignment"`
        
        BorderStyle string `json:"borderStyle"`
        
        BorderColor string `json:"borderColor"`
        
        BackgroundColor string `json:"backgroundColor"`
        
        Caption Documentbodytablecaptionblock `json:"caption"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

