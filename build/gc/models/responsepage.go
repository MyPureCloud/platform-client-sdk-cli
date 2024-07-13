package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsepageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsepageDud struct { 
    


    


    


    


    


    


    


    


    

}

// Responsepage
type Responsepage struct { 
    // Url - The page URL.
    Url string `json:"url"`


    // Title - Title of the page.
    Title string `json:"title"`


    // Domain - Domain of the page's URL.
    Domain string `json:"domain"`


    // Fragment - Fragment or hash of the page's URL.
    Fragment string `json:"fragment"`


    // Hostname - Hostname of the page's URL.
    Hostname string `json:"hostname"`


    // Keywords - Keywords from the HTML <meta> tag of the page.
    Keywords string `json:"keywords"`


    // Lang - ISO 639-1 language code for the page as defined in the <html> tag.
    Lang string `json:"lang"`


    // Pathname - Path name of the page for the event.
    Pathname string `json:"pathname"`


    // QueryString - Query string that is passed to the page in the current event.
    QueryString string `json:"queryString"`

}

// String returns a JSON representation of the model
func (o *Responsepage) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsepage) MarshalJSON() ([]byte, error) {
    type Alias Responsepage

    if ResponsepageMarshalled {
        return []byte("{}"), nil
    }
    ResponsepageMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Title string `json:"title"`
        
        Domain string `json:"domain"`
        
        Fragment string `json:"fragment"`
        
        Hostname string `json:"hostname"`
        
        Keywords string `json:"keywords"`
        
        Lang string `json:"lang"`
        
        Pathname string `json:"pathname"`
        
        QueryString string `json:"queryString"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

