package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentstatecountsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentstatecountsresponseDud struct { 
    


    


    


    


    


    


    

}

// Analyticsagentstatecountsresponse
type Analyticsagentstatecountsresponse struct { 
    // SegmentCounts - List of count by segment types
    SegmentCounts []Agentstatesegmenttypecount `json:"segmentCounts"`


    // PresenceCounts - List of count by presences
    PresenceCounts []Agentstatepresencecount `json:"presenceCounts"`


    // RoutingStatusCounts - List of count by routing statuses
    RoutingStatusCounts []Agentstateroutingstatuscount `json:"routingStatusCounts"`


    // IsOutOfOfficeCounts - List of count by out of office states
    IsOutOfOfficeCounts []Agentstateisoutofofficecount `json:"isOutOfOfficeCounts"`


    // AdherenceStateCounts - List of count by adherence state
    AdherenceStateCounts []Agentstateadherencestatecount `json:"adherenceStateCounts"`


    // ScheduledActivityCategoryCounts - List of count by scheduled activity category
    ScheduledActivityCategoryCounts []Agentstateactivitycategorycount `json:"scheduledActivityCategoryCounts"`


    // ActualActivityCategoryCounts - List of count by actual activity category
    ActualActivityCategoryCounts []Agentstateactivitycategorycount `json:"actualActivityCategoryCounts"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentstatecountsresponse) String() string {
     o.SegmentCounts = []Agentstatesegmenttypecount{{}} 
     o.PresenceCounts = []Agentstatepresencecount{{}} 
     o.RoutingStatusCounts = []Agentstateroutingstatuscount{{}} 
     o.IsOutOfOfficeCounts = []Agentstateisoutofofficecount{{}} 
     o.AdherenceStateCounts = []Agentstateadherencestatecount{{}} 
     o.ScheduledActivityCategoryCounts = []Agentstateactivitycategorycount{{}} 
     o.ActualActivityCategoryCounts = []Agentstateactivitycategorycount{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentstatecountsresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentstatecountsresponse

    if AnalyticsagentstatecountsresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentstatecountsresponseMarshalled = true

    return json.Marshal(&struct {
        
        SegmentCounts []Agentstatesegmenttypecount `json:"segmentCounts"`
        
        PresenceCounts []Agentstatepresencecount `json:"presenceCounts"`
        
        RoutingStatusCounts []Agentstateroutingstatuscount `json:"routingStatusCounts"`
        
        IsOutOfOfficeCounts []Agentstateisoutofofficecount `json:"isOutOfOfficeCounts"`
        
        AdherenceStateCounts []Agentstateadherencestatecount `json:"adherenceStateCounts"`
        
        ScheduledActivityCategoryCounts []Agentstateactivitycategorycount `json:"scheduledActivityCategoryCounts"`
        
        ActualActivityCategoryCounts []Agentstateactivitycategorycount `json:"actualActivityCategoryCounts"`
        *Alias
    }{

        
        SegmentCounts: []Agentstatesegmenttypecount{{}},
        


        
        PresenceCounts: []Agentstatepresencecount{{}},
        


        
        RoutingStatusCounts: []Agentstateroutingstatuscount{{}},
        


        
        IsOutOfOfficeCounts: []Agentstateisoutofofficecount{{}},
        


        
        AdherenceStateCounts: []Agentstateadherencestatecount{{}},
        


        
        ScheduledActivityCategoryCounts: []Agentstateactivitycategorycount{{}},
        


        
        ActualActivityCategoryCounts: []Agentstateactivitycategorycount{{}},
        

        Alias: (*Alias)(u),
    })
}

