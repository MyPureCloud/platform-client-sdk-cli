package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesactionlimitinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesactionlimitinfoDud struct { 
    FileAttachmentsPerResponseMax Capabilitieslimitcondition `json:"fileAttachmentsPerResponseMax"`


    InvocationsPerFlowMax Capabilitieslimitcondition `json:"invocationsPerFlowMax"`


    LoopCountMax Capabilitieslimitcondition `json:"loopCountMax"`


    NotificationTypeChatInvocationsPerFlowMax Capabilitieslimitcondition `json:"notificationTypeChatInvocationsPerFlowMax"`

}

// Capabilitiesactionlimitinfo
type Capabilitiesactionlimitinfo struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Capabilitiesactionlimitinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesactionlimitinfo) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesactionlimitinfo

    if CapabilitiesactionlimitinfoMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesactionlimitinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

