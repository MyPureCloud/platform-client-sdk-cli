package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailconversationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Emailconversation
type Emailconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Emailmediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    // RecentTransfers - The list of the most recent 20 transfer commands applied to this conversation.
    RecentTransfers []Transferresponse `json:"recentTransfers"`


    // UtilizationLabelId - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
    UtilizationLabelId string `json:"utilizationLabelId"`


    // Divisions - Identifiers of divisions associated with this conversation.
    Divisions []Conversationdivisionmembership `json:"divisions"`


    

}

// String returns a JSON representation of the model
func (o *Emailconversation) String() string {
    
     o.Participants = []Emailmediaparticipant{{}} 
     o.OtherMediaUris = []string{""} 
     o.RecentTransfers = []Transferresponse{{}} 
    
     o.Divisions = []Conversationdivisionmembership{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailconversation) MarshalJSON() ([]byte, error) {
    type Alias Emailconversation

    if EmailconversationMarshalled {
        return []byte("{}"), nil
    }
    EmailconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Participants []Emailmediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        
        RecentTransfers []Transferresponse `json:"recentTransfers"`
        
        UtilizationLabelId string `json:"utilizationLabelId"`
        
        Divisions []Conversationdivisionmembership `json:"divisions"`
        *Alias
    }{

        


        


        
        Participants: []Emailmediaparticipant{{}},
        


        
        OtherMediaUris: []string{""},
        


        
        RecentTransfers: []Transferresponse{{}},
        


        


        
        Divisions: []Conversationdivisionmembership{{}},
        


        

        Alias: (*Alias)(u),
    })
}

