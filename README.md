
# Under Development

# CSGO Demo Renderer
A CSGO Demo renderer using GraphQL. It takes a CSGO Demo in binary format and encodes it into JSON.
It only respects the attributes, which the user specifies in the query, therefore not over or underfetching
any data resulting in a bloated JSON.


## How to run

Run using
```bash
csgodemo --query query.query --demo demo.dem
```

where the **query.query** is a file containing a graphQL query.
```graphql
{
  demo(freq: 0.2) {
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
Where **freq** is the recording FPS. 

It creates a File named out.json containing all data. 
```bash
ls
+ out.json
```


For exploring purposes, you can use the interactive Graphiql tool available at

https://larskoelpin.github.io/csgo-demo-graphql/


## Considerations
If you want to send the json over the wire, try deflateing it using gzip

creates a File named out.json 
```bash
gzip out.json //out.json.gz
```
inflate using
```bash
gzip -d out.json.gz
```


