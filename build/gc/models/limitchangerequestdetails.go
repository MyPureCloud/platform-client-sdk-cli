package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitchangerequestdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitchangerequestdetailsDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    CreatedBy string `json:"createdBy"`


    Status string `json:"status"`


    CurrentValue float64 `json:"currentValue"`


    DateCreated time.Time `json:"dateCreated"`


    StatusHistory []Statuschange `json:"statusHistory"`


    DateCompleted time.Time `json:"dateCompleted"`


    LastChangedBy string `json:"lastChangedBy"`


    RejectReason string `json:"rejectReason"`


    SelfUri string `json:"selfUri"`

}

// Limitchangerequestdetails
type Limitchangerequestdetails struct { 
    


    // Key - Limit key to be overridden (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
    Key string `json:"key"`


    // Namespace - Namespace the key belongs to (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
    Namespace string `json:"namespace"`


    // RequestedValue - Requested limit value for a given key
    RequestedValue float64 `json:"requestedValue"`


    // Description - Description of the need for the limit change request
    Description string `json:"description"`


    // SupportCaseUrl - The support case url created by Care
    SupportCaseUrl string `json:"supportCaseUrl"`


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Limitchangerequestdetails) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitchangerequestdetails) MarshalJSON() ([]byte, error) {
    type Alias Limitchangerequestdetails

    if LimitchangerequestdetailsMarshalled {
        return []byte("{}"), nil
    }
    LimitchangerequestdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Namespace string `json:"namespace"`
        
        RequestedValue float64 `json:"requestedValue"`
        
        Description string `json:"description"`
        
        SupportCaseUrl string `json:"supportCaseUrl"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

