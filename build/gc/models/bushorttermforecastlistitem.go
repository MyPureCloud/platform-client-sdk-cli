package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BushorttermforecastlistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BushorttermforecastlistitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    Legacy bool `json:"legacy"`


    


    


    SelfUri string `json:"selfUri"`

}

// Bushorttermforecastlistitem
type Bushorttermforecastlistitem struct { 
    


    // WeekDate - The start week date of this forecast in yyyy-MM-dd.  Must fall on the start day of week for the associated business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks this forecast covers
    WeekCount int `json:"weekCount"`


    // CreationMethod - The method by which this forecast was created
    CreationMethod string `json:"creationMethod"`


    // Description - The description of this forecast
    Description string `json:"description"`


    


    // Metadata - Metadata for this forecast
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // CanUseForScheduling - Whether this forecast can be used for scheduling
    CanUseForScheduling bool `json:"canUseForScheduling"`


    

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastlistitem) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bushorttermforecastlistitem) MarshalJSON() ([]byte, error) {
    type Alias Bushorttermforecastlistitem

    if BushorttermforecastlistitemMarshalled {
        return []byte("{}"), nil
    }
    BushorttermforecastlistitemMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        
        CreationMethod string `json:"creationMethod"`
        
        Description string `json:"description"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        CanUseForScheduling bool `json:"canUseForScheduling"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

