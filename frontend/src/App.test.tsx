import React from 'react';
import { shallow } from 'enzyme';
import App from './App';

describe('App', () => {
  it('renders welcome text', () => {
    const component = shallow(<App />);
    expect(component.html()).toMatch(/welcome to leedprojects/i);
  });
});
