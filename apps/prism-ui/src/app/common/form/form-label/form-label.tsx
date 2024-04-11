import type { RegisterOptions, ValidationValueMessage } from "react-hook-form";
import styles from "./form-label.module.scss";

export interface FormLabelProps {
    htmlFor: string;
    label: string;
    validation?: RegisterOptions;
}

export function FormLabel({ htmlFor, label, validation }: FormLabelProps) {
    const required =
        (validation?.required as ValidationValueMessage<boolean>)?.value ??
        false;
    return (
        <div className={styles["container"]}>
            <label htmlFor={htmlFor}>
                {label}
                {required && <span className="mr-1 text-red-200">*</span>}
            </label>
        </div>
    );
}

export default FormLabel;
