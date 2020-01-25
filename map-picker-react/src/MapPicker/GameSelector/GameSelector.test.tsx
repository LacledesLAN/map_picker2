import React from 'react';
import { render } from '@testing-library/react';
import GameSelector from './GameSelector';

describe('GameSelector', () => {
  it('displays a background that is passed in as the prop', () => {
    const testBackground = 'test-background-url';

    render(<GameSelector logoUrl="" backgroundUrl={testBackground} />);
  });
});
