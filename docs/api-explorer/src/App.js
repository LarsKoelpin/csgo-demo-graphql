import React from 'react';
import "graphiql/graphiql.min.css";
import {GraphiQL} from "graphiql";
import {mockServer} from "graphql-tools";
import {buildClientSchema} from "graphql/utilities";
import schema from "./schema.json";
import data from './data.json';

const defaultQuery = `
{
  demo(freq: 0.2) {
    header {
      mapName
    }
    ticks {
      participants {
        name
        armor
        angleX
        angleY
        equipment {
          ammoInMagazine
          ammoReserve
          type
        }
      }
    }
  }
}

`;

const x = buildClientSchema(schema.data);

const myMockServer = mockServer(x, {
  demo: () => data.data.demo
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
