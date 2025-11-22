export const Verification = async (uuid: string, mac: string): Promise<{ data: string }> => {
    const res = await fetch("https://api.luola.me/api/order/vaild", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            uuid,
            mac
        })
    })
    return res.json()
}

export const Download = async (uuid: string, mac: string, version: string): Promise<any> => {
    const res = await fetch("https://api.luola.me/api/download/jetbra", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            uuid,
            mac,
            version
        })
    })
    // 检测返回的是json还是zip
    const contentType = res.headers.get('Content-Type');
    if (contentType && contentType.includes('application/json')) {
        return res.json();
    }
    return res.blob();
}

export const getConde = async (pluginName: string): Promise<any> => {
    const res = await (await fetch(`https://api.luola.me/api/plugin/code`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            code: pluginName,
        })
    })).json()
    return res
}