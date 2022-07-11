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
    // Ids - The set of ids to filter time off requests
    Ids []string `json:"ids"`


    // UserIds - The set of user ids to filter time off requests
    UserIds []string `json:"userIds"`


    // Statuses - The set of statuses to filter time off requests
    Statuses []string `json:"statuses"`


    // Substatuses - The set of substatuses to filter time off requests
    Substatuses []string `json:"substatuses"`


    // DateRange - The inclusive range of dates to filter time off requests
    DateRange Daterange `json:"dateRange"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestquerybody) String() string {
     o.Ids = []string{""} 
     o.UserIds = []string{""} 
     o.Statuses = []string{""} 
     o.Substatuses = []string{""} 
    

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
        
        Ids []string `json:"ids"`
        
        UserIds []string `json:"userIds"`
        
        Statuses []string `json:"statuses"`
        
        Substatuses []string `json:"substatuses"`
        
        DateRange Daterange `json:"dateRange"`
        *Alias
    }{

        
        Ids: []string{""},
        


        
        UserIds: []string{""},
        


        
        Statuses: []string{""},
        


        
        Substatuses: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

