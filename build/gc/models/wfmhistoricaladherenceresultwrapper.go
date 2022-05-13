package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherenceresultwrapperMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherenceresultwrapperDud struct { 
    


    


    

}

// Wfmhistoricaladherenceresultwrapper
type Wfmhistoricaladherenceresultwrapper struct { 
    // EntityId - The operation ID of the historical adherence query
    EntityId string `json:"entityId"`


    // Data - The list of historical adherence query results
    Data []Historicaladherencequeryresult `json:"data"`


    // LookupIdToSecondaryPresenceId - Map of secondary presence lookup ID to corresponding secondary presence ID
    LookupIdToSecondaryPresenceId map[string]string `json:"lookupIdToSecondaryPresenceId"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceresultwrapper) String() string {
    
     o.Data = []Historicaladherencequeryresult{{}} 
     o.LookupIdToSecondaryPresenceId = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherenceresultwrapper) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherenceresultwrapper

    if WfmhistoricaladherenceresultwrapperMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherenceresultwrapperMarshalled = true

    return json.Marshal(&struct {
        
        EntityId string `json:"entityId"`
        
        Data []Historicaladherencequeryresult `json:"data"`
        
        LookupIdToSecondaryPresenceId map[string]string `json:"lookupIdToSecondaryPresenceId"`
        *Alias
    }{

        


        
        Data: []Historicaladherencequeryresult{{}},
        


        
        LookupIdToSecondaryPresenceId: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

