package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RichlinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RichlinkDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Richlink
type Richlink struct { 
    


    // Header - Header for the Rich Link.
    Header Richlinkheader `json:"header"`


    // Footer - Footer text for the Rich Link.
    Footer string `json:"footer"`


    // Text - Text for the body of the Rich Link.
    Text string `json:"text"`


    // Url - Url for the Rich Link.
    Url string `json:"url"`


    // UrlLabel - Label for the URL of the Rich link.
    UrlLabel string `json:"urlLabel"`

}

// String returns a JSON representation of the model
func (o *Richlink) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Richlink) MarshalJSON() ([]byte, error) {
    type Alias Richlink

    if RichlinkMarshalled {
        return []byte("{}"), nil
    }
    RichlinkMarshalled = true

    return json.Marshal(&struct {
        
        Header Richlinkheader `json:"header"`
        
        Footer string `json:"footer"`
        
        Text string `json:"text"`
        
        Url string `json:"url"`
        
        UrlLabel string `json:"urlLabel"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

