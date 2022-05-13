package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserdetailsqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserdetailsqueryDud struct { 
    


    


    


    


    


    


    


    

}

// Userdetailsquery
type Userdetailsquery struct { 
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


    // PresenceAggregations - Include faceted search and aggregate roll-ups of presence data in your search results. This does not function as a filter, but rather, summary data about the presence results matching your filters
    PresenceAggregations []Analyticsqueryaggregation `json:"presenceAggregations"`


    // RoutingStatusAggregations - Include faceted search and aggregate roll-ups of agent routing status data in your search results. This does not function as a filter, but rather, summary data about the agent routing status results matching your filters
    RoutingStatusAggregations []Analyticsqueryaggregation `json:"routingStatusAggregations"`


    // Paging - Page size and number to control iterating through large result sets. Default page size is 25
    Paging Pagingspec `json:"paging"`

}

// String returns a JSON representation of the model
func (o *Userdetailsquery) String() string {
    
     o.UserFilters = []Userdetailqueryfilter{{}} 
     o.PresenceFilters = []Presencedetailqueryfilter{{}} 
     o.RoutingStatusFilters = []Routingstatusdetailqueryfilter{{}} 
    
     o.PresenceAggregations = []Analyticsqueryaggregation{{}} 
     o.RoutingStatusAggregations = []Analyticsqueryaggregation{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userdetailsquery) MarshalJSON() ([]byte, error) {
    type Alias Userdetailsquery

    if UserdetailsqueryMarshalled {
        return []byte("{}"), nil
    }
    UserdetailsqueryMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        UserFilters []Userdetailqueryfilter `json:"userFilters"`
        
        PresenceFilters []Presencedetailqueryfilter `json:"presenceFilters"`
        
        RoutingStatusFilters []Routingstatusdetailqueryfilter `json:"routingStatusFilters"`
        
        Order string `json:"order"`
        
        PresenceAggregations []Analyticsqueryaggregation `json:"presenceAggregations"`
        
        RoutingStatusAggregations []Analyticsqueryaggregation `json:"routingStatusAggregations"`
        
        Paging Pagingspec `json:"paging"`
        *Alias
    }{

        


        
        UserFilters: []Userdetailqueryfilter{{}},
        


        
        PresenceFilters: []Presencedetailqueryfilter{{}},
        


        
        RoutingStatusFilters: []Routingstatusdetailqueryfilter{{}},
        


        


        
        PresenceAggregations: []Analyticsqueryaggregation{{}},
        


        
        RoutingStatusAggregations: []Analyticsqueryaggregation{{}},
        


        

        Alias: (*Alias)(u),
    })
}

