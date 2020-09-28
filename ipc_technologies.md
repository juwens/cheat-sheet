Glossar

| Name | Explanation |
|--|--|
| Message | A (usually little complexity) data object which exists in your application code (ex: `var myCar = new Car { NumberOfWheels = 4, Weight = 1400 }`. Will be serialized/encoded to json/xml/protobuf |



# Layer 0 - Communication

Responsible for creating and maintaining a connection between the peers.
This Layer is processable by the kernel and network devices like router or switches.
Is usually held open for a longer time, to be reused for subsequently usage.

- TCP/IP
- unix socket
- named pipe (win)
- UDP/IP (theoretically)

# Layer 1

Responsible for 

- (raw)
- http
- http/2

Layer 2 - Serialization Format/Technology to serilize/deserialize the Message

- xml
- json
- bson
- protobuf
- MessagePack
- AVRO
- ASN.1

A lot more available: https://en.wikipedia.org/wiki/Comparison_of_data-serialization_formats
