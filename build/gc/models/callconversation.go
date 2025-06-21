package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallconversationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Callconversation
type Callconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Callmediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    // RecentTransfers - The list of the most recent 20 transfer commands applied to this conversation.
    RecentTransfers []Transferresponse `json:"recentTransfers"`


    // UtilizationLabelId - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
    UtilizationLabelId string `json:"utilizationLabelId"`


    // InactivityTimeout - The time in the future, after which this conversation would be considered inactive. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InactivityTimeout time.Time `json:"inactivityTimeout"`


    // Divisions - Identifiers of divisions associated with this conversation.
    Divisions []Conversationdivisionmembership `json:"divisions"`


    // RecordingState
    RecordingState string `json:"recordingState"`


    // MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
    MaxParticipants int `json:"maxParticipants"`


    // SecurePause - True when the recording of this conversation is in secure pause status.
    SecurePause bool `json:"securePause"`


    

}

// String returns a JSON representation of the model
func (o *Callconversation) String() string {
    
     o.Participants = []Callmediaparticipant{{}} 
     o.OtherMediaUris = []string{""} 
     o.RecentTransfers = []Transferresponse{{}} 
    
    
     o.Divisions = []Conversationdivisionmembership{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callconversation) MarshalJSON() ([]byte, error) {
    type Alias Callconversation

    if CallconversationMarshalled {
        return []byte("{}"), nil
    }
    CallconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Participants []Callmediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        
        RecentTransfers []Transferresponse `json:"recentTransfers"`
        
        UtilizationLabelId string `json:"utilizationLabelId"`
        
        InactivityTimeout time.Time `json:"inactivityTimeout"`
        
        Divisions []Conversationdivisionmembership `json:"divisions"`
        
        RecordingState string `json:"recordingState"`
        
        MaxParticipants int `json:"maxParticipants"`
        
        SecurePause bool `json:"securePause"`
        *Alias
    }{

        


        


        
        Participants: []Callmediaparticipant{{}},
        


        
        OtherMediaUris: []string{""},
        


        
        RecentTransfers: []Transferresponse{{}},
        


        


        


        
        Divisions: []Conversationdivisionmembership{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

