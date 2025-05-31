# Single Method Interface

This section describes techniques for using single method interfaces.

## Handler Patterns

In this example, our goal is to implement a way to separate out sqlite database operations using single method interfaces from specific domain data types.

First we have a package call `sqlops` that contains a collection of sql related interfaces. Refer to [this](../internal/sqlops/sqlops.go).

Second we have a package call `person` that contains a collection of data types related to person types. Refer to [this](../internal/person/person.go).

Third we have the main package that ties all the packages together. Refer to [this](../examples/single-method/ex1/main.go)
