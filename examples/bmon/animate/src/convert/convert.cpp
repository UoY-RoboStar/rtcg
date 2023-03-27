// Conversion functions for battery messages.

unsigned int fromBatteryInfo(const float msg)
{
  // Model is nat 0-100, ROS is float 0.0-1.0
  return static_cast<unsigned int>(msg * 100);
}

float toBatteryInfo(unsigned int value)
{
  return static_cast<float>(value) / 100.0;
}
