import React, { Component } from "react";
import { Route, Switch, withRouter } from "react-router-dom"

import Home from "./components/Home";
import ComicList from "./components/ComicList";
import ShowComic from "./components/ShowComic";

import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Switch>
          <Route exact path="/" component={Home} />
          <Route exact path="/comics" component={ComicList} />
          <Route exact path="/comics/:filter?" component={ShowComic} />
        </Switch>
      </div>
    )
  }
}

export default withRouter(App);
