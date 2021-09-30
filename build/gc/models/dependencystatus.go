package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencystatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencystatusDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dependencystatus
type Dependencystatus struct { 
    


    // Name
    Name string `json:"name"`


    // User - User that initiated the build.
    User User `json:"user"`


    // Client - OAuth client that initiated the build.
    Client Domainentityref `json:"client"`


    // BuildId
    BuildId string `json:"buildId"`


    // DateStarted - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // DateCompleted - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // Status
    Status string `json:"status"`


    // FailedObjects
    FailedObjects []Failedobject `json:"failedObjects"`


    

}

// String returns a JSON representation of the model
func (o *Dependencystatus) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FailedObjects = []Failedobject{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencystatus) MarshalJSON() ([]byte, error) {
    type Alias Dependencystatus

    if DependencystatusMarshalled {
        return []byte("{}"), nil
    }
    DependencystatusMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        User User `json:"user"`
        
        Client Domainentityref `json:"client"`
        
        BuildId string `json:"buildId"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        Status string `json:"status"`
        
        FailedObjects []Failedobject `json:"failedObjects"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        FailedObjects: []Failedobject{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

