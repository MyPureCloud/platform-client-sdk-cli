package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SingleworkdayaveragevaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SingleworkdayaveragevaluesDud struct { 
    DateWorkday time.Time `json:"dateWorkday"`


    Division Division `json:"division"`


    User Userreference `json:"user"`


    Timezone string `json:"timezone"`


    Results []Workdayvaluesmetricitem `json:"results"`

}

// Singleworkdayaveragevalues
type Singleworkdayaveragevalues struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragevalues) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Singleworkdayaveragevalues) MarshalJSON() ([]byte, error) {
    type Alias Singleworkdayaveragevalues

    if SingleworkdayaveragevaluesMarshalled {
        return []byte("{}"), nil
    }
    SingleworkdayaveragevaluesMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

