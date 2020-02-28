import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import './GameSelector.css';

declare interface GameSelectorProps {
  gameId?: number;
  numberOfRounds?: number;
  logo: string;
  link: string;
  background?: string;
}

const GameSelector: FC<GameSelectorProps> = ({ logo, background, link }) => {
  return (
    <div
      className="game-selector"
      style={{ backgroundImage: `url(${background})` }}
      data-testid="background"
    >
      <Link to={link}>
        <div>
          <img src={logo} className="game-selector" alt="logo" />
        </div>
      </Link>
    </div>
  );
};

export default GameSelector;
