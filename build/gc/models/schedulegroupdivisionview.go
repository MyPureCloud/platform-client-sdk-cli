package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegroupdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegroupdivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulegroupdivisionview - A schedule group.
type Schedulegroupdivisionview struct { 
    // Id - The schedule group identifier
    Id string `json:"id"`


    // Name - The schedule group name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Schedulegroupdivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegroupdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Schedulegroupdivisionview

    if SchedulegroupdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    SchedulegroupdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

