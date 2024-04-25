import { useEffect, useState } from "react";
import { useFormContext, type RegisterOptions } from "react-hook-form";
import FormLabel from "../form-label/form-label";
import InputError from "../input-error/input-error";
import styles from "./select-input.module.scss";

export interface SelectInputProps {
    label: string;
    validation?: RegisterOptions;
    options: { value: string; label: string }[];
}

export function SelectInput({
    id,
    label,
    validation,
    name,
    options,
}: SelectInputProps & Pick<HTMLSelectElement, "id" | "name">) {
    const {
        register,
        formState: { errors },
    } = useFormContext();
    const [errorMessages, setErrorMessages] = useState<string[]>([]);
    useEffect(() => {
        setErrorMessages(
            Object.keys(errors)
                .filter(inputName => inputName.startsWith(name))
                .reduce((cur, key) => {
                    cur.push(errors[key]?.message as string);
                    return cur;
                }, [] as string[]),
        );
    }, [errors, name, setErrorMessages]);

    return (
        <div className={styles["container"]}>
            {errorMessages.length > 0 &&
                errorMessages.map(errorMessage => (
                    <InputError message={errorMessage} />
                ))}
            <FormLabel
                htmlFor={id}
                label={label}
                validation={validation}
            ></FormLabel>
            <select
                id={id}
                {...register(name, validation)}
                className="border-blue-gray-200 text-blue-gray-700 placeholder-shown:border-blue-gray-200 placeholder-shown:border-t-blue-gray-200 disabled:bg-blue-gray-50 peer h-full w-full rounded-[7px] border border-t-transparent bg-transparent px-3 py-2.5 font-sans text-sm font-normal outline outline-0 transition-all placeholder-shown:border empty:!bg-gray-900 focus:border-2 focus:border-gray-900 focus:border-t-transparent focus:outline-0 disabled:border-0"
            >
                {options.map((option, index) => (
                    <option value={option.value} key={`${id}-${index}`}>
                        {option.label}
                    </option>
                ))}
            </select>
        </div>
    );
}

export default SelectInput;
