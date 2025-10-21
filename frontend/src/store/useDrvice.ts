import {defineStore} from "pinia";
import {GetDeviceID} from "../../wailsjs/go/main/App";
import {reactive, ref} from "vue";
import {ElMessage} from "element-plus";

interface DeviceIDType {
    uuid: string;
    boardSerialNumber: string;
    cpuSerialNumber: string;
    diskSerialNumber: string;
    mac: string;
}

export const useDrvice = defineStore('drvice', () => {
    const drviceInfo = reactive<DeviceIDType>(
        {
            uuid: '',
            boardSerialNumber: '',
            cpuSerialNumber: '',
            diskSerialNumber: '',
            mac: '',
        }
    );
    const code = ref<string>('');

    const getDrviceInfo = async () => {
        const data = await GetDeviceID()
        if (data?.error) {
            ElMessage.error(
                `获取设备信息失败: WMIC未安装，请查看问题列表的解决办法`
            )
            return
        }
        Object.assign(drviceInfo, data.machineCode)
    }

    const getCode = async () => {
        const res = await fetch("https://api.luola.me/api/code",{
            method: "GET",
        })
        code.value = (await res.json()).data.code
    }

    return {drviceInfo, getDrviceInfo,getCode,code}
})