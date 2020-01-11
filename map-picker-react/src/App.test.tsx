import React from "react";
import { render } from "@testing-library/react";
import App from "./App";

test("renders game logos", () => {
  const testObject = render(<App />);
  const linkElement = testObject.getByAltText("logo");
  expect(linkElement).toBeInTheDocument();
});
