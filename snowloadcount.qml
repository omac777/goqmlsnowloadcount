import QtQuick 2.0
import QtQuick.Window 2.0
import Ubuntu.Components 0.1
import "ui"

MainView {
    objectName: "mainView"

    // Note! applicationName needs to match the "name" field of the click manifest
    applicationName: "com.ubuntu.developer.uticdmarceau2007.uqttabtest1"

    width: Screen.width
    height: Screen.height

    Tabs {
        id: tabs

          OpenReport {
              objectName: "openReportTab"
              id: openReportTab
              width: Screen.width
              height: Screen.height
                }

          GuardShiftStartTime {
              objectName: "guardShiftStartTimeTab"
              id: guardShiftStartTimeTab
              width: Screen.width
              height: Screen.height
          }

          GuardShiftEndTime {
              objectName: "guardShiftEndTimeTab"
              id: guardShiftEndTimeTab
              width: Screen.width
              height: Screen.height
          }

          GuardShiftStartDate {
              objectName: "guardShiftStartDateTab"
              id: guardShiftStartDateTab
              width: Screen.width
              height: Screen.height
          }

        GuardShiftInfo {
            objectName: "guardshiftinfoTab"
            id: guardshiftinfoTab
            width: Screen.width
            height: Screen.height
        }

        GuardLocation {
            objectName: "guardlocationTab"
            id: guardlocationTab
            width: Screen.width
            height: Screen.height
        }

        AdmissionType {
            objectName: "admissiontypeTab"
            id: admissiontypeTab
            width: Screen.width
            height: Screen.height
        }

        TruckType {
            objectName: "trucktypeTab"
            id: trucktypeTab
            width: Screen.width
            height: Screen.height
        }
    }
}
