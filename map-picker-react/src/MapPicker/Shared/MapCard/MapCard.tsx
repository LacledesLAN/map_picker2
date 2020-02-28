import React, { FC } from 'react';
import Map from '../../../data/Map';

interface MapCardProps {
  gameMap: Map;
}

const MapCard: FC<MapCardProps> = ({ gameMap }) => {
  const cardStyle = {
    background: `url(${gameMap.imageUrl})`
  };

  return (
    <div key={gameMap.id} className={gameMap.name} style={cardStyle}>
      <p>{gameMap.name}</p>
    </div>
  );
};

export default MapCard;
