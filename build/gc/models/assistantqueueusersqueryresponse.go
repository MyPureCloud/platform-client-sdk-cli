package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueusersqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueusersqueryresponseDud struct { 
    Queue Assistantqueue `json:"queue"`


    Users []Userreference `json:"users"`

}

// Assistantqueueusersqueryresponse
type Assistantqueueusersqueryresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Assistantqueueusersqueryresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueueusersqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueueusersqueryresponse

    if AssistantqueueusersqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueusersqueryresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

