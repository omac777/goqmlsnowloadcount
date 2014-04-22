//snowloadcountmain.go
package main

import (
        "fmt"
        "gopkg.in/qml.v0"
        "os"
)

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
            Messagebt1:"bt1inithey"}
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
        //all the stuff in gotype could be here
        //to facilitate the setting/getting
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

func (ctrl *Control) OpenReportNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in OpenReportNextButtonClickedGo():")
}

func (ctrl *Control) OpenReportCloseAppButtonClickedGo(theButton qml.Object) {
        fmt.Println("in OpenReportCloseAppButtonClickedGo:")
}

func (ctrl *Control) OpenReportOpenReportButtonClickedGo(theButton qml.Object) {
        fmt.Println("in OpenReportOpenReportButtonClickedGo:")
}

func (ctrl *Control) GuardShiftStartTimePreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftStartTimePreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftStartTimeNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftStartTimeNextButtonClickedGo:")
}

func (ctrl *Control) GuardShiftEndTimePreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftEndTimePreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftEndTimeNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftEndTimeNextButtonClickedGo:")
}

func (ctrl *Control) GuardShiftStartDatePreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftStartDatePreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftStartDateNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftStartDateNextButtonClickedGo:")
}

func (ctrl *Control) GuardShiftInfoPreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftInfoPreviousButtonClickedGo:")
}

func (ctrl *Control) GuardShiftInfoNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardShiftInfoNextButtonClickedGo:")
}

func (ctrl *Control) GuardLocationPreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardLocationPreviousButtonClickedGo:")
}

func (ctrl *Control) GuardLocationNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in GuardLocationNextButtonClickedGo:")
}

func (ctrl *Control) AdmissionTypePreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in AdmissionTypePreviousButtonClickedGo:")
}

func (ctrl *Control) AdmissionTypeNextButtonClickedGo(theButton qml.Object) {
        fmt.Println("in AdmissionTypeNextButtonClickedGo:")
}

func (ctrl *Control) TruckTypePreviousButtonClickedGo(theButton qml.Object) {
        fmt.Println("in TruckTypePreviousButtonClickedGo:")
}

func (ctrl *Control) TruckTypeFinishButtonClickedGo(theButton qml.Object) {
        fmt.Println("in TruckTypeFinishButtonClickedGo:")
}
