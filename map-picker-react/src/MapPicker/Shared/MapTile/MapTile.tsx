import React, { FC } from 'react';
import './MapTile.css';
import Map from '../../../data/Map';

interface MapTile {
  gameMap: Map;
}

const MapTile: FC<MapTile> = ({ gameMap }) => {
  return (
    <div
      key={gameMap.id}
      className={(gameMap.name, 'MapTile')}
      style={{ background: `url(${gameMap.imageUrl})` }}
    >
      <p>{gameMap.name}</p>
    </div>
  );
};

export default MapTile;
