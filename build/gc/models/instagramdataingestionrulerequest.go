package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramdataingestionrulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramdataingestionrulerequestDud struct { 
    


    


    

}

// Instagramdataingestionrulerequest
type Instagramdataingestionrulerequest struct { 
    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // IntegrationId - The Integration Id from which public social posts are ingested. This entity is created using the /conversations/messaging/integrations/instagram resource
    IntegrationId string `json:"integrationId"`

}

// String returns a JSON representation of the model
func (o *Instagramdataingestionrulerequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramdataingestionrulerequest) MarshalJSON() ([]byte, error) {
    type Alias Instagramdataingestionrulerequest

    if InstagramdataingestionrulerequestMarshalled {
        return []byte("{}"), nil
    }
    InstagramdataingestionrulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        IntegrationId string `json:"integrationId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

