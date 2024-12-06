package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterdataingestionrulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterdataingestionrulerequestDud struct { 
    


    


    


    

}

// Twitterdataingestionrulerequest
type Twitterdataingestionrulerequest struct { 
    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // SearchTerms - Search terms for X (formally Twitter).
    SearchTerms string `json:"searchTerms"`


    // Countries - ISO 3166-1 alpha-2 country codes. Ingestion of matching tweets will be restricted to tweets posted in the countries specified here. Defaults to worldwide.
    Countries []string `json:"countries"`

}

// String returns a JSON representation of the model
func (o *Twitterdataingestionrulerequest) String() string {
    
    
    
     o.Countries = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterdataingestionrulerequest) MarshalJSON() ([]byte, error) {
    type Alias Twitterdataingestionrulerequest

    if TwitterdataingestionrulerequestMarshalled {
        return []byte("{}"), nil
    }
    TwitterdataingestionrulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        SearchTerms string `json:"searchTerms"`
        
        Countries []string `json:"countries"`
        *Alias
    }{

        


        


        


        
        Countries: []string{""},
        

        Alias: (*Alias)(u),
    })
}

