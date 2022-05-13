package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxsendrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxsendrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Faxsendrequest
type Faxsendrequest struct { 
    


    // Name
    Name string `json:"name"`


    // Addresses - A list of outbound fax dialing addresses. E.g. +13175555555 or 3175555555
    Addresses []string `json:"addresses"`


    // DocumentId - DocumentId of Content Management artifact. If Content Management document is not used for faxing, documentId should be null
    DocumentId string `json:"documentId"`


    // ContentType - The content type that is going to be uploaded. If Content Management document is used for faxing, contentType will be ignored
    ContentType string `json:"contentType"`


    // Workspace - Workspace in which the document should be stored. If Content Management document is used for faxing, workspace will be ignored
    Workspace Workspace `json:"workspace"`


    // CoverSheet - Data for coversheet generation.
    CoverSheet Coversheet `json:"coverSheet"`


    // TimeZoneOffsetMinutes - Time zone offset minutes from GMT
    TimeZoneOffsetMinutes int `json:"timeZoneOffsetMinutes"`


    

}

// String returns a JSON representation of the model
func (o *Faxsendrequest) String() string {
    
     o.Addresses = []string{""} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxsendrequest) MarshalJSON() ([]byte, error) {
    type Alias Faxsendrequest

    if FaxsendrequestMarshalled {
        return []byte("{}"), nil
    }
    FaxsendrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Addresses []string `json:"addresses"`
        
        DocumentId string `json:"documentId"`
        
        ContentType string `json:"contentType"`
        
        Workspace Workspace `json:"workspace"`
        
        CoverSheet Coversheet `json:"coverSheet"`
        
        TimeZoneOffsetMinutes int `json:"timeZoneOffsetMinutes"`
        *Alias
    }{

        


        


        
        Addresses: []string{""},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

