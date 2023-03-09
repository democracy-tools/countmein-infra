syntax = "proto3";
message ProtocolBuffer {
  string id = 1;
  string user_id = 2;
  string user_device_id = 3;
  string user_device_type = 4;
  string seen_device_id = 5;
  string seen_device_type = 6;
  float location_latitude = 7;
  float location_longitude = 8;
  int64 user_time = 9;
  int64 server_time = 10;
}
