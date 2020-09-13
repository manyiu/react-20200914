import React, { Component } from "react";
import Dots from "./components/dots";

import dummyResponse from "./data/dummy.json";

import "./App.css";

class App extends Component {
  render() {
    return <Dots data={dummyResponse.data} />;
  }
}

export default App;
