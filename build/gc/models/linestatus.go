package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LinestatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LinestatusDud struct { 
    


    


    


    


    

}

// Linestatus
type Linestatus struct { 
    // Id - The id of this line
    Id string `json:"id"`


    // Reachable - Indicates whether the edge can reach the line.
    Reachable bool `json:"reachable"`


    // AddressOfRecord - The line's address of record.
    AddressOfRecord string `json:"addressOfRecord"`


    // ContactAddresses - The addresses used to contact the line.
    ContactAddresses []string `json:"contactAddresses"`


    // ReachableStateTime - The time the line entered its current reachable state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReachableStateTime time.Time `json:"reachableStateTime"`

}

// String returns a JSON representation of the model
func (o *Linestatus) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ContactAddresses = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Linestatus) MarshalJSON() ([]byte, error) {
    type Alias Linestatus

    if LinestatusMarshalled {
        return []byte("{}"), nil
    }
    LinestatusMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Reachable bool `json:"reachable"`
        
        AddressOfRecord string `json:"addressOfRecord"`
        
        ContactAddresses []string `json:"contactAddresses"`
        
        ReachableStateTime time.Time `json:"reachableStateTime"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        ContactAddresses: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

