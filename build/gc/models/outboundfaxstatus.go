package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutboundfaxstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutboundfaxstatusDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Outboundfaxstatus
type Outboundfaxstatus struct { 
    


    // InitiatingUser - The user who sent the fax.
    InitiatingUser Addressableentityref `json:"initiatingUser"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // AuditTransactionId
    AuditTransactionId string `json:"auditTransactionId"`


    // ExpirationTime
    ExpirationTime int `json:"expirationTime"`


    // StatusCode - Lifecycle status of the outbound fax send (e.g. UPLOADING, TRANSMITTING, COMPLETE, TERMINATED).
    StatusCode string `json:"statusCode"`


    // Result - Transmission result of the fax. Does NOT indicate successful arrival to a workspace's inbox.
    Result string `json:"result"`


    

}

// String returns a JSON representation of the model
func (o *Outboundfaxstatus) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outboundfaxstatus) MarshalJSON() ([]byte, error) {
    type Alias Outboundfaxstatus

    if OutboundfaxstatusMarshalled {
        return []byte("{}"), nil
    }
    OutboundfaxstatusMarshalled = true

    return json.Marshal(&struct {
        
        InitiatingUser Addressableentityref `json:"initiatingUser"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        AuditTransactionId string `json:"auditTransactionId"`
        
        ExpirationTime int `json:"expirationTime"`
        
        StatusCode string `json:"statusCode"`
        
        Result string `json:"result"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

