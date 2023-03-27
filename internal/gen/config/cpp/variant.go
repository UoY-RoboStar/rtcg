package cpp

//go:generate go-enum --marshal -nocase

// Variant is the type of C++ template variants.
// It is defined using the following go-enum:
/*
ENUM(
animate, // CppVariant with a manually fed event loop.
ros,     // CppVariant targeting ROS1 Noetic.
)
*/
//
// There are currently two variants:
//
//   - 'VariantAnimate', which implements an (unfinished) manual animator for the test.  This depends
//     only on the C++11 standard library, and will have a loop that prompts the user for the SUT's
//     action at each stage;
//   - 'VariantRos', which targets ROS1 Noetic.
//
type Variant uint8
