package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulermessageseveritycountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulermessageseveritycountDud struct { 
    


    

}

// Schedulermessageseveritycount
type Schedulermessageseveritycount struct { 
    // Severity - The schedule message severity
    Severity string `json:"severity"`


    // Count - The number of schedule messages with the given severity
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Schedulermessageseveritycount) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulermessageseveritycount) MarshalJSON() ([]byte, error) {
    type Alias Schedulermessageseveritycount

    if SchedulermessageseveritycountMarshalled {
        return []byte("{}"), nil
    }
    SchedulermessageseveritycountMarshalled = true

    return json.Marshal(&struct { 
        Severity string `json:"severity"`
        
        Count int `json:"count"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

