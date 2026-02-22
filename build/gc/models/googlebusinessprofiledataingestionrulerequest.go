package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GooglebusinessprofiledataingestionrulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GooglebusinessprofiledataingestionrulerequestDud struct { 
    


    


    


    

}

// Googlebusinessprofiledataingestionrulerequest
type Googlebusinessprofiledataingestionrulerequest struct { 
    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // IntegrationId - The Integration Id from which to ingest public social posts. This entity is created using the /conversations/messaging/integrations/open/extensions/googlebusinessprofile resource
    IntegrationId string `json:"integrationId"`


    // ExternalSource - The external source associated with this data ingestion rule, which will be used when performing identity resolution
    ExternalSource Domainentityref `json:"externalSource"`

}

// String returns a JSON representation of the model
func (o *Googlebusinessprofiledataingestionrulerequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googlebusinessprofiledataingestionrulerequest) MarshalJSON() ([]byte, error) {
    type Alias Googlebusinessprofiledataingestionrulerequest

    if GooglebusinessprofiledataingestionrulerequestMarshalled {
        return []byte("{}"), nil
    }
    GooglebusinessprofiledataingestionrulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        IntegrationId string `json:"integrationId"`
        
        ExternalSource Domainentityref `json:"externalSource"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

