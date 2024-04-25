import { faWarning } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Alert } from "@material-tailwind/react";
import styles from "./input-error.module.scss";

export interface InputErrorProps {
    message: string;
}

export function InputError({ message }: InputErrorProps) {
    return (
        <div className={styles["container"]}>
            <Alert
                open={true}
                color={"red"}
                icon={<FontAwesomeIcon icon={faWarning} className="h-6 w-6" />}
            >
                {message}
            </Alert>
        </div>
    );
}

export default InputError;
