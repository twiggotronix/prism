import { useState } from "react";
import Breadcrumb from "../components/breadcrumb";
import type { BreadcrumbItem } from "../models/breadcrumb-item";

export const useBreadcrumb = (items?: BreadcrumbItem[]) => {
    const [breadcrumbItems, setBreadcumbItems] = useState<BreadcrumbItem[]>(
        items ?? [{ label: "Home", to: "/" }],
    );

    return {
        setBreadcumbItems,
        render: () => (
            <Breadcrumb breadcrumbItems={breadcrumbItems}></Breadcrumb>
        ),
    };
};
