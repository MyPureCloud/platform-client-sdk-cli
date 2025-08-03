package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClientpublicapiusageresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClientpublicapiusageresultsresponseDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Clientpublicapiusageresultsresponse
type Clientpublicapiusageresultsresponse struct { 
    // Name
    Name string `json:"name"`


    // QueryStatus - The status of the query.
    QueryStatus string `json:"queryStatus"`


    // ErrorBody - The reason for the failure. This will only be present if the query is in a FAILURE state but may not be included even if the state is FAILURE
    ErrorBody Errorbody `json:"errorBody"`


    // NextUri - The uri to get the next set of results. Will only be included if there is another page to retrieve.
    NextUri string `json:"nextUri"`


    // Entities - The results of the query.
    Entities []Clientpublicapiusage `json:"entities"`


    

}

// String returns a JSON representation of the model
func (o *Clientpublicapiusageresultsresponse) String() string {
    
    
    
    
     o.Entities = []Clientpublicapiusage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clientpublicapiusageresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Clientpublicapiusageresultsresponse

    if ClientpublicapiusageresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    ClientpublicapiusageresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        QueryStatus string `json:"queryStatus"`
        
        ErrorBody Errorbody `json:"errorBody"`
        
        NextUri string `json:"nextUri"`
        
        Entities []Clientpublicapiusage `json:"entities"`
        *Alias
    }{

        


        


        


        


        
        Entities: []Clientpublicapiusage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

