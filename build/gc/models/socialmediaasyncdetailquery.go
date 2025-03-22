package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaasyncdetailqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaasyncdetailqueryDud struct { 
    


    


    


    


    

}

// Socialmediaasyncdetailquery
type Socialmediaasyncdetailquery struct { 
    // Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // TimeZone - Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZone string `json:"timeZone"`


    // Filter - Behaves like a SQL WHERE clause. This is ANDed with the interval parameter. Expresses boolean logical predicates as well as dimensional filters
    Filter Socialmediaqueryfilter `json:"filter"`


    // PageSize - The number of results per page
    PageSize int `json:"pageSize"`


    // Order - Sorting of results based on time
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Socialmediaasyncdetailquery) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaasyncdetailquery) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaasyncdetailquery

    if SocialmediaasyncdetailqueryMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaasyncdetailqueryMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        TimeZone string `json:"timeZone"`
        
        Filter Socialmediaqueryfilter `json:"filter"`
        
        PageSize int `json:"pageSize"`
        
        Order string `json:"order"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

