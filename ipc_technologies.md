# Abstract

Inter-process communication (usually) consist of four components:

1. Definition of the service (method and exchanged data - example `int32 getCustomerId(string email)`) - **IDL** (**I**nterface **D**efinition **L**anguage). For SOAP it's WSDL (webservice description language).
2. Data encoder/serializer/marshaller, which converts an in-process/in-memory object to a string or byte-array which can be transfered over the network. For SOAP the data-format is XML and tons of (de)serializers are available for every modern language.
3. 

# Glossary

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

# Layer 2 - Serialization Format/Technology to serilize/deserialize the Message

- xml
- json
- bson
- protobuf
- MessagePack
- AVRO
- ASN.1

A lot more available: https://en.wikipedia.org/wiki/Comparison_of_data-serialization_formats

# IDL - Interface Definition Language

The 

- [Swagger/OpenAPI Specification](https://en.wikipedia.org/wiki/OpenAPI_Specification)
