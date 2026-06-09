package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CarouselcardMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CarouselcardDud struct { 
    


    


    

}

// Carouselcard
type Carouselcard struct { 
    // Header - Header for whatsApp carousel card
    Header Messageheader `json:"header"`


    // Buttons - List of buttons to be included in the whatsApp carousel messages template
    Buttons []Button `json:"buttons"`


    // CardBodyText - Optional card body text to be included in the whatsApp carousel messages template
    CardBodyText Cardbodytext `json:"cardBodyText"`

}

// String returns a JSON representation of the model
func (o *Carouselcard) String() string {
    
     o.Buttons = []Button{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Carouselcard) MarshalJSON() ([]byte, error) {
    type Alias Carouselcard

    if CarouselcardMarshalled {
        return []byte("{}"), nil
    }
    CarouselcardMarshalled = true

    return json.Marshal(&struct {
        
        Header Messageheader `json:"header"`
        
        Buttons []Button `json:"buttons"`
        
        CardBodyText Cardbodytext `json:"cardBodyText"`
        *Alias
    }{

        


        
        Buttons: []Button{{}},
        


        

        Alias: (*Alias)(u),
    })
}

