//snowloadcountmain.go
package main

import (
        "fmt"
        "gopkg.in/qml.v0"
        "os"
	"time"
	"bufio"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
	"bytes"
)

const SHIFTSTARTTIME = string("shiftstarttime")
const SHIFTENDTIME = string("shiftendtime")
const SHIFTSTARTDATE = string("shiftstartdate")
const GUARDNAME = string("guardname")
const LICENSENUMBER = string("licensenumber")
const SHIFTCOMMENT = string("shiftcomment")
const COUNTLOCATION = string("countlocation")
const COUNTFORITEMTYPE = string("countforitemtype")
const SINGLEAXLE = string("singleaxle")
const TANDEMAXLE = string("tandem2axle")
const TRIPLEAXLE = string("tripleaxle")
const COMBOTRUCK = string("combotruck")
const SEMITRAILER = string("semitrailer")

//json stuff has to have fields with names starting with uppercase letters
//this is a golang special wierdness
//to be honest, I'm not happy about this
//We should have complete control over our variable names
type SNLDB struct {
	SNLMap map[string]string `json:"SNLMap"`
	SingleL []time.Time `json:"SingleL"` 
	TandemL []time.Time `json:"TandemL"` 
	TripleL []time.Time `json:"TripleL"` 
	ComboL []time.Time `json:"ComboL"` 
	SemiL []time.Time `json:"SemiL"`
}

func NewSNLDB() *SNLDB {
	f := SNLDB{}
	f.SNLMap = make(map[string]string)
	f.SingleL = make([]time.Time, 0, 200) //average maximum number trucks counted in one night
	f.TandemL = make([]time.Time, 0, 200)
	f.TripleL = make([]time.Time, 0, 200)
	f.ComboL = make([]time.Time, 0, 200)
	f.SemiL = make([]time.Time, 0, 200)
	return &f
}

func (s *SNLDB) testSetAndGetDataFields() {
	s.setShiftStartTime("b");
	s.setShiftEndTime("c");
	s.setShiftStartDate("d");
	s.setGuardName("e");
	s.setLicenseNumber("f");
	s.setShiftComment("g");
	s.setCountLocation("h");
	s.setCountForItemType("i");
	s.singleAxleArrived();
	s.tandemAxleArrived();
	s.tandemAxleArrived();
	s.tripleAxleArrived();
	s.tripleAxleArrived();
	s.tripleAxleArrived();
	s.comboTruckArrived();
	s.comboTruckArrived();
	s.comboTruckArrived();
	s.comboTruckArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	s.semiTrailerArrived();
	fmt.Printf("%v\n", s.getShiftStartTime())
	fmt.Printf("%v\n", s.getShiftEndTime())
	fmt.Printf("%v\n", s.getShiftStartDate())
	fmt.Printf("%v\n", s.getGuardName())
	fmt.Printf("%v\n", s.getLicenseNumber())
	fmt.Printf("%v\n", s.getShiftComment())
	fmt.Printf("%v\n", s.getCountLocation())
	fmt.Printf("%v\n", s.getCountForItemType())
	fmt.Printf("single total: %v\n", s.getSingleAxleTotal())
	fmt.Printf("tandem total: %v\n", s.getTandemAxleTotal())
	fmt.Printf("triple total: %v\n", s.getTripleAxleTotal())
	fmt.Printf("combo total: %v\n" , s.getComboTruckTotal())
	fmt.Printf("semi total: %v\n", s.getSemiTrailerTotal())
}

func (s *SNLDB) debugDataFields() {
	fmt.Printf("%v\n", s.getShiftStartTime())
	fmt.Printf("%v\n", s.getShiftEndTime())
	fmt.Printf("%v\n", s.getShiftStartDate())
	fmt.Printf("%v\n", s.getGuardName())
	fmt.Printf("%v\n", s.getLicenseNumber())
	fmt.Printf("%v\n", s.getShiftComment())
	fmt.Printf("%v\n", s.getCountLocation())
	fmt.Printf("%v\n", s.getCountForItemType())
	fmt.Printf("single total: %v\n", s.getSingleAxleTotal())
	fmt.Printf("tandem total: %v\n", s.getTandemAxleTotal())
	fmt.Printf("triple total: %v\n", s.getTripleAxleTotal())
	fmt.Printf("combo total: %v\n" , s.getComboTruckTotal())
	fmt.Printf("semi total: %v\n", s.getSemiTrailerTotal())
}

func (s *SNLDB) getShiftStartTime() string {
	return string(s.SNLMap[SHIFTSTARTTIME])
}

func (s *SNLDB) getShiftEndTime() string {
	return string(s.SNLMap[SHIFTENDTIME])
}

func (s *SNLDB) getShiftStartDate() string {
	return string(s.SNLMap[SHIFTSTARTDATE])
}

func (s *SNLDB) getGuardName() string {
	return string(s.SNLMap[GUARDNAME])
}

func (s *SNLDB) getLicenseNumber() string {
	return string(s.SNLMap[LICENSENUMBER])
}

func (s *SNLDB) getShiftComment() string {
	return string(s.SNLMap[SHIFTCOMMENT])
}

func (s *SNLDB) getCountLocation() string {
	return string(s.SNLMap[COUNTLOCATION])
}

func (s *SNLDB) getCountForItemType() string {
	return string(s.SNLMap[COUNTFORITEMTYPE])
}

func (s *SNLDB) setShiftStartTime(s_ string) {
	s.SNLMap[SHIFTSTARTTIME] = s_
}

func (s *SNLDB) setShiftEndTime(s_ string) {
	s.SNLMap[SHIFTENDTIME] = s_
}

func (s *SNLDB) setShiftStartDate(s_ string) {
	s.SNLMap[SHIFTSTARTDATE] = s_
}

func (s *SNLDB) setGuardName(s_ string) {
	s.SNLMap[GUARDNAME] = s_
}

func (s *SNLDB) setLicenseNumber(s_ string) {
	s.SNLMap[LICENSENUMBER] = s_
}

func (s *SNLDB) setShiftComment(s_ string) {
	s.SNLMap[SHIFTCOMMENT] = s_
}

func (s *SNLDB) setCountLocation(s_ string) {
	s.SNLMap[COUNTLOCATION] = s_
}

func (s *SNLDB) setCountForItemType(s_ string) {
	s.SNLMap[COUNTFORITEMTYPE] = s_
}

func (s *SNLDB) getSingleAxleTotal() int {
	return len(s.SingleL)
}

func (s *SNLDB) getTandemAxleTotal() int {
	return len(s.TandemL)
}

func (s *SNLDB) getTripleAxleTotal() int {
	return len(s.TripleL)
}

func (s *SNLDB) getComboTruckTotal() int {
	return len(s.ComboL)
}

func (s *SNLDB) getSemiTrailerTotal() int {
	return len(s.SemiL)
}

func (s *SNLDB) singleAxleArrived() {
	s.SingleL = append(s.SingleL, time.Now())
}

func (s *SNLDB) tandemAxleArrived() {
	s.TandemL = append(s.TandemL, time.Now())
}

func (s *SNLDB) tripleAxleArrived() {
	s.TripleL = append(s.TripleL, time.Now())	
}

func (s *SNLDB) comboTruckArrived() {
	s.ComboL = append(s.ComboL, time.Now())	
}

func (s *SNLDB) semiTrailerArrived() {
	s.SemiL = append(s.SemiL, time.Now())	
}

func (s *SNLDB) clear() {
	s.SNLMap = make(map[string]string)
	s.SingleL = make([]time.Time, 0, 200) //average maximum number trucks counted in one night
	s.TandemL = make([]time.Time, 0, 200)
	s.TripleL = make([]time.Time, 0, 200)
	s.ComboL = make([]time.Time, 0, 200)
	s.SemiL = make([]time.Time, 0, 200)
}

func (ctrl *Control) saveJsonFileSNLDB(myFileName string, f *SNLDB) () {
	fo, err := os.Create(myFileName)
	if err != nil { panic(err) }
	defer fo.Close()
	w := bufio.NewWriter(fo)
 	var myjsonBytes bytes.Buffer
 	enc := json.NewEncoder(&myjsonBytes)
 	err = enc.Encode(f)
	//fmt.Println("myjsonbytes %v\n", myjsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(myjsonBytes.Bytes()) 
	if err != nil {
		log.Fatal(err)
	}
	if err = w.Flush(); err != nil { 
		log.Fatal(err) 
	}
}

func (ctrl *Control) readJsonFileSNLDB(myFileName string) (*SNLDB) {
	file, err2 := ioutil.ReadFile(myFileName)
	if err2 != nil {
	   	log.Fatal(err2)
	}
	jsonBytes := bytes.NewBuffer(file)
 	dec := json.NewDecoder(jsonBytes)
 	var res SNLDB
 	err3 := dec.Decode(&res)
	if err3 != nil {
	   	log.Fatal(err3)
	}
 	//fmt.Printf("Read Results: %+v %v\n", res, err3)
	return &res
}

func main() {
        if err := run(); err != nil {
                fmt.Fprintf(os.Stderr, "error: %v\n", err)
                os.Exit(1)
        }
}

func run() error {
        qml.Init(nil)
        engine := qml.NewEngine()
        engine.On("quit", func() { os.Exit(0) })
        component2, err := engine.LoadFile("snowloadcount.qml")
        if err != nil {
                return err
        }
        window := component2.CreateWindow(nil)
        ctrl := Control{
		Messagetf1:"tf1inithey",
		Messagebt1:"bt1inithey",
		SingleAxleMessage:"0",
		TandemAxleMessage:"0",
		TripleAxleMessage:"0",
		ComboAxleMessage:"0",
		SemiAxleMessage:"0",
		IsLoadedReport:false,
		LoadedReportFilename:"",
		Snldb:NewSNLDB()}
        context := engine.Context()
        context.SetVar("ctrl", &ctrl)
        ctrl.Root = window.Root()

        window.Show()
        window.Wait()
        return nil
}

type Control struct {
        Root    qml.Object
        Messagebt1 string
        Messagetf1 string
	SingleAxleMessage string
	TandemAxleMessage string
	TripleAxleMessage string
	ComboAxleMessage string
	SemiAxleMessage string
        //all the stuff in gotype could be here
        //to facilitate the setting/getting
	IsLoadedReport bool
	LoadedReportFilename string
	Snldb *SNLDB
}

//save one field from the gui wizard
func (ctrl *Control) Savetf1contents(text qml.Object) {
        fmt.Println("in Savetf1contents():")
        fmt.Println("text:", text.String("text"))
}

//save another field from the gui wizard
func (ctrl *Control) Savebt1contents(text qml.Object) {
        fmt.Println("in Savebt1contents():")
        fmt.Println("text:", text.String("text"))
}

func (ctrl *Control) Loadtf1contents(text qml.Object) {
        fmt.Println("in Loadtf1contents():")
        fmt.Println("text:", text.String("text"))
        go func() {
		ctrl.Messagetf1 = "loaded from tf1..."
		qml.Changed(ctrl, &ctrl.Messagetf1)
        }()
}

func (ctrl *Control) Loadbt1contents(text qml.Object) {
        fmt.Println("in Loadbt1contents():")
        fmt.Println("text:", text.String("text"))
        go func() {
		ctrl.Messagebt1 = "loaded from bt1"
		qml.Changed(ctrl, &ctrl.Messagebt1)
        }()
} 

func (ctrl *Control) OpenReportNextButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in OpenReportNextButtonClickedGo():")
}

func (ctrl *Control) OpenReportCloseAppButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in OpenReportCloseAppButtonClickedGo:")
}

func (ctrl *Control) OpenReportOpenReportButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in OpenReportOpenReportButtonClickedGo:")
	//openReport_FileNameEntry
	filename_ := theQmlObject.String("text")
	ctrl.Snldb = ctrl.readJsonFileSNLDB(filename_)
	ctrl.IsLoadedReport = true
	ctrl.LoadedReportFilename = filename_
	go func() {
		ctrl.loadReportDataIntoGui()
	}()
}

func (ctrl *Control) GuardShiftStartTimePreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftStartTimePreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftStartTimeNextButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftStartTimeNextButtonClickedGo:")
	//guardShiftStartTime_label
	starttime_ := theQmlObject.String("text")
	ctrl.Snldb.setShiftStartTime(starttime_);
}

func (ctrl *Control) GuardShiftEndTimePreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftEndTimePreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftEndTimeNextButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftEndTimeNextButtonClickedGo:")
	//guardShiftEndTime_label
	endtime_ := theQmlObject.String("text")
	ctrl.Snldb.setShiftEndTime(endtime_);
}

func (ctrl *Control) GuardShiftStartDatePreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftStartDatePreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftStartDateNextButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftStartDateNextButtonClickedGo:")
	//guardShiftStartDate_ShiftDateLabel
	startDate_ := theQmlObject.String("text")
	ctrl.Snldb.setShiftStartDate(startDate_);
}

func (ctrl *Control) GuardShiftInfoPreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftInfoPreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftInfoSNLDBSetGuardNameGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftInfoNextButtonClickedGo:")
	//guardShiftInfo_GuardNameEntry
	guardName_ := theQmlObject.String("text")
	ctrl.Snldb.setGuardName(guardName_);
}
func (ctrl *Control) GuardShiftInfoSNLDBSetLicenseNumberGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftInfoNextButtonClickedGo:")
	//guardShiftInfo_GuardLicenseNumberEntry
	licenseNumber_ := theQmlObject.String("text")
	ctrl.Snldb.setLicenseNumber(licenseNumber_);
}
func (ctrl *Control) GuardShiftInfoSNLDBSetShiftCommentGo(theQmlObject qml.Object) {
        fmt.Println("in GuardShiftInfoNextButtonClickedGo:")
	//guardShiftInfo_ShiftCommentsEntry
	shiftComment_ := theQmlObject.String("text")
	ctrl.Snldb.setShiftComment(shiftComment_);
}

func (ctrl *Control) GuardLocationPreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in GuardLocationPreviousButtonClickedGo:")
}

func (ctrl *Control) GuardLocationNextButtonClickedGo(theString string) {
        fmt.Println("in GuardLocationNextButtonClickedGo:")
	ctrl.Snldb.setCountLocation(theString);
}

func (ctrl *Control) AdmissionTypePreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in AdmissionTypePreviousButtonClickedGo:")
}

func (ctrl *Control) AdmissionTypeNextButtonClickedGo(theString string) {
        fmt.Println("in AdmissionTypeNextButtonClickedGo:")
	ctrl.Snldb.setCountForItemType(theString);
}


func (ctrl *Control) TruckTypeSingleAxleClickedGo(TruckType_SingleAxleLabel qml.Object) {
        fmt.Println("in TruckTypeSingleAxleClickedGo:")
	ctrl.Snldb.singleAxleArrived();
        go func() {
		ctrl.SingleAxleMessage = strconv.Itoa( ctrl.Snldb.getSingleAxleTotal() )
		qml.Changed(ctrl, &ctrl.SingleAxleMessage)
        }()
}

func (ctrl *Control) TruckTypeTandemAxleClickedGo(TruckType_TandemAxleLabel qml.Object) {
        fmt.Println("in TruckTypeTandemAxleClickedGo:")
	ctrl.Snldb.tandemAxleArrived();
        go func() {
		ctrl.TandemAxleMessage = strconv.Itoa( ctrl.Snldb.getTandemAxleTotal() )
		qml.Changed(ctrl, &ctrl.TandemAxleMessage)
        }()
}

func (ctrl *Control) TruckTypeTripleAxleClickedGo(TruckType_TripleAxleLabel qml.Object) {
        fmt.Println("in TruckTypeTripleAxleClickedGo:")
	ctrl.Snldb.tripleAxleArrived();
        go func() {
		ctrl.TripleAxleMessage = strconv.Itoa( ctrl.Snldb.getTripleAxleTotal() )
		qml.Changed(ctrl, &ctrl.TripleAxleMessage)
        }()
}

func (ctrl *Control) TruckTypeComboAxleClickedGo(TruckType_ComboAxleLabel qml.Object) {
        fmt.Println("in TruckTypeComboAxleClickedGo:")
	ctrl.Snldb.comboTruckArrived();
        go func() {
		ctrl.ComboAxleMessage = strconv.Itoa( ctrl.Snldb.getComboTruckTotal() )
		qml.Changed(ctrl, &ctrl.ComboAxleMessage)
        }()
}

func (ctrl *Control) TruckTypeSemiAxleClickedGo(TruckType_SemiAxleLabel qml.Object) {
        fmt.Println("in TruckTypeSemiAxleClickedGo:")
	ctrl.Snldb.semiTrailerArrived();
        go func() {
		ctrl.SemiAxleMessage = strconv.Itoa( ctrl.Snldb.getSemiTrailerTotal() )
		qml.Changed(ctrl, &ctrl.SemiAxleMessage)
        }()
}

func (ctrl *Control) TruckTypePreviousButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in TruckTypePreviousButtonClickedGo:")
}

func (ctrl *Control) TruckTypeFinishButtonClickedGo(theQmlObject qml.Object) {
        fmt.Println("in TruckTypeFinishButtonClickedGo:")
	ctrl.saveJsonFileSNLDB(ctrl.generateUniqueReportFilename(), ctrl.Snldb)
}

func (ctrl *Control) loadReportDataIntoGui() {
        //taken from the go-qml particle example 
        //important qml object navigation stuff
        //1)get the qml object class factory with id "emitterComponent"
	//component := ctrl.Root.Object("emitterComponent")
        
        //2)create an object instance of it
	//emitter := component.Create(nil)
        
        //3)set the property "x" with that object        
        //emitter.Set("x", x)

        //4)get the objectinstance "xAnim", then call its function named "start"
	//emitter.ObjectByName("xAnim").Call("start")

	//sa.shiftStartTimeEntry.SetText(ctrl.Snldb.getShiftStartTime())
	//ctrl.Root.ObjectByName("")
	//guardShiftStartTime_ShiftStartTimePicker

	//sa.shiftEndTimeEntry.SetText(ctrl.Snldb.getShiftEndTime())
	//ctrl.Root.ObjectByName("")
	//guardShiftEndTime_ShiftEndTimePicker

	//sa.shiftStartDateEntry.SetText(ctrl.Snldb.getShiftStartDate())
	//ctrl.Root.ObjectByName("")
	//guardShiftStartDate_ShiftStartDatePicker	

	//sa.guardNameEntry.SetText(ctrl.Snldb.getGuardName())
	//ctrl.Root.ObjectByName("")
	//guardShiftInfo_GuardNameEntry

	//sa.guardLicenceNumberEntry.SetText(ctrl.Snldb.getLicenseNumber())
	//ctrl.Root.ObjectByName("")
	//guardShiftInfo_GuardLicenseNumberEntry

	//sa.guardShiftCommentsEntry.SetText(ctrl.Snldb.getShiftComment())
	//ctrl.Root.ObjectByName("")
	//guardShiftInfo_ShiftCommentsEntry

	//sa.setActiveCountLocation(ctrl.Snldb.getCountLocation())
	//ctrl.Root.ObjectByName("")
	//countingAtLocationSelector
	//countingAtLocationSelector.model[countingAtLocationSelector.selectedIndex]

	//sa.setActiveCountForItemType(ctrl.Snldb.getCountForItemType())
	//ctrl.Root.ObjectByName("")
	//countingTotalsForOptionSelector
	//countingTotalsForOptionSelector.model[countingTotalsForOptionSelector.selectedIndex]

	//sa.SingleLabel.SetText(strconv.Itoa(ctrl.Snldb.getSingleAxleTotal()))
        // go func() {
	// 	ctrl.SingleAxleMessage = strconv.Itoa( ctrl.Snldb.getSingleAxleTotal() )
	// 	qml.Changed(ctrl, &ctrl.SingleAxleMessage)
        // }()

	//sa.TandemLabel.SetText(strconv.Itoa(ctrl.Snldb.getTandemAxleTotal()))
        // go func() {
	// 	ctrl.TandemAxleMessage = strconv.Itoa( ctrl.Snldb.getTandemAxleTotal() )
	// 	qml.Changed(ctrl, &ctrl.TandemAxleMessage)
        // }()

	//sa.TripleLabel.SetText(strconv.Itoa(ctrl.Snldb.getTripleAxleTotal()))
        // go func() {
	// 	ctrl.TripleAxleMessage = strconv.Itoa( ctrl.Snldb.getTripleAxleTotal() )
	// 	qml.Changed(ctrl, &ctrl.TripleAxleMessage)
        // }()

	//sa.ComboLabel.SetText(strconv.Itoa(ctrl.Snldb.getComboTruckTotal()))
        // go func() {
	// 	ctrl.ComboAxleMessage = strconv.Itoa( ctrl.Snldb.getComboTruckTotal() )
	// 	qml.Changed(ctrl, &ctrl.ComboAxleMessage)
        // }()

	//sa.SemiLabel.SetText(strconv.Itoa(ctrl.Snldb.getSemiTrailerTotal()))
        // go func() {
	// 	ctrl.SemiAxleMessage = strconv.Itoa( ctrl.Snldb.getSemiTrailerTotal() )
	// 	qml.Changed(ctrl, &ctrl.SemiAxleMessage)
        // }()

	//disable all the fields after filling them/changing them while in loaded report mode.
	//sa.shiftStartTimeEntry.SetSensitive(false)
	//sa.shiftEndTimeEntry.SetSensitive(false)
	//sa.shiftStartDateEntry.SetSensitive(false)
	//sa.guardNameEntry.SetSensitive(false)
	//sa.guardLicenceNumberEntry.SetSensitive(false)
	//sa.guardShiftCommentsEntry.SetSensitive(false)
	
	//sa.conroyRadio.SetSensitive(false)
	//sa.michaelRadio.SetSensitive(false)
	//sa.strandherdRadio.SetSensitive(false)
	//sa.innesRadio.SetSensitive(false)
	//sa.clydeRadio.SetSensitive(false)
	
	//sa.passesRadio.SetSensitive(false)
	//sa.ticketsRadio.SetSensitive(false)
	
	//sa.singleaxlebutton.SetSensitive(false)
	//sa.tandemaxlebutton.SetSensitive(false)
	//sa.tripleaxlebutton.SetSensitive(false)
	//sa.combotruckbutton.SetSensitive(false)
	//sa.semitrailerbutton.SetSensitive(false)
}

func (ctrl *Control) generateUniqueReportFilename() string {
	var myUniqueFilename string
	var t time.Time
	var myIsoTime string
	t = time.Now()
	myIsoTime = fmt.Sprintf("%d%02d%02dT%02d%02d%02d.%09d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	myUniqueFilename = "snowloadcount" + myIsoTime + ".json";
	return myUniqueFilename
}
