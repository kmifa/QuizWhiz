import { render, screen } from '@testing-library/react';

import { BaseTemplate } from './test';

describe('Base template', () => {
  describe('Render method', () => {
    it('data-testidを取得して中のtextを取得して比較する', () => {
      render(
          <BaseTemplate />
      );

      const element = screen.getByTestId('custom-element').textContent;
      expect(element).toBe('test');
    });
  });
});
