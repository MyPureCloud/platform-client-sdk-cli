package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoryentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoryentryDud struct { 
    


    


    


    


    


    


    


    

}

// Historyentry
type Historyentry struct { 
    // Action - The action performed
    Action string `json:"action"`


    // Resource - For actions performed not on the item itself, but on a sub-item, this field identifies the sub-item by name.  For example, for actions performed on prompt resources, this will be the prompt resource name.
    Resource string `json:"resource"`


    // Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // User - User associated with this entry.
    User User `json:"user"`


    // Client - OAuth client associated with this entry.
    Client Domainentityref `json:"client"`


    // Version
    Version string `json:"version"`


    // Secure
    Secure bool `json:"secure"`


    // VirtualAgentEnabled
    VirtualAgentEnabled bool `json:"virtualAgentEnabled"`

}

// String returns a JSON representation of the model
func (o *Historyentry) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historyentry) MarshalJSON() ([]byte, error) {
    type Alias Historyentry

    if HistoryentryMarshalled {
        return []byte("{}"), nil
    }
    HistoryentryMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        Resource string `json:"resource"`
        
        Timestamp time.Time `json:"timestamp"`
        
        User User `json:"user"`
        
        Client Domainentityref `json:"client"`
        
        Version string `json:"version"`
        
        Secure bool `json:"secure"`
        
        VirtualAgentEnabled bool `json:"virtualAgentEnabled"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

