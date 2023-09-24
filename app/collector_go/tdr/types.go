package tdr

type OperatingHours struct {
	OperatingHoursFromDate string `json:"OperatingHoursFromDate"`
	OperatingHoursFrom     string `json:"OperatingHoursFrom"`
	OperatingHoursToDate   string `json:"OperatingHoursToDate"`
	OperatingHoursTo       string `json:"OperatingHoursTo"`
	SunsetFlg              bool   `json:"SunsetFlg"`
	OperatingStatusCD      string `json:"OperatingStatusCD"`
	OperatingStatus        string `json:"OperatingStatus"`
	OperatingChgFlg        bool   `json:"OperatingChgFlg"`
}

type Facility struct {
	FacilityID             string           `json:"FacilityID"`
	FacilityName           string           `json:"FacilityName"`
	FacilityKanaName       string           `json:"FacilityKanaName"`
	NewFlg                 bool             `json:"NewFlg"`
	FacilityURLSP          interface{}      `json:"FacilityURLSP"`
	FacilityStatusCD       interface{}      `json:"FacilityStatusCD"`
	FacilityStatus         interface{}      `json:"FacilityStatus"`
	StandbyTime            string           `json:"StandbyTime"`
	OperatingHoursFromDate string           `json:"OperatingHoursFromDate"`
	OperatingHoursFrom     string           `json:"OperatingHoursFrom"`
	OperatingHoursToDate   string           `json:"OperatingHoursToDate"`
	OperatingHoursTo       string           `json:"OperatingHoursTo"`
	OperatingStatusCD      string           `json:"OperatingStatusCD"`
	OperatingStatus        string           `json:"OperatingStatus"`
	SunsetFlg              bool             `json:"SunsetFlg"`
	Fsflg                  bool             `json:"Fsflg"`
	FsStatusflg            interface{}      `json:"FsStatusflg"`
	FsStatus               interface{}      `json:"FsStatus"`
	FsStatusCD             interface{}      `json:"FsStatusCD"`
	FsStatusStartDate      interface{}      `json:"FsStatusStartDate"`
	FsStatusStartTime      interface{}      `json:"FsStatusStartTime"`
	FsStatusEndDate        interface{}      `json:"FsStatusEndDate"`
	FsStatusEndTime        interface{}      `json:"FsStatusEndTime"`
	UseLimitFlg            bool             `json:"UseLimitFlg"`
	UseStandbyTimeStyle    bool             `json:"UseStandbyTimeStyle"`
	OperatingChgFlg        bool             `json:"OperatingChgFlg"`
	UpdateTime             string           `json:"UpdateTime"`
	OperatingHours         []OperatingHours `json:"operatingHours"`
}

type GreetingFacility struct {
	FacilityID          string           `json:"FacilityID"`
	FacilityName        string           `json:"FacilityName"`
	FacilityKanaName    string           `json:"FacilityKanaName"`
	NewFlg              bool             `json:"NewFlg"`
	AreaJName           string           `json:"AreaJName"`
	AreaMName           string           `json:"AreaMName"`
	FacilityURLSP       interface{}      `json:"FacilityURLSP"`
	FacilityStatusCD    interface{}      `json:"FacilityStatusCD"`
	FacilityStatus      interface{}      `json:"FacilityStatus"`
	StandbyTime         string           `json:"StandbyTime"`
	OperatingHours      []OperatingHours `json:"operatinghours"`
	UseStandbyTimeStyle bool             `json:"UseStandbyTimeStyle"`
	UpdateTime          string           `json:"UpdateTime"`
}

type Cameraman struct {
	FacilityID       string      `json:"FacilityID"`
	FacilityName     string      `json:"FacilityName"`
	FacilityKanaName string      `json:"FacilityKanaName"`
	FacilityStatusCD string      `json:"FacilityStatusCD"`
	FacilityStatus   string      `json:"FacilityStatus"`
	OperatingHours   interface{} `json:"operatinghours"`
	UpdateTime       string      `json:"UpdateTime"`
}

type Greeting struct {
	Greeting  GreetingFacility `json:"greeting"`
	Cameraman Cameraman        `json:"cameraman"`
}
