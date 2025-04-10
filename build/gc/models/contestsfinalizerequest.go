package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestsfinalizerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestsfinalizerequestDud struct { 
    


    


    

}

// Contestsfinalizerequest
type Contestsfinalizerequest struct { 
    // Status - The Contest finalization status
    Status string `json:"status"`


    // Winners - The Contest finalization winners
    Winners []Contestwinnersrequest `json:"winners"`


    // DisqualifiedAgents - The Contest finalization disqualified agents
    DisqualifiedAgents []Contestdisqualifiedagents `json:"disqualifiedAgents"`

}

// String returns a JSON representation of the model
func (o *Contestsfinalizerequest) String() string {
    
     o.Winners = []Contestwinnersrequest{{}} 
     o.DisqualifiedAgents = []Contestdisqualifiedagents{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestsfinalizerequest) MarshalJSON() ([]byte, error) {
    type Alias Contestsfinalizerequest

    if ContestsfinalizerequestMarshalled {
        return []byte("{}"), nil
    }
    ContestsfinalizerequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Winners []Contestwinnersrequest `json:"winners"`
        
        DisqualifiedAgents []Contestdisqualifiedagents `json:"disqualifiedAgents"`
        *Alias
    }{

        


        
        Winners: []Contestwinnersrequest{{}},
        


        
        DisqualifiedAgents: []Contestdisqualifiedagents{{}},
        

        Alias: (*Alias)(u),
    })
}

