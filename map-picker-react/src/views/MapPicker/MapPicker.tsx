import React, { FC } from "react";
import "./MapPicker.css";
import GameSelector from "../../components/GameSelector/GameSelector";
import csgoLogo from "../../assets/images/map_picker/csgo/logo_white.png";
import csgoBackground from "../../assets/images/map_picker/csgo/logo_background.jpg";
import overwatchLogo from "../../assets/images/map_picker/overwatch/logo_white.png";
import overwatchBackground from "../../assets/images/map_picker/overwatch/logo_background.jpg";

const MapPicker: FC = () => {
  return (
    <div className="game-container">
      <GameSelector logoUrl={csgoLogo} backgroundUrl={csgoBackground} />
      <GameSelector
        logoUrl={overwatchLogo}
        backgroundUrl={overwatchBackground}
      />
    </div>
  );
};

export default MapPicker;
