import { NavLink } from "react-router-dom";
import type { BreadcrumbItem } from "../models/breadcrumb-item";
import styles from "./breadcrumb.module.scss";

export interface BreadcrumbProps {
    breadcrumbItems: BreadcrumbItem[];
}

export function Breadcrumb({ breadcrumbItems }: BreadcrumbProps) {
    return (
        <div className={styles["container"]}>
            <nav aria-label="breadcrumb" className="w-max">
                <ol className="bg-blue-gray-50 flex w-full flex-wrap items-center rounded-md bg-opacity-60 px-4 py-2">
                    {breadcrumbItems.map((breadcrumbItem, index) => (
                        <li
                            key={index}
                            className="text-blue-gray-900 hover:text-light-blue-500 flex cursor-pointer items-center font-sans text-sm font-normal leading-normal antialiased transition-colors duration-300"
                        >
                            <NavLink
                                to={breadcrumbItem.to}
                                className={isActive =>
                                    `opacity-60 ${isActive ? "acctive" : ""}`
                                }
                            >
                                {breadcrumbItem.label}
                            </NavLink>
                            {breadcrumbItems.length < index ? (
                                <span className="text-blue-gray-500 pointer-events-none mx-2 select-none font-sans text-sm font-normal leading-normal antialiased">
                                    /
                                </span>
                            ) : null}
                        </li>
                    ))}
                </ol>
            </nav>
        </div>
    );
}

export default Breadcrumb;
