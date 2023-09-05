package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulesdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulesdivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulesdivisionview - A schedule.
type Schedulesdivisionview struct { 
    // Id - The schedule identifier
    Id string `json:"id"`


    // Name - The schedule name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Schedulesdivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulesdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Schedulesdivisionview

    if SchedulesdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    SchedulesdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

