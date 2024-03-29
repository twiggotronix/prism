import { render } from "@testing-library/react";

import ProxyTable from "./proxy-table";

describe("ProxyTable", () => {
    it("should render successfully", () => {
        const { baseElement } = render(<ProxyTable tableData={[]} />);
        expect(baseElement).toBeTruthy();
    });
});
