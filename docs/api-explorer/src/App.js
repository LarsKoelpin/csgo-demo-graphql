import React from 'react';
import "graphiql/graphiql.min.css";
import {GraphiQL} from "graphiql";
import {mockServer} from "graphql-tools";
import {buildClientSchema, printSchema} from "graphql/utilities";
import schema from "./schema.json";

const defaultQuery = `
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
`;

const x = buildClientSchema(schema.data);

const myMockServer = mockServer(x, {
  demo: () => ({
    header: {
      mapName: "abc",
      snapshotRate: 0.2
    },
    ticks: [
      {
        participants: [
          {
            entityId: 7,
            position: {
              x: -1871.7869873046875
            },
            team: 2
          }
        ]
      }
    ]
  })
});

const fetcher = (abc) => {
  return myMockServer.query(abc.query);
};


export function App() {
  return (<GraphiQL
      schema={x}
      fetcher={fetcher}
      defaultQuery={defaultQuery}
    />
  );
}
