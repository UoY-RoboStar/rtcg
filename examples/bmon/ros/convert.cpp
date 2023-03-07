// Conversion functions for battery messages.
//
// The test driver models these as unsigned integers, whereas in ROS they are
// BatteryState messages.

#include "sensor_msgs/BatteryState.h"

// This is the same typedef that the test driver uses:
using BatteryInfoMsg = sensor_msgs::BatteryState;

unsigned int fromBatteryInfo(const BatteryInfoMsg::ConstPtr& msg)
{
  return static_cast<unsigned int>(msg->data); // multiply by 100?
}

BatteryInfoMsg toBatteryInfo(unsigned int value)
{
  BatteryInfoMsg msg;

  msg.percentage = static_cast<float>(value); // divide by 100?

  return msg;
}
