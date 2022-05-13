package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowmilestoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowmilestoneDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowmilestone
type Flowmilestone struct { 
    


    // Name - The flow milestone name.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - The flow milestone description.
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Flowmilestone) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowmilestone) MarshalJSON() ([]byte, error) {
    type Alias Flowmilestone

    if FlowmilestoneMarshalled {
        return []byte("{}"), nil
    }
    FlowmilestoneMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

