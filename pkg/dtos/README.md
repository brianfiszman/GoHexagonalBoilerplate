# Data Transfer Objects
As its name indicates, it is an object that will be used to ***transfer information*** and represents the object that will be sent to the client, this is the object that our API will return to the rest of the services, either for internal use or for third parties, so we can have multiple DTOs for each entity according to the use we need.
It is also used to define the type of objects to be received by the controllers

- The DTO should have only data, ***should not to have any type of business logic***.
- May have references to other DTOs

> [DTOs](https://en.wikipedia.org/wiki/Data_transfer_object) are essentially domain objects transformed to a shape that is more context aware by being filtered for private/[PII](https://en.wikipedia.org/wiki/Personal_data) data and absent of characteristics that are serialization barriers like circular dependencies. They are optimized for bandwidth usage as well, but more generally, **they include all and only the data the current read or write action requires to function.**

> They are ready for serialization into JSON or whatever transport protocol we choose to use and their primary purpose is to transfer data between two processes or systems according to a predefined spec or schema. We will be placing any DTO used for either input or output at pkg/dtos.