package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsuserdetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsuserdetailDud struct { 
    


    


    

}

// Analyticsuserdetail
type Analyticsuserdetail struct { 
    // UserId - The identifier for the user
    UserId string `json:"userId"`


    // PrimaryPresence - The presence records for the user
    PrimaryPresence []Analyticsuserpresencerecord `json:"primaryPresence"`


    // RoutingStatus - The ACD routing status records for the user
    RoutingStatus []Analyticsroutingstatusrecord `json:"routingStatus"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetail) String() string {
    
    
    
    
    
    
     o.PrimaryPresence = []Analyticsuserpresencerecord{{}} 
    
    
    
     o.RoutingStatus = []Analyticsroutingstatusrecord{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsuserdetail) MarshalJSON() ([]byte, error) {
    type Alias Analyticsuserdetail

    if AnalyticsuserdetailMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsuserdetailMarshalled = true

    return json.Marshal(&struct { 
        UserId string `json:"userId"`
        
        PrimaryPresence []Analyticsuserpresencerecord `json:"primaryPresence"`
        
        RoutingStatus []Analyticsroutingstatusrecord `json:"routingStatus"`
        
        *Alias
    }{
        

        

        

        
        PrimaryPresence: []Analyticsuserpresencerecord{{}},
        

        

        
        RoutingStatus: []Analyticsroutingstatusrecord{{}},
        

        
        Alias: (*Alias)(u),
    })
}

