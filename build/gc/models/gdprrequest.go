package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GdprrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GdprrequestDud struct { 
    Id string `json:"id"`


    


    CreatedBy Domainentityref `json:"createdBy"`


    


    


    CreatedDate time.Time `json:"createdDate"`


    Status string `json:"status"`


    


    ResultsUrl string `json:"resultsUrl"`


    SelfUri string `json:"selfUri"`

}

// Gdprrequest
type Gdprrequest struct { 
    


    // Name
    Name string `json:"name"`


    


    // ReplacementTerms - The replacement terms for the provided search terms, in the case of a GDPR_UPDATE request
    ReplacementTerms []Replacementterm `json:"replacementTerms"`


    // RequestType - The type of GDPR request
    RequestType string `json:"requestType"`


    


    


    // Subject - The subject of the GDPR request
    Subject Gdprsubject `json:"subject"`


    


    

}

// String returns a JSON representation of the model
func (o *Gdprrequest) String() string {
    
     o.ReplacementTerms = []Replacementterm{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gdprrequest) MarshalJSON() ([]byte, error) {
    type Alias Gdprrequest

    if GdprrequestMarshalled {
        return []byte("{}"), nil
    }
    GdprrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ReplacementTerms []Replacementterm `json:"replacementTerms"`
        
        RequestType string `json:"requestType"`
        
        Subject Gdprsubject `json:"subject"`
        *Alias
    }{

        


        


        


        
        ReplacementTerms: []Replacementterm{{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

