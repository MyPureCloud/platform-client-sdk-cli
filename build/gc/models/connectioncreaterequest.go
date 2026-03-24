package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectioncreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectioncreaterequestDud struct { 
    


    


    

}

// Connectioncreaterequest
type Connectioncreaterequest struct { 
    // Name - The name of the Connection
    Name string `json:"name"`


    // IntegrationId - Integration ID of the Connection
    IntegrationId string `json:"integrationId"`


    // RedirectUrl - Redirect Url for the Oauth flow
    RedirectUrl string `json:"redirectUrl"`

}

// String returns a JSON representation of the model
func (o *Connectioncreaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectioncreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Connectioncreaterequest

    if ConnectioncreaterequestMarshalled {
        return []byte("{}"), nil
    }
    ConnectioncreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        IntegrationId string `json:"integrationId"`
        
        RedirectUrl string `json:"redirectUrl"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

