package oForm

import "time"

type OForm struct {
	CaseNumber             string           `bson:"case_number"              json:"case_number"`
	Mobile                 string           `bson:"mobile"                   json:"mobile"`
	NameOfPatient          string           `bson:"name_of_patient"          json:"name_of_patient"`
	Amount                 string           `bson:"amount"                   json:"amount"`
	Address                string           `bson:"address"                  json:"address"`
	DateOfAdmission        time.Time        `bson:"date_of_admission"        json:"date_of_admission"`
	DateOfPtca             time.Time        `bson:"date_of_ptca"             json:"date_of_ptca"`
	DateOfDischarge        time.Time        `bson:"date_of_discharge"        json:"date_of_discharge"`
	PtcaHospital           string           `bson:"ptca_hospital"            json:"ptca_hospital"`
	DoctorName             string           `bson:"doctor_name"              json:"doctor_name"`
	History                string           `bson:"history"                  json:"history"`
	RiskFactors            []string         `bson:"risk_factors"             json:"risk_factors"`
	Investigations         []Investigations `bson:"investigations"           json:"investigations"`
	Ecg                    []string         `bson:"ecg"                      json:"ecg"`
	EchoComments           string           `bson:"echo_comments"            json:"echo_comments"`
	TypeOfVesselInvolved   []string         `bson:"type_of_vessel_involved"  json:"type_of_vessel_involved"`
	CagReport              string           `bson:"cag_report"               json:"cag_report"`
	ThrombusSeen           string           `bson:"thrombus_seen"            json:"thrombus_seen"`
	ThrombusAspirationDone string           `bson:"thrombus_aspiration_done" json:"thrombus_aspiration_done"`
	NumberOfVessels        string           `bson:"number_of_vessels"        json:"number_of_vessels"`
	Ptca                   []string         `bson:"ptca"                     json:"ptca"`
}

type Investigations struct {
	BloodPressure   string `json:"blood_pressure"`
	BloodSugar      string `json:"blood_sugar"`
	Ldl             string `json:"ldl"`
	Cholesterol     string `json:"cholesterol"`
	Lvef            string `json:"lvef"`
	Tg              string `json:"tg"`
	SerumCreatinine string `json:"serum_creatinine"`
	Hb              string `json:"hb"`
	Platelets       string `json:"platelets"`
}

type Test struct {
	CaseNumber string `bson:"case_number" json:"case_number"`
}

func SaveOFormDetailsBoundary(formDetails OForm) error {
	return saveOFormDetailsInteractor(formDetails)
}
