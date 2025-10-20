/**
 * 解密函数
 * @param ciphertext - base64 编码的密文
 * @param keyHex - 十六进制格式的密钥
 * @returns 解密后的明文
 */
export async function decrypt(ciphertext: string, keyHex: string): Promise<string> {
    try {
        // 将 base64 密文解码为 ArrayBuffer
        const data = base64ToArrayBuffer(ciphertext);

        // 将十六进制密钥转换为 ArrayBuffer
        const keyBuffer = hexToArrayBuffer(keyHex);

        // GCM nonce 大小为 12 字节
        const nonceSize = 12;

        // 提取 nonce（前 12 字节）
        const nonce = data.slice(0, nonceSize);

        // 提取密文和认证标签（剩余部分）
        const encryptedData = data.slice(nonceSize);

        // 导入密钥
        const cryptoKey = await crypto.subtle.importKey(
            'raw',
            keyBuffer,
            { name: 'AES-GCM' },
            false,
            ['decrypt']
        );

        // 解密
        const decryptedBuffer = await crypto.subtle.decrypt(
            {
                name: 'AES-GCM',
                iv: nonce,
                tagLength: 128 // 16 字节 = 128 位
            },
            cryptoKey,
            encryptedData
        );

        // 将 ArrayBuffer 转换为字符串
        const plaintext = new TextDecoder('utf-8').decode(decryptedBuffer);
        return plaintext;
    } catch (error) {
        throw new Error(`解密失败: ${error instanceof Error ? error.message : String(error)}`);
    }
}

/**
 * 将 base64 字符串转换为 ArrayBuffer
 */
function base64ToArrayBuffer(base64: string): ArrayBuffer {
    const binaryString = atob(base64);
    const bytes = new Uint8Array(binaryString.length);
    for (let i = 0; i < binaryString.length; i++) {
        bytes[i] = binaryString.charCodeAt(i);
    }
    return bytes.buffer;
}

/**
 * 将十六进制字符串转换为 ArrayBuffer
 */
function hexToArrayBuffer(hex: string): ArrayBuffer {
    const bytes = new Uint8Array(hex.length / 2);
    for (let i = 0; i < hex.length; i += 2) {
        bytes[i / 2] = parseInt(hex.substring(i, i + 2), 16);
    }
    return bytes.buffer;
}