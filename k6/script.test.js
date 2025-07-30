import http from "k6/http";
import { check, sleep } from "k6";
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export const options = {
  stages: [
    // { duration: "30s", target: 10 },
    // { duration: "1m", target: 100 },
    { duration: "10s", target: 500 },
    { duration: "1m", target: 3000 },
    // { duration: "30s", target: 0 },
  ],
  // duration: "1m30s",
  // vus: 2000,
  // vussss: 3000,
  // thresholds: {
  //   http_req_duration: ["p(95)<1000"],
  //   http_req_failed: ["rate<0.01"],
  // },
};

export default function () {
  const payload = JSON.stringify({
    fullname: `${uuidv4()} Test User`,
    salary: 40000,
    loan_amount: 150000,
  });

  const headers = {
    "Content-Type": "application/json",
  };

  const res = http.post("http://localhost:30090/api/vi/loan", payload, {
    headers,
  });
  // const res = http.post("http://localhost:4000/api/vi/loan", payload, {
  //   headers,
  // });

  check(res, {
    "status is 200": (r) => {
      const result = r.status === 200;
      // if (!result) {
      //   console.error(`Expected: 200`);
      //   console.error(`Actual: ${r.status}`);
      //   console.error(`Res ${res.body}`);
      // }
      return result;
    },
  });
}
