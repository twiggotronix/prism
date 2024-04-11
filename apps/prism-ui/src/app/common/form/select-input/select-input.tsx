import { useForm, type RegisterOptions } from "react-hook-form";
import FormLabel from "../form-label/form-label";
import styles from "./select-input.module.scss";

export interface SelectInputProps {
    label: string;
    defaultValue?: string;
    validation?: RegisterOptions;
    options: { value: string; label: string }[];
}

export function SelectInput({
    id,
    label,
    validation,
    defaultValue,
    name,
    options,
}: SelectInputProps & Pick<HTMLSelectElement, "id" | "name">) {
    const { register } = useForm();

    const methods = register(name, validation);
    return (
        <div className={styles["container"]}>
            <FormLabel
                htmlFor={id}
                label={label}
                validation={validation}
            ></FormLabel>
            <select
                id={id}
                defaultValue={defaultValue}
                {...methods}
                className="border-blue-gray-200 text-blue-gray-700 placeholder-shown:border-blue-gray-200 placeholder-shown:border-t-blue-gray-200 disabled:bg-blue-gray-50 peer h-full w-full rounded-[7px] border border-t-transparent bg-transparent px-3 py-2.5 font-sans text-sm font-normal outline outline-0 transition-all placeholder-shown:border empty:!bg-gray-900 focus:border-2 focus:border-gray-900 focus:border-t-transparent focus:outline-0 disabled:border-0"
            >
                {options.map(option => (
                    <option value={option.value}>{option.label}</option>
                ))}
            </select>
        </div>
    );
}

export default SelectInput;
