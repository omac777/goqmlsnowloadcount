import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1
import Ubuntu.Components.Pickers 0.1
import Ubuntu.Layouts 0.1

Tab {
    title: i18n.tr("Guard Shift Start Date")
    page: Page {        
        width: parent.width

        Grid {
            id: guardShiftStartDate_grid1
            width: parent.width
            columns : 1
            rows: 2

            Label {
                id: guardShiftStartDate_ShiftDateLabel
                text: "Shift Start Date: " + Qt.formatDate(guardShiftStartDate_ShiftStartDatePicker.date, "dd-MMM-yyyy")
                width: parent.width
            }

            DatePicker {
                id: guardShiftStartDate_ShiftStartDatePicker
                mode: "Years|Months|Days"
                width: parent.width
            }

	    function setGuardShiftStartDateToToday() {
	    	var d = new Date();
	    	guardShiftStartDate_ShiftStartDatePicker.year = d.getFullYear();
	    	guardShiftStartDate_ShiftStartDatePicker.month = d.getMonth();
	    	guardShiftStartDate_ShiftStartDatePicker.day = d.getDate();
	    }

            Grid {
		id: guardShiftStartDate_grid2
		y: guardShiftStartDate_grid1.height
		columns :2
		rows: 1
		Button {
                    objectName: "guardShiftStartDate_PrevButton"
                    id: guardShiftStartDate_PrevButton
                    text: i18n.tr("Previous")
                    onClicked: {
			ctrl.guardShiftStartDatePreviousButtonClickedGo(guardShiftStartDate_PrevButton)
			tabs.selectedTabIndex = tabs.selectedTabIndex - 1
			console.log("clicked prev..." + tabs.selectedTabIndex);
                    }
		}
		Button {
                    objectName: "guardShiftStartDate_NextButton"
                    id: guardShiftStartDate_NextButton
                    text: i18n.tr("Next")
                    onClicked: {
			ctrl.guardShiftStartDateNextButtonClickedGo(guardShiftStartDate_NextButton)
                        tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                        console.log("clicked next..." + tabs.selectedTabIndex);
                    }
		}
            }
	}
    }
}
