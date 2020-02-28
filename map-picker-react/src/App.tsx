import React, { FC } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { csgoMaps } from './data/csgoMaps';
import MapPicker from './MapPicker/MapPicker';
import PickBan from './MapPicker/PickBan/PickBan';
import './App.css';

const App: FC = () => {
  return (
    <Router>
      <Switch>
        <Route path="/" exact>
          <MapPicker />
        </Route>
        <Route path="/csgo">
          <PickBan gameMaps={csgoMaps} />
        </Route>
        <Route path="/overwatch">
          <Overwatch />
        </Route>
      </Switch>
    </Router>
  );
};

const Overwatch: FC = () => {
  return <h2>Overwatch map selection</h2>;
};

export default App;
