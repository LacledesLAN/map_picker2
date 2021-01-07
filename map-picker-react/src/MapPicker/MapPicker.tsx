import React, { FC } from "react";
import "./MapPicker.css";
import GameSelector from "./GameSelector/GameSelector";

const MapPicker: FC = () => {
  const csgoLogoUrl = "assets/images/map_picker/csgo/logo_white.png";
  const csgoBackgroundUrl = "assets/images/map_picker/csgo/logo_background.jpg";
  const overwatchLogoUrl = "assets/images/map_picker/overwatch/logo_white.png";
  const overwatchBackgroundUrl =
    "assets/images/map_picker/overwatch/logo_background.jpg";

  return (
    <div className="game-container">
      <GameSelector
        logo={csgoLogoUrl}
        background={csgoBackgroundUrl}
        link="/csgo"
      />
      <GameSelector
        logo={overwatchLogoUrl}
        background={overwatchBackgroundUrl}
        link="/overwatch"
      />
    </div>
  );
};

export default MapPicker;
