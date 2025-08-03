package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationpublicapiusageresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationpublicapiusageresultsresponseDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Organizationpublicapiusageresultsresponse
type Organizationpublicapiusageresultsresponse struct { 
    // Name
    Name string `json:"name"`


    // QueryStatus - The status of the query.
    QueryStatus string `json:"queryStatus"`


    // ErrorBody - The reason for the failure. This will only be present if the query is in a FAILURE state but may not be included even if the state is FAILURE
    ErrorBody Errorbody `json:"errorBody"`


    // NextUri - The uri to get the next set of results. Will only be included if there is another page to retrieve.
    NextUri string `json:"nextUri"`


    // Entities - The results of the query.
    Entities []Organizationpublicapiusage `json:"entities"`


    

}

// String returns a JSON representation of the model
func (o *Organizationpublicapiusageresultsresponse) String() string {
    
    
    
    
     o.Entities = []Organizationpublicapiusage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationpublicapiusageresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Organizationpublicapiusageresultsresponse

    if OrganizationpublicapiusageresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    OrganizationpublicapiusageresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        QueryStatus string `json:"queryStatus"`
        
        ErrorBody Errorbody `json:"errorBody"`
        
        NextUri string `json:"nextUri"`
        
        Entities []Organizationpublicapiusage `json:"entities"`
        *Alias
    }{

        


        


        


        


        
        Entities: []Organizationpublicapiusage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

