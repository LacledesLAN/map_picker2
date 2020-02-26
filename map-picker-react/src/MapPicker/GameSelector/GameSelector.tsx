import React, { FC } from 'react';
import './GameSelector.css';

declare interface GameSelectorProps {
  gameId?: number;
  numberOfRounds?: number;
  logo: string;
  background?: string;
}

const GameSelector: FC<GameSelectorProps> = ({ logo, background }) => {
  return (
    <div
      className="game-selector"
      style={{ backgroundImage: `url(${background})` }}
      data-testid="background"
    >
      <div>
        <img src={logo} className="game-selector" alt="logo" />
      </div>
    </div>
  );
};

export default GameSelector;
