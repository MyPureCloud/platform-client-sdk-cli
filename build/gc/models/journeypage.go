package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneypageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneypageDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Journeypage
type Journeypage struct { 
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


    // Keywords - Keywords from the HTML {@code <meta>} tag of the page.
    Keywords string `json:"keywords"`


    // Lang - ISO 639-1 language code for the page as defined in the {@code <html>} tag.
    Lang string `json:"lang"`


    // Pathname - Path name of the page for the event.
    Pathname string `json:"pathname"`


    // QueryString - Query string that is passed to the page in the current event.
    QueryString string `json:"queryString"`


    // Breadcrumb - Hierarchy of the current page in relation to the website's structure.
    Breadcrumb []string `json:"breadcrumb"`

}

// String returns a JSON representation of the model
func (o *Journeypage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Breadcrumb = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeypage) MarshalJSON() ([]byte, error) {
    type Alias Journeypage

    if JourneypageMarshalled {
        return []byte("{}"), nil
    }
    JourneypageMarshalled = true

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
        
        Breadcrumb []string `json:"breadcrumb"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Breadcrumb: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

