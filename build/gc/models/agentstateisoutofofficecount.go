package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateisoutofofficecountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateisoutofofficecountDud struct { 
    


    

}

// Agentstateisoutofofficecount
type Agentstateisoutofofficecount struct { 
    // IsOutOfOffice - The out of office state
    IsOutOfOffice bool `json:"isOutOfOffice"`


    // Count - Count of users with this out of office state
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstateisoutofofficecount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateisoutofofficecount) MarshalJSON() ([]byte, error) {
    type Alias Agentstateisoutofofficecount

    if AgentstateisoutofofficecountMarshalled {
        return []byte("{}"), nil
    }
    AgentstateisoutofofficecountMarshalled = true

    return json.Marshal(&struct {
        
        IsOutOfOffice bool `json:"isOutOfOffice"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

