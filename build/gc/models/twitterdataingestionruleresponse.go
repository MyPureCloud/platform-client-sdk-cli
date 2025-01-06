package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterdataingestionruleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterdataingestionruleresponseDud struct { 
    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Twitterdataingestionruleresponse
type Twitterdataingestionruleresponse struct { 
    // Id - ID of the data ingestion rule.
    Id string `json:"id"`


    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // SearchTerms - Search terms for X (formally Twitter).
    SearchTerms string `json:"searchTerms"`


    // Countries - ISO 3166-1 alpha-2 country codes where Data Ingestion Rules should apply. Defaults to worldwide.
    Countries []string `json:"countries"`


    // DateCreated - Timestamp indicating when the data ingestion rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Timestamp indicating when the data ingestion rule was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Status - The status of the data ingestion rule.
    Status string `json:"status"`


    // Version - The version number of the data ingestion rule.
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Twitterdataingestionruleresponse) String() string {
    
    
    
    
     o.Countries = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterdataingestionruleresponse) MarshalJSON() ([]byte, error) {
    type Alias Twitterdataingestionruleresponse

    if TwitterdataingestionruleresponseMarshalled {
        return []byte("{}"), nil
    }
    TwitterdataingestionruleresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        SearchTerms string `json:"searchTerms"`
        
        Countries []string `json:"countries"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Status string `json:"status"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        
        Countries: []string{""},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

