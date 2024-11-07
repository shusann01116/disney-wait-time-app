import { ParkType } from "./fetcher";
import { getLink } from "./types";
import { Link } from "./types";
import { describe, expect, it } from "vitest";

describe("getLink", () => {
  it("getLink for tdl", () => {
    const expected: Link = {
      attraction:
        "https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json",
      greeting: "https://www.tokyodisneyresort.jp/_/realtime/tdl_greeting.json",
    };
    expect(getLink("tdl")).toEqual(expected);
  });

  it("getLink for tds", () => {
    const expected: Link = {
      attraction:
        "https://www.tokyodisneyresort.jp/_/realtime/tds_attraction.json",
      greeting: "https://www.tokyodisneyresort.jp/_/realtime/tds_greeting.json",
    };
    expect(getLink("tds")).toEqual(expected);
  });

  it("getLink for invalid park", () => {
    expect(getLink("invalid" as ParkType)).toBeNull();
  });
});
