package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Timeoffrequestreference
type Timeoffrequestreference struct { 
    // Id - The id of the time off request
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Timeoffrequestreference) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestreference) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestreference

    if TimeoffrequestreferenceMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

