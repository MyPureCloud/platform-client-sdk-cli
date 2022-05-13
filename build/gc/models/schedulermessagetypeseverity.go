package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulermessagetypeseverityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulermessagetypeseverityDud struct { 
    


    

}

// Schedulermessagetypeseverity
type Schedulermessagetypeseverity struct { 
    // VarType - The type of the message
    VarType string `json:"type"`


    // Severity - The severity of the message
    Severity string `json:"severity"`

}

// String returns a JSON representation of the model
func (o *Schedulermessagetypeseverity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulermessagetypeseverity) MarshalJSON() ([]byte, error) {
    type Alias Schedulermessagetypeseverity

    if SchedulermessagetypeseverityMarshalled {
        return []byte("{}"), nil
    }
    SchedulermessagetypeseverityMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Severity string `json:"severity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

