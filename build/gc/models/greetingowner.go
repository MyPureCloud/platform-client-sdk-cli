package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GreetingownerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GreetingownerDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Greetingowner
type Greetingowner struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Greetingowner) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Greetingowner) MarshalJSON() ([]byte, error) {
    type Alias Greetingowner

    if GreetingownerMarshalled {
        return []byte("{}"), nil
    }
    GreetingownerMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

