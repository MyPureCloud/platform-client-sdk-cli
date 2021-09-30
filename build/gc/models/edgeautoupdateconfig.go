package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeautoupdateconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeautoupdateconfigDud struct { 
    


    


    


    

}

// Edgeautoupdateconfig
type Edgeautoupdateconfig struct { 
    // TimeZone - The timezone of the window in which any updates to the edges assigned to the site can be applied. The minimum size of the window is 2 hours.
    TimeZone string `json:"timeZone"`


    // Rrule - The recurrence rule for updating the Edges assigned to the site. The only supported frequencies are daily and weekly. Weekly frequencies require a day list with at least oneday specified. All other configurations are not supported.
    Rrule string `json:"rrule"`


    // Start - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
    Start time.Time `json:"start"`


    // End - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
    End time.Time `json:"end"`

}

// String returns a JSON representation of the model
func (o *Edgeautoupdateconfig) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeautoupdateconfig) MarshalJSON() ([]byte, error) {
    type Alias Edgeautoupdateconfig

    if EdgeautoupdateconfigMarshalled {
        return []byte("{}"), nil
    }
    EdgeautoupdateconfigMarshalled = true

    return json.Marshal(&struct { 
        TimeZone string `json:"timeZone"`
        
        Rrule string `json:"rrule"`
        
        Start time.Time `json:"start"`
        
        End time.Time `json:"end"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

