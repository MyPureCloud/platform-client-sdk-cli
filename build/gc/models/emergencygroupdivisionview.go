package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmergencygroupdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmergencygroupdivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Emergencygroupdivisionview - A group of call flows.
type Emergencygroupdivisionview struct { 
    // Id - The emergency group identifier
    Id string `json:"id"`


    // Name - The emergency group name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Emergencygroupdivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emergencygroupdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Emergencygroupdivisionview

    if EmergencygroupdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    EmergencygroupdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

