package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationusagequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationusagequeryresponseDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Organizationusagequeryresponse
type Organizationusagequeryresponse struct { 
    // Id - The jobId of the query.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // ResultsUri - The uri to get the results.
    ResultsUri string `json:"resultsUri"`


    

}

// String returns a JSON representation of the model
func (o *Organizationusagequeryresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationusagequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Organizationusagequeryresponse

    if OrganizationusagequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    OrganizationusagequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ResultsUri string `json:"resultsUri"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

