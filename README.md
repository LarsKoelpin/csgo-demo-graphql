
Run using
```bash
go run main.go demo.go header.go tick.go equipment.go position.go   
```

Execute query
```graphql
{
      demo {
			  header {
          mapName
        }
        ticks {
          tick
          players {
            entityId
            team
            position {
              x
            }
          }
        }
      }
		}
```

results in

```text~~~~
{"data":{"demo":{"header":{"mapName":"de_train"},"ticks":[{"players":[{"entityId":7,"position":{"x":-1871.7869873046875},"team":2},{"entityId":8,"position":{"x":-1914.048095703125},"team":2},{"entityId":9,"position":{"x":-447.09600830078125},"team":2},{"entityId":4,"position":{"x":-964.3233642578125},"team":2},{"entityId":6,"position":{"x":-759.9108276367188},"team":3},{"entityId":3,"position":{"x":-942.2088012695312},"team":3},{"entityId":5,"position":{"x":-1007.96875},"team":2},{"entityId":2,"position":{"x":-987.6046752929688},"team":3}],"tick":128},{"players":[{"entityId":2,"position":{"x":-670.0886840820312},"team":3},{"entityId":8,"position":{"x":-1655.302001953125},"team":2},{"entityId":4,"position":{"x":-964.3233642578125},"team":2},{"entityId":6,"position":{"x":1552.460693359375},"team":3},{"entityId":9,"position":{"x":-409.0330810546875},"team":2},{"entityId":3,"position":{"x":-790.9217529296875},"team":3},{"entityId":5,"position":{"x":-953.6282958984375},"team":2},{"entityId":7,"position":{"x":-1871.7869873046875},"team":2}],"tick":256},{"players":[{"entityId":2,"position":{"x":-181.8252410888672},"team":3},{
    ....
}
```
