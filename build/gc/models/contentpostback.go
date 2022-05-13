package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentpostbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentpostbackDud struct { 
    


    


    

}

// Contentpostback - Postback response object representing the click of a rich media button (Deprecated).
type Contentpostback struct { 
    // Id - An ID assigned to the button response.
    Id string `json:"id"`


    // Text - The response text from the button click.
    Text string `json:"text"`


    // Payload - The response payload associated with the clicked button.
    Payload string `json:"payload"`

}

// String returns a JSON representation of the model
func (o *Contentpostback) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentpostback) MarshalJSON() ([]byte, error) {
    type Alias Contentpostback

    if ContentpostbackMarshalled {
        return []byte("{}"), nil
    }
    ContentpostbackMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

