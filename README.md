OpenAPI Server Code Generator
-----------------------------

This package contains a set of utilities for generating Go boilerplate code for
services based on 
[OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md)
API definitions. When working with microservices, it's important to have an API
contract which servers and clients both imprement to minimize the chances of
incompatibilities. It's tedious to generate Go models which precisely correspond to
OpenAPI specifications, so let our code generator do that work for you, so that
you can focus on implementing the business logic for your service. 

We have chosen to use [Echo](https://github.com/labstack/echo) as
our HTTP routing engine, due to its speed and simplicity for the generated
stubs.

This package tries to be too simple rather than too generic, so we've made some
design decisions in favor of simplicity, knowing that we can't generate strongly
typed Go code for all possible OpenAPI Schemas.

## Overview

We're going to use the OpenAPI example of the
[Expanded Petstore](https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore-expanded.yaml)
in the descriptions below, please have a look at it.

In order to create a Go server to serve this exact schema, you would have to
write a lot of boilerplate code to perform all the marshalling and unmarshalling
into objects which match the OpenAPI 3.0 definition. The code generator in this
directory does a lot of that for you. You would run it like so:

    go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
    oapi-codegen petstore-expanded.yaml  > petstore.gen.go

Let's go through that `petstore.gen.go` file to show you everything which was
generated.


## Generated Server Boilerplate

The `/components/schemas` section in OpenAPI defines reusable objects, so Go
types are generated for these. The Pet Store example defines `Error`, `Pet`,
`Pets` and `NewPet`, so we do the same in Go:
```
// Type definition for component schema "Error"
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Type definition for component schema "NewPet"
type NewPet struct {
	Name string  `json:"name"`
	Tag  *string `json,omitempty:"tag"`
}

// Type definition for component schema "Pet"
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema
	Id int64 `json:"id"`
}

// Type definition for component schema "Pets"
type Pets []Pet
```

It's best to define objects under `/components` field in the schema, since
those will be turned into named Go types. If you use inline types in your
handler definitions, we will generate inline, anonymous Go types, but those
are more tedious to deal with since you will have to redeclare them at every
point of use.

For each element in the `paths` map in OpenAPI, we will generate a Go handler
function in an interface object. Here is the generated Go interface for our
server.

```
type ServerInterface interface {
    //  (GET /pets)
    FindPets(ctx echo.Context, params FindPetsParams) error
    //  (POST /pets)
    AddPet(ctx echo.Context) error
    //  (DELETE /pets/{id})
    DeletePet(ctx echo.Context, id int64) error
    //  (GET /pets/{id})
    FindPetById(ctx echo.Context, id int64) error
}
```

These are the functions which you will implement yourself in order to create
a server conforming to the API specification. Normally, all the arguments and
parameters are stored on the `echo.Context` in handlers, so we do the tedious
work of of unmarshaling the JSON automatically, simply passing values into
your handlers.

Notice that `FindPetById` takes a parameter `id int64`. All path arguments
will be passed as arguments to your function, since they are mandatory.

Remaining arguments can be passed in headers, query arguments or cookies. Those
will be written to a `params` object. Look at the `FindPets` function above, it
takes as input `FindPetsParams`, which is defined as follows:
 ```
// Parameters object for FindPets
type FindPetsParams struct {
	Tags  *[]string `json:"tags,omitempty"`
	Limit *int32   `json:"limit,omitempty"`
}
```

The HTTP query parameter `limit` turns into a Go field named `Limit`. It is
passed by pointer, since it is an optional parameter. If the parameter is
specified, the pointer will be non-`nil`, and you can read its value.

If you changed the OpenAPI specification to make the parameter required, the
`FindPetsParams` structure will contain the type by value:
```
type FindPetsParams struct {
	Tags  *[]string `json:"tags,omitempty"`
	Limit int32   `json:"limit"`
}
```

The usage of `Echo` is out of scope of this doc, but once you have an
echo instance, we generate a utility function to help you associate your handlers
with this autogenerated code. For the pet store, it looks like this:
```
func RegisterHandlers(router codegen.EchoRouter, si ServerInterface) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}
	router.GET("/pets", wrapper.FindPets)
	router.POST("/pets", wrapper.AddPet)
	router.DELETE("/pets/:id", wrapper.DeletePet)
	router.GET("/pets/:id", wrapper.FindPetById)
}
```

The wrapper functions referenced above contain generated code which pulls
parameters off the `Echo` request context, and unmarshals them into Go objects.

You would register the generated handlers as follows:
```
func SetupHandler() {
    var myApi PetStoreImpl  // This implements the pet store interface
    e := echo.New()
    petstore.RegisterHandlers(e, &myApi)
    ...
}
```

## What's missing

This code is still young, and not complete, since we're filling it in as we
need it. We've not yet implemented several things:

- `oneOf`, `anyOf` are not supported with strong Go typing. This schema:

        schema:
          oneOf:
            - $ref: '#/components/schemas/Cat'
            - $ref: '#/components/schemas/Dog'
    
    will result in a Go type of `interface{}`. It will be up to you
    to validate whether it conforms to `Cat` and/or `Dog`, depending on the
    keyword. It's not clear if we can do anything much better here given the
    limits of Go typing.
    
    `allOf` is supported, by taking the union of all the fields in all the
    component schemas. This is the most useful of these operations, and is
    commonly used to merge objects with an identifier, as in the
    `petstore-expanded` example.
    
- `additionalProperties` isn't supported, and will exit with an error. This
 should be possible to support in the future via a `map[string]interface{}` or
 `map[string]string`.
 
- `patternProperties` isn't yet supported and will exit with an error. This too
 should be possible to implement.

## Making changes to code generation

The code generator uses a tool to inline all the template definitions into
code, so that we don't have to deal with the location of the template files.
When you update any of the files under the `templates/` directory, you will
need to regenerate the template inlines:

    templates -s pkg/codegen/templates/ > pkg/codegen/templates/templates.gen.go

All this command does is inline the files ending in `.tmpl` into the specified
Go file. This command is found here:

    go get github.com/cyberdelia/templates
    
You can also run `go generate`, since we've set up those hooks.


