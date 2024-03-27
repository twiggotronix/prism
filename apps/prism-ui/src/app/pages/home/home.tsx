import Proxies from "../../proxies/components/proxies/proxies";
import Layout from "../common/layout/layout";
import styles from "./home.module.scss";

/* eslint-disable-next-line */
export interface HomeProps {}

export function Home(props: HomeProps) {
    return (
        <div className={styles["container"]}>
            <Layout pageTitle="Home page">
                <Proxies />
            </Layout>
        </div>
    );
}

export default Home;
