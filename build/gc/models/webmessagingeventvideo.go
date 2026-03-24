package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingeventvideoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingeventvideoDud struct { 
    


    


    

}

// Webmessagingeventvideo - A Video event.
type Webmessagingeventvideo struct { 
    // VarType - Describes the type of Video event.
    VarType string `json:"type"`


    // OfferingId - The Video offering ID.
    OfferingId string `json:"offeringId"`


    // Jwt - The Video offering JWT token.
    Jwt string `json:"jwt"`

}

// String returns a JSON representation of the model
func (o *Webmessagingeventvideo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingeventvideo) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingeventvideo

    if WebmessagingeventvideoMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingeventvideoMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        OfferingId string `json:"offeringId"`
        
        Jwt string `json:"jwt"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

