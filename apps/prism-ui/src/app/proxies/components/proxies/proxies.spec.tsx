import { render } from "@testing-library/react";

import Proxies from "./proxies";

describe("Proxies", () => {
    it("should render successfully", () => {
        const { baseElement } = render(<Proxies />);
        expect(baseElement).toBeTruthy();
    });
});
