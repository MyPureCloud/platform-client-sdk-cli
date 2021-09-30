package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresenceexpandMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresenceexpandDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Presenceexpand
type Presenceexpand struct { 
    


    // Name
    Name string `json:"name"`


    // Presences - An array of user presences
    Presences []Userpresence `json:"presences"`


    // OutOfOffices - An array of out of office statuses
    OutOfOffices []Outofoffice `json:"outOfOffices"`


    

}

// String returns a JSON representation of the model
func (o *Presenceexpand) String() string {
    
    
    
    
    
    
    
    
     o.Presences = []Userpresence{{}} 
    
    
    
     o.OutOfOffices = []Outofoffice{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presenceexpand) MarshalJSON() ([]byte, error) {
    type Alias Presenceexpand

    if PresenceexpandMarshalled {
        return []byte("{}"), nil
    }
    PresenceexpandMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Presences []Userpresence `json:"presences"`
        
        OutOfOffices []Outofoffice `json:"outOfOffices"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Presences: []Userpresence{{}},
        

        

        
        OutOfOffices: []Outofoffice{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

