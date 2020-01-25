import React, { FC, useState } from 'react';
import './GameSelector.css';

declare interface GameSelectorProps {
  gameId?: number;
  numberOfRounds?: number;
  logoUrl: string;
  backgroundUrl?: string;
}

const GameSelector: FC<GameSelectorProps> = ({ logoUrl, backgroundUrl }) => {
  const [logo] = useState(logoUrl);
  const [background] = useState(backgroundUrl);

  return (
    <div
      className="game-selector"
      style={{ backgroundImage: `url(${background})` }}
    >
      <div>
        <img src={logo} className="game-selector" alt="logo" />
      </div>
    </div>
  );
};

export default GameSelector;
