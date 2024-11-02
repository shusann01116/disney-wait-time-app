import { StandbyList } from "./standby-list";
import { Meta, StoryObj } from "@storybook/react";
import {
  mockTdlGreetingRespHandler,
  mockTdlRespHandler,
  mockTdsGreetingRespHandler,
  mockTdsRespHandler,
} from "@/lib/fetcher.mock";

const meta: Meta<typeof StandbyList> = {
  component: StandbyList,
};

export default meta;
type Story = StoryObj<typeof StandbyList>;

/**
 * this story is fails because of the some wierd storybook issue
 * https://github.com/storybookjs/storybook/issues/25891
 */
/* export */ const TokyoDisneyLand: Story = {
  args: {
    park: "tdl",
  },
  parameters: {
    msw: {
      handlers: [
        mockTdlRespHandler,
        mockTdsRespHandler,
        mockTdlGreetingRespHandler,
        mockTdsGreetingRespHandler,
      ],
    },
  },
};
