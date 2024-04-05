import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { type Proxy } from "../../models/proxy";
import { proxyKey } from "../../proxy.key";
import { ProxyService } from "../../services/proxy.service";
import ProxyTable from "../proxy-table/proxy-table";
import styles from "./proxies.module.scss";

export function Proxies() {
    const { data: proxyData, isLoading } = useQuery({
        queryKey: proxyKey.getProxies,
        queryFn: ProxyService.getProxies,
    });
    const [tableData, setTableData] = useState<Proxy[] | undefined>();
    useEffect(() => {
        setTableData(proxyData);
    }, [proxyData]);
    if (isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <div className={styles["container"]}>
            {tableData != null ? <ProxyTable tableData={tableData} /> : null}
        </div>
    );
}

export default Proxies;
