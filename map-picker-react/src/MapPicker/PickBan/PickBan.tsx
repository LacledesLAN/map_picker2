import React, { FC } from 'react';
import { Grid } from '@material-ui/core';
import Map from '../../data/Map';
import MapTile from '../Shared/MapTile/MapTile';
import './PickBan.css';

interface PickBanProps {
  gameMaps: Map[];
}

const PickBan: FC<PickBanProps> = ({ gameMaps }) => {
  return (
    <Grid container className="mapGrid" spacing={0} alignItems="stretch">
      {gameMaps.map(gameMap => (
        <Grid className="mapTile" item xs={3}>
          <MapTile key={gameMap.id} gameMap={gameMap} />
        </Grid>
      ))}
      <Grid className="controlTile" item xs={3}>
        <p>Controls</p>
      </Grid>
    </Grid>
  );
};

export default PickBan;
