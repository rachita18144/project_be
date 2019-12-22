package oForm

import (
	"gopkg.in/mgo.v2"
	"log"
	"project_be/base"
)

/*

func sabseBadiPagalHuMain() OForm {
	var temp OForm
	temp.CaseNumber = "sabka time waste hoga"
	temp.Mobile = "hello"
	temp.NameOfPatient = "hello"
	temp.Amount = "hello"
	temp.Address = "hello"
	temp.DateOfAdmission = time.Now()
	temp.DateOfPtca = time.Now()
	temp.DateOfDischarge = time.Now()
	temp.PtcaHospital = "hello"
	temp.DoctorName = "hello"
	temp.History = "hello"
	temp.RiskFactors = []string{"hello"}
	temp.Ecg = []string{"hello"}
	temp.EchoComments = "hello"
	temp.TypeOfVesselInvolved = []string{"hello"}
	temp.CagReport = "hello"
	temp.ThrombusSeen = "hello"
	temp.ThrombusAspirationDone = "hello"
	temp.NumberOfVessels = "hello"
	temp.Ptca = []string{"hello"}
	return temp
}*/

func saveOFormDetailsEntity(formDetails OForm) error {
	log.Println(formDetails.CaseNumber)

	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	defer session.Close()

	c := session.DB(base.DB_PORTAL).C(base.COL_OFORM)
	//data, err := bson.Marshal(&formDetails)
	//if err != nil {
	//	log.Println("Error:", err)
	//	return err
	//}
	if err = c.Insert(formDetails); err != nil {
		log.Println("ERROR:", err)
	}
	return err
}
