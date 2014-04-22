import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1

Tab {
    title: i18n.tr("AdmissionType")
    page: Page {
        width: parent.width
        OptionSelector {
            id: countingTotalsForOptionSelector
            width: parent.width
            text: i18n.tr("Counting Totals For:")
            model: [i18n.tr("Passes"),
                    i18n.tr("Tickets")]
        }
        Grid {
            y: countingTotalsForOptionSelector.height
            width: parent.width
            columns :2
            rows: 1
            Button {
                objectName: "admissionType_PrevButton"
                id: admissionType_PrevButton
                text: i18n.tr("Previous")
                onClicked: {
		    ctrl.admissionTypePreviousButtonClickedGo(admissionType_PrevButton)
                    tabs.selectedTabIndex = tabs.selectedTabIndex - 1
                    console.log("clicked prev..." + tabs.selectedTabIndex);
                }
            }
            Button {
                objectName: "admissionType_NextButton"
                id: admissionType_NextButton
                text: i18n.tr("Next")
                onClicked: {
		    ctrl.admissionTypeNextButtonClickedGo(admissionType_NextButton)
		    console.log("selected:" + countingTotalsForOptionSelector.model[countingTotalsForOptionSelector.selectedIndex]);
                    tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                    console.log("clicked next..." + tabs.selectedTabIndex);
                }
            }
        }
    }
}
