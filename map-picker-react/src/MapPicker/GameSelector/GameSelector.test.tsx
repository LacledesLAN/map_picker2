import React from 'react';
import { render } from '@testing-library/react';
import GameSelector from './GameSelector';

describe('GameSelector', () => {
  it('displays a background that is passed in as a prop', () => {
    const testBackground = 'test-background-url';

    const testObject = render(
      <GameSelector logo="" background={testBackground} />
    );

    expect(testObject.getByTestId('background')).toHaveStyle(
      `background-image: url(${testBackground})`
    );
  });

  it('displays a logo that is passed in as a prop', async () => {
    const testLogo = 'test-logo-url';

    const testObject = render(<GameSelector logo={testLogo} background="" />);
    const img = await testObject.findByAltText('logo');
    expect(img).toHaveAttribute('src', testLogo);
  });
});
