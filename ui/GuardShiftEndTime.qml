import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1
import Ubuntu.Components.Pickers 0.1
import Ubuntu.Layouts 0.1

Tab {
    title: i18n.tr("Guard Shift End Time")
    page: Page {        
        width: parent.width
        Grid {
            id: grid1
            width: parent.width
            columns :1
            rows: 2
            Label { //row1
                text: "Shift Start Time: " + Qt.formatTime(guardShiftEndTime_ShiftEndTimePicker.date, "hh:mm:ss")
                width: parent.width
            }
            DatePicker { //row2
                id: guardShiftEndTime_ShiftEndTimePicker
                mode: "Hours|Minutes|Seconds"
                width: parent.width
            }
        }


        Grid {
            id: grid2
            y: grid1.height
            columns :2
            rows: 1
            Button {
                objectName: "guardShiftEndTime_PrevButton"
                id: guardShiftEndTime_PrevButton
                text: i18n.tr("Previous")
                onClicked: {
		    ctrl.guardShiftEndTimePreviousButtonClickedGo(guardShiftEndTime_PrevButton)
                    tabs.selectedTabIndex = tabs.selectedTabIndex - 1
                    console.log("clicked prev..." + tabs.selectedTabIndex);
                }
            }
            Button {
                objectName: "guardShiftEndTime_NextButton"
                    id: guardShiftEndTime_NextButton
                    text: i18n.tr("Next")
                    onClicked: {
			ctrl.guardShiftEndTimeNextButtonClickedGo(guardShiftEndTime_NextButton)
                        tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                        console.log("clicked next..." + tabs.selectedTabIndex);
                    }
           }

        }

    }
}
