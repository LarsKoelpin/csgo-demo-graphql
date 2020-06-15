export const schema = {
  data: {
    __schema: {
      directives: [
        {
          args: [
            {
              defaultValue: null,
              description: "Included when true.",
              name: "if",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            }
          ],
          description:
            "Directs the executor to include this field or fragment only when the `if` argument is true.",
          locations: ["FIELD", "FRAGMENT_SPREAD", "INLINE_FRAGMENT"],
          name: "include"
        },
        {
          args: [
            {
              defaultValue: null,
              description: "Skipped when true.",
              name: "if",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            }
          ],
          description:
            "Directs the executor to skip this field or fragment when the `if` argument is true.",
          locations: ["FIELD", "FRAGMENT_SPREAD", "INLINE_FRAGMENT"],
          name: "skip"
        },
        {
          args: [
            {
              defaultValue: '"No longer supported"',
              description:
                "Explains why this element was deprecated, usually also including a suggestion for how to access supported similar data. Formattedin [Markdown](https://daringfireball.net/projects/markdown/).",
              name: "reason",
              type: { kind: "SCALAR", name: "String", ofType: null }
            }
          ],
          description:
            "Marks an element of a GraphQL schema as no longer supported.",
          locations: ["FIELD_DEFINITION", "ENUM_VALUE"],
          name: "deprecated"
        }
      ],
      mutationType: null,
      queryType: { name: "RootQuery" },
      subscriptionType: null,
      types: [
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "carrier",
              type: { kind: "OBJECT", name: "players", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "lastOnGroundPosition",
              type: { kind: "OBJECT", name: "position", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "bomb",
          possibleTypes: null
        },
        {
          description:
            "One possible value for a given Enum. Enum values are unique values, not a placeholder for a string or numeric value. However an Enum value is returned in a JSON response as a string.",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "deprecationReason",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "description",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "isDeprecated",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "name",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "String", ofType: null }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "__EnumValue",
          possibleTypes: null
        },
        {
          description:
            "A Directive provides a way to describe alternate runtime execution and type validation behavior in a GraphQL document. \n\nIn some cases, you need to provide options to alter GraphQL's execution behavior in ways field arguments will not suffice, such as conditionally including or skipping a field. Directives provide this by describing additional information to the executor.",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "args",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: {
                  kind: "LIST",
                  name: null,
                  ofType: {
                    kind: "NON_NULL",
                    name: null,
                    ofType: {
                      kind: "OBJECT",
                      name: "__InputValue",
                      ofType: null
                    }
                  }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "description",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "locations",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: {
                  kind: "LIST",
                  name: null,
                  ofType: {
                    kind: "NON_NULL",
                    name: null,
                    ofType: {
                      kind: "ENUM",
                      name: "__DirectiveLocation",
                      ofType: null
                    }
                  }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "name",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "String", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: "Use `locations`.",
              description: "",
              isDeprecated: true,
              name: "onField",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: "Use `locations`.",
              description: "",
              isDeprecated: true,
              name: "onFragment",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: "Use `locations`.",
              description: "",
              isDeprecated: true,
              name: "onOperation",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "__Directive",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [
                {
                  defaultValue: null,
                  description: "",
                  name: "freq",
                  type: {
                    kind: "NON_NULL",
                    name: null,
                    ofType: { kind: "SCALAR", name: "Float", ofType: null }
                  }
                },
                {
                  defaultValue: null,
                  description: "",
                  name: "demoFile",
                  type: { kind: "SCALAR", name: "String", ofType: null }
                }
              ],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "demo",
              type: { kind: "OBJECT", name: "demo", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "RootQuery",
          possibleTypes: null
        },
        {
          description:
            "The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point). ",
          enumValues: null,
          fields: null,
          inputFields: null,
          interfaces: null,
          kind: "SCALAR",
          name: "Float",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "angleX",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "angleY",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "armor",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "deaths",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "entityId",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "equipment",
              type: {
                kind: "LIST",
                name: null,
                ofType: { kind: "OBJECT", name: "equipment", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "flashDuration",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "hasDefuseKit",
              type: { kind: "SCALAR", name: "Boolean", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "hasHelmet",
              type: { kind: "SCALAR", name: "Boolean", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "hp",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "isDefusing",
              type: { kind: "SCALAR", name: "Boolean", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "isInBuyzone",
              type: { kind: "SCALAR", name: "Boolean", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "isNpc",
              type: { kind: "SCALAR", name: "Boolean", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "isPlanting",
              type: { kind: "SCALAR", name: "Boolean", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "kills",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "money",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "name",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "position",
              type: { kind: "OBJECT", name: "position", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "team",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "players",
          possibleTypes: null
        },
        {
          description:
            "A GraphQL Schema defines the capabilities of a GraphQL server. It exposes all available types and directives on the server, as well as the entry points for query, mutation, and subscription operations.",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "A list of all directives supported by this server.",
              isDeprecated: false,
              name: "directives",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: {
                  kind: "LIST",
                  name: null,
                  ofType: {
                    kind: "NON_NULL",
                    name: null,
                    ofType: {
                      kind: "OBJECT",
                      name: "__Directive",
                      ofType: null
                    }
                  }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description:
                "If this server supports mutation, the type that mutation operations will be rooted at.",
              isDeprecated: false,
              name: "mutationType",
              type: { kind: "OBJECT", name: "__Type", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "The type that query operations will be rooted at.",
              isDeprecated: false,
              name: "queryType",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "OBJECT", name: "__Type", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description:
                "If this server supports subscription, the type that subscription operations will be rooted at.",
              isDeprecated: false,
              name: "subscriptionType",
              type: { kind: "OBJECT", name: "__Type", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "A list of all types supported by this server.",
              isDeprecated: false,
              name: "types",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: {
                  kind: "LIST",
                  name: null,
                  ofType: {
                    kind: "NON_NULL",
                    name: null,
                    ofType: { kind: "OBJECT", name: "__Type", ofType: null }
                  }
                }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "__Schema",
          possibleTypes: null
        },
        {
          description:
            "The fundamental unit of any GraphQL Schema is the type. There are many kinds of types in GraphQL as represented by the `__TypeKind` enum.\n\nDepending on the kind of a type, certain fields describe information about that type. Scalar types provide no information beyond a name and description, while Enum types provide their values. Object and Interface types provide the fields they describe. Abstract types, Union and Interface, provide the Object types possible at runtime. List and NonNull types compose other types.",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "description",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [
                {
                  defaultValue: "false",
                  description: "",
                  name: "includeDeprecated",
                  type: { kind: "SCALAR", name: "Boolean", ofType: null }
                }
              ],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "enumValues",
              type: {
                kind: "LIST",
                name: null,
                ofType: {
                  kind: "NON_NULL",
                  name: null,
                  ofType: { kind: "OBJECT", name: "__EnumValue", ofType: null }
                }
              }
            },
            {
              args: [
                {
                  defaultValue: "false",
                  description: "",
                  name: "includeDeprecated",
                  type: { kind: "SCALAR", name: "Boolean", ofType: null }
                }
              ],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "fields",
              type: {
                kind: "LIST",
                name: null,
                ofType: {
                  kind: "NON_NULL",
                  name: null,
                  ofType: { kind: "OBJECT", name: "__Field", ofType: null }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "inputFields",
              type: {
                kind: "LIST",
                name: null,
                ofType: {
                  kind: "NON_NULL",
                  name: null,
                  ofType: { kind: "OBJECT", name: "__InputValue", ofType: null }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "interfaces",
              type: {
                kind: "LIST",
                name: null,
                ofType: {
                  kind: "NON_NULL",
                  name: null,
                  ofType: { kind: "OBJECT", name: "__Type", ofType: null }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "kind",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "ENUM", name: "__TypeKind", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "name",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "ofType",
              type: { kind: "OBJECT", name: "__Type", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "possibleTypes",
              type: {
                kind: "LIST",
                name: null,
                ofType: {
                  kind: "NON_NULL",
                  name: null,
                  ofType: { kind: "OBJECT", name: "__Type", ofType: null }
                }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "__Type",
          possibleTypes: null
        },
        {
          description:
            "An enum describing what kind of type a given `__Type` is",
          enumValues: [
            {
              deprecationReason: null,
              description:
                "Indicates this type is an interface. `fields` and `possibleTypes` are valid fields.",
              isDeprecated: false,
              name: "INTERFACE"
            },
            {
              deprecationReason: null,
              description:
                "Indicates this type is a union. `possibleTypes` is a valid field.",
              isDeprecated: false,
              name: "UNION"
            },
            {
              deprecationReason: null,
              description:
                "Indicates this type is an enum. `enumValues` is a valid field.",
              isDeprecated: false,
              name: "ENUM"
            },
            {
              deprecationReason: null,
              description:
                "Indicates this type is an input object. `inputFields` is a valid field.",
              isDeprecated: false,
              name: "INPUT_OBJECT"
            },
            {
              deprecationReason: null,
              description:
                "Indicates this type is a list. `ofType` is a valid field.",
              isDeprecated: false,
              name: "LIST"
            },
            {
              deprecationReason: null,
              description:
                "Indicates this type is a non-null. `ofType` is a valid field.",
              isDeprecated: false,
              name: "NON_NULL"
            },
            {
              deprecationReason: null,
              description: "Indicates this type is a scalar.",
              isDeprecated: false,
              name: "SCALAR"
            },
            {
              deprecationReason: null,
              description:
                "Indicates this type is an object. `fields` and `interfaces` are valid fields.",
              isDeprecated: false,
              name: "OBJECT"
            }
          ],
          fields: null,
          inputFields: null,
          interfaces: null,
          kind: "ENUM",
          name: "__TypeKind",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "header",
              type: { kind: "OBJECT", name: "header", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "ticks",
              type: {
                kind: "LIST",
                name: null,
                ofType: { kind: "OBJECT", name: "ticks", ofType: null }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "demo",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "bomb",
              type: { kind: "OBJECT", name: "bomb", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "participants",
              type: {
                kind: "LIST",
                name: null,
                ofType: { kind: "OBJECT", name: "players", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "tick",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "totalRoundsPlayed",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "ticks",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "mapName",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "snapshotRate",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "tickRate",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "header",
          possibleTypes: null
        },
        {
          description:
            "Arguments provided to Fields or Directives and the input fields of an InputObject are represented as Input Values which describe their type and optionally a default value.",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description:
                "A GraphQL-formatted string representing the default value for this input value.",
              isDeprecated: false,
              name: "defaultValue",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "description",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "name",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "String", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "type",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "OBJECT", name: "__Type", ofType: null }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "__InputValue",
          possibleTypes: null
        },
        {
          description:
            "A Directive can be adjacent to many parts of the GraphQL language, a __DirectiveLocation describes one such possible adjacencies.",
          enumValues: [
            {
              deprecationReason: null,
              description: "Location adjacent to a union definition.",
              isDeprecated: false,
              name: "UNION"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to an enum definition.",
              isDeprecated: false,
              name: "ENUM"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to an enum value definition.",
              isDeprecated: false,
              name: "ENUM_VALUE"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a mutation operation.",
              isDeprecated: false,
              name: "MUTATION"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a schema definition.",
              isDeprecated: false,
              name: "SCHEMA"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a scalar definition.",
              isDeprecated: false,
              name: "SCALAR"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a object definition.",
              isDeprecated: false,
              name: "OBJECT"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a fragment definition.",
              isDeprecated: false,
              name: "FRAGMENT_DEFINITION"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a field definition.",
              isDeprecated: false,
              name: "FIELD_DEFINITION"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to an interface definition.",
              isDeprecated: false,
              name: "INTERFACE"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a subscription operation.",
              isDeprecated: false,
              name: "SUBSCRIPTION"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a fragment spread.",
              isDeprecated: false,
              name: "FRAGMENT_SPREAD"
            },
            {
              deprecationReason: null,
              description:
                "Location adjacent to an input object type definition.",
              isDeprecated: false,
              name: "INPUT_OBJECT"
            },
            {
              deprecationReason: null,
              description:
                "Location adjacent to an input object field definition.",
              isDeprecated: false,
              name: "INPUT_FIELD_DEFINITION"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a query operation.",
              isDeprecated: false,
              name: "QUERY"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to a field.",
              isDeprecated: false,
              name: "FIELD"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to an inline fragment.",
              isDeprecated: false,
              name: "INLINE_FRAGMENT"
            },
            {
              deprecationReason: null,
              description: "Location adjacent to an argument definition.",
              isDeprecated: false,
              name: "ARGUMENT_DEFINITION"
            }
          ],
          fields: null,
          inputFields: null,
          interfaces: null,
          kind: "ENUM",
          name: "__DirectiveLocation",
          possibleTypes: null
        },
        {
          description:
            "The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.",
          enumValues: null,
          fields: null,
          inputFields: null,
          interfaces: null,
          kind: "SCALAR",
          name: "String",
          possibleTypes: null
        },
        {
          description:
            "The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1. ",
          enumValues: null,
          fields: null,
          inputFields: null,
          interfaces: null,
          kind: "SCALAR",
          name: "Int",
          possibleTypes: null
        },
        {
          description:
            "The `Boolean` scalar type represents `true` or `false`.",
          enumValues: null,
          fields: null,
          inputFields: null,
          interfaces: null,
          kind: "SCALAR",
          name: "Boolean",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "ammoInMagazine",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "ammoReserve",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "type",
              type: { kind: "SCALAR", name: "Int", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "equipment",
          possibleTypes: null
        },
        {
          description: "",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "x",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "y",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "z",
              type: { kind: "SCALAR", name: "Float", ofType: null }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "position",
          possibleTypes: null
        },
        {
          description:
            "Object and Interface types are described by a list of Fields, each of which has a name, potentially a list of arguments, and a return type.",
          enumValues: null,
          fields: [
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "args",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: {
                  kind: "LIST",
                  name: null,
                  ofType: {
                    kind: "NON_NULL",
                    name: null,
                    ofType: {
                      kind: "OBJECT",
                      name: "__InputValue",
                      ofType: null
                    }
                  }
                }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "deprecationReason",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "description",
              type: { kind: "SCALAR", name: "String", ofType: null }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "isDeprecated",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "Boolean", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "name",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "SCALAR", name: "String", ofType: null }
              }
            },
            {
              args: [],
              deprecationReason: null,
              description: "",
              isDeprecated: false,
              name: "type",
              type: {
                kind: "NON_NULL",
                name: null,
                ofType: { kind: "OBJECT", name: "__Type", ofType: null }
              }
            }
          ],
          inputFields: null,
          interfaces: [],
          kind: "OBJECT",
          name: "__Field",
          possibleTypes: null
        }
      ]
    }
  }
};
