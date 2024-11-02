import type { Meta, StoryObj } from "@storybook/react";

import { Badge } from "./badge";

const meta: Meta<typeof Badge> = {
  component: Badge,
  tags: ["autodocs"],
};

export default meta;
type Story = StoryObj<typeof Badge>;

export const Default: Story = {
  args: {
    variant: "default",
    children: "Badge",
  },
};

export const Secondary: Story = {
  args: {
    variant: "secondary",
    children: "Badge",
  },
};

export const Destructive: Story = {
  args: {
    variant: "destructive",
    children: "Badge",
  },
};

export const Outline: Story = {
  args: {
    variant: "outline",
    children: "Badge",
  },
};
