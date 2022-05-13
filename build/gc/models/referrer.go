package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReferrerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReferrerDud struct { 
    


    


    


    


    


    


    


    


    

}

// Referrer
type Referrer struct { 
    // Url - Referrer URL.
    Url string `json:"url"`


    // Domain - Referrer URL domain.
    Domain string `json:"domain"`


    // Hostname - Referrer URL hostname.
    Hostname string `json:"hostname"`


    // Keywords - Referrer keywords.
    Keywords string `json:"keywords"`


    // Pathname - Referrer URL pathname.
    Pathname string `json:"pathname"`


    // QueryString - Referrer URL querystring.
    QueryString string `json:"queryString"`


    // Fragment - Referrer URL fragment.
    Fragment string `json:"fragment"`


    // Name - Name of referrer (e.g. Yahoo!, Google, InfoSpace).
    Name string `json:"name"`


    // Medium - Type of referrer (e.g. search, social).
    Medium string `json:"medium"`

}

// String returns a JSON representation of the model
func (o *Referrer) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Referrer) MarshalJSON() ([]byte, error) {
    type Alias Referrer

    if ReferrerMarshalled {
        return []byte("{}"), nil
    }
    ReferrerMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Domain string `json:"domain"`
        
        Hostname string `json:"hostname"`
        
        Keywords string `json:"keywords"`
        
        Pathname string `json:"pathname"`
        
        QueryString string `json:"queryString"`
        
        Fragment string `json:"fragment"`
        
        Name string `json:"name"`
        
        Medium string `json:"medium"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

