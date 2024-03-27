import styles from "./breadcrumb.module.scss";

/* eslint-disable-next-line */
export interface BreadcrumbProps {}

export function Breadcrumb(props: BreadcrumbProps) {
    return (
        <div className={styles["container"]}>
            <nav aria-label="breadcrumb" className="w-max">
                <ol className="bg-blue-gray-50 flex w-full flex-wrap items-center rounded-md bg-opacity-60 px-4 py-2">
                    <li className="text-blue-gray-900 hover:text-light-blue-500 flex cursor-pointer items-center font-sans text-sm font-normal leading-normal antialiased transition-colors duration-300">
                        <a href="#" className="opacity-60">
                            Docs
                        </a>
                        <span className="text-blue-gray-500 pointer-events-none mx-2 select-none font-sans text-sm font-normal leading-normal antialiased">
                            /
                        </span>
                    </li>
                    <li className="text-blue-gray-900 hover:text-light-blue-500 flex cursor-pointer items-center font-sans text-sm font-normal leading-normal antialiased transition-colors duration-300">
                        <a href="#" className="opacity-60">
                            Components
                        </a>
                        <span className="text-blue-gray-500 pointer-events-none mx-2 select-none font-sans text-sm font-normal leading-normal antialiased">
                            /
                        </span>
                    </li>
                    <li className="text-blue-gray-900 hover:text-light-blue-500 flex cursor-pointer items-center font-sans text-sm font-normal leading-normal antialiased transition-colors duration-300">
                        <a href="#">Breadcrumbs</a>
                    </li>
                </ol>
            </nav>
        </div>
    );
}

export default Breadcrumb;
