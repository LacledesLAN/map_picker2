import React, { FC } from "react";
import "./GameSelector.css";

declare interface GameSelectorProps {
  gameId?: number;
  numberOfRounds?: number;
  logoUrl: string;
  backgroundUrl?: string;
}

const GameSelector: FC<GameSelectorProps> = ({ logoUrl, backgroundUrl }) => {
  return (
    <div
      className="game-selector"
      style={{ backgroundImage: `url(${backgroundUrl})` }}
    >
      <div>
        <img src={logoUrl} className="game-selector" alt="logo" />
      </div>
    </div>
  );
};

export default GameSelector;
