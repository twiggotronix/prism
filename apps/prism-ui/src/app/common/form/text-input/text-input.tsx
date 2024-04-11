import { faWarning } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Alert, Input } from "@material-tailwind/react";
import { useEffect, useState } from "react";
import { useFormContext, type RegisterOptions } from "react-hook-form";
import FormLabel from "../form-label/form-label";
import styles from "./text-input.module.scss";

export interface TextInputProps {
    label: string;
    validation?: RegisterOptions;
}

export function TextInput({
    placeholder,
    id,
    label,
    validation,
    defaultValue,
    name,
}: TextInputProps &
    Pick<HTMLInputElement, "placeholder" | "id" | "defaultValue" | "name">) {
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
            <FormLabel htmlFor={id} label={label} validation={validation} />
            {errorMessages.length > 0 &&
                errorMessages.map(errorMessage => (
                    <InputError message={errorMessage} />
                ))}
            <Input
                placeholder={placeholder}
                id={id}
                defaultValue={defaultValue}
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

const InputError = ({ message }: { message: string }) => {
    return (
        <Alert
            open={true}
            color={"red"}
            icon={<FontAwesomeIcon icon={faWarning} className="h-6 w-6" />}
        >
            {message}
        </Alert>
    );
};
