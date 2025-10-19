import {defineStore} from "pinia";
import {GetDeviceID} from "../../wailsjs/go/main/App";
import {reactive} from "vue";

export const useDrvice = defineStore('drvice', () => {
    const drviceInfo = reactive<DeviceIDType>(
        {
            uuid: '',
            boardSerialNumber: '',
            cpuSerialNumber: '',
            diskSerialNumber: '',
            mac: ''
        }
    );

    const getDrviceInfo = async () => {
        const data = await GetDeviceID()
        if (data?.error) {
            throw new Error(data.error)
        }
        Object.assign(drviceInfo, data.machineCode)
    }

    return {drviceInfo, getDrviceInfo}
})