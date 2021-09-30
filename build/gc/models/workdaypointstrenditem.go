package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdaypointstrenditemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdaypointstrenditemDud struct { 
    DateWorkday time.Time `json:"dateWorkday"`


    Points float64 `json:"points"`

}

// Workdaypointstrenditem
type Workdaypointstrenditem struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Workdaypointstrenditem) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdaypointstrenditem) MarshalJSON() ([]byte, error) {
    type Alias Workdaypointstrenditem

    if WorkdaypointstrenditemMarshalled {
        return []byte("{}"), nil
    }
    WorkdaypointstrenditemMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

