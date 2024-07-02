import { expect } from '@storybook/jest';
import type { Meta, StoryObj } from '@storybook/react';
import { userEvent, within } from '@storybook/test';
import { UserForm } from './UserForm';

const meta = {
  component: UserForm,
} satisfies Meta<typeof UserForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

// nameの入力
export const InputName: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    const nameInput = canvas.getByRole('textbox', {
      name: 'Name',
    });
    await userEvent.type(nameInput, 'example-name', {
      delay: 100,
    });
  },
};
// ageの入力
export const InputAge: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    const ageInput = canvas.getByRole('spinbutton', {
      name: 'Age',
    });
    await userEvent.type(ageInput, '20', {
      delay: 100,
    });
  },
};

// emailの入力
export const InputEmail: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    const emailInput = canvas.getByRole('textbox', {
      name: 'Email',
    });
    await userEvent.type(emailInput, 'example-email@email.com', {
      delay: 100,
    });
  },
};

export const InputAll: Story = {
  play: async (ctx) => {
    await InputName.play?.(ctx);
    await InputAge.play?.(ctx);
    await InputEmail.play?.(ctx);
    // submit
    const canvas = within(ctx.canvasElement);
    const submitButton = canvas.getByRole('button', {
      name: 'Submit',
    });
    await userEvent.click(submitButton);

    const listItems = canvas.getAllByRole('listitem');
    // ↓追加
    expect(listItems[0]).toHaveTextContent('Name: example-name');
    expect(listItems[1]).toHaveTextContent('Age: 20');
    expect(listItems[2]).toHaveTextContent('Email: example-email');
  },
};
