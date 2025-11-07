export const getVersion = async (code: string): Promise<any> => {
    const res = await fetch(`https://data.services.jetbrains.com/products/releases?code=${code}&latest=false&type=release`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
    return res.json()
}
