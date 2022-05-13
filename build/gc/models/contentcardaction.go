package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentcardactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentcardactionDud struct { 
    


    


    


    

}

// Contentcardaction - A card action that a user can take.
type Contentcardaction struct { 
    // VarType - Describes the type of action.
    VarType string `json:"type"`


    // Text - The response text from the button click.
    Text string `json:"text"`


    // Payload - Text to be returned as the payload from a ButtonResponse when a button is clicked. The payload and text are a combination which will have to be unique across each card and carousel in order to determine which button was clicked in that card or carousel.
    Payload string `json:"payload"`


    // Url - A URL of a web page to direct the user to.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Contentcardaction) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentcardaction) MarshalJSON() ([]byte, error) {
    type Alias Contentcardaction

    if ContentcardactionMarshalled {
        return []byte("{}"), nil
    }
    ContentcardactionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

