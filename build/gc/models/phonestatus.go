package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonestatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonestatusDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Phonestatus
type Phonestatus struct { 
    


    // Name
    Name string `json:"name"`


    // OperationalStatus - The Operational Status of this phone
    OperationalStatus string `json:"operationalStatus"`


    // EdgesStatus - The status of the primary or secondary Edges assigned to the phone lines.
    EdgesStatus string `json:"edgesStatus"`


    // EventCreationTime - Event Creation Time represents an ISO-8601 string. For example: UTC, UTC+01:00, or Europe/London
    EventCreationTime string `json:"eventCreationTime"`


    // Provision - Provision information for this phone
    Provision Provisioninfo `json:"provision"`


    // LineStatuses - A list of LineStatus information for each of the lines of this phone
    LineStatuses []Linestatus `json:"lineStatuses"`


    // PhoneAssignmentToEdgeType - The phone status's edge assignment type.
    PhoneAssignmentToEdgeType string `json:"phoneAssignmentToEdgeType"`


    // Edge - The URI of the edge that provided this status information.
    Edge Domainentityref `json:"edge"`


    

}

// String returns a JSON representation of the model
func (o *Phonestatus) String() string {
    
    
    
    
    
     o.LineStatuses = []Linestatus{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonestatus) MarshalJSON() ([]byte, error) {
    type Alias Phonestatus

    if PhonestatusMarshalled {
        return []byte("{}"), nil
    }
    PhonestatusMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        OperationalStatus string `json:"operationalStatus"`
        
        EdgesStatus string `json:"edgesStatus"`
        
        EventCreationTime string `json:"eventCreationTime"`
        
        Provision Provisioninfo `json:"provision"`
        
        LineStatuses []Linestatus `json:"lineStatuses"`
        
        PhoneAssignmentToEdgeType string `json:"phoneAssignmentToEdgeType"`
        
        Edge Domainentityref `json:"edge"`
        *Alias
    }{

        


        


        


        


        


        


        
        LineStatuses: []Linestatus{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

