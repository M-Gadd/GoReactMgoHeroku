import React from "react";
import PingComponent from "./components/PingComponent";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>Deploy React + Go to Heroku using Docker</p>
        <PingComponent />
      </header>
    </div>
  );
}

export default App;
