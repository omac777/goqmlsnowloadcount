import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1
import Ubuntu.Components.Pickers 0.1
import Ubuntu.Layouts 0.1
//import "../components"

Tab {
    title: i18n.tr("GuardShiftInfo")
    page: Page {        
        width: parent.width

        Grid {
            id: guardShiftInfo_grid1
            columns :2
            rows: 4            
            Label { //row 1
                objectName: "guardShiftInfo_GuardNameLabel"
                id: guardShiftInfo_GuardNameLabel
                text: i18n.tr("Guard Name")
            }
            TextField {
                objectName: "guardShiftInfo_GuardNameEntry"
                id: guardShiftInfo_GuardNameEntry
                text: i18n.tr("")
            }            
            Label {//row 2
                objectName: "guardShiftInfo_GuardLicenseNumberLabel"
                id: guardShiftInfo_GuardLicenseNumberLabel
                text: i18n.tr("Guard License#")
            }
            TextField {
                objectName: "guardShiftInfo_GuardLicenseNumberEntry"
                    id: guardShiftInfo_GuardLicenseNumberEntry
                    text: i18n.tr("")
            }            
            Label {//row 3
                objectName: "guardShiftInfo_ShiftCommentsLabel"
                id: guardShiftInfo_ShiftCommentsLabel
                text: i18n.tr("Shift Comments")
            }
            TextField {
                objectName: "guardShiftInfo_ShiftCommentsEntry"
                    id: guardShiftInfo_ShiftCommentsEntry
                    text: i18n.tr("")
            }
            Button {//row4
                objectName: "guardShiftInfo_PrevButton"
                id: guardShiftInfo_PrevButton
                text: i18n.tr("Previous")
                onClicked: {
		    ctrl.guardShiftInfoPreviousButtonClickedGo(guardShiftInfo_PrevButton)
                    tabs.selectedTabIndex = tabs.selectedTabIndex - 1
                    console.log("clicked prev..." + tabs.selectedTabIndex);
                }
            }
            Button {
                objectName: "guardShiftInfo_NextButton"
                    id: guardShiftInfo_NextButton
                    text: i18n.tr("Next")
                    onClicked: {
		    ctrl.guardShiftInfoNextButtonClickedGo(guardShiftInfo_NextButton)
                        tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                        console.log("clicked next..." + tabs.selectedTabIndex);
                    }
           }
        }
    }
}
