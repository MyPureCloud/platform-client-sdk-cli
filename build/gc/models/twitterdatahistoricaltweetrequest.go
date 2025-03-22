package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterdatahistoricaltweetrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterdatahistoricaltweetrequestDud struct { 
    


    

}

// Twitterdatahistoricaltweetrequest
type Twitterdatahistoricaltweetrequest struct { 
    // SearchTerms - Search terms to use in the query
    SearchTerms string `json:"searchTerms"`


    // Countries - ISO 3166-1 alpha-2 country codes to use for the search. Defaults to worldwide.
    Countries []string `json:"countries"`

}

// String returns a JSON representation of the model
func (o *Twitterdatahistoricaltweetrequest) String() string {
    
     o.Countries = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterdatahistoricaltweetrequest) MarshalJSON() ([]byte, error) {
    type Alias Twitterdatahistoricaltweetrequest

    if TwitterdatahistoricaltweetrequestMarshalled {
        return []byte("{}"), nil
    }
    TwitterdatahistoricaltweetrequestMarshalled = true

    return json.Marshal(&struct {
        
        SearchTerms string `json:"searchTerms"`
        
        Countries []string `json:"countries"`
        *Alias
    }{

        


        
        Countries: []string{""},
        

        Alias: (*Alias)(u),
    })
}

