package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentactionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentactionsDud struct { 
    


    


    

}

// Contentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Contentactions struct { 
    // Url - A URL of a web page to direct the user to.
    Url string `json:"url"`


    // UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
    UrlTarget string `json:"urlTarget"`


    // Textback - Text to be returned as the payload from a ButtonResponse when a button is clicked. The textback and title are a combination which will have to be unique across each card and carousel in order to determine which button was clicked in that card or carousel.
    Textback string `json:"textback"`

}

// String returns a JSON representation of the model
func (o *Contentactions) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentactions) MarshalJSON() ([]byte, error) {
    type Alias Contentactions

    if ContentactionsMarshalled {
        return []byte("{}"), nil
    }
    ContentactionsMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        UrlTarget string `json:"urlTarget"`
        
        Textback string `json:"textback"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

