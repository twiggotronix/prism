import { ReactNode } from "react";
import Breadcrumb from "../breadcrumb/breadcrumb";
import SideMenu from "../sideMenu/sideMenu";

export interface LayoutProps {
    pageTitle: string;
    children: ReactNode;
}

export function Layout({ children, pageTitle }: LayoutProps) {
    return (
        <div>
            <SideMenu />
            <div className="capitalize">
                <div className="p-4 xl:ml-80">
                    <div className="mb-2 block w-full max-w-full rounded-xl bg-transparent px-0 py-1 text-white shadow-none transition-all">
                        <Breadcrumb></Breadcrumb>
                        <h6 className="text-blue-gray-900 mt-2 block font-sans text-base font-semibold leading-relaxed tracking-normal antialiased">
                            {pageTitle}
                        </h6>
                    </div>
                    {children}
                </div>
            </div>
        </div>
    );
}

export default Layout;
