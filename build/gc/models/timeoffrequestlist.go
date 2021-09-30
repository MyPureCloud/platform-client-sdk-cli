package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestlistDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Timeoffrequestlist
type Timeoffrequestlist struct { 
    


    // Name
    Name string `json:"name"`


    // TimeOffRequests
    TimeOffRequests []Timeoffrequestresponse `json:"timeOffRequests"`


    

}

// String returns a JSON representation of the model
func (o *Timeoffrequestlist) String() string {
    
    
    
    
    
    
    
    
     o.TimeOffRequests = []Timeoffrequestresponse{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestlist) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestlist

    if TimeoffrequestlistMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestlistMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        TimeOffRequests []Timeoffrequestresponse `json:"timeOffRequests"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        TimeOffRequests: []Timeoffrequestresponse{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

