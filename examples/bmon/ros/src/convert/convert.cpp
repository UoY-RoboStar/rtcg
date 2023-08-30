// Conversion functions for battery messages.
//
// The test driver models these as unsigned integers, whereas in ROS they are
// BatteryState messages.

#include "convert.h"
#include "battery_monitor/BatteryState.h"
#include "battery_monitor/BatteryStatus.h"

unsigned int fromBatteryInfo(const BatteryInfoMsg msg)
{
  // Model is nat 0-100, ROS is float 0.0-1.0
  return static_cast<unsigned int>(msg->percentage * 100);
}

BatteryInfoVal toBatteryInfo(unsigned int value)
{
  BatteryInfoVal msg;

  msg.percentage = static_cast<float>(value) / 100.0;

  return msg;
}

std::string fromBatteryStatus(const BatteryStatusMsg msg)
{
  if (!msg->status.compare("OK")) {
    return "Ok";
  } else if (!msg->status.compare("MC")) {
    return "MissionCritical";
  } else if (!msg->status.compare("SC")) {
    return "SafetyCritical";
  } else {
    return "Invalid";
  }
  // switch (msg->status) {
  //   case battery_monitor::BatteryStatus::OK: 
  //     return "Ok";
  //   case battery_monitor::BatteryStatus::MISSION_CRITICAL: 
  //     return "MissionCritical";
  //   case battery_monitor::BatteryStatus::SAFETY_CRITICAL:
  //     return "SafetyCritical";
  //   default:
  //     return "Invalid";
  // }
  //return std::string(msg->data.c_str());
}

BatteryStatusVal toBatteryStatus(const std::string value)
{
  BatteryStatusVal msg;
  //std_msgs::String msg;
  //msg.data = value.c_str();
  //msg.status = BatteryStatusVal::UNSET;
  msg.status = "UNSET";

  if (value.compare("Ok")) {
    //msg.status = BatteryStatusVal::OK;
    msg.status = "OK";
  } else if (value.compare("MissionCritical")) {
    //msg.status = BatteryStatusVal::MISSION_CRITICAL;
    msg.status = "MC";
  } else if (value.compare("SafetyCritical")) {
    //msg.status = BatteryStatusVal::SAFETY_CRITICAL;
    msg.status = "SC";
  }

  return msg;
}
