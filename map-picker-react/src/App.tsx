import React, { FC } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import MapPicker from './MapPicker/MapPicker';
import './App.css';

const App: FC = () => {
  return (
    <Router>
      <div>
        <Switch>
          <Route path="/" exact>
            <MapPicker />
          </Route>
          <Route path="/csgo">
            <CsGo />
          </Route>
          <Route path="/overwatch">
            <Overwatch />
          </Route>
        </Switch>
      </div>
    </Router>
  );
};

const CsGo: FC = () => {
  return <h2>CS Go map selection</h2>;
};

const Overwatch: FC = () => {
  return <h2>Overwatch map selection</h2>;
};

export default App;
