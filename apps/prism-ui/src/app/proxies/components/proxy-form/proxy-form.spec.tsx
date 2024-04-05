import { render } from "@testing-library/react";

import ProxyForm from "./proxy-form";

describe("ProxyForm", () => {
    it("should render successfully", () => {
        const { baseElement } = render(<ProxyForm />);
        expect(baseElement).toBeTruthy();
    });
});
