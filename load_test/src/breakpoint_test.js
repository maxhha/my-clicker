import http from "k6/http";
import { check } from "k6";

/** @type {import("k6/options").Options} */
export const options = {
  scenarios: {
    clicking: {
      executor: "ramping-arrival-rate",
      startRate: 100,
      timeUnit: "1s",

      preAllocatedVUs: 50,
      maxVUs: 100,

      stages: [
        { target: 100, duration: "1m" },
        { target: 200, duration: "1m" },
        { target: 400, duration: "1m" },
        { target: 800, duration: "1m" },
        { target: 1600, duration: "1m" },
      ],
    },
  },
};

export function setup() {
  const res = http.post(
    `${__ENV.CLICKER_URL}/api/link/create`,
    JSON.stringify({ redirect: "http://localhost/" }),
    {
      headers: { "Content-Type": "application/json" },
    }
  );

  const link_id = res.json("link_id");
  console.log("Link created", link_id);
  return { link_id };
}

export default function ({ link_id }) {
  const res = http.get(`${__ENV.CLICKER_URL}/l/${link_id}`, { redirects: 0 });
  check(res, {
    "is status 302": (r) => r.status === 302,
  });
}

export function teardown({ link_id }) {
  const res = http.get(`${__ENV.CLICKER_URL}/api/link/${link_id}/counter`);

  const total = res.json("total");

  console.log("Counter for", link_id, "is", total);
}
