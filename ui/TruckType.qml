import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1

Tab {
    title: i18n.tr("TruckType Count")

    page: Page {
        width: parent.width

        Grid {
            width: parent.width
            columns :2
            rows: 6
            Button {//row 1
                objectName: "TruckType_SingleAxleButton"
		id: truckType_SingleAxleButton
                text: i18n.tr("singleaxle")
                onClicked: {
		    ctrl.truckTypeSingleAxleClickedGo(truckType_SingleAxleLabel)
                }
            }
            Label {
                objectName: "TruckType_SingleAxleLabel"
		id: truckType_SingleAxleLabel
                text: i18n.tr("0")
            }
            Button {//row 2
                objectName: "TruckType_TandemAxleButton"
		id: truckType_TandemAxleButton
                text: i18n.tr("tandemaxle")
                onClicked: {
		    ctrl.truckTypeTandemAxleClickedGo(truckType_TandemAxleLabel)
                }
            }
            Label {
                objectName: "TruckType_TandemAxleLabel"
		id: truckType_TandemAxleLabel
                text: i18n.tr("0")
            }            
            Button {//row 3
                objectName: "TruckType_TripleAxleButton"
		id: truckType_TripleAxleButton
                text: i18n.tr("tripleaxle")
                onClicked: {
		    ctrl.truckTypeTripleAxleClickedGo(truckType_TripleAxleLabel)
                }
            }
            Label {
                objectName: "TruckType_TripleAxleLabel"
		id: truckType_TripleAxleLabel
                text: i18n.tr("0")
            }            
            Button {//row 4
                objectName: "TruckType_ComboAxleButton"
		id: truckType_ComboAxleButton
                text: i18n.tr("comboaxle")
                onClicked: {
		    ctrl.truckTypeComboAxleClickedGo(truckType_ComboAxleLabel)
                }
            }
            Label {
                objectName: "TruckType_ComboAxleLabel"
		id: truckType_ComboAxleLabel
                text: i18n.tr("0")
            }            
            Button {//row 5
                objectName: "TruckType_SemiAxleButton"
		id: truckType_SemiAxleButton
                text: i18n.tr("semiaxle")
                onClicked: {
		    ctrl.truckTypeSemiAxleClickedGo(truckType_SemiAxleLabel)
                }
            }
            Label {
                objectName: "TruckType_SemiAxleLabel"
		id: truckType_SemiAxleLabel
                text: i18n.tr("0")
            }
            Button {//row 6
                objectName: "TruckType_PrevButton"
                id: truckType_PrevButton
                text: i18n.tr("Previous")
                onClicked: {
		    ctrl.truckTypePreviousButtonClickedGo(truckType_PrevButton)
                    tabs.selectedTabIndex = tabs.selectedTabIndex - 1
                    console.log("clicked prev..." + tabs.selectedTabIndex);
                }
            }
            Button {
                objectName: "TruckType_FinishButton"
                id: truckType_FinishButton
                text: i18n.tr("Finish")
                onClicked: {
		    ctrl.truckTypeFinishButtonClickedGo(truckType_FinishButton)
                    console.log("clicked finish...");
                    Qt.quit()
                }
            }
        }
    }
}
