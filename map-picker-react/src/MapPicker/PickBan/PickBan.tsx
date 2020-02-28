import React, { FC } from 'react';
import { GridList, GridListTile } from '@material-ui/core';
import Map from '../../data/Map';
import MapCard from '../Shared/MapCard/MapCard';

interface PickBanProps {
  gameMaps: Map[];
}

const PickBan: FC<PickBanProps> = ({ gameMaps }) => {
  return (
    <GridList cellHeight={'auto'} cols={4}>
      {gameMaps.map(gameMap => (
        <GridListTile>
          <MapCard key={gameMap.id} gameMap={gameMap}></MapCard>
        </GridListTile>
      ))}
      <GridListTile>
        <p>Controls</p>
      </GridListTile>
    </GridList>
  );
};

export default PickBan;
