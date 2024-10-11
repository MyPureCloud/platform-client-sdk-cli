package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdivisionmembershipMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdivisionmembershipDud struct { 
    


    

}

// Conversationdivisionmembership
type Conversationdivisionmembership struct { 
    // Division - A division the conversation belongs to.
    Division Domainentityref `json:"division"`


    // Entities - The entities on the conversation within the division. These are the users, queues, work flows, etc. that can be on conversations and and be assigned to different divisions.
    Entities []Divisionentityref `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Conversationdivisionmembership) String() string {
    
     o.Entities = []Divisionentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdivisionmembership) MarshalJSON() ([]byte, error) {
    type Alias Conversationdivisionmembership

    if ConversationdivisionmembershipMarshalled {
        return []byte("{}"), nil
    }
    ConversationdivisionmembershipMarshalled = true

    return json.Marshal(&struct {
        
        Division Domainentityref `json:"division"`
        
        Entities []Divisionentityref `json:"entities"`
        *Alias
    }{

        


        
        Entities: []Divisionentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

