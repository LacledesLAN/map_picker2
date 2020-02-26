import React, { FC, useState } from 'react';
import './MapPicker.css';
import GameSelector from './GameSelector/GameSelector';

const csgoLogoUrl = 'assets/images/map_picker/csgo/logo_white.png';
const csgoBackgroundUrl = 'assets/images/map_picker/csgo/logo_background.jpg';
const overwatchLogoUrl = 'assets/images/map_picker/overwatch/logo_white.png';
const overwatchBackgroundUrl =
  'assets/images/map_picker/overwatch/logo_background.jpg';

const MapPicker: FC = () => {
  const [csgoLogo] = useState(csgoLogoUrl);
  const [csgoBackground] = useState(csgoBackgroundUrl);
  const [overwatchLogo] = useState(overwatchLogoUrl);
  const [overwatchBackground] = useState(overwatchBackgroundUrl);

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
