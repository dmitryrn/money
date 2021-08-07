import React, {useEffect} from 'react';
import logo from './logo.svg';
import {BudgetingAppClient} from "./proto/base_pb_service";
import {GetAccountsRequest} from "./proto/base_pb";

const client = new BudgetingAppClient("http://localhost:3001", {
  debug: true,
})

function App() {
  useEffect(() => {
    client.getAccounts(new GetAccountsRequest(), (e, resp) => {
      if (e) {
        console.log("error getting accounts", e)
        return
      }
      if (!resp) {
        console.log("resp is undefined")
        return
      }

      console.log(resp.getAccountsList())
    })
  })

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
