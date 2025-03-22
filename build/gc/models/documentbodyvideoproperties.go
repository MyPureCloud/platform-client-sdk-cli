package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyvideopropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyvideopropertiesDud struct { 
    


    


    


    


    

}

// Documentbodyvideoproperties
type Documentbodyvideoproperties struct { 
    // BackgroundColor - The background color for the video. The valid values in hex color code representation. For example black color - #000000
    BackgroundColor string `json:"backgroundColor"`


    // Align - The align type for the video.
    Align string `json:"align"`


    // Indentation - The indentation for the video. The valid values in 'em'.
    Indentation float32 `json:"indentation"`


    // Width - The width of the video in the specified unit.
    Width Documentelementlength `json:"width"`


    // Height - The height of the video in the specified unit.
    Height Documentelementlength `json:"height"`

}

// String returns a JSON representation of the model
func (o *Documentbodyvideoproperties) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyvideoproperties) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyvideoproperties

    if DocumentbodyvideopropertiesMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyvideopropertiesMarshalled = true

    return json.Marshal(&struct {
        
        BackgroundColor string `json:"backgroundColor"`
        
        Align string `json:"align"`
        
        Indentation float32 `json:"indentation"`
        
        Width Documentelementlength `json:"width"`
        
        Height Documentelementlength `json:"height"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

