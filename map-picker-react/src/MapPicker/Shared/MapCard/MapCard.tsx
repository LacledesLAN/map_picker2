import React, { FC } from 'react';
import './MapCard.css';
import Map from '../../../data/Map';

interface MapCardProps {
  gameMap: Map;
}

const MapCard: FC<MapCardProps> = ({ gameMap }) => {
  return (
    <div
      key={gameMap.id}
      className={(gameMap.name, 'map-card')}
      style={{ background: `url(${gameMap.imageUrl})` }}
    >
      <p>{gameMap.name}</p>
    </div>
  );
};

export default MapCard;
