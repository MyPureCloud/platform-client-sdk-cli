package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainnetworkcommandresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainnetworkcommandresponseDud struct { 
    


    


    


    

}

// Domainnetworkcommandresponse
type Domainnetworkcommandresponse struct { 
    // CorrelationId
    CorrelationId string `json:"correlationId"`


    // CommandName
    CommandName string `json:"commandName"`


    // Acknowledged
    Acknowledged bool `json:"acknowledged"`


    // ErrorInfo
    ErrorInfo *Errordetails `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkcommandresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainnetworkcommandresponse) MarshalJSON() ([]byte, error) {
    type Alias Domainnetworkcommandresponse

    if DomainnetworkcommandresponseMarshalled {
        return []byte("{}"), nil
    }
    DomainnetworkcommandresponseMarshalled = true

    return json.Marshal(&struct {
        
        CorrelationId string `json:"correlationId"`
        
        CommandName string `json:"commandName"`
        
        Acknowledged bool `json:"acknowledged"`
        
        ErrorInfo *Errordetails `json:"errorInfo"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

