package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncuserdetailsqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncuserdetailsqueryDud struct { 
    


    


    


    


    


    

}

// Asyncuserdetailsquery
type Asyncuserdetailsquery struct { 
    // Interval - Specifies the date and time range of data being queried. Conversations MUST have started within this time range to potentially be included within the result set. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // UserFilters - Filters that target the users to retrieve data for
    UserFilters []Userdetailqueryfilter `json:"userFilters"`


    // PresenceFilters - Filters that target system and organization presence-level data
    PresenceFilters []Presencedetailqueryfilter `json:"presenceFilters"`


    // RoutingStatusFilters - Filters that target agent routing status-level data
    RoutingStatusFilters []Routingstatusdetailqueryfilter `json:"routingStatusFilters"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`


    // Limit - Specify number of results to be returned
    Limit int `json:"limit"`

}

// String returns a JSON representation of the model
func (o *Asyncuserdetailsquery) String() string {
    
     o.UserFilters = []Userdetailqueryfilter{{}} 
     o.PresenceFilters = []Presencedetailqueryfilter{{}} 
     o.RoutingStatusFilters = []Routingstatusdetailqueryfilter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncuserdetailsquery) MarshalJSON() ([]byte, error) {
    type Alias Asyncuserdetailsquery

    if AsyncuserdetailsqueryMarshalled {
        return []byte("{}"), nil
    }
    AsyncuserdetailsqueryMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        UserFilters []Userdetailqueryfilter `json:"userFilters"`
        
        PresenceFilters []Presencedetailqueryfilter `json:"presenceFilters"`
        
        RoutingStatusFilters []Routingstatusdetailqueryfilter `json:"routingStatusFilters"`
        
        Order string `json:"order"`
        
        Limit int `json:"limit"`
        *Alias
    }{

        


        
        UserFilters: []Userdetailqueryfilter{{}},
        


        
        PresenceFilters: []Presencedetailqueryfilter{{}},
        


        
        RoutingStatusFilters: []Routingstatusdetailqueryfilter{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

