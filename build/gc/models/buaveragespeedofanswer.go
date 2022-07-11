package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuaveragespeedofanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuaveragespeedofanswerDud struct { 
    


    

}

// Buaveragespeedofanswer
type Buaveragespeedofanswer struct { 
    // Include - Whether to include average speed of answer (ASA) in the associated configuration
    Include bool `json:"include"`


    // Seconds - The target average speed of answer (ASA) in seconds. Required if include == true
    Seconds int `json:"seconds"`

}

// String returns a JSON representation of the model
func (o *Buaveragespeedofanswer) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buaveragespeedofanswer) MarshalJSON() ([]byte, error) {
    type Alias Buaveragespeedofanswer

    if BuaveragespeedofanswerMarshalled {
        return []byte("{}"), nil
    }
    BuaveragespeedofanswerMarshalled = true

    return json.Marshal(&struct {
        
        Include bool `json:"include"`
        
        Seconds int `json:"seconds"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

