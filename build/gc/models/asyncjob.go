package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncjobDud struct { 
    


    


    

}

// Asyncjob
type Asyncjob struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // State
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Asyncjob) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncjob) MarshalJSON() ([]byte, error) {
    type Alias Asyncjob

    if AsyncjobMarshalled {
        return []byte("{}"), nil
    }
    AsyncjobMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

