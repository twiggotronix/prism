import { render } from "@testing-library/react";

import SideMenu from "./sideMenu";

describe("SideMenu", () => {
    it("should render successfully", () => {
        const { baseElement } = render(<SideMenu />);
        expect(baseElement).toBeTruthy();
    });
});
