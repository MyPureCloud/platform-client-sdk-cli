package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdjacentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdjacentsDud struct { 
    


    


    

}

// Adjacents
type Adjacents struct { 
    // Superiors
    Superiors []User `json:"superiors"`


    // Siblings
    Siblings []User `json:"siblings"`


    // DirectReports
    DirectReports []User `json:"directReports"`

}

// String returns a JSON representation of the model
func (o *Adjacents) String() string {
    
    
     o.Superiors = []User{{}} 
    
    
    
     o.Siblings = []User{{}} 
    
    
    
     o.DirectReports = []User{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adjacents) MarshalJSON() ([]byte, error) {
    type Alias Adjacents

    if AdjacentsMarshalled {
        return []byte("{}"), nil
    }
    AdjacentsMarshalled = true

    return json.Marshal(&struct { 
        Superiors []User `json:"superiors"`
        
        Siblings []User `json:"siblings"`
        
        DirectReports []User `json:"directReports"`
        
        *Alias
    }{
        

        
        Superiors: []User{{}},
        

        

        
        Siblings: []User{{}},
        

        

        
        DirectReports: []User{{}},
        

        
        Alias: (*Alias)(u),
    })
}

