package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchclosebuttonstylepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchclosebuttonstylepropertiesDud struct { 
    


    

}

// Patchclosebuttonstyleproperties
type Patchclosebuttonstyleproperties struct { 
    // Color - Color of button. (eg. #FF0000)
    Color string `json:"color"`


    // Opacity - Opacity of button.
    Opacity float32 `json:"opacity"`

}

// String returns a JSON representation of the model
func (o *Patchclosebuttonstyleproperties) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchclosebuttonstyleproperties) MarshalJSON() ([]byte, error) {
    type Alias Patchclosebuttonstyleproperties

    if PatchclosebuttonstylepropertiesMarshalled {
        return []byte("{}"), nil
    }
    PatchclosebuttonstylepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Color string `json:"color"`
        
        Opacity float32 `json:"opacity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

