import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1
import Ubuntu.Components.Pickers 0.1
import Ubuntu.Layouts 0.1

Tab {
    title: i18n.tr("OpenReport")
    page: Page {        
        width: parent.width
        Grid {
            id: grid1
            width: parent.width
            columns :2
            rows: 2
            Button {//row 1
                    objectName: "openReport_OpenReportButton"
                    id: openReport_OpenReportButton
                    text: i18n.tr("Open Report")
		onClicked: {
                        ctrl.openReportOpenReportButtonClickedGo(openReport_OpenReportButton)
		}
            }
            TextField {
                    objectName: "openReport_FileNameEntry"
                    text: i18n.tr("")
                    width: parent.width - openReport_OpenReportButton.width
            }

            Button {
                        objectName: "openReport_CloseAppButton"
                        id: openReport_CloseAppButton
                        text: i18n.tr("Close App")
                        onClicked: {
                            console.log("clicked Close App...");
			    ctrl.openReportCloseAppButtonClickedGo(openReport_CloseAppButton)			    
                            Qt.quit()
                        }
            }
            Button {
                objectName: "openReport_NextButton"
                    id: openReport_NextButton
                    text: i18n.tr("Next")
                    onClicked: {
                        ctrl.openReportNextButtonClickedGo(openReport_NextButton)
                        tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                        console.log("clicked next..." + tabs.selectedTabIndex);
                    }
           }
        }
    }
}
