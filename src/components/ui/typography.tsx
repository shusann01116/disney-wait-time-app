import { cva } from "class-variance-authority";

export const Typography = cva([], {
  variants: {
    variant: {
      large: ["text-lg", "font-semibold"],
    },
  },
});
