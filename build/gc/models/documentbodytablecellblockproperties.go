package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablecellblockpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablecellblockpropertiesDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Documentbodytablecellblockproperties
type Documentbodytablecellblockproperties struct { 
    // CellType - The type of the table cell.
    CellType string `json:"cellType"`


    // Width - The width of the table cell converted to em unit.
    Width float32 `json:"width"`


    // Height - The height for the table cell.
    Height float32 `json:"height"`


    // HorizontalAlign - The horizontal alignment for the table cell.
    HorizontalAlign string `json:"horizontalAlign"`


    // VerticalAlign - The vertical alignment for the table cell.
    VerticalAlign string `json:"verticalAlign"`


    // BorderWidth - The border width for the table cell. The valid values in 'em'
    BorderWidth float32 `json:"borderWidth"`


    // BorderStyle - The border style for the table cell.
    BorderStyle string `json:"borderStyle"`


    // BorderColor - The border color for the table cell. For example black color - #000000
    BorderColor string `json:"borderColor"`


    // BackgroundColor - The background color for the table cell. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`


    // Scope - The scope for the table cell.
    Scope string `json:"scope"`


    // ColSpan - The colSpan for the table cell.
    ColSpan int `json:"colSpan"`


    // RowSpan - The rowSpan for the table cell.
    RowSpan int `json:"rowSpan"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablecellblockproperties) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablecellblockproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablecellblockproperties

    if DocumentbodytablecellblockpropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablecellblockpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        CellType string `json:"cellType"`
        
        Width float32 `json:"width"`
        
        Height float32 `json:"height"`
        
        HorizontalAlign string `json:"horizontalAlign"`
        
        VerticalAlign string `json:"verticalAlign"`
        
        BorderWidth float32 `json:"borderWidth"`
        
        BorderStyle string `json:"borderStyle"`
        
        BorderColor string `json:"borderColor"`
        
        BackgroundColor string `json:"backgroundColor"`
        
        Scope string `json:"scope"`
        
        ColSpan int `json:"colSpan"`
        
        RowSpan int `json:"rowSpan"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

