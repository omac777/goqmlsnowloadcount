import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1

Tab {
    title: i18n.tr("GuardLocation")
    page: Page {
        width: parent.width

        //row 1
        OptionSelector {
            id: countingAtLocationSelector
            width: parent.width
            text: i18n.tr("Counting At Location:")
            model: [i18n.tr("Conroy"),
                    i18n.tr("Michael"),
                    i18n.tr("Strandherd"),
                    i18n.tr("Innes"),
                    i18n.tr("Clyde")]
        }

        Grid {
            y: countingAtLocationSelector.height
            width: parent.width
            columns :2
            rows: 1
            Button {
                objectName: "GuardLocation_PrevButton"
                id: guardLocation_PrevButton
                text: i18n.tr("Previous")
                onClicked: {
		    ctrl.guardLocationPreviousButtonClickedGo(guardLocation_PrevButton)
                    tabs.selectedTabIndex = tabs.selectedTabIndex - 1
                    console.log("clicked prev..." + tabs.selectedTabIndex);
                }
            }
            Button {
                objectName: "GuardLocation_NextButton"
                id: guardLocation_NextButton
                text: i18n.tr("Next")
                onClicked: {
		    ctrl.guardLocationNextButtonClickedGo(guardLocation_NextButton)
		    console.log("selected:" + countingAtLocationSelector.model[countingAtLocationSelector.selectedIndex]);
                    tabs.selectedTabIndex = tabs.selectedTabIndex + 1
                    console.log("clicked next..." + tabs.selectedTabIndex);
                }
            }
        }
    }
}
