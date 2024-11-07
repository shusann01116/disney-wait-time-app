import type { Meta } from "@storybook/react";
import StandbyByArea from "./standby-by-area";

const meta: Meta<typeof StandbyByArea> = {
  component: StandbyByArea,
  args: {
    standbyList: [
      {
        id: "173",
        name: "アリスのティーパーティー",
        operatingStatus: {
          id: "open",
          name: "Open",
        },
        operatingHour: {
          from: new Date("2024-11-02T09:00:00"),
          to: new Date("2024-11-02T21:00:00"),
        },
        standbyTime: 5,
        updatedAt: new Date("2024-11-02T16:54:00"),
      },
    ],
  },
};

export default meta;

export const Default = {};
