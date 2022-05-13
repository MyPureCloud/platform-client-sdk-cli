package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxsummaryDud struct { 
    


    


    

}

// Faxsummary
type Faxsummary struct { 
    // ReadCount
    ReadCount int `json:"readCount"`


    // UnreadCount
    UnreadCount int `json:"unreadCount"`


    // TotalCount
    TotalCount int `json:"totalCount"`

}

// String returns a JSON representation of the model
func (o *Faxsummary) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxsummary) MarshalJSON() ([]byte, error) {
    type Alias Faxsummary

    if FaxsummaryMarshalled {
        return []byte("{}"), nil
    }
    FaxsummaryMarshalled = true

    return json.Marshal(&struct {
        
        ReadCount int `json:"readCount"`
        
        UnreadCount int `json:"unreadCount"`
        
        TotalCount int `json:"totalCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

