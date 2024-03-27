import styles from "./proxies.module.scss";

/* eslint-disable-next-line */
export interface ProxiesProps {}

export function Proxies(props: ProxiesProps) {
    return (
        <div className={styles["container"]}>
            <h1>Welcome to Proxies!</h1>
        </div>
    );
}

export default Proxies;
