import { Input } from "@material-tailwind/react";
import { useEffect, useState, type HTMLInputTypeAttribute } from "react";
import { useFormContext, type RegisterOptions } from "react-hook-form";
import FormLabel from "../form-label/form-label";
import InputError from "../input-error/input-error";
import styles from "./text-input.module.scss";

export interface TextInputProps {
    label: string;
    validation?: RegisterOptions;
    type?: HTMLInputTypeAttribute;
}

export function TextInput({
    placeholder,
    id,
    label,
    validation,
    name,
    type,
}: TextInputProps & Pick<HTMLInputElement, "placeholder" | "id" | "name">) {
    const {
        register,
        formState: { errors },
    } = useFormContext();
    const [errorMessages, setErrorMessages] = useState<string[]>([]);

    useEffect(() => {
        const err = Object.keys(errors)
            .filter(inputName => inputName.startsWith(name))
            .reduce((cur, key) => {
                cur.push(errors[key]?.message as string);
                return cur;
            }, [] as string[]);
        setErrorMessages(err);
    }, [errors, name, setErrorMessages]);

    return (
        <div className={styles["container"]}>
            <FormLabel htmlFor={id} label={label} validation={validation} />
            {errorMessages.length > 0 &&
                errorMessages.map(errorMessage => (
                    <InputError message={errorMessage} />
                ))}
            <Input
                type={type ?? "text"}
                placeholder={placeholder}
                id={id}
                size="md"
                className="!border-t-blue-gray-200 focus:!border-t-gray-900"
                labelProps={{
                    className: "before:content-none after:content-none",
                }}
                crossOrigin={undefined}
                {...register(name, validation)}
            ></Input>
        </div>
    );
}

export default TextInput;
