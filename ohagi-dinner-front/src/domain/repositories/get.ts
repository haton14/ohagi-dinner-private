export const get = <T>(
    input: RequestInfo,
    init?: RequestInit<RequestInitCfProperties>
): Promise<T> => {
    return new Promise((resolve, reject) => {
        fetch(input, init).then((res) => {
            resolve(res.json<T>())
        }).catch((error) => {
            reject(error)
        })
    })
}
