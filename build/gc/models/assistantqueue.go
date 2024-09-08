package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueDud struct { 
    


    


    Assistant Assistant `json:"assistant"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Assistantqueue
type Assistantqueue struct { 
    // Id - The globally unique identifier for the queue.
    Id string `json:"id"`


    // MediaTypes - List of media Types in which the assistant is activated for this queue.
    MediaTypes []string `json:"mediaTypes"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Assistantqueue) String() string {
    
     o.MediaTypes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueue) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueue

    if AssistantqueueMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        MediaTypes []string `json:"mediaTypes"`
        *Alias
    }{

        


        
        MediaTypes: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

