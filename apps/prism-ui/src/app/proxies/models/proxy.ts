export enum Method {
    Get = "GET"
}
export type Proxy = {
    id: number
    name: string
    path: string
    method: Method
    source: string
}