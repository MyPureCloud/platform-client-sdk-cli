package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueusersjoberrorinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueusersjoberrorinfoDud struct { 
    


    

}

// Assistantqueueusersjoberrorinfo
type Assistantqueueusersjoberrorinfo struct { 
    // Message
    Message string `json:"message"`


    // Code
    Code string `json:"code"`

}

// String returns a JSON representation of the model
func (o *Assistantqueueusersjoberrorinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueueusersjoberrorinfo) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueueusersjoberrorinfo

    if AssistantqueueusersjoberrorinfoMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueusersjoberrorinfoMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Code string `json:"code"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

