
# Under Development

# CSGO Demo Renderer
A CSGO Demo renderer using GraphQL. It takes a CSGO Demo in binary format and encodes it into JSON.
It only respects the attributes, which the user specifies in the query, therefore not over or underfetching
any data resulting in a bloated JSON.


# How to run

Run using
```bash
csgodemo path/to/query.query
```

where the **query.query** is a file containing a graphQL query.
```graphql
{
  demo(freq: 0.2, demoFile: "asd") {
    header {
      mapName
    }
    ticks {
      participants {
        entityId
      }
    }
  }
}
```
Where 
* **freq** is the recording FPS. 
* **demoFile**  is the path to the demoFile.

For creation purposes, you can use the interactive Graphiql tool available at

https://larskoelpin.github.io/csgo-demo-graphql/

