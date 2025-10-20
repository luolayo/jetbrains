import {defineStore} from "pinia";
import {GetDeviceID} from "../../wailsjs/go/main/App";
import {reactive, ref} from "vue";

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
            throw new Error(data.error)
        }
        Object.assign(drviceInfo, data.machineCode)
    }

    const getCode = async () => {
        const res = await fetch("http://localhost:8080/api/code",{
            method: "GET",
        })
        code.value = (await res.json()).data.code
    }

    return {drviceInfo, getDrviceInfo,getCode,code}
})