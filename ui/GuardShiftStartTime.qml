import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1
import Ubuntu.Components.Pickers 0.1
import Ubuntu.Layouts 0.1

Tab {
    title: i18n.tr("Guard Shift Start Time")
    page: Page {        
        width: parent.width
        Grid {
            id: grid1
            width: parent.width
            columns :1
            rows: 2
            Label { //row1
	        id: guardShiftStartTime_label
                text: Qt.formatTime(guardShiftStartTime_ShiftStartTimePicker.date, "hh:mm:ss")
                width: parent.width
            }
            DatePicker { //row2
	    	objectName: "guardShiftStartTime_ShiftStartTimePicker"
                id: guardShiftStartTime_ShiftStartTimePicker
                mode: "Hours|Minutes|Seconds"
                width: parent.width
            }

	    //function setGuardShiftStartTime(hour_, minute_, second_) {
		//var d = new Date();
		//guardShiftStartTime_ShiftStartTimePicker.date = new Date(d.year, d.month - 1, d.day , hour_, minute_, second_)
	    	//var d = new Date();
	    	// guardShiftStartDate_ShiftStartDatePicker.hour = d.hour();
	    	// guardShiftStartDate_ShiftStartDatePicker.minute = d.minute();
	    	// guardShiftStartDate_ShiftStartDatePicker.second = d.second();
		//var mylocale: Qt.locale()
		//string mydateTimeString: "2013-09-17 10:56:06"
		//var d = Date.fromLocaleString(mylocale, mydateTimeString, "yyyy-MM-dd hh:mm:ss")
	    //}
        }

        Grid {
            id: grid2
            y: grid1.height
            columns :2
            rows: 1
            Button {
                objectName: "guardShiftStartTime_PrevButton"
                id: guardShiftStartTime_PrevButton
                text: i18n.tr("Previous")
                onClicked: {
                    ctrl.guardShiftStartTimePreviousButtonClickedGo(guardShiftStartTime_PrevButton)
                    tabs.selectedTabIndex = tabs.selectedTabIndex - 1
                    console.log("clicked prev..." + tabs.selectedTabIndex);
                }
            }
            Button {
                objectName: "guardShiftStartTime_NextButton"
                id: guardShiftStartTime_NextButton
                text: i18n.tr("Next")
                onClicked: {
                    ctrl.guardShiftStartTimeNextButtonClickedGo(guardShiftStartTime_label)
                    tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                    console.log("clicked next..." + tabs.selectedTabIndex);
                }
            }
        }
    }
}
