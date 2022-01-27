# Repositories

>According to Martin Fowler, the Repository Pattern can be described as:
Mediates between the domain and data mapping layers using a collection-like interface for accessing domain objects.
So a repository will generally have methods such as findOne, findAll, remove, and such. Your service layer will manipulate the collection through these repository methods. Their implementation is not a concern to our business logic.

The repository is an intermediary between the domain and the data source and ***provides the services with the basic extraction operations (CRUD)*** (findOne, findAll, updateOne, remove).