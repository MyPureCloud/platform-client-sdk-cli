package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestmodetriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestmodetriggerDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Testmodetrigger - Basic identifying information about a trigger
type Testmodetrigger struct { 
    


    // Name - The name of the trigger
    Name string `json:"name"`


    // Enabled - Whether or not the trigger is enabled
    Enabled bool `json:"enabled"`


    

}

// String returns a JSON representation of the model
func (o *Testmodetrigger) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testmodetrigger) MarshalJSON() ([]byte, error) {
    type Alias Testmodetrigger

    if TestmodetriggerMarshalled {
        return []byte("{}"), nil
    }
    TestmodetriggerMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

