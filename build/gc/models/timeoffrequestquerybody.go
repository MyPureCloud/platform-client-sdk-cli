package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestquerybodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestquerybodyDud struct { 
    


    


    

}

// Timeoffrequestquerybody
type Timeoffrequestquerybody struct { 
    // UserIds - The set of user ids to filter time off requests
    UserIds []string `json:"userIds"`


    // Statuses - The set of statuses to filter time off requests
    Statuses []string `json:"statuses"`


    // DateRange - The inclusive range of dates to filter time off requests
    DateRange Daterange `json:"dateRange"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestquerybody) String() string {
    
    
     o.UserIds = []string{""} 
    
    
    
     o.Statuses = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestquerybody) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestquerybody

    if TimeoffrequestquerybodyMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestquerybodyMarshalled = true

    return json.Marshal(&struct { 
        UserIds []string `json:"userIds"`
        
        Statuses []string `json:"statuses"`
        
        DateRange Daterange `json:"dateRange"`
        
        *Alias
    }{
        

        
        UserIds: []string{""},
        

        

        
        Statuses: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

