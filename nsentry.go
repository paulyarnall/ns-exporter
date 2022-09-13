package main

import "time"

type NsEntry struct {
	OpenAps struct {
		Suggested struct {
			Temp             string    `json:"temp" bson:"temp"`
			Bg               float64   `json:"bg" bson:"bg"`
			Tick             float64   `json:"-" bson:"-"`
			EventualBG       float64   `json:"eventualBG" bson:"eventualBG"`
			TargetBG         float64   `json:"targetBG" bson:"targetBG"`
			InsulinReq       float64   `json:"insulinReq" bson:"insulinReq"`
			DeliverAt        time.Time `json:"deliverAt" bson:"deliverAt"`
			SensitivityRatio float64   `json:"sensitivityRatio" bson:"sensitivityRatio"`
			PredBGs          struct {
				IOB []float64 `json:"IOB"`
				ZT  []float64 `json:"ZT"`
				COB []float64 `json:"COB"`
				UAM []float64 `json:"UAM"`
			} `json:"predBGs"`
			COB       float64   `json:"COB"`
			IOB       float64   `json:"IOB"`
			Reason    string    `json:"reason"`
			Units     float64   `json:"units"`
			Rate      float64   `json:"rate"`
			Duration  int       `json:"duration"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"suggested,omitempty" bson:"suggested,omitempty"`
		IOB struct {
			IOB      float64   `json:"iob" bson:"iob"`
			BasalIOB float64   `json:"basaliob" bson:"basaliob"`
			Activity float64   `json:"activity" bson:"activity"`
			Time     time.Time `json:"time" bson:"time"`
		} `json:"iob" bson:"iob"`
	} `json:"openaps" bson:"openaps"`
	Pump struct {
		Clock     time.Time `json:"clock"`
		Reservoir float64   `json:"reservoir"`
		Status    struct {
			Status    string `json:"status"`
			Timestamp int64  `json:"-" bson:"-"`
		} `json:"status"`
		Extended struct {
			Version               string  `json:"Version"`
			ActiveProfile         string  `json:"ActiveProfile"`
			TempBasalAbsoluteRate float64 `json:"TempBasalAbsoluteRate"`
			TempBasalPercent      int     `json:"TempBasalPercent"`
			TempBasalRemaining    int     `json:"TempBasalRemaining"`
		} `json:"extended"`
		Battery struct {
			Percent int `json:"percent"`
		} `json:"battery"`
	} `json:"pump"`
	User string `json:"-"`
}

type NsTreatment struct {
	CreatedAt    time.Time `json:"created_at"`
	EnteredBy    string    `json:"enteredBy"`
	EventType    string    `json:"eventType"`
	Carbs        int       `json:"carbs,omitempty"`
	Duration     int       `json:"duration,omitempty"`
	Insulin      float64   `json:"insulin,omitempty"`
	IsSMB        bool      `json:"isSMB,omitempty"`
	Notes        string    `json:"notes,omitempty"`
	Percent      int       `json:"percent,omitempty"`
	TargetTop    float64   `json:"targetTop,omitempty"`
	TargetBottom float64   `json:"targetBottom,omitempty"`
	Reason       string    `json:"reason,omitempty"`
	Rate         float64   `json:"rate,omitempty"`
	Units        string    `json:"units,omitempty"`
	User         string    `json:"-"`
}

type Config struct {
	NsUri        string `json:"ns-uri,omitempty"`
	NsToken      string `json:"ns-token,omitempty"`
	MongoUri     string `json:"mongo-uri,omitempty"`
	MongoDb      string `json:"mongo-db,omitempty"`
	Limit        int64  `json:"limit,omitempty"`
	Skip         int64  `json:"skip,omitempty"`
	InfluxUri    string `json:"influx-uri,omitempty"`
	InfluxToken  string `json:"influx-token,omitempty"`
	InfluxOrg    string `json:"influx-org,omitempty"`
	InfluxBucket string `json:"influx-bucket,omitempty"`
	Imports      []struct {
		NsUri    string `json:"ns-uri,omitempty"`
		NsToken  string `json:"ns-token,omitempty"`
		MongoUri string `json:"mongo-uri,omitempty"`
		MongoDb  string `json:"mongo-db,omitempty"`
		User     string `json:"user"`
	} `json:"imports,omitempty"`
}
