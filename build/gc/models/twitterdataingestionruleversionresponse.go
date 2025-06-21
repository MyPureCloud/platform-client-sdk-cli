package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterdataingestionruleversionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterdataingestionruleversionresponseDud struct { 
    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Platform string `json:"platform"`


    


    


    SelfUri string `json:"selfUri"`

}

// Twitterdataingestionruleversionresponse
type Twitterdataingestionruleversionresponse struct { 
    // Id - ID of the data ingestion rule.
    Id string `json:"id"`


    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // Status - The status of the data ingestion rule.
    Status string `json:"status"`


    // Version - The version number of the data ingestion rule.
    Version int `json:"version"`


    


    


    


    // Countries - ISO 3166-1 alpha-2 country codes where Data Ingestion Rules should apply. Defaults to worldwide.
    Countries []string `json:"countries"`


    // SearchTerms - Search terms for X (formally Twitter).
    SearchTerms string `json:"searchTerms"`


    

}

// String returns a JSON representation of the model
func (o *Twitterdataingestionruleversionresponse) String() string {
    
    
    
    
    
     o.Countries = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterdataingestionruleversionresponse) MarshalJSON() ([]byte, error) {
    type Alias Twitterdataingestionruleversionresponse

    if TwitterdataingestionruleversionresponseMarshalled {
        return []byte("{}"), nil
    }
    TwitterdataingestionruleversionresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        Version int `json:"version"`
        
        Countries []string `json:"countries"`
        
        SearchTerms string `json:"searchTerms"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Countries: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

