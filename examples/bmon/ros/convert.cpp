// Conversion functions for battery messages.
//
// The test driver models these as unsigned integers, whereas in ROS they are
// BatteryState messages.

#include "sensor_msgs/BatteryState.h"

unsigned int fromBatteryInfo(const sensor_msgs::BatteryState::ConstPtr& msg)
{
  // Model is nat 0-100, ROS is float 0.0-1.0
  return static_cast<unsigned int>(msg->data * 100);
}

sensor_msgs::BatteryState toBatteryInfo(unsigned int value)
{
  sensor_msgs::BatteryState msg;

  msg.percentage = static_cast<float>(value) / 100.0;

  return msg;
}
