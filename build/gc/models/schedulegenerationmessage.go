package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegenerationmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegenerationmessageDud struct { 
    


    

}

// Schedulegenerationmessage
type Schedulegenerationmessage struct { 
    // VarType - The type of the message
    VarType string `json:"type"`


    // Arguments - The arguments describing the message
    Arguments []Schedulermessageargument `json:"arguments"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationmessage) String() string {
    
    
    
    
    
    
     o.Arguments = []Schedulermessageargument{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegenerationmessage) MarshalJSON() ([]byte, error) {
    type Alias Schedulegenerationmessage

    if SchedulegenerationmessageMarshalled {
        return []byte("{}"), nil
    }
    SchedulegenerationmessageMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Arguments []Schedulermessageargument `json:"arguments"`
        
        *Alias
    }{
        

        

        

        
        Arguments: []Schedulermessageargument{{}},
        

        
        Alias: (*Alias)(u),
    })
}

