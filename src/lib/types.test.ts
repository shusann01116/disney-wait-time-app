import { getLink } from "./types";
import { Link } from "./types";
import { expect, test } from "vitest";

test("getLink for tdl", () => {
  const expected: Link = {
    attraction:
      "https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json",
    greeting: "https://www.tokyodisneyresort.jp/_/realtime/tdl_greeting.json",
  };
  expect(getLink("tdl")).toEqual(expected);
});

test("getLink for tds", () => {
  const expected2: Link = {
    attraction:
      "https://www.tokyodisneyresort.jp/_/realtime/tds_attraction.json",
    greeting: "https://www.tokyodisneyresort.jp/_/realtime/tds_greeting.json",
  };
  expect(getLink("tds")).toEqual(expected2);
});

test("getLink for invalid park", () => {
  expect(getLink("invalid" as any)).toBeNull();
});
