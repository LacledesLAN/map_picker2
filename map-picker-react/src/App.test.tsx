import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

describe('App', () => {
  it('renders two game logos', () => {
    const testObject = render(<App />);
    const linkElement = testObject.getAllByAltText('logo');
    expect(linkElement).toHaveLength(2);
  });
});
