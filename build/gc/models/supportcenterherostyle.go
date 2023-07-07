package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterherostyleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterherostyleDud struct { 
    


    


    

}

// Supportcenterherostyle
type Supportcenterherostyle struct { 
    // BackgroundColor - Background color for hero section, in hexadecimal format, eg #ffffff
    BackgroundColor string `json:"backgroundColor"`


    // TextColor - Text color for hero section, in hexadecimal format, eg #ffffff
    TextColor string `json:"textColor"`


    // Image - Background image for hero section
    Image Supportcenterimage `json:"image"`

}

// String returns a JSON representation of the model
func (o *Supportcenterherostyle) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterherostyle) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterherostyle

    if SupportcenterherostyleMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterherostyleMarshalled = true

    return json.Marshal(&struct {
        
        BackgroundColor string `json:"backgroundColor"`
        
        TextColor string `json:"textColor"`
        
        Image Supportcenterimage `json:"image"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

