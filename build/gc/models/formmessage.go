package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormmessageDud struct { 
    


    


    

}

// Formmessage - Form message with title, subtitle, and optional image
type Formmessage struct { 
    // Title - Title of the message
    Title string `json:"title"`


    // Subtitle - Subtitle of the message
    Subtitle string `json:"subtitle"`


    // ImageUrl - URL of the image to display
    ImageUrl string `json:"imageUrl"`

}

// String returns a JSON representation of the model
func (o *Formmessage) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Formmessage) MarshalJSON() ([]byte, error) {
    type Alias Formmessage

    if FormmessageMarshalled {
        return []byte("{}"), nil
    }
    FormmessageMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ImageUrl string `json:"imageUrl"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

