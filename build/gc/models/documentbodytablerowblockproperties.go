package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablerowblockpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablerowblockpropertiesDud struct { 
    


    


    


    


    


    

}

// Documentbodytablerowblockproperties
type Documentbodytablerowblockproperties struct { 
    // RowType - The type of the table row.
    RowType string `json:"rowType"`


    // Alignment - The alignment for the table row.
    Alignment string `json:"alignment"`


    // Height - The height for the table row.
    Height float32 `json:"height"`


    // BorderStyle - The border style for the table row.
    BorderStyle string `json:"borderStyle"`


    // BorderColor - The border color for the table row. For example black color - #000000
    BorderColor string `json:"borderColor"`


    // BackgroundColor - The background color for the table row. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablerowblockproperties) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablerowblockproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablerowblockproperties

    if DocumentbodytablerowblockpropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablerowblockpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        RowType string `json:"rowType"`
        
        Alignment string `json:"alignment"`
        
        Height float32 `json:"height"`
        
        BorderStyle string `json:"borderStyle"`
        
        BorderColor string `json:"borderColor"`
        
        BackgroundColor string `json:"backgroundColor"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

