import { Button } from "@material-tailwind/react";

/* eslint-disable-next-line */
export interface SideMenuProps {}

export function SideMenu(props: SideMenuProps) {
    return (
        <div className="border-blue-gray-100 fixed inset-0 z-50 my-4 ml-4 h-[calc(100vh-32px)] w-72 -translate-x-80 rounded-xl border bg-gradient-to-br from-gray-800 to-gray-900 transition-transform duration-300 xl:translate-x-0">
            <h1 className="block text-center font-sans text-base font-semibold leading-relaxed tracking-normal text-white antialiased">
                Prism
            </h1>
            <div className="m-4">
                <ul className="flex flex-col gap-1">
                    <li>
                        <Button className="flex w-full select-none items-center gap-4 rounded-lg bg-gradient-to-tr from-gray-900 to-gray-800 py-3 px-4 text-center align-middle font-sans text-xs font-bold capitalize text-white shadow-md shadow-gray-900/10 transition-all hover:shadow-lg hover:shadow-gray-900/20 active:opacity-[0.85] disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none">
                            <i className="fa-solid fa-house"></i> Dashboard
                        </Button>
                    </li>
                </ul>
            </div>
        </div>
    );
}

export default SideMenu;
