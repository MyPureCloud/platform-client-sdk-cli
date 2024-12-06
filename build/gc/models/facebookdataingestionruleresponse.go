package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookdataingestionruleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookdataingestionruleresponseDud struct { 
    


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Facebookdataingestionruleresponse
type Facebookdataingestionruleresponse struct { 
    // Id - ID of the Facebook data ingestion rule.
    Id string `json:"id"`


    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // Status - The status of the data ingestion rule.
    Status string `json:"status"`


    // Version - The version number of the data ingestion rule.
    Version int `json:"version"`


    // IntegrationId - The Integration Id from which to ingest public social posts. This entity is created using the /conversations/messaging/integrations/facebook resource
    IntegrationId string `json:"integrationId"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Facebookdataingestionruleresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookdataingestionruleresponse) MarshalJSON() ([]byte, error) {
    type Alias Facebookdataingestionruleresponse

    if FacebookdataingestionruleresponseMarshalled {
        return []byte("{}"), nil
    }
    FacebookdataingestionruleresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        Version int `json:"version"`
        
        IntegrationId string `json:"integrationId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

