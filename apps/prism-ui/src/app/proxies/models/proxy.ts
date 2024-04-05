export enum Method {
    Get = "GET",
    Post = "POST",
    Put = "PUT",
    Delete = "DELETE",
}
export type Proxy = {
    id: number;
    name: string;
    path: string;
    method: Method;
    source: string;
};
