package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestpageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestpageDud struct { 
    


    


    


    

}

// Requestpage
type Requestpage struct { 
    // Url - The page URL.
    Url string `json:"url"`


    // Title - Title of the page.
    Title string `json:"title"`


    // Keywords - Keywords from the HTML <meta> tag of the page.
    Keywords string `json:"keywords"`


    // Lang - ISO 639-1 language code for the page as defined in the <html> tag.
    Lang string `json:"lang"`

}

// String returns a JSON representation of the model
func (o *Requestpage) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestpage) MarshalJSON() ([]byte, error) {
    type Alias Requestpage

    if RequestpageMarshalled {
        return []byte("{}"), nil
    }
    RequestpageMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Title string `json:"title"`
        
        Keywords string `json:"keywords"`
        
        Lang string `json:"lang"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

